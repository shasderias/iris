package evt

import (
	"github.com/shasderias/iris/beatsaber"
	"github.com/shasderias/iris/iter"
)

func (g *RotationEventGroup) LightRotationEventGroupV320() []beatsaber.LightRotationGroupV320 {
	events := []beatsaber.LightRotationGroupV320{}
	for _, gid := range g.Group {
		events = append(events, beatsaber.LightRotationGroupV320{
			Beat:    g.Beat,
			GroupID: gid,
			Boxes: iter.MapSlice(g.Boxes, func(b *RotationEventBox) beatsaber.LightRotationBoxV320 {
				return b.LightRotationEventBoxV320()
			}),
		})
	}
	return events
}

func (b *RotationEventBox) LightRotationEventBoxV320() beatsaber.LightRotationBoxV320 {
	return beatsaber.LightRotationBoxV320{
		IndexFilter:               b.IndexFilter.IndexFilterV320(),
		BeatDistributionParam:     b.BeatDistribution.Param,
		BeatDistributionType:      int(b.BeatDistribution.Type),
		RotationDistributionParam: b.RotationDistribution.Param,
		RotationDistributionType:  int(b.RotationDistribution.Type),
		Axis:                      int(b.Axis),
		FlipRotation:              boolToIntBool(b.Flip),
		AffectFirst:               boolToIntBool(b.AffectFirst),
		Easing:                    int(b.RotationDistribution.Easing),

		Events: iter.MapSlice(b.Events, func(e *RotationEvent) beatsaber.LightRotationEventV320 {
			return e.LightRotationEventV320()
		}),
	}
}

func (e *RotationEvent) LightRotationEventV320() beatsaber.LightRotationEventV320 {
	return beatsaber.LightRotationEventV320{
		Beat:                          e.Beat,
		UsePreviousEventRotationValue: int(e.TransitionType),
		EaseType:                      int(e.Easing),
		LoopsCount:                    e.LoopsCount,
		Rotation:                      e.Rotation,
		RotationDirection:             int(e.Direction),
	}
}
