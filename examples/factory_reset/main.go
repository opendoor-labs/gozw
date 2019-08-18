package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gozwave/gozw"
)

var networkKey = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

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

	client.FactoryReset()

	for _, node := range client.Nodes() {
		fmt.Println(node.String())
	}

	time.Sleep(2 * time.Second)

}
