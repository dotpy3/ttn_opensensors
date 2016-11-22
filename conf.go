package main

import (
	"encoding/json"
	"os"
)

// OpenSensorsConf object wraps OpenSensors.io configuration variables
type OpenSensorsConf struct {
	apiKey         string
	apiUrl         string
	deviceID       string
	devicePassword string
	topicName      string
	username       string
}

// TTNConf object wraps TTN configuration variables
type TTNConf struct {
	accessKey     string
	applicationID string
	region        string
}

// Conf object wraps the configuration file into different configuration objects
type Conf struct {
	OpenSensors OpenSensorsConf
	TTN         TTNConf
}

// confFileReader returns a configuration object from a conf.json file in the execution folder
func confFileReader() (error, Conf) {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	conf := Conf{}
	err := decoder.Decode(&conf)
	if err != nil {
		return err, conf
	}
	return nil, conf
}
