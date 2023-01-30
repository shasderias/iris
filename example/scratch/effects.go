package main

import (
	"github.com/shasderias/iris/beat"
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/evt"
	"github.com/shasderias/iris/fx"
)

func spotlightGhost(ctx context.Context,
	seq beat.Sequence, colRng, rotRng beat.Range,
	sRot, eRot, m, c float64,
	options ...any) {
	ctx.WSeq(seq, func(ctx context.Context) {
		//var (
		//	sRot = 10.0
		//	eRot = -10.0
		//	m    = -13.0
		//	c    = -12.0
		//)

		resetG, resetB := evt.RotationGroupWithBox(ctx, groupOptions(options...)...)
		resetG.Beat -= 0.1
		ctx.WSeq(beat.Seq(0), func(ctx context.Context) {
			resetB.AddEvent(ctx, evt.ORotation(sRot), evt.EasingNone)
		})

		ctx.WOpt(options...).Do(func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx)
			ctx.WRng(colRng, func(ctx context.Context) {
				cb.AddEvent(ctx, fx.ExtendTransit)
			})

			_, rb := evt.RotationGroupWithBox(ctx, evt.ODistWave(m*ctx.OrdinalF()+c))
			ctx.WRng(rotRng, func(ctx context.Context) {
				rb.AddEvent(ctx, evt.ORotate(sRot, eRot, evt.CCW))
			})
		})
	})
}

func groupOptions(options ...any) []any {
	ret := []any{}
	for _, opt := range options {
		switch opt.(type) {
		case evt.ColorEventGroupOption:
			ret = append(ret, opt)
		case evt.RotationEventGroupOption:
			ret = append(ret, opt)
		}
	}
	return ret
}
