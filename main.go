package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/TheThingsNetwork/ttn/core/types"
	TTNMQTT "github.com/TheThingsNetwork/ttn/mqtt"
)

func main() {
	// Reading the configuration file
	conf, err := confFileReader()
	if err != nil {
		fmt.Println("Couldn't get configuration information: " + err.Error() + " // Please make sure that there is a configuration file conf.json in the current directory that respects the format of the example.")
		os.Exit(0)
	}

	// Connection to the broker
	client := TTNMQTT.NewClient(nil, "ttn-opensensors", conf.TTN.ApplicationID, conf.TTN.AccessKey, "tcp://"+conf.TTN.Region+".thethings.network:1883")
	if err := client.Connect(); err != nil {
		fmt.Println("Couldn't connect: " + err.Error())
		os.Exit(0)
	}
	fmt.Println("Connection complete.")

	// Initializing OpenSensors Client
	var OSCli OpenSensorsClient
	OSCli.conf = conf.OpenSensors

	// Saving message handling function
	messageHandler := func(client TTNMQTT.Client, appID string, devID string, req types.UplinkMessage) {
		fmt.Print("\n\n=========================\nMESSAGE RECEPTION\n\n")

		if rep, err := OSCli.postMQTTPayload(req.PayloadFields); err != nil || ErrorResponse(rep) {
			// Error handling
			if ErrorResponse(rep) {
				fmt.Println("Error " + strconv.Itoa(rep.StatusCode) + ": " + ReaderToString(rep.Body))
			} else if err != nil {
				fmt.Println("Error transmitting the message to OpenSensors: " + err.Error())
			} else {
				fmt.Println("No response transmitted")
			}
		} else {
			fmt.Println("Message successfuly retransmitted to OpenSensors.")
		}
	}

	// Subscribing
	token := client.SubscribeDeviceUplink(conf.TTN.ApplicationID, conf.TTN.DeviceID, messageHandler)
	fmt.Println("Subscription complete - waiting for messages...")
	token.Wait()
	if err := token.Error(); err != nil {
		fmt.Println("Could not subscribe: " + err.Error())
		os.Exit(0)
	}

	// Launching program main process as idle mode
	select {}
}
