package main

import (
	"github.com/shasderias/iris/beat"
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/ease"
	"github.com/shasderias/iris/env/thesecond"
	"github.com/shasderias/iris/evt"
	"github.com/shasderias/iris/fx"
	"github.com/shasderias/iris/opt"
)

func spotlightGhost(ctx context.Context,
	seq beat.Sequence, colRng, rotRng beat.Range,
	sRot, eRot, m, c float64,
	options ...any) {

	reset := opt.Get[bool]("reset", true, options)

	ctx.WSeq(seq, func(ctx context.Context) {
		if reset {
			resetG, resetB := evt.RotationGroupWithBox(ctx, groupOptions(options...)...)
			resetG.Beat -= 0.1
			ctx.WSeq(beat.Seq(0), func(ctx context.Context) {
				resetB.AddEvent(ctx, evt.ORotation(sRot), evt.EasingNone)
			})
		}

		ctx.WOpt(options...).Do(func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx)
			ctx.WRng(colRng, func(ctx context.Context) {
				cb.AddEvent(ctx, fx.ExtendTransit)
			})

			_, rb := evt.RotationGroupWithBox(ctx, evt.ODistWave(m*ctx.OrdinalF()+c))
			ctx.WRng(rotRng, func(ctx context.Context) {
				rb.AddEvent(ctx, fx.ORotation(0, 1, sRot, eRot, ease.OutCirc))
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

func clock(ctx context.Context, b, interval, speed, sr, er, wave float64, dir evt.RotationDirection, options ...any) {
	alternate := opt.Get[bool]("alternate", false, options)

	fx.RotReset(ctx, b-0.1, append(groupOptions(options...), evt.ORotation(sr))...)

	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(beat.Seq(b), func(ctx context.Context) {
			_, b := evt.RotationGroupWithBox(ctx, evt.OBeatDistStep(interval), evt.ODistWave(wave))

			if alternate {
				if dir == evt.CW {
					dir = evt.CCW
				} else if dir == evt.CCW {
					dir = evt.CW
				}
			}

			ctx.WRng(beat.RngStep(0, speed, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.ORotation(sr), evt.EasingNone)
				} else {
					b.AddEvent(ctx, evt.ORotation(er), evt.EasingInOutQuad, dir, evt.OLoop(1))
				}
			})
		})
	})
}

func gesture(ctx context.Context, seq beat.Sequence, sr, er, rotWave, rotBeatStep float64, options ...any) {
	var (
		reset      = opt.Get[bool]("reset", true, options)
		easing     = opt.Get[ease.Ing]("easing", ease.InCirc, options)
		brightness = opt.Get[evt.ColorEventOption]("brightness", fx.OBrightness(0, 1, 3.2, 1, ease.InCirc), options)
	)

	if reset {
		fx.RotReset(ctx, seq.Beats[0]-0.1, append(options, evt.ORotation(sr), evt.EasingNone)...)
	}

	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx, evt.OBeatDistStep(0.1))
			ctx.WRng(beat.RngStep(0, 0.8, 5), func(ctx context.Context) {
				cb.AddEvent(ctx, fx.InstantTransit, brightness)
			})

			_, rb := evt.RotationGroupWithBox(ctx, evt.OBeatDistStep(rotBeatStep), evt.ODistWave(rotWave))
			ctx.WRng(beat.RngStep(0, 0.5, 8), func(ctx context.Context) {
				if ctx.First() {
					if reset {
						rb.AddEvent(ctx, evt.EasingNone, evt.ORotation(sr))
					} else {
						rb.AddEvent(ctx, evt.Extend)
					}
				} else {
					rb.AddEvent(ctx, fx.ORotation(0, 1, sr, er, easing))
				}
			})
		})
	})
}

func trill(ctx context.Context, seq beat.Sequence, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx)
			ctx.WRng(beat.RngStep(0, 2, 2), func(ctx context.Context) {
				if ctx.First() {
					cb.AddEvent(ctx, evt.Extend)
				} else {
					cb.AddEvent(ctx, evt.Red, evt.Transition)
				}
			})
			_, rb := evt.RotationGroupWithBox(ctx, evt.OBeatDistStep(0.2), evt.ODistWave(37))

			ctx.WRng(beat.RngStep(0, 2, 30), func(ctx context.Context) {
				if ctx.First() {
					rb.AddEvent(ctx, evt.Extend)
				} else {
					rb.AddEvent(ctx, fx.ORotation(0, 1, 17, 84, ease.OutCirc), evt.CW)
				}
			})
		})
	})
}

func stdColor(ctx context.Context, seq beat.Sequence, rng beat.Range, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx, options...)
			ctx.WRng(rng, func(ctx context.Context) {
				cb.AddEvent(ctx, opt.Of[evt.ColorEventOption](options...)...)
			})
		})
	})
}

func smolEscalation(ctx context.Context, seq beat.Sequence, rng beat.Range, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, rb := evt.RotationGroupWithBox(ctx, thesecond.SmallRing, evt.OBeatDistWave(1.2), evt.ODistStep(15+30*ctx.OrdinalF()))

			ctx.WRng(rng, func(ctx context.Context) {
				rb.AddEvent(ctx, fx.ORotation(0, 1, 0+ctx.SeqOrdinalF()*30, 60+ctx.SeqOrdinalF()*120, ease.OutCirc))
			})

			evt.Basic(ctx, thesecond.RingZoom, evt.OValue(ctx.Ordinal()*5))
		})
	})
}

func smolReduction(ctx context.Context, seq beat.Sequence, rng beat.Range, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, rb := evt.RotationGroupWithBox(ctx, thesecond.SmallRing, evt.OBeatDistWave(0.6), evt.ODistStep(30*(3-ctx.OrdinalF())))

			ctx.WRng(rng, func(ctx context.Context) {
				if ctx.First() {
					rb.AddEvent(ctx, evt.ORotation(0-90*ctx.SeqOrdinalF()), evt.EasingNone)
				} else {
					rb.AddEvent(ctx, evt.ORotation(0-77*ctx.SeqOrdinalF()))
				}
			})

			evt.Basic(ctx, thesecond.RingZoom, evt.OValue((4-ctx.Ordinal())*3))
		})
	})
}

func incantation(ctx context.Context, seq beat.Sequence, cRng, bRng beat.Range, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx)
			_, rb := evt.RotationGroupWithBox(ctx)
			ctx.WRng(cRng, func(ctx context.Context) {
				cb.AddEvent(ctx, fx.InstantTransit)
			})
			ctx.WRng(bRng, func(ctx context.Context) {
				rb.AddEvent(ctx)
			})
		})
	})
}

func twinkleOut(ctx context.Context, seq beat.Sequence, rng beat.Range, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx)
			ctx.WRng(rng, func(ctx context.Context) {
				cb.AddEvent(ctx, fx.InstantTransit)
			})
		})
	})
}
