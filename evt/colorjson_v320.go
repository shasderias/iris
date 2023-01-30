package evt

import (
	"github.com/shasderias/iris/beatsaber"
	"github.com/shasderias/iris/iter"
)

func (g *ColorEventGroup) LightColorEventGroupV320() []beatsaber.LightColorEventGroupV320 {
	events := []beatsaber.LightColorEventGroupV320{}
	for _, gid := range g.Group {
		events = append(events, beatsaber.LightColorEventGroupV320{
			Beat:    g.Beat,
			GroupID: gid,
			Boxes: iter.MapSlice(g.Boxes, func(b *ColorEventBox) beatsaber.LightColorBoxV320 {
				return b.LightColorEventBoxV320()
			}),
		})
	}
	return events
}
func (b *ColorEventBox) LightColorEventBoxV320() beatsaber.LightColorBoxV320 {
	return beatsaber.LightColorBoxV320{
		IndexFilter:                       b.IndexFilter.IndexFilterV320(),
		BeatDistributionParam:             b.BeatDistribution.Param,
		BeatDistributionType:              int(b.BeatDistribution.Type),
		BrightnessDistributionParam:       b.BrightnessDistribution.Param,
		BrightnessDistributionType:        int(b.BrightnessDistribution.Type),
		BrightnessDistributionAffectFirst: boolToIntBool(b.BrightnessDistribution.AffectFirst),
		BrightnessDistributionEasing:      int(b.BrightnessDistribution.Easing),

		Events: iter.MapSlice(b.Events, func(e *ColorEvent) beatsaber.LightColorEventV320 {
			return e.LightColorEventV320()
		}),
	}
}

func (e *ColorEvent) LightColorEventV320() beatsaber.LightColorEventV320 {
	return beatsaber.LightColorEventV320{
		Beat:                 e.Beat,
		TransitionType:       int(e.TransitionType),
		EnvironmentColorType: int(e.Color),
		Brightness:           e.Brightness,
		StrobeFrequency:      e.StrobeFrequency,
	}
}
