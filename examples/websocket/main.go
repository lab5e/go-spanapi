package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lab5e/go-spanapi/v4"
	"github.com/lab5e/go-spanapi/v4/apitools"
)

var (
	token        = flag.String("token", "", "API token for the Span service")
	collectionID = flag.String("collection-id", "", "The collection to query")
	deviceID     = flag.String("device-id", "", "The device to query")
)

func main() {
	flag.Parse()

	if *token == "" || *collectionID == "" {
		log.Fatal("Missing parameter(s)")
	}

	config := spanapi.NewConfiguration()
	// config.Debug = true // uncomment if you want to see what's going on

	// Create a Context that has our auth token
	ctx := apitools.ContextWithAuth(*token)

	// if deviceID is set we want messages from a specific device
	if *deviceID != "" {
		ds, err := apitools.NewDeviceDataStream(ctx, config, *collectionID, *deviceID)
		if err != nil {
			log.Fatalf("Error connecting data stream: %v", err.Error())
		}
		readDataStream(ds)
		return
	}

	// deviceID is "" so we will listen to messages from the collection
	ds, err := apitools.NewCollectionDataStream(ctx, config, *collectionID)
	if err != nil {
		fmt.Println("Error connecting data stream: ", err.Error())
		return
	}
	readDataStream(ds)
}

func readDataStream(ds apitools.DataStream) {
	defer ds.Close()
	fmt.Println("Waiting for data...")
	for {
		msg, err := ds.Recv()
		if err != nil {
			log.Fatalf("Error reading message: %v", err.Error())
		}

		fmt.Printf("Type: %s, Device ID: %s, Transport: %s, Payload: %s Message ID: %s\n",
			*msg.Type,
			*msg.Device.DeviceId,
			*msg.Transport,
			*msg.Payload,
			*msg.MessageId)
	}
}
