package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gozwave/gozw"
)

var networkKey = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*2)

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
		log.Println(node.String())
	}

	log.Println("removing node, put device in unpairing mode")
	if _, err := client.RemoveNode(); err != nil {
		log.Fatalf("failed to remove node: %v", err)
	}

	log.Println("adding node, put device in pairing mode")

	progressChan := make(chan gozw.PairingProgressUpdate)

	go func() {
		for {
			select {
			case newProgress := <-progressChan:
				log.Printf("pairing progress update: %d/%d", newProgress.InterviewedCommandClassCount, newProgress.ReportedCommandClassCount)
			case <-ctx.Done():
				log.Fatalln("pairing failed", ctx.Err())
				return
			}
		}
	}()

	node, err := client.AddNodeWithProgress(ctx, progressChan)
	if err != nil {
		log.Fatalf("failed to add node: %v", err)
	}

	log.Println(node.String())
}
