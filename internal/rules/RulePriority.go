package rules

import (
	ics "github.com/arran4/golang-ical"
	"github.com/rs/zerolog/log"
	"slices"
	"strconv"
)

type RulePriority struct {
	Priority []uint8
}

func (rp *RulePriority) Complies(event *ics.VEvent) bool {
	prio := event.GetProperty(ics.ComponentPropertyPriority)
	if prio == nil {
		return false
	}

	priority, err := strconv.ParseInt(prio.Value, 10, 8)
	if err != nil {
		log.Fatal().Err(err).Str("Priority", prio.Value).Msg("Error parsing priority")
	}

	p := uint8(priority)
	return slices.Contains(rp.Priority, p)
}

func (rp *RulePriority) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp struct {
		Max uint8 `yaml:"max"`
		Min uint8 `yaml:"min"`
	}
	if err := unmarshal(&tmp); err == nil {
		rp.Priority = make([]uint8, tmp.Max-tmp.Min)
		for i := 0; i < int(tmp.Max-tmp.Min); i++ {
			rp.Priority[i] = tmp.Min
		}
	}

	var tmp2 []uint8
	if err := unmarshal(&tmp); err != nil {
		var tmp3 *uint8
		if err := unmarshal(&tmp2); err == nil {
			tmp2 = append(tmp2, *tmp3)
			return nil
		} else {
			return err
		}
	}
	rp.Priority = tmp2
	return nil
}
