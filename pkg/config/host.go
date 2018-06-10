package config

import (
	"encoding/json"
	"io/ioutil"

	"gopx.io/gopx-web/pkg/log"
)

// HostConfigPath holds service host configuration file path of whole GoPX system.
const HostConfigPath = "./config/host.json"

// HostConfig represents host names of all GoPX services.
type HostConfig struct {
	Web string `json:"web"`
	API string `json:"api"`
	VCS string `json:"vcs"`
}

// Host holds host names of all GoPX services.
var Host = new(HostConfig)

func init() {
	bytes, err := ioutil.ReadFile(HostConfigPath)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	err = json.Unmarshal(bytes, Host)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
}
