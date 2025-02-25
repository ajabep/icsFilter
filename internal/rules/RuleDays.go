package rules

import (
	"errors"
	ics "github.com/arran4/golang-ical"
	"github.com/rs/zerolog/log"
	"slices"
	"strings"
	"time"
)

var MapEnumStringToWeekDay = func() map[string]time.Weekday {
	m := make(map[string]time.Weekday)
	for i := time.Sunday; i <= time.Saturday; i++ {
		m[i.String()] = i
	}
	return m
}()
var MapEnumLStringToWeekday = func() map[string]time.Weekday {
	m := make(map[string]time.Weekday, len(MapEnumStringToWeekDay))
	for k, v := range MapEnumStringToWeekDay {
		m[strings.ToLower(k)] = v
	}
	return m
}()

type RuleDays struct {
	Days []time.Weekday
}

func (rd *RuleDays) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp []string
	if err := unmarshal(&tmp); err != nil {
		return err
	}
	if len(tmp) == 0 {
		return errors.New("days should not be empty")
	}
	for _, day := range tmp {
		value, found := MapEnumLStringToWeekday[strings.ToLower(day)]
		if !found {
			return errors.New("day enum value not found")
		}
		rd.Days = append(rd.Days, value)
	}
	return nil
}

func (rd *RuleDays) Complies(event *ics.VEvent) bool {
	start := event.GetProperty(ics.ComponentPropertyDtStart)
	if start == nil {
		return false
	}
	startTime, err := parseCalDateTime(start.Value, time.UTC)
	if err != nil {
		log.Fatal().Str("eventID", event.Id()).Err(err).Msg("Error parsing start time")
	}

	return slices.Contains(rd.Days, startTime.Weekday())
}
