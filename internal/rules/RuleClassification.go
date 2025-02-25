package rules

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"github.com/rs/zerolog/log"
	"slices"
	"strings"
)

type Classification uint8

//go:generate stringer -type=Classification
const (
	Private Classification = iota
	Public
	Confidential
)

var MapEnumStringToClassification = func() map[string]Classification {
	m := make(map[string]Classification)
	for i := Private; i <= Confidential; i++ {
		m[i.String()] = i
	}
	return m
}()
var MapEnumLStringToClassification = func() map[string]Classification {
	m := make(map[string]Classification, len(MapEnumStringToClassification))
	for k, v := range MapEnumStringToClassification {
		m[strings.ToLower(k)] = v
	}
	return m
}()

type RuleClassification struct {
	Classifications []Classification
}

func (rc *RuleClassification) Complies(event *ics.VEvent) bool {
	class := event.GetProperty(ics.ComponentPropertyClass)
	if class == nil {
		return false
	}
	classification, ok := MapEnumLStringToClassification[strings.ToLower(class.Value)]
	if !ok {
		log.Fatal().Str("classification", class.Value).Msg("unknown classification")
	}
	return slices.Contains(rc.Classifications, classification)
}

func (rc *RuleClassification) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp []string
	if err := unmarshal(&tmp); err != nil {
		var tmp2 *string
		if err := unmarshal(&tmp2); err == nil {
			tmp = append(tmp, *tmp2)
			return nil
		} else {
			return err
		}
	}
	for _, cStr := range tmp {
		classification, ok := MapEnumLStringToClassification[strings.ToLower(cStr)]
		if !ok {
			return fmt.Errorf("unknown classification: %s", cStr)
		}
		rc.Classifications = append(rc.Classifications, classification)
	}
	return nil
}
