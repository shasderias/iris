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
	fx.RotReset(ctx, b-0.1, append(groupOptions(options...), evt.ORotation(sr))...)

	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(beat.Seq(b), func(ctx context.Context) {
			_, b := evt.RotationGroupWithBox(ctx, evt.OBeatDistStep(interval), evt.ODistWave(wave))

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

func gesture(ctx context.Context, seq beat.Sequence, sr, er, rotWave, rotBeatWave float64, options ...any) {
	var (
		reset      = opt.Get[bool]("reset", true, options)
		easing     = opt.Get[ease.Ing]("easing", ease.OutElastic, options)
		brightness = opt.Get[evt.ColorEventOption]("brightness", fx.OBrightness(0, 1, 3.2, 1, ease.InCirc), options)
		cRng       = opt.Get[beat.Range]("colorRng", beat.RngStep(0, 0.8, 7), options)
	)

	if reset {
		fx.RotReset(ctx, seq.Beats[0]-0.1, append(options, evt.ORotation(sr), evt.EasingNone)...)
	}

	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx, evt.OBeatDistStep(0.1))
			ctx.WRng(cRng, func(ctx context.Context) {
				cb.AddEvent(ctx, fx.InstantTransit, brightness)
			})

			_, rb := evt.RotationGroupWithBox(ctx, evt.OBeatDistWave(rotBeatWave), evt.ODistWave(rotWave))
			ctx.WRng(beat.RngStep(0, 0.5, 12), func(ctx context.Context) {
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

func smolSimple(ctx context.Context, seq beat.Sequence, rng beat.Range, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, rb := evt.RotationGroupWithBox(ctx, thesecond.SmallRing)
			ctx.WRng(rng, func(ctx context.Context) {
				rb.AddEvent(ctx)
			})
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
				} else if ctx.Last() {
					rb.AddEvent(ctx, evt.ORotation(90))
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

func starfield(ctx context.Context, seq beat.Sequence, rng1, rng2 beat.Range, peakBri float64, options ...any) {
	bri := opt.Combine(
		fx.OBrightness(0.0, 0.2, 0, 0, ease.OutCirc),
		fx.OBrightness(0.2, 0.5, 0, peakBri, ease.OutCirc),
		fx.OBrightness(0.5, 0.8, peakBri, 0, ease.InCirc),
		fx.OBrightness(0.8, 1.0, 0, 0, ease.InCirc),
	)
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			cg := evt.ColorGroup(ctx, thesecond.BigRing)
			cb1 := cg.AddBox(ctx,
				evt.OStepAndOffsetFilter(0, 2, evt.OOrder(evt.FilterOrderRandom, evt.SeedRand)))
			cb2 := cg.AddBox(ctx,
				evt.OStepAndOffsetFilter(1, 2))
			ctx.WRng(rng1, func(ctx context.Context) {
				cb1.AddEvent(ctx, bri, opt.SeqOrdinal(evt.White, evt.Blue))
			})
			ctx.WRng(rng2, func(ctx context.Context) {
				cb2.AddEvent(ctx, bri, opt.SeqOrdinal(evt.Blue, evt.White))
			})
		})
	})
}

func starfieldSpin(ctx context.Context, sb, eb float64, options ...any) {
	rngEB := eb - sb
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(beat.Seq(sb), func(ctx context.Context) {
			g := evt.RotationGroup(ctx, thesecond.BigRing)
			b1 := g.AddBox(ctx,
				evt.OStepAndOffsetFilter(0, 2, evt.OOrder(evt.FilterOrderRandom, evt.SeedRand)),
				evt.ODistWave(91),
			)
			b2 := g.AddBox(ctx,
				evt.OStepAndOffsetFilter(1, 2),
				evt.ODistWave(171),
			)
			ctx.WRng(beat.RngStep(0, rngEB, 80), func(ctx context.Context) {
				b1.AddEvent(ctx, fx.ORotation(0, 1, 31, 360*1+31, ease.InOutQuad), evt.CW)
			})
			ctx.WRng(beat.RngStep(0, rngEB, 80), func(ctx context.Context) {
				b2.AddEvent(ctx, fx.ORotation(0, 1, 360*0.5+43, 43, ease.InOutQuad), evt.CCW)
			})
		})
	})
}

func splitSmolRing(ctx context.Context, options ...any) {
	//ctx.WOpt(options...).Do(func(ctx context.Context) {
	//	ctx.WSeq(seq, func(ctx context.Context) {
	//		_, b1 := evt.ColorGroupWithBox(ctx,
	//			thesecond.SmallRing,
	//			evt.OStepAndOffsetFilter()
	//		)
	//		_, b2 := evt.ColorGroupWithBox(ctx,
	//			thesecond.SmallRing,
	//		)
	//		ctx.WRng(rng, func(ctx context.Context) {
	//
	//		})
	//	})
	//})
}
func sparkleFadeReset(ctx context.Context, b float64) {
	ctx.WSeq(beat.Seq(b), func(ctx context.Context) {
		_, b := evt.RotationGroupWithBox(ctx,
			thesecond.TopLasers, thesecond.BottomLasers,
			evt.ODistWave(33), evt.AffectFirst)
		ctx.WRng(beat.RngStep(0, 1, 1), func(ctx context.Context) {
			b.AddEvent(ctx, evt.ORotation(0), evt.EasingNone)
		})
	})
}

func sparkleFade(ctx context.Context, seq beat.Sequence, rng beat.Range, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx,
				evt.OSectionFilter(0, 0, evt.OOrder(evt.FilterOrderRandom, evt.SeedRand)),
			)
			ctx.WRng(rng, func(ctx context.Context) {
				cb.AddEvent(ctx)
			})
		})
	})
}

func sparkleFadeMotion(ctx context.Context, seq beat.Sequence, rng beat.Range, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, rb := evt.RotationGroupWithBox(ctx,
				evt.OBeatDistWave(6),
				evt.ODistWave(67.5),
			)

			ctx.WRng(rng, func(ctx context.Context) {
				rb.AddEvent(ctx, fx.ORotation(0, 1, 0, 12.5, ease.OutCirc))
			})
		})
	})
}

func darkPeak(delay, peak, peakBrightness float64, easeIn, easeOut ease.Ing) opt.Combined {
	return opt.Combine(
		fx.OBrightness(0.0, delay, 0, 0, easeIn),
		fx.OBrightness(delay, peak, 0, peakBrightness, easeIn),
		fx.OBrightness(peak, 1.0-delay, peakBrightness, 0, easeOut),
		fx.OBrightness(1.0-delay, 1.0, 0, 0, easeOut),
	)
}

func rotHold(ctx context.Context, b float64, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(beat.Seq(b), func(ctx context.Context) {
			_, b := evt.RotationGroupWithBox(ctx)
			ctx.WRng(beat.RngStep(0, 1, 1), func(ctx context.Context) {
				b.AddEvent(ctx, evt.Extend)
			})
		})
	})
}
