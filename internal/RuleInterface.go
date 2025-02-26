package internal

import (
	ics "github.com/arran4/golang-ical"
)

type RuleInterface interface {
	Complies(event *ics.VEvent) bool
}
