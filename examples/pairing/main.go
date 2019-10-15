package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gozwave/gozw"
)

var networkKey = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func unpair(ctx context.Context, client *gozw.Client) {

	fmt.Println("removing node, put device in unpairing mode")
	if _, err := client.RemoveNode(); err != nil {
		log.Fatalf("failed to remove node: %v", err)
	}

}

func pair(ctx context.Context, client *gozw.Client) *gozw.Node {

	fmt.Println("adding node, put device in pairing mode")

	node, err := client.AddNode()
	if err != nil {
		log.Fatalf("failed to add node: %v", err)
	}

	return node
}

func pairAsync(ctx context.Context, client *gozw.Client) *gozw.Node {

	fmt.Println("adding node, put device in pairing mode")

	progressChan := make(chan gozw.PairingProgressUpdate)

	go func() {
		for {
			select {
			case newProgress := <-progressChan:
				fmt.Printf("pairing progress update: %d/%d", newProgress.InterviewedCommandClassCount, newProgress.ReportedCommandClassCount)
			case <-ctx.Done():
				fmt.Println("pairing failed", ctx.Err())
				return
			}
		}
	}()

	//node, err := client.AddNodeWithProgress(ctx, progressChan)
	node, err := client.AddNode()
	if err != nil {
		log.Fatalf("failed to add node: %v", err)
	}

	return node
}

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
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""
	client.SetLogger(zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.Lock(os.Stdout), atom)))

	spew.Dump(client.Controller)

	for _, node := range client.Nodes() {
		fmt.Println(node.String())
	}

	for i := 0; i < 10; i++ {
		log.Println("starting unpair")
		unpair(ctx, client)

		log.Println("Unpair complete - starting pairing in")
		for i := 3; i > 0; i-- {
			log.Printf("%d", i)
			time.Sleep(time.Second)
		}

		node := pair(ctx, client)
		if node == nil {
			log.Fatal("didn't get a new node (try 1/2)")
		}

		fmt.Println(node.String())
		log.Println("Pair complete")
		for i := 3; i > 0; i-- {
			log.Printf("%d", i)
			time.Sleep(time.Second)
		}
	}
}
