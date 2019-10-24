package gozw

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gozwave/gozw/cc"
	"github.com/gozwave/gozw/cc/association"
	"github.com/gozwave/gozw/cc/battery"
	manufacturerspecific "github.com/gozwave/gozw/cc/manufacturer-specific"
	manufacturerspecificv2 "github.com/gozwave/gozw/cc/manufacturer-specific-v2"
	"github.com/gozwave/gozw/cc/security"
	"github.com/gozwave/gozw/cc/version"
	versionv2 "github.com/gozwave/gozw/cc/version-v2"
	versionv3 "github.com/gozwave/gozw/cc/version-v3"
	"github.com/gozwave/gozw/protocol"
	"github.com/gozwave/gozw/serialapi"
	"github.com/gozwave/gozw/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

type GoZWNode interface {
	loadFromDb() error
	removeFromDb() error
	saveToDb() error
}

type PairingProgressUpdate struct {
	// How command classes we've heard from
	InterviewedCommandClassCount int

	// How many command classes have been reported from NIF
	ReportedCommandClassCount int
}

type Node struct {
	GoZWNode
	NodeID byte

	Capability          byte
	BasicDeviceClass    byte
	GenericDeviceClass  byte
	SpecificDeviceClass byte

	Failing bool

	InterviewedCommandClassCount int
	CommandClasses               cc.CommandClassSet

	NetworkKeySent bool

	ManufacturerID uint16
	ProductTypeID  uint16
	ProductID      uint16

	QueryStageSecurity     bool
	QueryStageManufacturer bool
	QueryStageVersions     bool

	queryStageSecurityComplete     chan bool
	queryStageManufacturerComplete chan bool
	queryStageVersionsComplete     chan bool

	client *Client
}

func NewNode(client *Client, nodeID byte) (*Node, error) {
	node := &Node{
		NodeID: nodeID,

		CommandClasses: cc.CommandClassSet{},

		QueryStageSecurity:     false,
		QueryStageManufacturer: false,
		QueryStageVersions:     false,

		queryStageSecurityComplete:     make(chan bool),
		queryStageManufacturerComplete: make(chan bool),
		queryStageVersionsComplete:     make(chan bool),

		client: client,
	}

	err := node.loadFromDb()
	if err != nil {
		initErr := node.initialize()
		if initErr != nil {
			return nil, initErr
		}

		err = node.saveToDb()
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}

func (n *Node) loadFromDb() error {
	var data []byte
	err := n.client.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		data = bucket.Get([]byte{n.NodeID})

		if len(data) == 0 {
			return errors.New("Node not found")
		}

		return nil
	})

	if err != nil {
		return err
	}

	err = msgpack.Unmarshal(data, n)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) initialize() error {
	nodeInfo, err := n.client.serialAPI.GetNodeProtocolInfo(n.NodeID)
	if err != nil {
		fmt.Println(err)
	} else {
		n.client.l.Debug("setting from node protocol info")
		n.setFromNodeProtocolInfo(nodeInfo)
	}

	if n.NodeID == 1 {
		// self is never failing
		n.Failing = false
	} else {
		failing, err := n.client.serialAPI.IsFailedNode(n.NodeID)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		n.Failing = failing
	}

	return n.saveToDb()

}

func (n *Node) removeFromDb() error {
	return n.client.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		err := bucket.Delete([]byte{n.NodeID})
		return err
	})
}

func (n *Node) saveToDb() error {
	data, err := msgpack.Marshal(n)
	if err != nil {
		return err
	}

	return n.client.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		return bucket.Put([]byte{n.NodeID}, data)
	})
}

func (n *Node) IsSecure() bool {
	return n.CommandClasses.Supports(cc.Security)
}

func (n *Node) IsListening() bool {
	return n.Capability&0x80 == 0x80
}

func (n *Node) GetBasicDeviceClassName() string {
	return protocol.GetBasicDeviceTypeName(n.BasicDeviceClass)
}

func (n *Node) GetGenericDeviceClassName() string {
	return protocol.GetGenericDeviceTypeName(n.GenericDeviceClass)
}

func (n *Node) GetSpecificDeviceClassName() string {
	return protocol.GetSpecificDeviceTypeName(n.GenericDeviceClass, n.SpecificDeviceClass)
}

func (n *Node) SendCommand(command cc.Command) error {
	commandClass := cc.CommandClassID(command.CommandClassID())

	if commandClass == cc.Security {
		switch command.(type) {
		case *security.CommandsSupportedGet, *security.CommandsSupportedReport:
			return n.client.SendDataSecure(n.NodeID, command)
		}
	}

	if !n.CommandClasses.Supports(commandClass) {
		return errors.New("Command class not supported")
	}

	if n.CommandClasses.IsSecure(commandClass) {
		return n.client.SendDataSecure(n.NodeID, command)
	}

	return n.client.SendData(n.NodeID, command)
}

func (n *Node) SendRawCommand(payload []byte) error {
	commandClass := cc.CommandClassID(payload[0])

	if !n.CommandClasses.Supports(commandClass) {
		return errors.New("Command class not supported")
	}

	if n.CommandClasses.IsSecure(commandClass) {
		return n.client.SendDataSecure(n.NodeID, util.ByteMarshaler(payload))
	}

	return n.client.SendData(n.NodeID, util.ByteMarshaler(payload))
}

func (n *Node) AddAssociation(groupID byte, nodeIDs ...byte) error {
	// sort of an arbitrary limit for now, but I'm not sure what it should be
	if len(nodeIDs) > 20 {
		return errors.New("Too many associated nodes")
	}

	fmt.Println("Associating")

	return n.SendCommand(&association.Set{
		GroupingIdentifier: groupID,
		NodeId:             nodeIDs,
	})
}

func (n *Node) LoadSupportedSecurityCommands() error {
	return n.client.SendDataSecure(n.NodeID, &security.CommandsSupportedGet{})
}

func (n *Node) RequestNodeInformationFrame() error {

	_, err := n.client.serialAPI.RequestNodeInfo(n.NodeID)
	return err
}

// To optimize the interview process we'll be interview command classes
// in the following order:
// 1) SecurityCommandClass
// 2) VersionCommandClass
// 3) Remaining insecure command classes
// 4) Secure command classes
func commandClassesInOrderToInterview(set cc.CommandClassSet) []cc.CommandClassID {
	insecure := set.ListBySecureStatus(false)

	securityccidx := -1
	versionccidx := -1
	remaininginsecure := []cc.CommandClassID{}
	for idx, ccid := range insecure {
		if isVersionCommandClassID(ccid) {
			versionccidx = idx
			continue
		}
		if isSecurityCommandClassID(ccid) {
			securityccidx = idx
			continue
		}
		remaininginsecure = append(remaininginsecure, ccid)
	}

	if securityccidx > -1 && versionccidx > -1 {
		insecure = append([]cc.CommandClassID{
			insecure[securityccidx],
			insecure[versionccidx],
		},
			remaininginsecure...)
	}

	secure := set.ListBySecureStatus(true)
	return append(insecure, secure...)

}

func isVersionCommandClassID(id cc.CommandClassID) bool {
	// in practice these three command class IDs are the same
	return id == cc.Version ||
		id == cc.VersionV2 ||
		id == cc.VersionV3
}

func isSecurityCommandClassID(id cc.CommandClassID) bool {
	// TODO(zacatac): These are different command class IDs.
	// Need to refer to documentation on Security2 command class.
	//
	// It's possible that newer devices which support
	// S2 security report that they support both S0 and S2
	return id == cc.Security ||
		id == cc.Security2
}

func (n *Node) LoadCommandClassVersions() error {
	// The pairing process is incredibly sensitive to timeouts or
	// delays. Certain values may cause the interview process to
	// fail intermittently.

	// Defines how long to wait between checking whether a given CC
	// has reported a version
	versionReportedQueryPeriod := time.Millisecond * 10

	inOrder := commandClassesInOrderToInterview(n.CommandClasses)
	for _, commandClassID := range inOrder {
		var cmd cc.Command
		switch n.CommandClasses.GetVersion(cc.Version) {
		case 0x02:
			cmd = &versionv2.CommandClassGet{
				RequestedCommandClass: byte(commandClassID),
			}
		case 0x03:
			cmd = &versionv3.CommandClassGet{
				RequestedCommandClass: byte(commandClassID),
			}
		default:
			cmd = &version.CommandClassGet{
				RequestedCommandClass: byte(commandClassID),
			}
		}

		var err error
		if n.IsSecure() {
			err = n.client.SendDataSecure(n.NodeID, cmd)
		} else {
			err = n.client.SendData(n.NodeID, cmd)
		}
		if err != nil {
			return err
		}

		// Record time that we started waiting for a version report
		waitingForVersionSince := time.Now()

		// Wait until either the version has been reported, or we've waited too long
		// 5 seconds is used to be sure any collisions (and thus CAN timeouts) are over
		for {
			// Check to see if this command class has been reported
			if v := n.CommandClasses.GetVersion(commandClassID); v != 0 {
				break
			}

			if waitingForVersionSince.Add(time.Second * 5).Before(time.Now()) {
				n.client.l.Warn("stopped waiting for confirmation of version command class",
					zap.String("command class ID", fmt.Sprintf("%02x", commandClassID)),
				)
				break
			}

			time.Sleep(versionReportedQueryPeriod)
		}

	}

	return nil
}

func (n *Node) LoadManufacturerInfo() error {
	return n.SendCommand(&manufacturerspecific.Get{})
}

func (n *Node) nextQueryStage() {
	if !n.QueryStageSecurity && n.IsSecure() {
		err := n.LoadSupportedSecurityCommands()
		if err != nil {
			n.client.l.Error("load-supported-commands-failure", zap.Error(err))
		}
		return
	}

	if !n.QueryStageVersions {
		err := n.LoadCommandClassVersions()
		if err != nil {
			n.client.l.Error("load-supported-commands-failure", zap.Error(err))
		}
		return
	}

	if !n.QueryStageManufacturer {
		err := n.LoadManufacturerInfo()
		if err != nil {
			n.client.l.Error("load-supported-commands-failure", zap.Error(err))
		}
		return
	}
}

func (n *Node) emitNodeEvent(event cc.Command) {
	n.client.EventCallback(n.client, n.NodeID, event)
	// buf, err := event.MarshalBinary()
	// if err != nil {
	// 	fmt.Printf("error encoding: %v\n", err)
	// 	return
	// }

	// n.client.EventBus.Publish("nodeCommand", n.NodeID, buf)
}

func (n *Node) receiveControllerUpdate(update serialapi.ControllerUpdate) {
	n.setFromApplicationControllerUpdate(update)
	err := n.saveToDb()
	if err != nil {
		n.client.l.Error("load-supported-commands-failure", zap.Error(err))
	}
}

func (n *Node) setFromAddNodeCallback(nodeInfo *serialapi.AddRemoveNodeCallback) {
	n.NodeID = nodeInfo.Source
	n.BasicDeviceClass = nodeInfo.Basic
	n.GenericDeviceClass = nodeInfo.Generic
	n.SpecificDeviceClass = nodeInfo.Specific

	for _, cmd := range nodeInfo.CommandClasses {
		n.CommandClasses.Add(cc.CommandClassID(cmd))
	}

	err := n.saveToDb()
	if err != nil {
		n.client.l.Error("save-to-db-failed", zap.Error(err))
	}
}

func (n *Node) setFromApplicationControllerUpdate(nodeInfo serialapi.ControllerUpdate) {
	n.BasicDeviceClass = nodeInfo.Basic
	n.GenericDeviceClass = nodeInfo.Generic
	n.SpecificDeviceClass = nodeInfo.Specific

	for _, cmd := range nodeInfo.CommandClasses {
		n.CommandClasses.Add(cc.CommandClassID(cmd))
	}

	err := n.saveToDb()
	if err != nil {
		n.client.l.Error("save-to-db-failed", zap.Error(err))
	}
}

func (n *Node) setFromNodeProtocolInfo(nodeInfo *serialapi.NodeProtocolInfo) {
	n.Capability = nodeInfo.Capability
	n.BasicDeviceClass = nodeInfo.BasicDeviceClass
	n.GenericDeviceClass = nodeInfo.GenericDeviceClass
	n.SpecificDeviceClass = nodeInfo.SpecificDeviceClass

	err := n.saveToDb()
	if err != nil {
		n.client.l.Error("save-to-db-failed", zap.Error(err))
	}
}

func (n *Node) receiveSecurityCommandsSupportedReport(cmd security.CommandsSupportedReport) {
	for _, supported := range cmd.CommandClassSupport {
		n.CommandClasses.SetSecure(cc.CommandClassID(supported), true)
	}

	select {
	case n.queryStageSecurityComplete <- true:
	default:
	}

	n.QueryStageSecurity = true
	err := n.saveToDb()
	if err != nil {
		n.client.l.Error("save-to-db-failed", zap.Error(err))
	}
	n.nextQueryStage()
}

func (n *Node) receiveManufacturerInfo(mfgId, productTypeId, productId uint16) {
	n.ManufacturerID = mfgId
	n.ProductTypeID = productTypeId
	n.ProductID = productId

	select {
	case n.queryStageManufacturerComplete <- true:
	default:
	}

	n.QueryStageManufacturer = true
	err := n.saveToDb()
	if err != nil {
		n.client.l.Error("save-to-db-failed", zap.Error(err))
	}
	n.nextQueryStage()
}

func (n *Node) receiveCommandClassVersion(id cc.CommandClassID, version uint8) {
	n.CommandClasses.SetVersion(id, version)
	n.InterviewedCommandClassCount++

	if n.CommandClasses.AllVersionsReceived() {
		select {
		case n.queryStageVersionsComplete <- true:
		default:
		}

		n.QueryStageVersions = true
		defer n.nextQueryStage()
	}

	err := n.saveToDb()
	if err != nil {
		n.client.l.Error("save-to-db-failed", zap.Error(err))
	}
}

func (n *Node) receiveApplicationCommand(cmd serialapi.ApplicationCommand) {
	commandClassID := cc.CommandClassID(cmd.CommandData[0])
	ver := n.CommandClasses.GetVersion(commandClassID)
	if ver == 0 {
		ver = 1

		if !(commandClassID == cc.Version || commandClassID == cc.Security) {
			n.client.l.Error("no version loaded", zap.String("commandClass", commandClassID.String()))
		}
	}

	n.client.l.Info("command data passed to parser", zap.String("data", fmt.Sprintf("%x", cmd.CommandData)))
	command, err := cc.Parse(ver, cmd.CommandData)
	if err != nil {
		n.client.l.Error("error parsing command class", zap.Error(err))
		return
	}

	n.client.l.Debug("device command received", zap.String("commandClass", command.CommandClassID().String()), zap.String("command", command.CommandIDString()))

	switch command.(type) {

	case *battery.Report:
		if cmd.CommandData[2] == 0xFF {
			fmt.Printf("Node %d: low battery alert\n", n.NodeID)
		} else {
			fmt.Printf("Node %d: battery level is %d\n", n.NodeID, command.(*battery.Report))
		}
		n.emitNodeEvent(command)

	case *security.CommandsSupportedReport:
		fmt.Println("security commands supported report")
		n.receiveSecurityCommandsSupportedReport(*command.(*security.CommandsSupportedReport))
		fmt.Println(n.GetSupportedSecureCommandClassStrings())

	case *manufacturerspecific.Report:
		spew.Dump(command.(*manufacturerspecific.Report))
		report := *command.(*manufacturerspecific.Report)
		n.receiveManufacturerInfo(report.ManufacturerId, report.ProductTypeId, report.ProductId)
		n.emitNodeEvent(command)

	case *manufacturerspecificv2.Report:
		spew.Dump(command.(*manufacturerspecificv2.Report))
		report := *command.(*manufacturerspecificv2.Report)
		n.receiveManufacturerInfo(report.ManufacturerId, report.ProductTypeId, report.ProductId)
		n.emitNodeEvent(command)

	case *version.CommandClassReport:
		spew.Dump(command.(*version.CommandClassReport))
		report := command.(*version.CommandClassReport)
		n.receiveCommandClassVersion(cc.CommandClassID(report.RequestedCommandClass), report.CommandClassVersion)
		err := n.saveToDb()
		if err != nil {
			n.client.l.Error("save-to-db-failed", zap.Error(err))
		}

	case *versionv2.CommandClassReport:
		spew.Dump(command.(*versionv2.CommandClassReport))
		report := command.(*versionv2.CommandClassReport)
		n.receiveCommandClassVersion(cc.CommandClassID(report.RequestedCommandClass), report.CommandClassVersion)
		err := n.saveToDb()
		if err != nil {
			n.client.l.Error("save-to-db-failed", zap.Error(err))
		}

		// case alarm.Report:
		// 	spew.Dump(command.(alarm.Report))
		//
		// case usercode.Report:
		// 	spew.Dump(command.(usercode.Report))
		//
		// case doorlock.OperationReport:
		// 	spew.Dump(command.(doorlock.OperationReport))
		//
		// case thermostatmode.Report:
		// 	spew.Dump(command.(thermostatmode.Report))
		//
		// case thermostatoperatingstate.Report:
		// 	spew.Dump(command.(thermostatoperatingstate.Report))
		//
		// case thermostatsetpoint.Report:
		// 	spew.Dump(command.(thermostatsetpoint.Report))

	default:
		n.emitNodeEvent(command)
	}
}

func (n *Node) String() string {
	str := fmt.Sprintf("Node %d: \n", n.NodeID)
	str += fmt.Sprintf("  Failing? %t\n", n.Failing)
	str += fmt.Sprintf("  Is listening? %t\n", n.IsListening())
	str += fmt.Sprintf("  Is secure? %t\n", n.IsSecure())
	str += fmt.Sprintf("  Basic device class: %s\n", n.GetBasicDeviceClassName())
	str += fmt.Sprintf("  Generic device class: %s\n", n.GetGenericDeviceClassName())
	str += fmt.Sprintf("  Specific device class: %s\n", n.GetSpecificDeviceClassName())
	str += fmt.Sprintf("  Manufacturer ID: %#x\n", n.ManufacturerID)
	str += fmt.Sprintf("  Product Type ID: %#x\n", n.ProductTypeID)
	str += fmt.Sprintf("  Product ID: %#x\n", n.ProductID)
	str += fmt.Sprintf("  Supported command classes:\n")

	for _, cmd := range n.CommandClasses {
		if cmd.Secure {
			str += fmt.Sprintf("    - %s (v%d) (secure)\n", cmd.CommandClass.String(), cmd.Version)
		} else {
			str += fmt.Sprintf("    - %s (v%d)\n", cmd.CommandClass.String(), cmd.Version)
		}
	}

	return str
}

func (n *Node) GetSupportedCommandClassStrings() []string {
	strings := commandClassSetToStrings(n.CommandClasses.ListBySecureStatus(false))
	if len(strings) == 0 {
		return []string{
			"None (probably not loaded; need to request a NIF)",
		}
	}

	return strings
}

func (n *Node) GetSupportedSecureCommandClassStrings() []string {
	strings := commandClassSetToStrings(n.CommandClasses.ListBySecureStatus(true))
	return strings
}

func commandClassSetToStrings(commandClasses []cc.CommandClassID) []string {
	if len(commandClasses) == 0 {
		return []string{}
	}

	ccStrings := []string{}

	for _, cmd := range commandClasses {
		ccStrings = append(ccStrings, cmd.String())
	}

	return ccStrings
}
