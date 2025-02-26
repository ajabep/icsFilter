package internal

import (
	"github.com/ajabep/icsFilter/internal/rules"
	ics "github.com/arran4/golang-ical"
)

type Rule struct {
	Conditions []RuleInterface
}

func (r *Rule) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp struct {
		Title *rules.RuleTitle `yaml:"title"`
		Time  *rules.RuleTime  `yaml:"time"`
		Days  *rules.RuleDays  `yaml:"days"`
	}
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	if tmp.Title != nil {
		r.Conditions = append(r.Conditions, tmp.Title)
	}
	if tmp.Time != nil {
		r.Conditions = append(r.Conditions, tmp.Time)
	}
	if tmp.Days != nil {
		r.Conditions = append(r.Conditions, tmp.Days)
	}
	return nil
}

func (r *Rule) Complies(event *ics.VEvent) bool {
	for _, condition := range r.Conditions {
		if !condition.Complies(event) {
			return false
		}
	}
	return true
}
