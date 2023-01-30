package main

import (
	"github.com/shasderias/iris/beat"
	"github.com/shasderias/iris/beatsaber"
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/ease"
	"github.com/shasderias/iris/env/thesecond"
	"github.com/shasderias/iris/evt"
	"github.com/shasderias/iris/fx"
	"github.com/shasderias/iris/opt"
	"github.com/shasderias/iris/scale"
)

func ifFirst[T any](ctx context.Context, t T, f T) T {
	if ctx.First() {
		return t
	}
	return f
}

func beatSeq[T any](ctx context.Context, items ...T) T {
	return items[ctx.SeqOrdinal()%len(items)]
}

var extendTransit = opt.IfFirst(evt.Extend, evt.Transition)
var instantTransit = opt.IfFirst(evt.Instant, evt.Transition)

func main() {
	const mapPath = `D:\Beat Saber Data\CustomWIPLevels\Otona No Okite`

	m, err := beatsaber.Open(mapPath)
	if err != nil {
		panic(err)
	}
	diff, err := m.OpenDifficulty(beatsaber.CharacteristicStandard, beatsaber.BeatmapDifficultyExpertPlus)
	if err != nil {
		panic(err)
	}

	ctx := context.NewBase()

	Intro(ctx)
	Verse1(ctx)
	Verse2(ctx)

	diff.SetEvents(*ctx.Events())

	if err := diff.Save(); err != nil {
		panic(err)
	}
}

func Intro(ctx context.Context) {
	spotlightGhost(ctx,
		beat.Seq(6, 14, 22, 26, 30),
		beat.RngInterval(0, 1, 10),
		beat.RngStep(0, 0.5, 4),
		10, -10, -13, -12,
		thesecond.SpotlightLeft, thesecond.SpotlightRight,
		fx.OBrightness(0.0, 0.4, 0, 1.6, ease.InOutCirc),
		fx.OBrightness(0.4, 1.0, 1.6, 0, ease.InOutCirc),
		opt.SeqOrdinal(evt.Red, evt.Blue),
		opt.ColorBoxOnly(evt.OBeatDistWave(1.45)),
		opt.RotationBoxOnly(evt.OBeatDistWave(1.2)),
	)
	//ctx.WSeq(beat.Seq(6, 14, 22, 26, 30), func(ctx context.Context) {
	//	var (
	//		spotGroup = evt.ColorGroup(ctx,
	//			thesecond.SpotlightLeft, thesecond.SpotlightRight)
	//		spotBox        = spotGroup.AddBox(ctx, evt.OBeatDistWave(1.75))
	//		spotBrightness = fx.NewCurve(0, 0.2, 0.8, 2.4, 2.0, 1.6, 1.2, 0.4, 0)
	//		spotColor      = opt.SeqOrdinal(evt.Red, evt.Blue)(ctx)
	//	)
	//
	//	ctx.WRng(beat.RngInterval(0, 1, 10), func(ctx context.Context) {
	//		spotBox.AddEvent(ctx, spotColor, evt.OBrightness(spotBrightness.Lerp(ctx.T())), extendTransit(ctx))
	//	})
	//
	//	spotRotReset := evt.RotationGroup(ctx, thesecond.SpotlightLeft, thesecond.SpotlightRight)
	//	spotRotReset.Beat -= 0.1
	//
	//	spotRotResetBox := spotRotReset.AddBox(ctx)
	//
	//	spotRot := evt.RotationGroup(ctx, thesecond.SpotlightLeft, thesecond.SpotlightRight)
	//	spotRotBox := spotRot.AddBox(ctx, evt.OBeatDistWave(1.2), evt.ODistWave(-12-float64(ctx.Ordinal())*13))
	//
	//	ctx.WRng(beat.RngStep(0, 0.5, 4), func(ctx context.Context) {
	//		if ctx.First() {
	//			spotRotResetBox.AddEvent(ctx, evt.ORotation(10), evt.EasingNone, evt.RotationTransitionTypeTransition)
	//		}
	//		spotRotBox.AddEvent(ctx, evt.ORotate(10, -10, evt.RotationDirectionCounterClockwise))
	//	})
	//})

	ctx.WSeq(beat.Seq(30.6), func(ctx context.Context) {
		var (
			runwayGroup = evt.ColorGroup(ctx, thesecond.RunwayLeft, thesecond.RunwayRight)
			runwayBox   = runwayGroup.AddBox(ctx,
				evt.OSectionFilter(0, 0, evt.OReverse(true)),
				evt.OBeatDistWave(1.5))
			runwayBrightness = fx.NewCurve(0, 0.6, 0)
			runwayColor      = opt.Ordinal(evt.Blue, evt.White)
		)

		ctx.WRng(beat.RngStep(0, 1.6, 3), func(ctx context.Context) {
			runwayBox.AddEvent(ctx, runwayColor(ctx), evt.OBrightness(runwayBrightness.Lerp(ease.InOutQuad.Func(ctx.T()))), extendTransit(ctx))
		})
	})
	ctx.WSeq(beat.Seq(38, 46, 54, 62), func(ctx context.Context) {
		var (
			spotGroup = evt.ColorGroup(ctx,
				thesecond.SpotlightLeft, thesecond.SpotlightRight)
			spotBox        = spotGroup.AddBox(ctx, evt.OBeatDistWave(2.2))
			spotBrightness = fx.NewCurve(0, 2.4, 0)
			spotColor      = opt.SeqOrdinal(
				[]evt.ColorEventColor{evt.Red, evt.Blue, evt.Red},
				[]evt.ColorEventColor{evt.Blue, evt.Blue, evt.White},
				[]evt.ColorEventColor{evt.Red, evt.Blue, evt.White},
				[]evt.ColorEventColor{evt.White, evt.Red, evt.Blue},
			)(ctx)
		)
		ctx.WRng(beat.RngStep(0, 1, 3), func(ctx context.Context) {
			spotBox.AddEvent(ctx, spotColor[ctx.Ordinal()], evt.OBrightness(spotBrightness.Lerp(ctx.T())), extendTransit(ctx))
		})

		spotRotReset := evt.RotationGroup(ctx, thesecond.SpotlightLeft, thesecond.SpotlightRight)
		spotRotReset.Beat -= 0.1

		spotRotResetBox := spotRotReset.AddBox(ctx)

		spotRot := evt.RotationGroup(ctx, thesecond.SpotlightLeft, thesecond.SpotlightRight)
		spotRotBox := spotRot.AddBox(ctx, evt.OBeatDistWave(1.2), evt.ODistWave(-90-float64(ctx.Ordinal())*80))

		ctx.WRng(beat.RngStep(0, 2.0, 2), func(ctx context.Context) {
			if ctx.First() {
				spotRotResetBox.AddEvent(ctx, evt.ORotation(20), evt.EasingNone, evt.RotationTransitionTypeTransition)
				spotRotBox.AddEvent(ctx, evt.ORotate(20, -20, evt.RotationDirectionCounterClockwise), evt.OLoop(0))
			} else {
				spotRotBox.AddEvent(ctx, evt.ORotate(20, -20, evt.RotationDirectionCounterClockwise), evt.OLoop(0))
			}
		})
	})

	ctx.WSeq(beat.Seq(64), func(ctx context.Context) {
		var (
			spotGroup = evt.ColorGroup(ctx,
				thesecond.SpotlightLeft, thesecond.SpotlightRight)
			spotBox        = spotGroup.AddBox(ctx, evt.OBeatDistWave(0.5))
			spotBrightness = fx.NewCurve(0, 2.4, 2.4, 2.4, 1.2)
			spotColor      = opt.SeqOrdinal(
				[]evt.ColorEventColor{evt.Red, evt.Blue, evt.White, evt.White},
			)(ctx)
		)
		ctx.WRng(beat.RngStep(0, 1, 3), func(ctx context.Context) {
			spotBox.AddEvent(ctx, spotColor[ctx.Ordinal()], evt.OBrightness(spotBrightness.Lerp(ctx.T())), evt.Transition)
		})

	})

	ctx.WSeq(beat.Seq(66), func(ctx context.Context) {
		spotRot := evt.RotationGroup(ctx, thesecond.SpotlightLeft, thesecond.SpotlightRight)
		spotRotBox := spotRot.AddBox(ctx, evt.OBeatDistWave(1.0), evt.ODistWave(0))

		ctx.WRng(beat.RngStep(0, 0.9, 1), func(ctx context.Context) {
			spotRotBox.AddEvent(ctx, evt.OLoop(1), evt.ORotation(0))
		})
	})
}

func Verse1(ctx context.Context) {
	fx.Off(ctx, 70, thesecond.SpotlightLeft, thesecond.SpotlightRight)

	fx.RotReset(ctx, 70-0.1, thesecond.TopLasersLeftTop, thesecond.TopLasersRightTop, evt.ORotation(90))
	fx.RotReset(ctx, 78-0.1, thesecond.BottomLasersLeftTop, thesecond.BottomLasersRightTop, evt.ORotation(90))

	ctx.WOpt(
		evt.OBeatDistWave(1.35),
		opt.IfSeqLast[evt.ColorEventOption](
			evt.Red,
			opt.Ordinal(evt.White, evt.Red, evt.White),
		),
		evt.OBrightnessT(ease.InCirc.Func, scale.FromUnitClamp(3, 1.2)),
		fx.InstantTransit,
	).Do(func(ctx context.Context) {
		ctx.WSeq(beat.Seq(70, 78, 86, 94), func(ctx context.Context) {
			box := evt.ColorGroup(ctx, thesecond.TopLasers).AddBox(ctx)
			ctx.WRng(beat.RngStep(0, 0.75, 3), func(ctx context.Context) {
				box.AddEvent(ctx)
			})
		})

		ctx.WSeq(beat.Seq(78, 86, 94, 95), func(ctx context.Context) {
			box := evt.ColorGroup(ctx, thesecond.BottomLasers).AddBox(ctx)
			ctx.WRng(beat.RngStep(0, 0.75, 3), func(ctx context.Context) {
				box.AddEvent(ctx)
			})
		})
	})

	ctx.WSeq(beat.Seq(70), func(ctx context.Context) {
		{
			b := evt.RotationGroup(ctx, thesecond.TopLasersLeftTop, thesecond.TopLasersRightTop).
				AddBox(ctx, evt.OBeatDistStep(2), evt.ODistWave(-45))

			ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.ORotation(90), evt.EasingNone)
				} else {
					b.AddEvent(ctx, evt.ORotation(270), evt.EasingInOutQuad, evt.CW, evt.OLoop(1))
				}
			})
		}

		{
			b := evt.RotationGroup(ctx, thesecond.TopLasersLeftBottom, thesecond.TopLasersRightBottom).
				AddBox(ctx, evt.OBeatDistStep(2), evt.ODistWave(45.0/2.0))

			ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.ORotation(90), evt.EasingNone)
				} else {
					b.AddEvent(ctx, evt.ORotation(270), evt.EasingInOutQuad, evt.CCW, evt.OLoop(1))
				}
			})

		}
	})
	ctx.WSeq(beat.Seq(78), func(ctx context.Context) {
		{
			b := evt.RotationGroup(ctx, thesecond.BottomLasersLeftTop, thesecond.BottomLasersRightTop).
				AddBox(ctx, evt.OBeatDistStep(2), evt.ODistWave(-45))

			ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.ORotation(90), evt.EasingNone)
				} else {
					b.AddEvent(ctx, evt.ORotation(270), evt.EasingInOutQuad, evt.CCW, evt.OLoop(1))
				}
			})
		}

		{
			b := evt.RotationGroup(ctx, thesecond.BottomLasersLeftBottom, thesecond.BottomLasersRightBottom).
				AddBox(ctx, evt.OBeatDistStep(2), evt.ODistWave(45.0/2.0))

			ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.ORotation(90), evt.EasingNone)
				} else {
					b.AddEvent(ctx, evt.ORotation(270), evt.EasingInOutQuad, evt.CW, evt.OLoop(1))
				}
			})

		}
	})
	ctx.WSeq(beat.Seq(86), func(ctx context.Context) {
		{
			b := evt.RotationGroup(ctx, thesecond.TopLasers, thesecond.BottomLasers).
				AddBox(ctx, evt.OBeatDistStep(2), evt.ODistWave(0))

			ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.Extend)
				} else {
					b.AddEvent(ctx, evt.ORotation(180), evt.EasingInOutQuad, evt.OLoop(1))
				}
			})
		}
	})

	ctx.WSeq(beat.Seq(94), func(ctx context.Context) {
		{
			b := evt.RotationGroup(ctx, thesecond.BottomLasers).
				AddBox(ctx, evt.OBeatDistWave(1), evt.ODistWave(71))

			ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.Extend)
				} else {
					b.AddEvent(ctx, evt.ORotation(33), evt.EasingInOutQuad)
				}
			})
		}
		{
			b := evt.RotationGroup(ctx, thesecond.TopLasers).
				AddBox(ctx, evt.OBeatDistWave(1), evt.ODistWave(33))

			ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.Extend)
				} else {
					b.AddEvent(ctx, evt.ORotation(71), evt.EasingInOutQuad)
				}
			})
		}
	})
	ctx.WSeq(beat.Seq(98.5), func(ctx context.Context) {
		{
			b := evt.ColorGroup(ctx, thesecond.BottomLasers, thesecond.TopLasers).
				AddBox(ctx, evt.OBeatDistStep(0.5))
			ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.Extend)
				} else {
					b.AddEvent(ctx, evt.White)
				}
			})
		}
		{
			b := evt.RotationGroup(ctx, thesecond.BottomLasers).
				AddBox(ctx, evt.OBeatDistStep(0.5))

			ctx.WRng(beat.RngStep(0, 0.75, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.Extend)
				} else {
					b.AddEvent(ctx, evt.ORotation(270), evt.EasingInOutQuad, evt.OLoop(0))
				}
			})
		}
		{
			g := evt.RotationGroup(ctx, thesecond.TopLasers)
			g.Beat += 0.5
			b := g.AddBox(ctx, evt.OBeatDistStep(0.5))
			ctx.WRng(beat.RngStep(0, 0.75, 2), func(ctx context.Context) {
				if ctx.First() {
					b.AddEvent(ctx, evt.Extend)
				} else {
					b.AddEvent(ctx, evt.ORotation(270), evt.EasingOutQuad, evt.OLoop(0))
				}
			})
		}
	})

	ctx.WOpt(
		evt.OBeatDistWave(1.35),
		opt.IfSeqLast[evt.ColorEventOption](
			evt.Red,
			opt.Ordinal(evt.White, evt.Red, evt.White),
		),
		evt.OBrightnessT(ease.InCirc.Func, scale.FromUnitClamp(3, 1.2)),
		fx.InstantTransit,
	).Do(func(ctx context.Context) {
		ctx.WSeq(beat.Seq(102, 110, 118), func(ctx context.Context) {
			box := evt.ColorGroup(ctx, thesecond.TopLasers, thesecond.BottomLasers).AddBox(ctx)
			ctx.WRng(beat.RngStep(0, 0.75, 3), func(ctx context.Context) {
				box.AddEvent(ctx)
			})
		})
	})

	ctx.WSeq(beat.Seq(102, 110, 118), func(ctx context.Context) {
		g := evt.RotationGroup(ctx, thesecond.TopLasers, thesecond.BottomLasers)
		b := g.AddBox(ctx, evt.ODistWave(opt.SeqOrdinal(15.0, 47.0, 7.0)(ctx)))

		ctx.WRng(beat.RngStep(0, 0.5, 2), func(ctx context.Context) {
			if ctx.First() {
				b.AddEvent(ctx, evt.Extend)
			} else {
				b.AddEvent(ctx, evt.ORotation(opt.SeqOrdinal(260.0, 96.0, 260.0)(ctx)), evt.EasingInOutQuad)
			}
		})

		g2 := evt.RotationGroup(ctx, thesecond.TopLasers, thesecond.BottomLasers)
		g2.Beat += 0.1
		b2 := g2.AddBox(ctx,
			evt.OBeatDistStep(2),
			evt.ODistWave(opt.SeqOrdinal(43.0, 13.0, -37.0)(ctx)),
			evt.ODistAffectFirst(true),
		)

		ctx.WRng(beat.RngStep(0, 0.5, 2), func(ctx context.Context) {
			if ctx.First() {
				b2.AddEvent(ctx, evt.Extend)
			} else {
				b2.AddEvent(ctx,
					evt.ORotation(opt.SeqOrdinal(43.0, 290.0, 11.0)(ctx)),
					evt.EasingInOutQuad, evt.CCW, evt.OLoop(1))
			}
		})
	})

	ctx.WOpt(
		evt.OBeatDistWave(0.8),
		opt.Ordinal(evt.White, evt.Red, evt.Blue, evt.White),
		evt.OBrightnessT(ease.OutCirc.Func, scale.FromUnitClamp(3, 0)),
		fx.InstantTransit,
	).Do(func(ctx context.Context) {
		ctx.WSeq(beat.Seq(125), func(ctx context.Context) {
			box := evt.ColorGroup(ctx, thesecond.TopLasers, thesecond.BottomLasers).AddBox(ctx)
			ctx.WRng(beat.RngStep(0, 2.6, 4), func(ctx context.Context) {
				box.AddEvent(ctx)
			})
		})
	})

	ctx.WSeq(beat.Seq(125), func(ctx context.Context) {
		g := evt.RotationGroup(ctx, thesecond.TopLasers, thesecond.BottomLasers)
		b := g.AddBox(ctx, evt.ODistWave(31), evt.OBeatDistStep(0.4))

		ctx.WSeq(beat.Seq(0, 2.1), func(ctx context.Context) {
			if ctx.First() {
				b.AddEvent(ctx, evt.Extend)
			} else {
				b.AddEvent(ctx,
					opt.SeqOrdinal(evt.RotAuto, evt.CCW, evt.RotAuto, evt.RotAuto),
					evt.ORotation(opt.SeqOrdinal(0.0, -70.0, -75.0, -65.0)(ctx)),
					opt.SeqOrdinal(evt.EasingNone, evt.EasingInOutQuad, evt.EasingLinear, evt.EasingOutQuad),
				)
			}
		})
	})

	ctx.WSeq(beat.Seq(129, 129.5), func(ctx context.Context) {
		ctx.WOpt(opt.SeqOrdinal(thesecond.TopLasers, thesecond.BottomLasers)).Do(func(ctx context.Context) {
			g := evt.ColorGroup(ctx)
			b := g.AddBox(ctx)
			ctx.WRng(beat.RngStep(0, 1.0, 10), func(ctx context.Context) {
				b.AddEvent(ctx, fx.InstantTransit, evt.White,
					evt.OBrightness(scale.FromUnitClamp(3, 0)(ease.OutCirc.Func(ctx.T()))))
			})

			rg := evt.RotationGroup(ctx)
			rb := rg.AddBox(ctx, evt.ODistWave(31))

			ctx.WSeq(beat.Seq(0, 0.7), func(ctx context.Context) {
				if ctx.First() {
					rb.AddEvent(ctx, evt.EasingNone, evt.ORotation(23))
				} else {
					rb.AddEvent(ctx, evt.EasingLinear, evt.ORotation(33))
				}
			})
		})
	})
}

func Verse2(ctx context.Context) {
	ctx.WSeq(beat.Seq(131, 132, 133), func(ctx context.Context) {
		_, b := evt.ColorGroupWithBox(ctx, opt.SeqOrdinal[evt.ColorEventGroupOption](thesecond.SmallRing, thesecond.Runway, thesecond.SmallRing))

		ctx.WRng(beat.RngStep(0, 1.5, 10), func(ctx context.Context) {
			b.AddEvent(ctx, fx.InstantTransit, evt.Red,
				evt.OBrightness(scale.FromUnitClamp(1.8, 0)(ease.OutCirc.Func(ctx.T()))))
		})

		if ctx.First() {
			rg := evt.RotationGroup(ctx, thesecond.SmallRing, thesecond.BigRing)
			rb := rg.AddBox(ctx, evt.ODistWave(-71), evt.OBeatDistWave(1.2))

			ctx.WSeq(beat.Seq(0, 3), func(ctx context.Context) {
				if ctx.First() {
					rb.AddEvent(ctx, evt.EasingNone, evt.ORotation(-15))
				} else {
					rb.AddEvent(ctx, evt.EasingLinear, evt.ORotation(60))
				}
			})
		}
	})

	gesture(ctx, beat.Seq(134), 43, 47, 13, 0.1, true, thesecond.TopLasersLeftTop)
	gesture(ctx, beat.Seq(136), 131, 138, 19, 0.1, true, thesecond.TopLasersRightTop)
	gesture(ctx, beat.Seq(137), 31, 39, 21, 0.1, true, thesecond.BottomLasersRightTop)
	gesture(ctx, beat.Seq(139), 23, 25, 6, 0.1, true, thesecond.TopLasersLeftBottom)
	gesture(ctx, beat.Seq(140), 24, 26, 7, 0.1, true, thesecond.BottomLasersLeftTop)
	gesture(ctx, beat.Seq(141), 20, 30, 17, 0.15, true, thesecond.BottomLasersLeftBottom)

	gesture(ctx, beat.Seq(142), 47, 49, 13, 0.1, true, thesecond.TopLasersLeftTop, evt.White)
	gesture(ctx, beat.Seq(142), 138, 140, 19, 0.1, true, thesecond.TopLasersRightTop, evt.White)
	gesture(ctx, beat.Seq(142), 39, 41, 21, 0.1, true, thesecond.BottomLasersRightTop, evt.Blue)
	gesture(ctx, beat.Seq(142), 25, 27, 6, 0.1, true, thesecond.TopLasersLeftBottom, evt.White)
	gesture(ctx, beat.Seq(142), 26, 28, 7, 0.1, true, thesecond.BottomLasersLeftTop, evt.Blue)
	gesture(ctx, beat.Seq(142), 30, 32, 17, 0.1, true, thesecond.BottomLasersLeftBottom, evt.Blue)

	gesture(ctx, beat.Seq(144), 71, 74, 9, 0.1, true, thesecond.TopLasersRightBottom, evt.Red)
	gesture(ctx, beat.Seq(145), 73, 81, 9, 0.1, true, thesecond.BottomLasersRightBottom, evt.Red)

	gesture(ctx, beat.Seq(147), 9, 12, 13, 0.1, false, thesecond.TopLasers, evt.Red)
	gesture(ctx, beat.Seq(148), 13, 21, 26, 0.1, false, thesecond.BottomLasers, evt.Blue)
	gesture(ctx, beat.Seq(149), 10, 17, 5, 0.05, false, thesecond.TopLasers, thesecond.BottomLasers, evt.White)

	trill(ctx, beat.Seq(150), thesecond.TopLasers, thesecond.BottomLasers)

	gesture(ctx, beat.Seq(152), 14, 16, 13, 0.1, true, thesecond.TopLasersLeftTop, thesecond.TopLasersRightBottom, evt.Blue)
	gesture(ctx, beat.Seq(153), 16, 14, 13, 0.1, true, thesecond.TopLasersLeftBottom, thesecond.TopLasersRightTop, evt.Blue)

	stdColor(ctx, beat.Seq(155), beat.RngStep(0, 1, 10),
		thesecond.TopLasersLeftTop, thesecond.TopLasersRightBottom,
		evt.White,
		fx.OBrightness(0, 1, 3.5, 0, ease.OutCirc),
	)
	stdColor(ctx, beat.Seq(156), beat.RngStep(0, 1, 10),
		thesecond.TopLasersLeftBottom, thesecond.TopLasersRightTop,
		evt.White,
		fx.OBrightness(0, 1, 3.5, 0, ease.OutCirc),
	)
	stdColor(ctx, beat.Seq(157), beat.RngStep(0, 1, 10),
		thesecond.BottomLasersTop,
		evt.White,
		fx.OBrightness(0, 1, 3.5, 0, ease.OutCirc),
	)
	stdColor(ctx, beat.Seq(158), beat.RngStep(0, 1, 10),
		thesecond.BottomLasersBottom,
		evt.White,
		fx.OBrightness(0, 1, 3.5, 0, ease.OutCirc),
	)
}

func gesture(ctx context.Context, seq beat.Sequence, sr, er, rotWave, rotBeatStep float64, reset bool, options ...any) {
	if reset {
		fx.RotReset(ctx, seq.Beats[0]-0.1, append(options, evt.ORotation(sr), evt.EasingNone)...)
	}

	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(seq, func(ctx context.Context) {
			_, cb := evt.ColorGroupWithBox(ctx, evt.OBeatDistStep(0.1))
			ctx.WRng(beat.RngStep(0, 0.8, 5), func(ctx context.Context) {
				cb.AddEvent(ctx, fx.InstantTransit,
					fx.OBrightness(0, 1, 3.2, 1, ease.InCirc))
			})

			_, rb := evt.RotationGroupWithBox(ctx, evt.OBeatDistStep(rotBeatStep), evt.ODistWave(rotWave))
			ctx.WSeq(beat.Seq(0, 0.3), func(ctx context.Context) {
				if ctx.First() {
					if reset {
						rb.AddEvent(ctx, evt.EasingNone, evt.ORotation(sr))
					} else {
						rb.AddEvent(ctx, evt.Extend)
					}
				} else {
					rb.AddEvent(ctx, evt.EasingOutQuad, evt.ORotation(er))
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
				cb.AddEvent(ctx, opt.For[evt.ColorEventOption](options...)...)
			})
		})
	})
}
