package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	// Conf is the global configuration
	Conf Configuration
)

// Configuration stores application configuration
type Configuration struct {
	APIKey         string `json:"apikey"`
	NetworkID      string `json:"networkid"`
	SSIDNum        string `json:"ssidnum"`
	BTCEnabled     bool   `json:"btc"`
	ETHEnabled     bool   `json:"eth"`
	Currency       string `json:"currency"`
	UpdateInterval int    `json:"interval"`
}

// ReadConfig reads config.json into memory for global configuration
func ReadConfig() Configuration {
	// Check to make sure config.json exists
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		LogErr("Configuration Error! config.json improperly loaded or does not exist :: " + err.Error())
		os.Exit(1)
	}

	// Unmarshal the raw json config into a Configuration object
	var c Configuration
	if errUnm := json.Unmarshal(raw, &c); errUnm != nil {
		LogErr("Configuration Error! config.json improperly unmarshaled :: " + errUnm.Error())
		os.Exit(1)
	}

	return c
}
