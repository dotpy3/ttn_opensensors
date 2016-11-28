package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/TheThingsNetwork/ttn/core/types"
	TTNMQTT "github.com/TheThingsNetwork/ttn/mqtt"
)

func main() {
	conf, err := confFileReader()
	if err != nil {
		fmt.Println("Couldn't get configuration information: " + err.Error())
		os.Exit(0)
	}

	fmt.Println("Connecting...")
	client := TTNMQTT.NewClient(nil, "ttn-opensensors", conf.TTN.ApplicationID, conf.TTN.AccessKey, "tcp://"+conf.TTN.Region+".thethings.network:1883")
	if err := client.Connect(); err != nil {
		fmt.Println("Couldn't connect: " + err.Error())
		os.Exit(0)
	}
	fmt.Println("Connection complete.")

	var OSCli OpenSensorsClient
	OSCli.conf = conf.OpenSensors
	fmt.Println("Client initialized.")

	token := client.SubscribeDeviceUplink(conf.TTN.ApplicationID, conf.TTN.DeviceID, func(client TTNMQTT.Client, appID string, devID string, req types.UplinkMessage) {
		fmt.Println("\n\n=========================\nMESSAGE RECEPTION\n")
		if rep, err := OSCli.postMQTTPayload(req.PayloadFields); err != nil || ErrorResponse(rep) {
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
	})

	fmt.Println("Subscription complete - waiting for messages...")
	token.Wait()
	if err := token.Error(); err != nil {
		fmt.Println("Could not subscribe")
		os.Exit(0)
	}

	select {}
}
