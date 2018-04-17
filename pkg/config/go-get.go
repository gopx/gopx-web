package config

import (
	"encoding/json"
	"io/ioutil"

	"gopx.io/gopx-web/pkg/log"
)

// GoGetConfigPath holds configuration file path for go-get query by go tools.
const GoGetConfigPath = "./config/go-get.json"

// GoGetConfig represents configuration for go-get query by go tools.
type GoGetConfig struct {
	VCS                string `json:"vcs"`
	VCSProtocol        string `json:"vcsProtocol"`
	VCSProtocolPrivate string `json:"vcsProtocolPrivate"`
	VCSUserPrivate     string `json:"vcsUserPrivate"`
}

// GoGet holds configs for go-get query i.e. ?go-get=1
var GoGet = new(GoGetConfig)

func init() {
	bytes, err := ioutil.ReadFile(GoGetConfigPath)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	err = json.Unmarshal(bytes, GoGet)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
}
