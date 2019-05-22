package main

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gozwave/gozw"
	"github.com/gozwave/gozw/cc"
)

var networkKey = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func main() {
	client, err := gozw.NewDefaultClient("/tmp/data.db", "/dev/tty.usbmodem1431", 115200, networkKey)
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
		fmt.Printf("event received: %s\n", e)
	})

	time.Sleep(time.Second * 10)
	fmt.Println("done listening, exiting")
}
