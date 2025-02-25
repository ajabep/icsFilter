package internal

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
)

type RulesFile struct {
	Endpoints []Endpoint `yaml:"endpoint"`
}

func (rf *RulesFile) Load(path string) error {
	ruleFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(ruleFile, rf)
	return err
}

func (rf *RulesFile) InitHttp() {
	for _, endpoint := range rf.Endpoints {
		logger := log.With().Str("endpointId", endpoint.ID).Str("endpointUrl", endpoint.URL).Logger()
		logger.Trace().Msg("Handling Endpoint")

		urlEndpoint := fmt.Sprintf("/%s", endpoint.ID)
		http.HandleFunc(urlEndpoint, endpoint.HandleIcs)
		logger.Info().Str("URL", urlEndpoint).Msg("Endpoint initialized")
	}
}
