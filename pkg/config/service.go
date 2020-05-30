package config

import (
	"encoding/json"
	"io/ioutil"

	"gopx.io/gopx-web/pkg/log"
)

// ServiceConfigPath holds service related configuration file path.
const ServiceConfigPath = "./config/service.json"

// ServiceConfig represents service related configurations.
type ServiceConfig struct {
	Host      string `json:"host"`
	UseHTTP   bool   `json:"useHTTP"`
	HTTPPort  int    `json:"HTTPPort"`
	UseHTTPS  bool   `json:"useHTTPS"`
	HTTPSPort int    `json:"HTTPSPort"`
	CertFile  string `json:"certFile"`
	KeyFile   string `json:"keyFile"`
}

// Service holds loaded service related configurations.
var Service = new(ServiceConfig)

func init() {
	bytes, err := ioutil.ReadFile(ServiceConfigPath)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	err = json.Unmarshal(bytes, Service)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
}
