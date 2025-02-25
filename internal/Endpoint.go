package internal

import (
	ics "github.com/arran4/golang-ical"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Endpoint struct {
	ID     string `yaml:"id"`
	URL    string `yaml:"url"`
	Delete []Rule `yaml:"delete"`
	Edit   []Rule `yaml:"edit"`
}

func (e *Endpoint) HandleIcs(w http.ResponseWriter, _ *http.Request) {
	// TODO Implement cache
	logger := log.With().Str("url", e.URL).Logger()
	cal, err := ics.ParseCalendarFromUrl(e.URL)
	if err != nil {
		logger.Fatal().Err(err).Msg("Loading remote calendar")
	}

	var idToDelete []string
	for _, event := range cal.Events() {
		for _, rule := range e.Delete {
			if rule.Complies(event) {
				idToDelete = append(idToDelete, event.Id())
			}
		}
		for _, rule := range e.Edit {
			logger.Fatal().Interface("rule", rule).Msg("TODO")
		}
	}

	for _, id := range idToDelete {
		cal.RemoveEvent(id)
	}

	err = cal.SerializeTo(w)
	if err != nil {
		logger.Fatal().Err(err).Msg("Serializing calendar")
	}
	logger.Info().Msg("Request Handled successfully")
}
