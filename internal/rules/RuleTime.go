package rules

import (
	ics "github.com/arran4/golang-ical"
	"github.com/rs/zerolog/log"
	"time"
)

type TimeCond struct {
	Hours, Minutes int
}

func parseTime(str string) (TimeCond, error) {
	var t TimeCond
	var tParsed time.Time
	var err error
	for _, patternHours := range []string{
		"15:04:05.0",
		"15:04:05",
		"15:04",
		"150405.0",
		"150405",
		"1504",
	} {
		for _, patternSep := range []string{
			"",
			" ",
		} {
			for _, patternZone := range []string{
				"",
				"MST",
				"-07:00:00",
				"-070000",
				"-07:00",
				"-0700",
				"-07",
				"-07",
				"Z07:00:00",
				"Z070000",
				"Z07:00",
				"Z0700",
				"Z07",
				"Z07",
			} {
				pattern := patternHours + patternSep + patternZone
				tParsed, err = time.Parse(pattern, str)
				if err == nil {
					t.Hours = tParsed.Hour()
					t.Minutes = tParsed.Minute()
					return t, nil
				}
			}
		}
	}
	return t, err
}

func parseCalDateTime(str string, loc *time.Location) (time.Time, error) {
	var t time.Time
	var err error
	for _, patternDate := range []string{
		"20060102",
		"2006-01-02",
	} {
		for _, patternSep := range []string{
			"",
			"T",
			" ",
		} {
			for _, patternHours := range []string{
				"",
				"150405.0",
				"150405",
				"1504",
			} {
				for _, patternSep2 := range []string{
					"",
					" ",
				} {
					for _, patternZone := range []string{
						"",
						" MST",
						"-07:00:00",
						"-070000",
						"-07:00",
						"-0700",
						"-07",
						"-07",
						"Z07:00:00",
						"Z070000",
						"Z07:00",
						"Z0700",
						"Z07",
						"Z07",
					} {
						pattern := patternDate + patternSep + patternHours + patternSep2 + patternZone
						t, err = time.ParseInLocation(pattern, str, loc)
						if err == nil {
							return t, nil
						}
					}
				}
			}
		}
	}
	return t, err
}

type RuleTime struct {
	AllTheDay bool
	StartTime TimeCond
	EndTime   TimeCond
	Location  *time.Location
}

func (rt *RuleTime) Complies(event *ics.VEvent) bool {
	// TODO add the dates
	start := event.GetProperty(ics.ComponentPropertyDtStart)
	if start == nil {
		return false
	}
	startTime, err := parseCalDateTime(start.Value, rt.Location)
	if err != nil {
		log.Fatal().Str("eventID", event.Id()).Err(err).Msg("Error parsing start time")
	}

	end := event.GetProperty(ics.ComponentPropertyDtEnd)
	if end == nil {
		return false
	}
	endTime, err := parseCalDateTime(end.Value, rt.Location)
	if err != nil {
		log.Fatal().Str("eventID", event.Id()).Err(err).Msg("Error parsing start time")
	}

	if rt.AllTheDay {
		log.Fatal().Interface("RuleTime", rt).Str("eventID", event.Id()).Msg("TODO")
	}
	return startTime.Hour() >= rt.StartTime.Hours && startTime.Minute() >= rt.StartTime.Minutes &&
		endTime.Hour() <= rt.StartTime.Hours && endTime.Minute() <= rt.StartTime.Minutes
}

func (rt *RuleTime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp struct {
		//AllTheDay *bool  `yaml:"allDay"`
		StartTime string `yaml:"start"`
		EndTime   string `yaml:"end"`
		Zone      string `yaml:"zone"`
	}
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	var err error
	//if tmp.AllTheDay != nil {
	//	rt.AllTheDay = *tmp.AllTheDay
	//}
	if tmp.StartTime != "" {
		rt.StartTime, err = parseTime(tmp.StartTime)
		if err != nil {
			return err
		}
	}
	if tmp.EndTime != "" {
		rt.EndTime, err = parseTime(tmp.EndTime)
		if err != nil {
			return err
		}
	}
	if tmp.Zone != "" {
		zone, err := time.LoadLocation(tmp.Zone)
		if err != nil {
			return err
		}
		rt.Location = zone
	} else {
		rt.Location = time.UTC
	}
	return nil
}
