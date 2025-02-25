package rules

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"github.com/rs/zerolog/log"
	"slices"
	"strings"
)

type Status uint8

//go:generate stringer -type=Status
const (
	Tentative Status = iota
	Confirmed
	Cancelled
	NeedsAction
	Completed
	InProcess
	Draft
	Final
)

var MapEnumStringToStatus = func() map[string]Status {
	m := make(map[string]Status)
	for i := Tentative; i <= Final; i++ {
		m[i.String()] = i
	}
	return m
}()
var MapEnumLStringToStatus = func() map[string]Status {
	m := make(map[string]Status, len(MapEnumStringToStatus))
	for k, v := range MapEnumStringToStatus {
		m[strings.ToLower(k)] = v
	}
	return m
}()

type RuleStatus struct {
	Status []Status
}

func (rs *RuleStatus) Complies(event *ics.VEvent) bool {
	stat := event.GetProperty(ics.ComponentPropertyStatus)
	if stat == nil {
		return false
	}
	status, ok := MapEnumLStringToStatus[strings.ToLower(stat.Value)]
	if !ok {
		log.Fatal().Str("Status", stat.Value).Msg("unknown Status")
	}
	return slices.Contains(rs.Status, status)
}

func (rs *RuleStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
	for _, sStr := range tmp {
		status, ok := MapEnumLStringToStatus[strings.ToLower(sStr)]
		if !ok {
			return fmt.Errorf("unknown status: %s", sStr)
		}
		rs.Status = append(rs.Status, status)
	}
	return nil
}
