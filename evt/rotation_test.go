package evt_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/shasderias/iris/beat"
	"github.com/shasderias/iris/beatsaber"
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/evt"
)

var specimenRotationGroup = []beatsaber.LightRotationGroupV320{{
	Beat:    7.1,
	GroupID: 8,
	Boxes: []beatsaber.LightRotationBoxV320{
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
			BeatDistributionType:      int(evt.DistributionTypeStep),
			BeatDistributionParam:     15.1,
			RotationDistributionType:  int(evt.DistributionTypeStep),
			RotationDistributionParam: 16.1,
			Axis:                      int(evt.AxisY),
			FlipRotation:              1,
			AffectFirst:               1,
			Easing:                    int(evt.EasingInOutQuad),
			Events: []beatsaber.LightRotationEventV320{
				{
					Beat:                          17.1,
					UsePreviousEventRotationValue: int(evt.RotationTransitionTypeExtend),
					EaseType:                      int(evt.EasingInOutQuad),
					LoopsCount:                    18,
					Rotation:                      19.1,
					RotationDirection:             int(evt.RotationDirectionCounterClockwise),
				},
			},
		},
	}},
}

func TestRotationBox(t *testing.T) {
	ctx := context.NewBase()

	var rotationGroup *evt.RotationEventGroup

	ctx.WSeq(beat.Seq(7.1), func(ctx context.Context) {
		rg := evt.RotationGroup(ctx, evt.OEventGroup(8))
		rotationGroup = rg

		lane1 := rg.AddBox(ctx,
			evt.OStepAndOffsetFilter(9, 10,
				evt.OReverse(true),
				evt.OChunks(12),
				evt.OOrder(evt.FilterOrderRandom, 13),
				evt.OLimit(14.1, evt.LimitAffectsDurationAndDistribution),
			),
			evt.OBeatDistStep(15.1),
			evt.ODistStep(16.1),
			evt.ODistAffectFirst(true),
			evt.AxisY,
			evt.ODistFlip(true),
			evt.EasingInOutQuad,
		)

		ctx.WRng(beat.RngStep(17.1, 18.1, 1), func(ctx context.Context) {
			lane1.AddEvent(ctx, evt.Extend, evt.OLoop(18), evt.ORotation(19.1), evt.EasingInOutQuad, evt.RotationDirectionCounterClockwise)
		})
	})

	if diff := cmp.Diff(rotationGroup.LightRotationEventGroupV320(), specimenRotationGroup); diff != "" {
		t.Fatal(diff)
	}
}
