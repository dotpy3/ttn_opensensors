package main

import (
	"encoding/json"
	"io/ioutil"
)

// OpenSensorsConf object wraps OpenSensors.io configuration variables
type OpenSensorsConf struct {
	APIKEY         string `json:"apiKey"`
	APIURL         string `json:"apiURL"`
	DeviceID       string `json:"deviceID"`
	DevicePassword string `json:"devicePassword"`
	TopicName      string `json:"topicName"`
	Username       string `json:"username"`
}

// TTNConf object wraps TTN configuration variables
type TTNConf struct {
	AccessKey     string `json:"accessKey"`
	ApplicationID string `json:"applicationID"`
	DeviceID      string `json:"deviceID"`
	Region        string `json:"region"`
}

// Conf object wraps the configuration file into different configuration objects
type Conf struct {
	OpenSensors OpenSensorsConf `json:"OpenSensors"`
	TTN         TTNConf         `json:"TTN"`
}

// confFileReader returns a configuration object from a conf.json file in the execution folder
func confFileReader() (Conf, error) {
	var conf Conf
	file, err := ioutil.ReadFile("conf.json")
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
