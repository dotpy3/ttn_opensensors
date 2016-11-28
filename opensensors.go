package main

import "net/http"
import "encoding/json"
import "bytes"
import "github.com/mitchellh/mapstructure"
import "fmt"

// OpenSensorsClient provides an API to interact with the OpenSensorsClient HTTP API
type OpenSensorsClient struct {
	conf OpenSensorsConf
}

func (cli OpenSensorsClient) getAPICallURL() string {
	return cli.conf.APIURL + "topics//users/" + cli.conf.Username + "/" +
		cli.conf.TopicName + "?client-id=" + cli.conf.DeviceID + "&password=" +
		cli.conf.DevicePassword
}

func (cli OpenSensorsClient) postToAPI(data []byte) (*http.Response, error) {
	httpClient := &http.Client{}
	fmt.Println("Payload to be expected:" + string(data[:]))
	req, err := http.NewRequest("POST", cli.getAPICallURL(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "api-key "+cli.conf.APIKEY)
	fmt.Println("Body of the request: " + ReaderToString(req.Body))
	return httpClient.Do(req)
}

func (cli OpenSensorsClient) postMQTTPayload(data map[string]interface{}) (*http.Response, error) {
	fmt.Printf("dump1:\n%+v\n\n", data)
	var rawData interface{}
	mapstructure.Decode(data, rawData)
	fmt.Printf("dump2:\n%+v\n\n", rawData)
	payload, err := json.Marshal(rawData)
	if err != nil {
		return nil, err
	}
	return cli.postToAPI(payload)
}
