package config

import (
	"encoding/json"
	"io/ioutil"

	"gopx.io/gopx-web/pkg/log"
)

// WebConfigPath holds web specific configuration file path.
const WebConfigPath = "./config/web.json"

// WebConfig represents web specific configurations.
type WebConfig struct {
	StaticAssetPath       string `json:"staticAssetPath"`
	PageTemplatePath      string `json:"pageTemplatePath"`
	PageTemplateExtension string `json:"pageTemplateExtension"`
}

// Web holds loaded web specific configurations.
var Web = new(WebConfig)

func init() {
	bytes, err := ioutil.ReadFile(WebConfigPath)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	err = json.Unmarshal(bytes, Web)
	if err != nil {
		log.Fatal("Error: %s", err)
	}
}
