package rules

import (
	ics "github.com/arran4/golang-ical"
	"github.com/rs/zerolog/log"
	"strings"
)

type Transparency uint8

//go:generate stringer -type=Transparency
const (
	Transp Transparency = iota
	Opaque
	Transparent = Transp
)

var MapEnumStringToTransparency = func() map[string]Transparency {
	m := make(map[string]Transparency)
	for i := Transp; i <= Opaque; i++ {
		m[i.String()] = i
	}
	return m
}()
var MapEnumLStringToTransparency = func() map[string]Transparency {
	m := make(map[string]Transparency, len(MapEnumStringToTransparency))
	for k, v := range MapEnumStringToTransparency {
		m[strings.ToLower(k)] = v
	}
	return m
}()

type RuleTransparency struct {
	Transparency Transparency
}

func (rt *RuleTransparency) Complies(event *ics.VEvent) bool {
	transp := event.GetProperty(ics.ComponentPropertyTransp)
	if transp == nil {
		return false
	}
	transparency, ok := MapEnumLStringToTransparency[strings.ToLower(transp.Value)]
	if !ok {
		log.Fatal().Str("Transparency", transp.Value).Msg("unknown Transparency")
	}
	return rt.Transparency == transparency
}

func (rt *RuleTransparency) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp *string
	if err := unmarshal(&tmp); err != nil {
		return err
	}
	rt.Transparency = MapEnumLStringToTransparency[*tmp]
	return nil
}
