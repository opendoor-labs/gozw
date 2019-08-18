package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gozwave/gozw"
)

var networkKey = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func main() {
	devicePath := "/dev/tty.usbmodem141101"
	if p := os.Getenv("GOZW_DEVICE_PATH"); p != "" {
		devicePath = p
	}
	fmt.Println("Adding the same node twice will update the network appropriately, but the client will keep duplicate entries of the nodes in its internal map. PruneNodes removes the node from the client's internal map and DB")

	client, err := gozw.NewDefaultClient("/tmp/data.db", devicePath, 115200, networkKey)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	spew.Dump(client.Controller)

	for _, node := range client.Nodes() {
		fmt.Println(node.String())
	}

	fmt.Println("removing node, put device in unpairing mode")
	if _, err := client.RemoveNode(); err != nil {
		log.Fatalf("failed to remove node: %v", err)
	}
	fmt.Println("-----------------------------------------")

	fmt.Println("adding node, put device in pairing mode")
	_, err = client.AddNode()
	if err != nil {
		log.Fatalf("failed to add node: %v", err)
	}

	fmt.Printf("%+v", client.Nodes())
	client.PruneNodes()
	fmt.Printf("%+v", client.Nodes())
	fmt.Println("=========================================")

	fmt.Println("removing node, put device in unpairing mode")
	if _, err = client.RemoveNode(); err != nil {
		log.Fatalf("failed to remove node: %v", err)
	}
	fmt.Println("-----------------------------------------")

	fmt.Println("adding node, put device in pairing mode")
	_, err = client.AddNode()
	if err != nil {
		log.Fatalf("failed to add node: %v", err)
	}

	fmt.Printf("%+v", client.Nodes())
	client.PruneNodes()
	fmt.Printf("%+v", client.Nodes())
	fmt.Println("=========================================")
}
