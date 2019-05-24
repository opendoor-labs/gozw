package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gozwave/gozw"
	"github.com/gozwave/gozw/cc"
	"github.com/gozwave/gozw/cc/allclasses"
)

var networkKey = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func init() {
	allclasses.Noop()
}

func main() {
	devicePath := "/dev/ttyACM0"
	if p := os.Getenv("GOZW_DEVICE_PATH"); p != "" {
		devicePath = p
	}
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

	fmt.Println("listening for events")
	client.SetEventCallback(func(c *gozw.Client, nodeID byte, e cc.Command) {
		node, err := c.Node(nodeID)
		if err != nil {
			log.Printf("failed to get node: %v", node)
		}
		spew.Dump(e)

		fmt.Printf("event received: %v\n", e)
	})

	time.Sleep(time.Second * 600)
	fmt.Println("done listening, exiting")
}
