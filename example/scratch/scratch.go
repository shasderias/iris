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
	Verse3(ctx)
	Verse4(ctx)

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
	spotlightGhost(ctx,
		beat.Seq(38, 46, 54, 62),
		beat.RngStep(0, 1, 3),
		beat.RngStep(0, 2.0, 2),
		22, -20, -80, -90,
		thesecond.SpotlightLeft, thesecond.SpotlightRight,
		fx.OBrightness(0.0, 0.5, 0, 2.4, ease.InOutCirc),
		fx.OBrightness(0.5, 1.0, 2.4, 0, ease.InOutCirc),
		opt.SeqOrdinal(
			opt.Ordinal(evt.Red, evt.Blue, evt.Red),
			opt.Ordinal(evt.Blue, evt.Blue, evt.White),
			opt.Ordinal(evt.Red, evt.Blue, evt.White),
			opt.Ordinal(evt.White, evt.Red, evt.Blue),
		),
		opt.ColorBoxOnly(evt.OBeatDistWave(2.1)),
		opt.RotationBoxOnly(evt.OBeatDistWave(1.2)),
	)
	//ctx.WSeq(beat.Seq(38, 46, 54, 62), func(ctx context.Context) {
	//	var (
	//		spotGroup = evt.ColorGroup(ctx,
	//			thesecond.SpotlightLeft, thesecond.SpotlightRight)
	//		spotBox        = spotGroup.AddBox(ctx, evt.OBeatDistWave(2.2))
	//		spotBrightness = fx.NewCurve(0, 2.4, 0)
	//		spotColor      = opt.SeqOrdinal(
	//			[]evt.ColorEventColor{evt.Red, evt.Blue, evt.Red},
	//			[]evt.ColorEventColor{evt.Blue, evt.Blue, evt.White},
	//			[]evt.ColorEventColor{evt.Red, evt.Blue, evt.White},
	//			[]evt.ColorEventColor{evt.White, evt.Red, evt.Blue},
	//		)(ctx)
	//	)
	//	ctx.WRng(beat.RngStep(0, 1, 3), func(ctx context.Context) {
	//		spotBox.AddEvent(ctx, spotColor[ctx.Ordinal()], evt.OBrightness(spotBrightness.Lerp(ctx.T())), extendTransit(ctx))
	//	})
	//
	//	spotRotReset := evt.RotationGroup(ctx, thesecond.SpotlightLeft, thesecond.SpotlightRight)
	//	spotRotReset.Beat -= 0.1
	//
	//	spotRotResetBox := spotRotReset.AddBox(ctx)
	//
	//	spotRot := evt.RotationGroup(ctx, thesecond.SpotlightLeft, thesecond.SpotlightRight)
	//	spotRotBox := spotRot.AddBox(ctx, evt.OBeatDistWave(1.2), evt.ODistWave(-90-float64(ctx.Ordinal())*80))
	//
	//	ctx.WRng(beat.RngStep(0, 2.0, 2), func(ctx context.Context) {
	//		if ctx.First() {
	//			spotRotResetBox.AddEvent(ctx, evt.ORotation(20), evt.EasingNone, evt.RotationTransitionTypeTransition)
	//			spotRotBox.AddEvent(ctx, evt.ORotate(20, -20, evt.RotationDirectionCounterClockwise), evt.OLoop(0))
	//		} else {
	//			spotRotBox.AddEvent(ctx, evt.ORotate(20, -20, evt.RotationDirectionCounterClockwise), evt.OLoop(0))
	//		}
	//	})
	//})

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

	clock(ctx, 70, 2, 1, 90, 270, -45, evt.CW,
		thesecond.TopLasersTop)
	clock(ctx, 70, 2, 1, 90, 270, 22.5, evt.CCW,
		thesecond.TopLasersBottom)
	clock(ctx, 78, 2, 1, 90, 270, 45, evt.CW,
		thesecond.BottomLasersTop)
	clock(ctx, 78, 2, 1, 90, 270, -22.5, evt.CCW,
		thesecond.BottomLasersBottom)

	//ctx.WSeq(beat.Seq(70), func(ctx context.Context) {
	//	{
	//		b := evt.RotationGroup(ctx, thesecond.TopLasersLeftTop, thesecond.TopLasersRightTop).
	//			AddBox(ctx, evt.OBeatDistStep(2), evt.ODistWave(-45))
	//
	//		ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
	//			if ctx.First() {
	//				b.AddEvent(ctx, evt.ORotation(90), evt.EasingNone)
	//			} else {
	//				b.AddEvent(ctx, evt.ORotation(270), evt.EasingInOutQuad, evt.CW, evt.OLoop(1))
	//			}
	//		})
	//	}
	//
	//	{
	//		b := evt.RotationGroup(ctx, thesecond.TopLasersLeftBottom, thesecond.TopLasersRightBottom).
	//			AddBox(ctx, evt.OBeatDistStep(2), evt.ODistWave(45.0/2.0))
	//
	//		ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
	//			if ctx.First() {
	//				b.AddEvent(ctx, evt.ORotation(90), evt.EasingNone)
	//			} else {
	//				b.AddEvent(ctx, evt.ORotation(270), evt.EasingInOutQuad, evt.CCW, evt.OLoop(1))
	//			}
	//		})
	//
	//	}
	//})
	//ctx.WSeq(beat.Seq(78), func(ctx context.Context) {
	//	{
	//		b := evt.RotationGroup(ctx, thesecond.BottomLasersLeftTop, thesecond.BottomLasersRightTop).
	//			AddBox(ctx, evt.OBeatDistStep(2), evt.ODistWave(-45))
	//
	//		ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
	//			if ctx.First() {
	//				b.AddEvent(ctx, evt.ORotation(90), evt.EasingNone)
	//			} else {
	//				b.AddEvent(ctx, evt.ORotation(270), evt.EasingInOutQuad, evt.CCW, evt.OLoop(1))
	//			}
	//		})
	//	}
	//
	//	{
	//		b := evt.RotationGroup(ctx, thesecond.BottomLasersLeftBottom, thesecond.BottomLasersRightBottom).
	//			AddBox(ctx, evt.OBeatDistStep(2), evt.ODistWave(45.0/2.0))
	//
	//		ctx.WRng(beat.RngStep(0, 1, 2), func(ctx context.Context) {
	//			if ctx.First() {
	//				b.AddEvent(ctx, evt.ORotation(90), evt.EasingNone)
	//			} else {
	//				b.AddEvent(ctx, evt.ORotation(270), evt.EasingInOutQuad, evt.CW, evt.OLoop(1))
	//			}
	//		})
	//
	//	}
	//})
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
			ctx.WRng(beat.RngStep(0, 1.8, 10), func(ctx context.Context) {
				b.AddEvent(ctx, fx.InstantTransit, evt.White,
					evt.OBrightness(scale.FromUnitClamp(3, 0)(ease.InCirc.Func(ctx.T()))))
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

	resetF := opt.Set("reset", false)

	gesture(ctx, beat.Seq(134), 43, 47, 13, 0.1, thesecond.TopLasersLeftTop)
	gesture(ctx, beat.Seq(136), 131, 138, 19, 0.1, thesecond.TopLasersRightTop)
	gesture(ctx, beat.Seq(137), 31, 39, 21, 0.1, thesecond.BottomLasersRightTop)
	gesture(ctx, beat.Seq(139), 23, 25, 6, 0.1, thesecond.TopLasersLeftBottom)
	gesture(ctx, beat.Seq(140), 24, 26, 7, 0.1, thesecond.BottomLasersLeftTop)
	gesture(ctx, beat.Seq(141), 20, 30, 17, 0.15, thesecond.BottomLasersLeftBottom)

	gesture(ctx, beat.Seq(142), 47, 49, 13, 0.1, thesecond.TopLasersLeftTop, evt.White)
	gesture(ctx, beat.Seq(142), 138, 140, 19, 0.1, thesecond.TopLasersRightTop, evt.White)
	gesture(ctx, beat.Seq(142), 39, 41, 21, 0.1, thesecond.BottomLasersRightTop, evt.Blue)
	gesture(ctx, beat.Seq(142), 25, 27, 6, 0.1, thesecond.TopLasersLeftBottom, evt.White)
	gesture(ctx, beat.Seq(142), 26, 28, 7, 0.1, thesecond.BottomLasersLeftTop, evt.Blue)
	gesture(ctx, beat.Seq(142), 30, 32, 17, 0.1, thesecond.BottomLasersLeftBottom, evt.Blue)

	gesture(ctx, beat.Seq(144), 71, 74, 9, 0.1, thesecond.TopLasersRightBottom, evt.Red)
	gesture(ctx, beat.Seq(145), 73, 81, 9, 0.1, thesecond.BottomLasersRightBottom, evt.Red)

	kimochiEasing := opt.Set("easing", ease.OutQuad)
	gesture(ctx, beat.Seq(147), 9, 12, 7, 0.05, thesecond.TopLasers, evt.Red, kimochiEasing, resetF)
	gesture(ctx, beat.Seq(148), 13, 21, 7, 0.05, thesecond.BottomLasers, evt.Blue, kimochiEasing, resetF)
	gesture(ctx, beat.Seq(149), 10, 17, 7, 0.05, thesecond.TopLasers, thesecond.BottomLasers, evt.White, kimochiEasing, resetF)

	trill(ctx, beat.Seq(150), thesecond.TopLasers, thesecond.BottomLasers)

	gesture(ctx, beat.Seq(152), 14, 16, 13, 0.1, thesecond.TopLasersLeftTop, thesecond.TopLasersRightBottom, evt.Blue)
	gesture(ctx, beat.Seq(153), 16, 14, 13, 0.1, thesecond.TopLasersLeftBottom, thesecond.TopLasersRightTop, evt.Blue)

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

	ctx.WSeq(beat.Seq(160, 162), func(ctx context.Context) {
		light := opt.SeqOrdinal(thesecond.SpotlightLeft, thesecond.SpotlightRight)(ctx)
		spotlightGhost(ctx,
			beat.Seq(0),
			beat.RngStep(0, 1.6, 4),
			beat.RngStep(0, 2.4, 10),
			20, -20, 0, -45,
			light,
			fx.OBrightness(0.0, 0.5, 0, 2.4, ease.InOutCirc),
			fx.OBrightness(0.5, 1.0, 2.4, 0, ease.InOutCirc),
			opt.Ordinal(evt.Red, evt.Red, evt.White, evt.White),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
			evt.CCW,
		)
	})
}
func Verse3(ctx context.Context) {
	spotlightGhost(ctx,
		beat.Seq(164),
		beat.RngStep(0, 2, 4),
		beat.RngStep(0, 2.4, 10),
		-20, 20, 0, -75,
		thesecond.SpotlightLeft, thesecond.SpotlightRight,
		fx.OBrightness(0.0, 1.0, 0, 2.4, ease.OutCirc),
		opt.Ordinal(evt.White, evt.Red),
		opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
		opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
		evt.CW,
	)

	ctx.WSeq(beat.Seq(164), func(ctx context.Context) {
		stdColor(ctx, beat.Seq(0), beat.RngStep(0, 1, 2),
			thesecond.Runway, thesecond.BigRing,
			opt.Ordinal(evt.White, evt.Red),
			opt.Ordinal(evt.OBrightness(0), evt.OBrightness(2.8)),
			fx.InstantTransit,
			evt.OBeatDistStep(0.08),
			evt.OStepAndOffsetFilter(0, 1, evt.OOrder(evt.FilterOrderRandom, 123)),
		)
	})
	ctx.WSeq(beat.Seq(164), func(ctx context.Context) {
		stdColor(ctx, beat.Seq(0), beat.RngStep(0, 1, 2),
			thesecond.SmallRing,
			opt.Ordinal(evt.Red, evt.White),
			opt.Ordinal(evt.OBrightness(0), evt.OBrightness(1.0)),
			fx.InstantTransit,
			evt.OBeatDistStep(0.04),
			evt.OStepAndOffsetFilter(0, 1, evt.OOrder(evt.FilterOrderRandom, 123)),
		)
	})

	fx.Off(ctx, 166, thesecond.Spotlight)

	smolEscSeq := beat.Seq(166, 168, 170, 172)

	spotlightGhost(ctx,
		smolEscSeq,
		beat.RngStep(0, 1.4, 3),
		beat.RngStep(0, 1.4, 10),
		-20+15*ctx.SeqOrdinalF(), 20+15*ctx.SeqOrdinalF(), 15, -15-27*ctx.SeqOrdinalF(),
		thesecond.SpotlightLeft, thesecond.SpotlightRight,
		opt.IfSeqFirst[evt.ColorEventOption](
			fx.OBrightness(0.0, 1.0, 2.4, 0, ease.OutCirc),
			opt.Combine(
				fx.OBrightness(0.0, 0.5, 0, 3.2, ease.OutCirc),
				fx.OBrightness(0.5, 1.0, 3.2, 0, ease.OutCirc),
			),
		),
		opt.Ordinal(evt.Blue, evt.White, evt.Blue),
		opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
		opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
		opt.IfFirst(evt.RotAuto, evt.CW),
		opt.Set("reset", false),
	)

	smolEscalation(ctx,
		smolEscSeq,
		beat.RngStep(0, 1, 10),
	)

	stdColor(ctx,
		smolEscSeq,
		beat.RngStep(0, 1, 10),
		thesecond.SmallRing,
		opt.SeqOrdinal(evt.Blue, evt.White),
		opt.IfSeqFirst[evt.ColorEventOption](
			fx.OBrightness(0.0, 1.0, 2.4, 0, ease.OutCirc),
			opt.Combine(
				fx.OBrightness(0.0, 0.5, 0, 3.2, ease.OutCirc),
				fx.OBrightness(0.5, 1.0, 3.2, 0, ease.OutCirc),
			),
		),
		fx.InstantTransit,
		evt.OBeatDistStep(0.04),
	)

	stdColor(ctx,
		smolEscSeq,
		beat.RngStep(0, 1, 10),
		thesecond.BigRing,
		opt.SeqOrdinal(evt.White, evt.Red),
		opt.IfSeqFirst[evt.ColorEventOption](
			fx.OBrightness(0.0, 1.0, 2.4, 0, ease.OutCirc),
			opt.Combine(
				fx.OBrightness(0.0, 0.5, 0, 3.2, ease.OutCirc),
				fx.OBrightness(0.5, 1.0, 3.2, 0, ease.OutCirc),
			),
		),
		fx.InstantTransit,
		evt.OBeatDistStep(0.04),
	)

	stdColor(ctx,
		smolEscSeq,
		beat.RngStep(0, 1, 10),
		thesecond.Runway,
		evt.Red,
		opt.IfSeqFirst[evt.ColorEventOption](
			fx.OBrightness(0.0, 1.0, 2.4, 0, ease.OutCirc),
			opt.Combine(
				fx.OBrightness(0.0, 0.5, 0, 3.2, ease.OutCirc),
				fx.OBrightness(0.5, 1.0, 3.2, 0, ease.OutCirc),
			),
		),
		fx.InstantTransit,
		evt.OBeatDistStep(0.04),
	)

	// 174, 175, 176, 177
	fusawashiSeq := beat.Seq(174, 175, 176, 177)

	smolReduction(ctx,
		fusawashiSeq,
		beat.RngStep(0, 0.31, 2),
	)

	ctx.WSeq(fusawashiSeq, func(ctx context.Context) {
		rng := opt.IfSeqLast(
			beat.RngStep(0, 2.2, 10),
			beat.RngStep(0, 0.8, 10),
		)(ctx)
		col := opt.IfSeqLast(
			evt.Red,
			opt.SeqOrdinal(evt.Blue, evt.White)(ctx),
		)(ctx)
		bri := opt.IfSeqLast(
			fx.OBrightness(0.0, 1.0, 2.4, 0, ease.OutCirc),
			fx.OBrightness(0.0, 1.0, 1.2, 0, ease.InCirc),
		)(ctx)
		stdColor(ctx,
			beat.Seq(0),
			rng,
			thesecond.SmallRing,
			col, bri,
			fx.InstantTransit,
		)
	})

	stdColor(ctx,
		fusawashiSeq,
		beat.RngStep(0, 1, 10),
		thesecond.BigRing,
		evt.OStepAndOffsetFilter(0, 1, evt.OReverse(true)),
		opt.SeqOrdinal(evt.White, evt.Red),
		opt.IfSeqFirst[evt.ColorEventOption](
			fx.OBrightness(0.0, 1.0, 2.4, 0, ease.OutCirc),
			opt.Combine(
				fx.OBrightness(0.0, 0.5, 0, 3.2, ease.OutCirc),
				fx.OBrightness(0.5, 1.0, 3.2, 0, ease.OutCirc),
			),
		),
		fx.InstantTransit,
		evt.OBeatDistStep(0.04),
	)

	stdColor(ctx,
		fusawashiSeq,
		beat.RngStep(0, 1, 10),
		thesecond.Runway,
		evt.OStepAndOffsetFilter(0, 1, evt.OReverse(true)),
		evt.Red,
		opt.IfSeqFirst[evt.ColorEventOption](
			fx.OBrightness(0.0, 1.0, 2.4, 0, ease.OutCirc),
			opt.Combine(
				fx.OBrightness(0.0, 0.5, 0, 3.2, ease.OutCirc),
				fx.OBrightness(0.5, 1.0, 3.2, 0, ease.OutCirc),
			),
		),
		fx.InstantTransit,
		evt.OBeatDistStep(0.04),
	)
	// 179, 180, 181, 182
	horobiSeq := beat.Seq(179, 180, 181)
	stdColor(ctx,
		horobiSeq,
		beat.RngStep(0, 0.8, 3),
		thesecond.Runway,
		//evt.OBeatDistWave(1.0),
		evt.OSectionFilter(0, 0, evt.OReverse(true)),
		evt.Blue,
		opt.Ordinal(evt.Blue, evt.White, evt.Blue),
		opt.Ordinal(evt.OBrightness(0), evt.OBrightness(0.2), evt.OBrightness(0.1)),
		fx.ExtendTransit,
		//fx.OBrightness(0, 0.5, 0, 0.4, ease.InOutQuad),
		//fx.OBrightness(0.5, 1, 0.4, 0, ease.InOutQuad),
	)

	// 182, 183, 184, 185,
	// 186, 187, 188, 189
	cRng := beat.RngStep(0, 1.2, 2)
	incantation(ctx,
		beat.Seq(182),
		cRng,
		beat.RngStep(0, 6.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersLeftTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
		fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(183),
		cRng,
		beat.RngStep(0, 5.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersRightTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
		fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(184),
		cRng,
		beat.RngStep(0, 4.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersLeftBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(185),
		cRng,
		beat.RngStep(0, 3.5, 30),
		opt.Ordinal(evt.Blue, evt.Red),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersRightBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(186),
		cRng,
		beat.RngStep(0, 2.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.BottomLasersLeftBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(1), evt.ODistWave(121)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(187),
		cRng,
		beat.RngStep(0, 1.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.BottomLasersRightBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(1), evt.ODistWave(121)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(188),
		cRng,
		beat.RngStep(0, 0.5, 30),
		opt.Ordinal(evt.White, evt.Red),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.BottomLasersTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(0.4), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)

	ctx.WSeq(beat.Seq(189), func(ctx context.Context) {
		_, cb := evt.ColorGroupWithBox(ctx, thesecond.TopLasers, thesecond.BottomLasers)
		ctx.WRng(beat.RngStep(0, 0.4, 2), func(ctx context.Context) {
			cb.AddEvent(ctx,
				opt.Ordinal(evt.Red, evt.White),
				opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
				fx.InstantTransit,
			)
		})

		_, rb1 := evt.RotationGroupWithBox(ctx, thesecond.TopLasersTop, thesecond.BottomLasersBottom)
		ctx.WRng(beat.RngStep(0, 1, 1), func(ctx context.Context) {
			rb1.AddEvent(ctx, evt.Transition, evt.ORotation(180))
		})
		_, rb2 := evt.RotationGroupWithBox(ctx, thesecond.TopLasersBottom, thesecond.BottomLasersTop)
		ctx.WRng(beat.RngStep(0, 1, 1), func(ctx context.Context) {
			rb2.AddEvent(ctx, evt.Transition, evt.ORotation(0))
		})
	})

	fx.Off(ctx, 189, thesecond.Runway)
}

func Verse4(ctx context.Context) {
	stdColor(ctx,
		beat.Seq(198),
		beat.RngStep(0, 8, 20),
		thesecond.TopLasersLeftTop,
		evt.OBeatDistWave(6),
		evt.Blue,
		fx.ExtendTransit,
		fx.OBrightness(0, 1.0, 1.8, 0.7, ease.InOutCirc),
	)
	stdColor(ctx,
		beat.Seq(206),
		beat.RngStep(0, 8, 20),
		thesecond.TopLasersLeftBottom,
		evt.OBeatDistWave(6),
		evt.Blue,
		fx.ExtendTransit,
		fx.OBrightness(0, 1.0, 1.8, 0.7, ease.InOutCirc),
	)
	stdColor(ctx,
		beat.Seq(214),
		beat.RngStep(0, 8, 20),
		thesecond.BottomLasersLeftTop,
		evt.OBeatDistWave(6),
		evt.White,
		fx.ExtendTransit,
		fx.OBrightness(0, 1.0, 1.8, 0.7, ease.InOutCirc),
	)
	stdColor(ctx,
		beat.Seq(222),
		beat.RngStep(0, 8, 20),
		thesecond.BottomLasersLeftBottom,
		evt.OBeatDistWave(6),
		evt.White,
		fx.ExtendTransit,
		fx.OBrightness(0, 1.0, 1.8, 0.7, ease.InOutCirc),
	)
	clock(ctx, 198, 2, 1, 90, 260, -19, evt.CW,
		thesecond.TopLasersLeftTop)
	clock(ctx, 206, 2, 1, -90, -260, -19, evt.CCW,
		thesecond.TopLasersLeftBottom)
	clock(ctx, 214, 2, 1, 90, 71, 19, evt.CW, opt.Set("alternate", true),
		thesecond.BottomLasersLeftTop)
	clock(ctx, 222, 2, 1, 193, -111, 18, evt.CCW,
		thesecond.BottomLasersLeftBottom)

	// sekai 229, 230
	// fusekai 231.50, 232, 233
	gesture(ctx, beat.Seq(229), 260, 270, 45, 0.25, thesecond.TopLasersLeftTop, evt.White, opt.Set("reset", true))
	gesture(ctx, beat.Seq(229), -260, -270, 45, 0.25, thesecond.TopLasersLeftBottom, evt.Blue, opt.Set("reset", true))
	gesture(ctx, beat.Seq(229), 71, 90, 45, 0.25, thesecond.BottomLasersLeftTop, evt.Blue, opt.Set("reset", true))
	gesture(ctx, beat.Seq(229), -111, -90, 45, 0.25, thesecond.BottomLasersLeftBottom, evt.White, opt.Set("reset", true))

	gesture(ctx, beat.Seq(231.5), 270, 273, -45, 1.5/4, thesecond.TopLasersLeftTop, evt.Blue, opt.Set("reset", false))
	gesture(ctx, beat.Seq(231.5), -270, -273, -45, 1.5/4, thesecond.TopLasersLeftBottom, evt.White, opt.Set("reset", false))
	gesture(ctx, beat.Seq(231.5), 90, 93, -45, 1.5/4, thesecond.BottomLasersLeftTop, evt.White, opt.Set("reset", false))
	gesture(ctx, beat.Seq(231.5), -90, -93, -45, 1.5/4, thesecond.BottomLasersLeftBottom, evt.Blue, opt.Set("reset", false))

	// dochira 235, 236, 237

	stdColor(ctx,
		beat.Seq(235, 237),
		beat.RngStep(0, 1.4, 2),
		thesecond.TopLasersLeftTop, thesecond.BottomLasersLeftBottom,
		opt.Ordinal(evt.Red, evt.White),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.8), evt.OBrightness(0.7)),
	)

	stdColor(ctx,
		beat.Seq(236),
		beat.RngStep(0, 1.4, 2),
		thesecond.TopLasersLeftBottom, thesecond.BottomLasersLeftTop,
		opt.Ordinal(evt.Red, evt.Blue),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.8), evt.OBrightness(0.7)),
	)

	briNoReset := opt.Combine(
		opt.Set("brightness", fx.OBrightness(0, 1, 1.6, 0.7, ease.InCirc)),
		opt.Set("reset", false),
	)
	// kawo 239, 240
	gesture(ctx, beat.Seq(239), 273, 267, -12, 0.02,
		thesecond.TopLasersLeftTop, evt.Blue, briNoReset)
	gesture(ctx, beat.Seq(239), -273, -267, -12, 0.02,
		thesecond.TopLasersLeftBottom, evt.White, briNoReset)
	gesture(ctx, beat.Seq(239), 93, 87, -12, 0.02,
		thesecond.BottomLasersLeftTop, evt.White, briNoReset)
	gesture(ctx, beat.Seq(239), -93, -87, -12, 0.02,
		thesecond.BottomLasersLeftBottom, evt.Blue, briNoReset)

	gesture(ctx, beat.Seq(240), 267, 273, 12, 0.02,
		thesecond.TopLasersLeftTop, evt.White, briNoReset)
	gesture(ctx, beat.Seq(240), -267, -273, 12, 0.02,
		thesecond.TopLasersLeftBottom, evt.Blue, briNoReset)
	gesture(ctx, beat.Seq(240), 87, 93, 12, 0.02,
		thesecond.BottomLasersLeftTop, evt.Blue, briNoReset)
	gesture(ctx, beat.Seq(240), -87, -93, 12, 0.02,
		thesecond.BottomLasersLeftBottom, evt.White, briNoReset)

	// eraberu 243.5, 244, 244.5, 245
	eraberuOpts := opt.Combine(
		evt.Red,
		fx.InstantTransit,
		evt.OStepAndOffsetFilter(0, 1, evt.OOrder(evt.FilterOrderRandom, evt.SeedRand)),
		fx.OBrightness(0, 1, 3.5, 0, ease.OutCirc),
	)
	stdColor(ctx,
		beat.Seq(243.5),
		beat.RngStep(0, 5, 20),
		thesecond.TopLasersLeftTop,
		evt.OBeatDistStep(0.55),
		eraberuOpts,
	)

	stdColor(ctx,
		beat.Seq(243.5),
		beat.RngStep(0, 4.5, 20),
		thesecond.TopLasersLeftBottom,
		evt.OBeatDistStep(0.45),
		eraberuOpts,
	)
	stdColor(ctx,
		beat.Seq(244),
		beat.RngStep(0, 4, 20),
		thesecond.BottomLasersLeftTop,
		evt.Red,
		evt.OBeatDistStep(0.35),
		eraberuOpts,
	)
	stdColor(ctx,
		beat.Seq(244.5),
		beat.RngStep(0, 4, 20),
		thesecond.BottomLasersLeftBottom,
		evt.OBeatDistStep(0.35),
		eraberuOpts,
	)

	clock(ctx, 246, 2, 1, 17, 19, 19, evt.CW,
		thesecond.TopLasersRight)
}
