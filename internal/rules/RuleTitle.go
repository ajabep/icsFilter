package rules

import (
	ics "github.com/arran4/golang-ical"
	"github.com/rs/zerolog/log"
	"strings"
)

type TextCondition uint8

//go:generate stringer -type=TextCondition
const (
	Contains TextCondition = iota
	NotContains
	Exact
	NotExact
)

var MapEnumStringToTextCondition = func() map[string]TextCondition {
	m := make(map[string]TextCondition)
	for i := Contains; i <= NotExact; i++ {
		m[i.String()] = i
	}
	return m
}()
var MapEnumLStringToTextCondition = func() map[string]TextCondition {
	m := make(map[string]TextCondition, len(MapEnumStringToTextCondition))
	for k, v := range MapEnumStringToTextCondition {
		m[strings.ToLower(k)] = v
	}
	return m
}()

type RuleTitle struct {
	Condition TextCondition
	Value     string
}

func (rt *RuleTitle) Complies(event *ics.VEvent) bool {
	title := event.GetProperty(ics.ComponentPropertySummary)
	if title == nil {
		return false
	}
	switch rt.Condition {
	case Contains:
		return strings.Contains(title.Value, rt.Value)
	case NotContains:
		return !strings.Contains(title.Value, rt.Value)
	case Exact:
		return title.Value == rt.Value
	case NotExact:
		return title.Value != rt.Value

	default:
		log.Fatal().Msg("Unknown condition")
		return false
	}
}

func (rt *RuleTitle) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp struct {
		Condition string `yaml:"condition"`
		Value     string `yaml:"value"`
	}
	if err := unmarshal(&tmp); err == nil {
		rt.Condition = MapEnumLStringToTextCondition[strings.ToLower(tmp.Condition)]
		rt.Value = tmp.Value
		return nil
	}

	var tmp2 *string
	if err := unmarshal(&tmp2); err == nil {
		rt.Condition = Exact
		rt.Value = *tmp2
		return nil
	} else {
		return err
	}
}
