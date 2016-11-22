package main

import (
	"bytes"

	"github.com/TheThingsNetwork/ttn/core/types"

	"fmt"
	"os"

	TTNMQTT "github.com/TheThingsNetwork/ttn/mqtt"
)

func main() {
	err, conf := confFileReader()
	if err != nil {
		fmt.Println("Couldn't get configuration information: " + err.Error())
		os.Exit(0)
	}
	client := TTNMQTT.NewClient(nil, "ttnctl", conf.TTN.applicationID, conf.TTN.accessKey, "tcp://"+conf.TTN.region+".thethings.network:1883")
	if err := client.Connect(); err != nil {
		fmt.Println("Couldn't connect: " + err.Error())
		os.Exit(0)
	}
	OSCli := OpenSensorsClient{conf.OpenSensors}
	token := client.SubscribeDeviceUplink(conf.TTN.applicationID, conf.TTN.deviceID, func(client TTNMQTT.Client, appID string, devID string, req types.UplinkMessage) {
		length := bytes.IndexByte(req.PayloadRaw, 0)
		payload := string(req.PayloadRaw[:length])
		fmt.Println("Uplink message received: " + payload)
		if err := OSCli.post(payload); err != nil {
			fmt.Println("Error transmitting the message to OpenSensors: " + err.Error())
		} else {
			fmt.Println("Message successfuly transmitted to OpenSensors.")
		}
	})
	token.Wait()
	if err := token.Error(); err != nil {
		fmt.Println("Could not subscribe")
		os.Exit(0)
	}
}
