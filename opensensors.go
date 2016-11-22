package main

import "net/http"
import "encoding/json"
import "bytes"

// OpenSensorsClient provides an API to interact with the OpenSensorsClient HTTP API
type OpenSensorsClient struct {
	conf OpenSensorsConf
}

func (cli OpenSensorsClient) getAPICallURL() string {
	return cli.conf.apiUrl + "topics//users/" + cli.conf.username + "/" +
		cli.conf.topicName + "?client-id=" + cli.conf.deviceID + "&password=" +
		cli.conf.devicePassword
}

func (cli OpenSensorsClient) postToAPI(data string) (*http.Response, error) {
	var f interface{}
	var payload struct {
		data interface{}
	}
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		return nil, err
	}
	payload.data = f
	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return http.Post(cli.getAPICallURL(), "application/json", bytes.NewBuffer(encodedPayload))
}

func (cli OpenSensorsClient) post(data string) bool {
	_, err := cli.postToAPI(data)
	return err == nil
}
