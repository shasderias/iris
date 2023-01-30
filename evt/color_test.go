package evt_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/shasderias/iris/beat"
	"github.com/shasderias/iris/beatsaber"
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/evt"
)

var (
	specimenColorEventGroup = []beatsaber.LightColorEventGroupV320{{
		Beat:    7.1,
		GroupID: 8,
		Boxes: []beatsaber.LightColorBoxV320{
			{
				IndexFilter: beatsaber.IndexFilterV320{
					FilterType:   int(evt.IndexFilterTypeStepAndOffset),
					ParamP:       9,
					ParamT:       10,
					Reverse:      1,
					Chunks:       12,
					Order:        int(evt.FilterOrderRandom),
					Seed:         13,
					Limit:        14.1,
					LimitAffects: int(evt.LimitAffectsDurationAndDistribution),
				},
				BeatDistributionType:              int(evt.DistributionTypeStep),
				BeatDistributionParam:             15.1,
				BrightnessDistributionType:        int(evt.DistributionTypeStep),
				BrightnessDistributionParam:       16.1,
				BrightnessDistributionAffectFirst: 1,
				BrightnessDistributionEasing:      int(evt.EasingInOutQuad),
				Events: []beatsaber.LightColorEventV320{
					{
						Beat:                 17.1,
						TransitionType:       int(evt.ColorEventTransitionExtend),
						EnvironmentColorType: int(evt.ColorEventWhite),
						Brightness:           18.1,
						StrobeFrequency:      19.1,
					},
				},
			},
		},
	}}
)

func TestColorBox(t *testing.T) {
	ctx := context.NewBase()

	var colorGroup *evt.ColorEventGroup

	ctx.WSeq(beat.Seq(7.1), func(ctx context.Context) {
		colorGroup = evt.ColorGroup(ctx, evt.OColorEventGroup(8))
		cg := colorGroup

		lane1 := cg.AddBox(ctx,
			evt.OStepAndOffsetFilter(9, 10,
				evt.OReverse(true),
				evt.OChunks(12),
				evt.OOrder(evt.FilterOrderRandom, 13),
				evt.OLimit(14.1, evt.LimitAffectsDurationAndDistribution),
			),
			evt.OBeatDistStep(15.1),
			evt.ODistStep(16.1),
			evt.ODistAffectFirst(true),
			evt.EasingInOutQuad,
		)

		ctx.WRng(beat.RngStep(17.1, 18.1, 1), func(ctx context.Context) {
			lane1.AddEvent(ctx, evt.Extend, evt.White, evt.OBrightness(18.1), evt.OStrobe(19.1))
		})
	})

	if diff := cmp.Diff(colorGroup.LightColorEventGroupV320(), specimenColorEventGroup); diff != "" {
		t.Fatal(diff)
	}
}
