package main

import "net/http"
import "encoding/json"
import "bytes"

// OpenSensorsClient provides an API to interact with the OpenSensorsClient HTTP API
type OpenSensorsClient struct {
	conf OpenSensorsConf
}

// getAPICallURL returns the URI to post the data to
func (cli OpenSensorsClient) getAPICallURL() string {
	return cli.conf.APIURL + "topics//users/" + cli.conf.Username + "/" +
		cli.conf.TopicName + "?client-id=" + cli.conf.DeviceID + "&password=" +
		cli.conf.DevicePassword
}

// postToAPI takes a data payload, whatever its format, and posts it to the OpenSensorsAPI
func (cli OpenSensorsClient) postToAPI(data []byte) (*http.Response, error) {
	httpClient := &http.Client{}
	bytesData := bytes.NewReader(data)
	req, err := http.NewRequest("POST", cli.getAPICallURL(), bytesData)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "api-key "+cli.conf.APIKEY)
	return httpClient.Do(req)
}

// OSPayload allows us to encapsulate the data in a OS-readable format
type OSPayload struct {
	data map[string]interface{}
}

// encapsulateIntoData takes a JSON unmarshalled object, and wraps it as a string in another JSON unmarshalled object
func encapsulateIntoData(data map[string]interface{}) (map[string]interface{}, error) {
	stringifiedData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	payload := make(map[string]interface{})
	payload["data"] = string(stringifiedData[:])
	return payload, nil
}

// postMQTTPayload takes a JSON unmarshalled object containing the data to post, and posts it to OpenSensors
func (cli OpenSensorsClient) postMQTTPayload(data map[string]interface{}) (*http.Response, error) {
	encapsulatedData, err := encapsulateIntoData(data)
	if err != nil {
		return nil, err
	}
	payload, err := json.Marshal(encapsulatedData)
	if err != nil {
		return nil, err
	}
	return cli.postToAPI(payload)
}
