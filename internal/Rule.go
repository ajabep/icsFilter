package internal

import (
	ics "github.com/arran4/golang-ical"
	"icsFilter/internal/rules"
)

type RuleInterface interface {
	Complies(event *ics.VEvent) bool
}

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

	/*
		if tmp.StartTime == "" {
			r.StartTime = nil
		} else {
			t, err := parseTime(tmp.StartTime)
			if err != nil {
				return err
			}
			r.StartTime = &t
		}
		if tmp.EndTime == "" {
			r.EndTime = nil
		} else {
			t, err := parseTime(tmp.EndTime)
			if err != nil {
				return err
			}
			r.EndTime = &t
		}
		r.Days = make([]Day, len(tmp.Days))
		for i, d := range tmp.Days {
			r.Days[i] = MapEnumStringToDay[d]
		}
	*/
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
