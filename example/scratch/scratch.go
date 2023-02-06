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

	Reset(ctx)
	Intro(ctx)
	Verse1(ctx)
	Build1(ctx)
	Chorus1(ctx)
	Verse2(ctx)
	Build2(ctx)
	Chorus2(ctx)
	Bridge1(ctx)
	Bridge2(ctx)
	SilentBridge(ctx)
	Chorus3(ctx)
	Chorus4(ctx)
	Outro(ctx)

	diff.SetEvents(*ctx.Events())

	if err := diff.Save(); err != nil {
		panic(err)
	}
}

func Reset(ctx context.Context) {
	ctx.WSeq(beat.Seq(0), func(ctx context.Context) {
		evt.Basic(ctx, thesecond.RingZoom, evt.OValue(9))
		_, rb := evt.RotationGroupWithBox(ctx,
			thesecond.SmallRing, thesecond.BigRing,
			thesecond.Lasers,
			thesecond.Spotlight,
		)
		ctx.WRng(beat.RngStep(0, 1, 1), func(ctx context.Context) {
			rb.AddEvent(ctx, evt.ORotation(0))
		})
	})
}

func Intro(ctx context.Context) {
	spotlightGhost(ctx,
		beat.Seq(6, 14, 22, 26, 30),
		beat.RngStep(0, 1.2, 12),
		beat.RngStep(0, 1.2, 12),
		10, -10, -13, -12,
		thesecond.Spotlight,
		darkPeak(0, 0.4, 1.2, ease.OutCirc, ease.InCirc),
		opt.SeqOrdinal(evt.Red, evt.Blue),
		opt.ColorBoxOnly(evt.OBeatDistWave(2)),
		opt.RotationBoxOnly(evt.OBeatDistWave(2)),
	)

	stdColor(ctx, beat.Seq(30.6), beat.RngStep(0, 1.6, 3),
		thesecond.Runway,
		evt.OSectionFilter(0, 0, evt.OReverse(true)),
		evt.OBeatDistWave(1.5),
		opt.Ordinal(evt.OBrightness(0), evt.OBrightness(0.6), evt.OBrightness(0)),
		opt.Ordinal(evt.Blue, evt.White),
		fx.ExtendTransit,
	)

	spotlightGhost(ctx,
		beat.Seq(38, 46, 54, 62),
		beat.RngStep(0, 1, 3),
		beat.RngStep(0, 1.8, 2),
		22, -20, -80, -90,
		thesecond.Spotlight,
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

	stdColor(ctx, beat.Seq(64), beat.RngStep(0, 1, 3),
		thesecond.Spotlight,
		evt.OBeatDistWave(0.5),
		fx.OBrightness(0, 0.3, 0.0, 2.4, ease.InOutCirc),
		fx.OBrightness(0.3, 1, 2.4, 1.2, ease.InOutCirc),
		opt.Ordinal(evt.Red, evt.Blue, evt.White),
		fx.InstantTransit,
	)

	stdRotation(ctx, beat.Seq(66), beat.RngStep(0, 0.9, 1),
		thesecond.Spotlight,
		evt.OBeatDistWave(1.0), evt.ODistWave(0),
		evt.OLoop(1), evt.ORotation(0),
	)

	fx.Off(ctx, 70, thesecond.Spotlight)
}

func Verse1(ctx context.Context) {
	fx.RotReset(ctx, 70-0.1, thesecond.TopLasersLeftTop, thesecond.TopLasersRightTop, evt.ORotation(90))
	fx.RotReset(ctx, 78-0.1, thesecond.BottomLasersLeftTop, thesecond.BottomLasersRightTop, evt.ORotation(90))

	ctx.WOpt(
		evt.OBeatDistWave(1.35),
		opt.IfSeqLast[evt.ColorEventOption](
			evt.Red,
			opt.Ordinal(evt.White, evt.Red, evt.White),
		),
		fx.OBrightness(0, 1, 3, 1.2, ease.InCirc),
		fx.InstantTransit,
	).Do(func(ctx context.Context) {
		stdColor(ctx,
			beat.Seq(70, 78, 86, 94),
			beat.RngStep(0, 0.75, 3),
			thesecond.TopLasers,
		)
		stdColor(ctx,
			beat.Seq(78, 86, 94, 94.5),
			beat.RngStep(0, 0.75, 3),
			thesecond.BottomLasers,
		)
	})

	clock(ctx, 70, 2, 1, 90, 270, -30, evt.CW,
		thesecond.TopLasersTop)
	clock(ctx, 70, 2, 1, 90, 270, 22.5, evt.CCW,
		thesecond.TopLasersBottom)
	clock(ctx, 78, 2, 1, 90, 270, 30, evt.CW,
		thesecond.BottomLasersTop)
	clock(ctx, 78, 2, 1, 90, 270, -22.5, evt.CCW,
		thesecond.BottomLasersBottom)

	clock(ctx, 86, 2, 1, 270, 180, 0, evt.RotAuto,
		thesecond.Lasers, clockTransit)
	clock(ctx, 94, 0, 1, 180, 33, 71, evt.RotAuto,
		thesecond.BottomLasers, clockTransit, clockBDistWave(1), clockLoop(0))
	clock(ctx, 94.5, 0, 1, 180, 33, 71, evt.RotAuto,
		thesecond.TopLasers, clockTransit, clockBDistWave(1), clockLoop(0))

	stdColor(ctx, beat.Seq(98.5), beat.RngStep(0, 1, 2),
		thesecond.Lasers, evt.OBeatDistStep(0.5), evt.White, fx.ExtendTransit)

	clock(ctx, 98.5, 0.5, 0.75, 33, 270, 0, evt.RotAuto,
		thesecond.BottomLasers,
		clockTransit, clockLoop(0))
	clock(ctx, 99.0, 0.5, 0.75, 33, 270, 0, evt.RotAuto,
		thesecond.TopLasers,
		clockTransit, clockLoop(0))

	stdColor(ctx, beat.Seq(102, 110, 118), beat.RngStep(0, 0.75, 3),
		thesecond.Lasers,
		evt.OBeatDistWave(1.35),
		opt.IfSeqLast[evt.ColorEventOption](
			evt.Red,
			opt.Ordinal(evt.White, evt.Red, evt.White),
		),
		fx.OBrightness(0, 1, 3, 1.2, ease.InCirc),
		fx.InstantTransit,
	)

	clock2(ctx, 102, 2, 0.5, 260, 43, 15, 43, thesecond.Lasers)
	clock2(ctx, 110, 2, 0.5, 96, 290, 47, 13, thesecond.Lasers)
	clock2(ctx, 118, 2, 0.5, 260, 35, 7, -37, thesecond.Lasers)

	stdColor(ctx, beat.Seq(123), beat.RngStep(0, 2.5, 12), thesecond.TopLasersLeft,
		fx.OBrightness(0, 1, 1.2, 0, ease.InCirc),
		opt.T(evt.White, evt.Blue, evt.Blue),
		fx.InstantTransit,
	)
	fx.RotHold(ctx, 123-0.1, thesecond.TopLasersLeft)
	stdRotation(ctx, beat.Seq(123), beat.RngStep(0, 3, 20),
		thesecond.TopLasersLeft,
		evt.OBeatDistWave(1), evt.ODistWave(-37),
		fx.ORotation(0, 1, 35, 60, ease.OutCirc),
	)

	stdColor(ctx, beat.Seq(124), beat.RngStep(0, 2.5, 12), thesecond.TopLasersRight,
		fx.OBrightness(0, 1, 1.2, 0, ease.InCirc),
		opt.T(evt.White, evt.Blue, evt.Blue),
		fx.InstantTransit,
	)
	fx.RotHold(ctx, 124-0.1, thesecond.TopLasersRight)
	stdRotation(ctx, beat.Seq(124), beat.RngStep(0, 3, 20),
		thesecond.TopLasersRight,
		evt.OBeatDistWave(1), evt.ODistWave(-37),
		fx.ORotation(0, 1, 35, 60, ease.OutCirc),
	)

	stdColor(ctx, beat.Seq(125), beat.RngStep(0, 1.9, 12), thesecond.BottomLasers,
		fx.OBrightness(0, 1, 1.2, 0, ease.InCirc),
		opt.T(evt.White, evt.Red, evt.Red),
		fx.InstantTransit,
	)
	fx.RotHold(ctx, 125-0.1, thesecond.BottomLasers)
	stdRotation(ctx, beat.Seq(125), beat.RngStep(0, 3, 20),
		thesecond.BottomLasers,
		evt.ODistWave(10),
		fx.ORotation(0, 1, 35, 60, ease.OutCirc),
	)

	fx.RotHold(ctx, 129-0.2, thesecond.Lasers)
	stdColor(ctx,
		beat.Seq(128.9, 129.3),
		beat.RngStep(0, 1.3, 10),
		opt.SeqOrdinal(thesecond.TopLasers, thesecond.BottomLasers),
		evt.White, fx.InstantTransit,
		//evt.OBeatDistWave(0.9),
		fx.OBrightness(0, 1, 1.2, 0, ease.InOutCirc),
	)

	ctx.WSeq(beat.Seq(128.9, 129.3), func(ctx context.Context) {
		_, b := evt.RotationGroupWithBox(ctx,
			opt.SeqOrdinal(thesecond.TopLasers, thesecond.BottomLasers),
			evt.ODistWave(91), evt.OBeatDistWave(1.5))
		ctx.WRng(beat.RngStep(0, 1.6, 20), func(ctx context.Context) {
			b.AddEvent(ctx, fx.ORotation(0, 1, 23, 33, ease.InCirc))
		})
	})
}

func Build1(ctx context.Context) {
	softWave(ctx, beat.Seq(131, 132, 133), thesecond.Runway)

	stdColor(ctx,
		beat.SeqInterval(134, 158, 0.5),
		beat.RngStep(0, 1.3, 10),
		opt.SeqOrdinal(thesecond.SpotlightLeft, thesecond.SpotlightRight),
		opt.SeqOrdinal(evt.Blue, evt.Blue, evt.White, evt.White),
		evt.OBeatDistWave(2),
		fx.InstantTransit,
		fx.OBrightness(0, 1, 2.4, 0, ease.InCirc),
	)

	sideBounce(ctx,
		beat.SeqInterval(133.6, 157.6, 0.5),
		beat.RngStep(0, 1.6, 20),
		evt.OBeatDistWave(1.8),
		sideBounceAlt,
		fx.ORotation(0, 1, 335, 305, ease.OutBounce),
	)

	//resetF := opt.Set("reset", false)

	gesture(ctx, beat.Seq(134), 40, 100, 13, 1.2,
		thesecond.TopLasersLeftTop, thesecond.BottomLasersLeftBottom,
		opt.T(evt.White, evt.Red, evt.Red),
	)

	gesture(ctx, beat.Seq(136), 74, 39, 38, 0.6,
		thesecond.BottomLasersRightTop, evt.Red)
	gesture(ctx, beat.Seq(137), 74, 39, 19, 0.6,
		thesecond.TopLasersRightTop, evt.Blue)

	ctx.WSeq(beat.Seq(139, 140, 141), func(ctx context.Context) {
		gesture(ctx, beat.Seq(0), 74, 70, -13, 0.1,
			opt.SeqOrdinal(evt.Blue, evt.Red, evt.Red),
			opt.SeqOrdinal(thesecond.TopLasersLeftBottom, thesecond.TopLasersRightBottom, thesecond.BottomLasersLeftTop),
		)
	})

	ctx.WSeq(beat.Seq(142), func(ctx context.Context) {
		gesture(ctx, beat.Seq(0), 74+180, 70+180, -13, 0.1,
			thesecond.TopLasersLeftTop,
			opt.T(evt.White, evt.Red, evt.Red))
		gesture(ctx, beat.Seq(0), 74+180, 70+180, -13, 0.1,
			thesecond.TopLasersRightTop,
			opt.T(evt.White, evt.Blue, evt.Blue))
		gesture(ctx, beat.Seq(0), 74+180, 70+180, -13, 0.1,
			thesecond.BottomLasersLeftBottom,
			opt.T(evt.White, evt.Blue, evt.Blue))

		gesture(ctx, beat.Seq(0), 0, 0, 0, 0.1,
			thesecond.BottomLasersRightTop,
			evt.Blue)
		gesture(ctx, beat.Seq(0), 180, 180, 0, 0.1,
			thesecond.BottomLasersRightBottom,
			evt.Red)
	})

	gesture(ctx, beat.Seq(144), 70, 75, -30, 0.1,
		thesecond.TopLasersLeftBottom,
		opt.T(evt.Red, evt.White, evt.White))
	gesture(ctx, beat.Seq(145), 100, 135, -30, 0.1,
		thesecond.BottomLasersRightBottom,
		opt.T(evt.Blue, evt.White, evt.White))

	// kimochi - 147, 148, 149
	ctx.WSeq(beat.Seq(147, 148, 149), func(ctx context.Context) {
		lightColor := opt.SeqOrdinal(
			opt.Combine(thesecond.TopLasersTop, opt.Ordinal(evt.Blue, evt.White)),
			opt.Combine(thesecond.TopLasersBottom, opt.Ordinal(evt.Red, evt.White)),
			opt.Combine(thesecond.BottomLasers, opt.Ordinal(evt.Blue, evt.White)),
		)
		fx.RotReset(ctx, -0.1, lightColor, evt.ORotation(0))
		ctx.WOpt(
			opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
			opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
			fx.ORotation(0, 1, 0, 31, ease.OutCirc),
		).Do(func(ctx context.Context) {
			incantation(ctx, beat.Seq0,
				beat.RngStep(0, 1.2, 2),
				beat.RngStep(0, 6.5, 30),
				lightColor,
			)
		})
	})

	trill(ctx, beat.Seq(150), thesecond.TopLasers, evt.Red)
	trill(ctx, beat.Seq(150), thesecond.BottomLasers, evt.Red)

	gesture(ctx, beat.Seq(152), 14, 16, 13, 0.1,
		thesecond.TopLasersLeftTop, thesecond.TopLasersRightBottom, evt.Blue)
	gesture(ctx, beat.Seq(153), 16, 14, 13, 0.1,
		thesecond.TopLasersLeftBottom, thesecond.TopLasersRightTop, evt.Blue)

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
		light := opt.SeqOrdinal(thesecond.SpotlightLeft, thesecond.SpotlightRight)
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

func Chorus1(ctx context.Context) {
	ctx.WSeq(beat.Seq(164), func(ctx context.Context) {
		spotlightGhost(ctx, beat.Seq0,
			beat.RngStep(0, 1.2, 4),
			beat.RngStep(0, 1.2, 10),
			-80, 40, 0, 120,
			thesecond.SpotlightLeft, thesecond.SpotlightRight,
			fx.OBrightness(0.0, 1.0, 0, 2.4, ease.OutCirc),
			opt.T(evt.White, evt.Red),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(2.0)),
			evt.CW,
		)
		randFill(ctx, beat.Seq0,
			beat.RngStep(0, 0.8, 2),
			thesecond.Runway, thesecond.BigRing,
			fx.OBrightness(0, 1, 0, 2.4, ease.InOutCirc),
			opt.T(evt.White, evt.Red),
			evt.OBeatDistWave(1.7),
		)
		randFill(ctx, beat.Seq0,
			beat.RngStep(0, 0.8, 2),
			thesecond.SmallRing,
			fx.OBrightness(0, 1, 0, 1.0, ease.InOutCirc),
			opt.T(evt.Red, evt.White),
			evt.OBeatDistWave(1.7),
		)
	})

	fx.Off(ctx, 166, thesecond.Spotlight)

	ctx.WSeq(beat.Seq(166, 168, 170, 172), func(ctx context.Context) {
		spotlightGhost(ctx, beat.Seq0,
			beat.RngStep(0, 1.4, 3),
			beat.RngStep(0, 1.4, 10),
			-20+15*ctx.SeqOrdinalF(), 20+15*ctx.SeqOrdinalF(), 15, -15-27*ctx.SeqOrdinalF(),
			thesecond.Spotlight,
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

		//smolEscalation2(ctx, beat.Seq0,
		//	beat.RngStep(0, 1, 10),
		//	1.5,
		//	30, 15,
		//	30, 0,
		//	120, 60,
		//	ease.OutCirc,
		//)

		smolEscalation3(ctx, beat.Seq0,
			beat.RngStep(0, 1, 10),
			scale.FromUnitClamp(1.2, 1.8),
			scale.FromUnitClamp(15, 85),
			scale.FromUnitClamp(0, 90),
			scale.FromUnitClamp(60, 420),
			ease.OutCirc,
		)

		ctx.WOpt(
			opt.IfSeqFirst[evt.ColorEventOption](
				fx.OBrightness(0.0, 1.0, 2.4, 0, ease.OutCirc),
				opt.Combine(
					fx.OBrightness(0.0, 0.5, 0, 3.2, ease.OutCirc),
					fx.OBrightness(0.5, 1.0, 3.2, 0, ease.OutCirc),
				),
			),
			fx.InstantTransit,
		).Do(func(ctx context.Context) {
			stdColor(ctx, beat.Seq0,
				beat.RngStep(0, 1, 10),
				thesecond.SmallRing,
				opt.SeqOrdinal(evt.Blue, evt.White),
				evt.OBeatDistStep(0.04),
			)

			stdColor(ctx, beat.Seq0,
				beat.RngStep(0, 1, 10),
				thesecond.BigRing,
				opt.SeqOrdinal(evt.White, evt.Red),
				evt.OBeatDistWave(1.4),
			)

			stdColor(ctx, beat.Seq0,
				beat.RngStep(0, 1, 10),
				thesecond.Runway,
				opt.SeqOrdinal(evt.Blue, evt.Red),
				evt.OBeatDistWave(1.4),
			)
		})
	})

	// fusawashi - 174, 175, 176, 177
	ctx.WSeq(beat.Seq(174, 175, 176, 177), func(ctx context.Context) {
		if ctx.SeqFirst() || ctx.SeqLast() {
			fx.RotReset(ctx, 0, thesecond.BigRing, evt.ORotation(0))
		}

		smolReduction2(ctx, beat.Seq0,
			beat.RngStep(0, 0.31, 10),
			scale.FromUnitClamp(0.6, 0.6),
			scale.FromUnitClamp(111, 90),
			scale.FromUnitClamp(420, 90),
			scale.FromUnitClamp(420, 90),
			ease.OutCirc,
		)

		ctx.WSeq(beat.Seq0, func(ctx context.Context) {
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

		stdColor(ctx, beat.Seq0,
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

		stdColor(ctx, beat.Seq0,
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
	})

	// horobi -  179, 180, 181
	softWave(ctx, beat.Seq(179, 180, 181), thesecond.Runway)

	// 182, 183, 184, 185,
	// 186, 187, 188, 189

	incantation2(ctx, 182, 0, 5.5, 2.4, 1, 90, 180, 3, 45,
		thesecond.TopLasersRightTop, opt.T(evt.Red, evt.Blue))
	incantation2(ctx, 183, 0, 6.5, 2.4, 1, 90, 180, 3, 45,
		thesecond.TopLasersLeftTop, opt.T(evt.White, evt.Red))
	incantation2(ctx, 184, 0, 4.5, 2.4, 1, 135, 270, 3, 45,
		thesecond.TopLasersLeftBottom, opt.T(evt.White, evt.Blue))
	incantation2(ctx, 185, 0, 3.5, 2.4, 1, 135, 270, 3, 45,
		thesecond.TopLasersRightBottom, opt.T(evt.Blue, evt.Red))
	incantation2(ctx, 186, 0, 2.5, 2.4, 1, 180, 260, 1, 30,
		thesecond.BottomLasersLeftBottom, opt.T(evt.Red, evt.White))
	incantation2(ctx, 187, 0, 1.5, 2.4, 1, 180, 260, 1, 30,
		thesecond.BottomLasersRightBottom, opt.T(evt.Red, evt.White))
	incantation2(ctx, 188, 0, 0.5, 2.4, 1, 180, 270, 0.4, 90,
		thesecond.BottomLasersTop, opt.T(evt.White, evt.Red))

	fx.RotHold(ctx, 189-0.1, thesecond.Lasers)

	ctx.WSeq(beat.Seq(189), func(ctx context.Context) {
		_, lasersCb := evt.ColorGroupWithBox(ctx, thesecond.Lasers)
		_, ringCb := evt.ColorGroupWithBox(ctx, thesecond.SmallRing, evt.OBeatDistWave(0.6))
		ctx.WRng(beat.RngStep(0, 0.4, 2), func(ctx context.Context) {
			lasersCb.AddEvent(ctx,
				opt.Ordinal(evt.Red, evt.White),
				opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
				fx.InstantTransit,
			)
			ringCb.AddEvent(ctx,
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

func Verse2(ctx context.Context) {
	ctx = context.WOptions(ctx, gestureEasing(ease.OutCirc))

	resetFalse := opt.Set("reset", false)

	ctx.WSeq(beat.Seq(198, 206, 214, 222), func(ctx context.Context) {
		stdColor(ctx, beat.Seq0,
			beat.RngStep(0, 8, 20),
			opt.SeqOrdinal(thesecond.TopLasersLeftTop, thesecond.TopLasersRightTop, thesecond.BottomLasersLeftTop, thesecond.BottomLasersRightTop),
			opt.SeqOrdinal(evt.Blue, evt.White, evt.Blue, evt.White),
			evt.OBeatDistWave(6),
			fx.ExtendTransit,
			fx.OBrightness(0, 1.0, 1.8, 0.7, ease.InOutCirc),
		)
	})
	clock(ctx, 198, 2, 1, 90, 260, -19, evt.CW,
		thesecond.TopLasersLeftTop)
	clock(ctx, 206, 2, 1, 90, 260, -19, evt.CCW,
		thesecond.TopLasersRightTop)
	clock(ctx, 214, 2, 1, 0, 71, 19, evt.CW,
		thesecond.BottomLasersLeftTop)
	clock(ctx, 222, 2, 1, 0, 69, 18, evt.CCW,
		thesecond.BottomLasersRightTop)

	// sekai 229, 230
	ctx.WSeq(beat.Seq(229), func(ctx context.Context) {
		gesture(ctx, beat.Seq0, 260, 270, 45, 0.25,
			thesecond.TopLasersLeftTop, evt.White)
		gesture(ctx, beat.Seq0, 260, 270, 45, 0.25,
			thesecond.TopLasersRightTop, evt.Blue)
		gesture(ctx, beat.Seq0, 71, 90, 45, 0.25,
			thesecond.BottomLasersLeftTop, evt.Blue)
		gesture(ctx, beat.Seq0, 69, 90, 45, 0.25,
			thesecond.BottomLasersRightTop, evt.White)
	})

	// fusekai 231.50, 232, 233
	ctx.WSeq(beat.Seq(231.5), func(ctx context.Context) {
		gesture(ctx, beat.Seq0, 270, 273, -25, 1.5/4,
			thesecond.TopLasersLeftTop, evt.Blue, resetFalse)
		gesture(ctx, beat.Seq0, 270, 267, -25, 1.5/4,
			thesecond.TopLasersRightTop, evt.White, resetFalse)
		gesture(ctx, beat.Seq0, 90, 93, -25, 1.5/4,
			thesecond.BottomLasersLeftTop, evt.White, resetFalse)
		gesture(ctx, beat.Seq0, 90, 87, -25, 1.5/4,
			thesecond.BottomLasersRightTop, evt.Blue, resetFalse)
	})

	ctx.WSeq(beat.Seq(233), func(ctx context.Context) {
		gesture(ctx, beat.Seq0, 273, 267, -25, 1.5/4,
			thesecond.TopLasersLeftTop, evt.Blue, resetFalse)
		gesture(ctx, beat.Seq0, 267, 273, -25, 1.5/4,
			thesecond.TopLasersRightTop, evt.White, resetFalse)
		gesture(ctx, beat.Seq0, 93, 87, -25, 1.5/4,
			thesecond.BottomLasersLeftTop, evt.White, resetFalse)
		gesture(ctx, beat.Seq0, 87, 93, -25, 1.5/4,
			thesecond.BottomLasersRightTop, evt.Blue, resetFalse)
	})

	// dochira 235, 236, 237
	stdColor(ctx,
		beat.Seq(235, 237),
		beat.RngStep(0, 1.4, 2),
		thesecond.TopLasersRightTop, thesecond.BottomLasersLeftTop,
		opt.Ordinal(evt.Red, evt.White),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.8), evt.OBrightness(0.7)),
	)

	stdColor(ctx,
		beat.Seq(236),
		beat.RngStep(0, 1.4, 2),
		thesecond.TopLasersLeftTop, thesecond.BottomLasersRightTop,
		opt.Ordinal(evt.Red, evt.Blue),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.8), evt.OBrightness(0.7)),
	)

	briNoReset := opt.Combine(
		opt.Set("brightness", fx.OBrightness(0, 1, 1.6, 0.7, ease.InCirc)),
		opt.Set("reset", false),
	)
	// kawo 239, 240
	ctx.WSeq(beat.Seq(239), func(ctx context.Context) {
		gesture(ctx, beat.Seq0, 267, 273, -12, 0.02,
			thesecond.TopLasersLeftTop, evt.Blue, briNoReset)
		gesture(ctx, beat.Seq0, 273, 267, -12, 0.02,
			thesecond.TopLasersRightTop, evt.White, briNoReset)
		gesture(ctx, beat.Seq0, 87, 93, -12, 0.02,
			thesecond.BottomLasersLeftTop, evt.White, briNoReset)
		gesture(ctx, beat.Seq0, 93, 87, -12, 0.02,
			thesecond.BottomLasersRightTop, evt.Blue, briNoReset)
	})

	ctx.WSeq(beat.Seq(240), func(ctx context.Context) {
		gesture(ctx, beat.Seq0, 273, 270, 12, 0.02,
			thesecond.TopLasersLeftTop, evt.White, briNoReset)
		gesture(ctx, beat.Seq0, 267, 270, 12, 0.02,
			thesecond.TopLasersRightTop, evt.Blue, briNoReset)
		gesture(ctx, beat.Seq0, 93, 90, 12, 0.02,
			thesecond.BottomLasersLeftTop, evt.Blue, briNoReset)
		gesture(ctx, beat.Seq0, 87, 90, 12, 0.02,
			thesecond.BottomLasersRightTop, evt.White, briNoReset)
	})

	// eraberu 243.5, 244, 244.5, 245
	ctx.WOpt(
		fx.InstantTransit,
		evt.OStepAndOffsetFilter(0, 1, evt.OOrder(evt.FilterOrderRandom, evt.SeedRand)),
		fx.OBrightness(0, 1, 3.5, 0, ease.OutCirc),
	).Do(func(ctx context.Context) {
		stdColor(ctx,
			beat.Seq(243.5),
			beat.RngStep(0, 6, 20),
			thesecond.TopLasersLeftTop, evt.Red, evt.OBeatDistStep(0.55),
		)
		stdColor(ctx,
			beat.Seq(244),
			beat.RngStep(0, 5.5, 20),
			thesecond.BottomLasersLeftTop, evt.Red, evt.OBeatDistStep(0.45),
		)
		stdColor(ctx,
			beat.Seq(244.5),
			beat.RngStep(0, 5, 20),
			thesecond.BottomLasersRightTop, evt.White, evt.OBeatDistStep(0.45),
		)
		stdColor(ctx,
			beat.Seq(245),
			beat.RngStep(0, 4.5, 20),
			thesecond.TopLasersRightTop, evt.White, evt.OBeatDistStep(0.35),
		)
	})

	clock(ctx, 246, 2, 1, 17, 90, -12, evt.CW,
		thesecond.TopLasersBottom)
	stdColor(ctx, beat.Seq(246),
		beat.RngStep(0, 8, 20),
		thesecond.TopLasersBottom, evt.Blue,
		evt.OBeatDistWave(6),
		fx.ExtendTransit,
		fx.OBrightness(0, 1.0, 1.4, 0.7, ease.OutCirc),
	)

	clock(ctx, 254, 2, 1, 115, 90, -12, evt.CCW,
		thesecond.BottomLasersBottom)
	stdColor(ctx, beat.Seq(254),
		beat.RngStep(0, 8, 20),
		thesecond.BottomLasersBottom, evt.White,
		evt.OBeatDistWave(6),
		fx.ExtendTransit,
		fx.OBrightness(0, 1.0, 1.4, 0.7, ease.OutCirc),
	)

	stdColor(ctx,
		beat.Seq(257),
		beat.RngStep(0, 1.5, 2),
		thesecond.BottomLasersBottom,
		opt.Ordinal(evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.4), evt.OBrightness(0)),
	)
	stdColor(ctx,
		beat.Seq(257.5),
		beat.RngStep(0, 1.5, 2),
		thesecond.TopLasersBottom,
		opt.Ordinal(evt.Red, evt.White),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.4), evt.OBrightness(0)),
	)
}

func Build2(ctx context.Context) {
	stdColor(ctx,
		beat.Seq(259, 260, 261),
		beat.RngStep(0, 0.8, 2),
		thesecond.Runway,
		opt.Ordinal(evt.Blue, evt.White),
		fx.InstantTransit,
		evt.ODistWave(0.4),
		opt.IfSeqLast(
			opt.T(evt.OBrightness(0.7), evt.OBrightness(0.0)),
			opt.T(evt.OBrightness(0.7), evt.OBrightness(0.4)),
		),
	)

	ctx.WSeq(beat.Seq(262), func(ctx context.Context) {
		stdColor(ctx, beat.Seq0,
			beat.RngStep(0, 1.3, 10),
			thesecond.Spotlight,
			evt.OBeatDistWave(1.8),
			opt.T(evt.Red, evt.White),
			fx.OBrightness(0, 1, 2.4, 0, ease.InCirc),
			fx.InstantTransit,
		)

		stdRotation(ctx, beat.Seq0,
			beat.RngStep(0, 1.4, 20),
			thesecond.Spotlight,
			evt.OBeatDistWave(1.8), evt.ODistWave(-53),
			fx.ORotation(0, 1, 323, 423, ease.OutCirc),
		)
		//ctx.WSeq(beat.SeqInterval(262.7, 286.7, 0.5), func(ctx context.Context) {
		//	_, b := evt.RotationGroupWithBox(ctx,
		//		opt.SeqOrdinal(thesecond.SpotlightLeft, thesecond.SpotlightRight),
		//		evt.OBeatDistWave(1.8), evt.ODistWave(11),
		//	)
		//	ctx.WRng(beat.RngStep(0, 1.4, 20), func(ctx context.Context) {
		//		if ctx.SeqFirst() {
		//			b.AddEvent(ctx, evt.ORotation(325), evt.EasingNone)
		//		} else {
		//			b.AddEvent(ctx, fx.ORotation(0, 1.0, 323, 307, ease.OutBounce))
		//		}
		//	})
		//})
	})

	fx.RotReset(ctx, 264-0.1, thesecond.Spotlight, evt.ORotation(323))
	stdColor(ctx,
		beat.SeqInterval(264, 286, 0.5),
		beat.RngStep(0, 1.3, 10),
		opt.SeqOrdinal(thesecond.SpotlightLeft, thesecond.SpotlightRight),
		opt.SeqOrdinal(evt.Blue, evt.Blue, evt.White, evt.White),
		evt.OBeatDistWave(2),
		fx.InstantTransit,
		fx.OBrightness(0, 1, 2.4, 0, ease.InCirc),
	)

	ctx.WSeq(beat.SeqInterval(264.7, 286.7, 0.5), func(ctx context.Context) {
		_, b := evt.RotationGroupWithBox(ctx,
			opt.SeqOrdinal(thesecond.SpotlightLeft, thesecond.SpotlightRight),
			evt.OBeatDistWave(1.8), evt.ODistWave(11),
		)
		ctx.WRng(beat.RngStep(0, 1.4, 20), func(ctx context.Context) {
			if ctx.SeqFirst() {
				b.AddEvent(ctx, evt.ORotation(325), evt.EasingNone)
			} else {
				b.AddEvent(ctx, fx.ORotation(0, 1.0, 323, 307, ease.OutBounce))
			}
		})
	})

	resetF := opt.Set("reset", false)
	fx.RotHold(ctx, 262-0.1, thesecond.TopLasersTop, thesecond.BottomLasersBottom)

	incantation2(ctx, 262, 0, 6, 1.8, 0, 45, 120, 2, 30,
		thesecond.TopLasersTop, opt.T(evt.Blue, evt.Red), inc2ColorRange(beat.RngStep(0, 2, 2)))
	gesture(ctx, beat.Seq(264), 70, 50, 11, 0.3,
		thesecond.BottomLasersLeftTop, opt.T(evt.Blue, evt.White, evt.White))
	gesture(ctx, beat.Seq(265), 70, 50, 11, 0.3,
		thesecond.BottomLasersRightTop, opt.T(evt.Red, evt.White, evt.White))

	gesture(ctx, beat.Seq(267), 230, 240, 11, 0.3,
		thesecond.BottomLasersRightBottom, evt.White)
	gesture(ctx, beat.Seq(268), 230, 240, 11, 0.3,
		thesecond.BottomLasersLeftBottom, evt.White)
	gesture(ctx, beat.Seq(269), 230, 240, 11, 0.3,
		thesecond.BottomLasersRightBottom, evt.Blue)

	incantation2(ctx, 270, 0, 6, 1.8, 0, 80, 90, 2, 15,
		thesecond.TopLasersBottom, opt.T(evt.Blue, evt.Red), inc2ColorRange(beat.RngStep(0, 1.4, 2)))

	gesture(ctx, beat.Seq(272), 95, 100, -13, 0.4,
		thesecond.TopLasersLeftTop, thesecond.BottomLasersLeftTop, opt.T(evt.Red, evt.White))
	gesture(ctx, beat.Seq(273), 85, 80, 13, 0.4,
		thesecond.TopLasersRightTop, thesecond.BottomLasersRightTop, opt.T(evt.Blue, evt.White))

	fx.RotHold(ctx, 275-0.1, thesecond.TopLasers)
	incantation2(ctx, 275, 0, 6, 1.8, 0, 70, 25, 2, 30,
		thesecond.TopLasersLeftBottom, opt.T(evt.Blue, evt.Red), inc2ColorRange(beat.RngStep(0, 3, 2)))
	incantation2(ctx, 276, 0, 6, 1.8, 0, 65, 10, 2, 20,
		thesecond.TopLasersRightBottom, opt.T(evt.Red, evt.White), inc2ColorRange(beat.RngStep(0, 3, 2)))
	incantation2(ctx, 277, 0, 6, 1.8, 0, 100, 140, 2, 10,
		thesecond.TopLasersTop, opt.T(evt.White, evt.Blue), inc2ColorRange(beat.RngStep(0, 3, 2)))

	fx.RotHold(ctx, 278-0.1, thesecond.BottomLasers)
	gesture(ctx, beat.Seq(278), 0, 45, 60, 2.0,
		thesecond.TopLasersBottom, evt.Blue, gestureEasing(ease.OutCirc))
	gesture(ctx, beat.Seq(278), 0, 45, -45, 1.4,
		thesecond.TopLasersTop, evt.Red, gestureEasing(ease.OutCirc))
	incantation2(ctx, 278, 0, 3, 1.8, 0, 45, 0, 2, 45,
		thesecond.BottomLasersTop, opt.T(evt.Red, evt.White), inc2ColorRange(beat.RngStep(0, 2, 2)))
	incantation2(ctx, 278, 0, 3, 1.8, 0, 225, 180, 2, 45,
		thesecond.BottomLasersBottom, opt.T(evt.White, evt.Blue), inc2ColorRange(beat.RngStep(0, 2, 2)))

	fx.RotHold(ctx, 281-0.1, thesecond.BottomLasers)
	gesture(ctx, beat.Seq(280), 280, 280, 80, 0.7,
		thesecond.TopLasersBottom, evt.Blue, resetF)
	gesture(ctx, beat.Seq(280), 53, 53, -35, 0.7,
		thesecond.TopLasersTop, evt.Red, resetF)
	gesture(ctx, beat.Seq(281), 33, 33, 35, 0.7,
		thesecond.BottomLasersTop, evt.White, resetF)
	gesture(ctx, beat.Seq(281), 30, 30, 35, 0.7,
		thesecond.BottomLasersBottom, evt.White, resetF)

	gesture(ctx, beat.Seq(283), 45, 45+22.5, 22.5, 0.4,
		thesecond.TopLasersLeftTop, opt.T(evt.Blue, evt.White, evt.White))
	gesture(ctx, beat.Seq(284), 112.5, 90, -22.5, 0.4,
		thesecond.TopLasersLeftBottom, opt.T(evt.White, evt.Red, evt.Red))
	gesture(ctx, beat.Seq(285), 270, 292.5, 22.5, 0.4,
		thesecond.TopLasersRightBottom, opt.T(evt.Red, evt.Blue, evt.Blue))
	gesture(ctx, beat.Seq(286), 247.5, 225, -22.5, 0.4,
		thesecond.TopLasersRightTop, opt.T(evt.Blue, evt.White, evt.White))

	fx.Off(ctx, 288-0.15, thesecond.BottomLasersRightTop)
	fx.RotReset(ctx, 288-0.1, thesecond.BottomLasers, evt.ORotation(70))
	incantation2(ctx, 288, 0, 6, 1.8, 0, 70, 30, 2, 45,
		thesecond.BottomLasersRightTop, opt.T(evt.White, evt.Blue), inc2ColorRange(beat.RngStep(0, 2, 2)))
	incantation2(ctx, 290, 0, 6, 1.8, 0, 70, 30, 2, 45,
		thesecond.BottomLasersLeftTop, opt.T(evt.White, evt.Blue), inc2ColorRange(beat.RngStep(0, 2, 2)))

	// 288, 290
	ctx.WSeq(beat.Seq(288, 290), func(ctx context.Context) {
		light := opt.SeqOrdinal(thesecond.SpotlightRight, thesecond.SpotlightLeft)(ctx)
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

	softWave(ctx, beat.Seq(288), thesecond.Runway, evt.Red)
}

func Chorus2(ctx context.Context) {
	//fx.RotReset(ctx, 292-0.1, thesecond.BigRing, evt.OStepAndOffsetFilter(0, 2), evt.ORotation(45), evt.ODistStep(90), evt.EasingLinear)
	//fx.RotReset(ctx, 317-0.1, thesecond.BigRing, evt.OStepAndOffsetFilter(0, 2), evt.ORotation(90), evt.ODistStep(23), evt.EasingLinear)
	//fx.RotReset(ctx, 292-0.1, thesecond.BigRing, evt.OStepAndOffsetFilter(1, 2), evt.ORotation(45), evt.ODistStep(90))
	//ctx.WSeq(beat.Seq(292), func(ctx context.Context) {
	//	g := evt.RotationGroup(ctx, thesecond.BigRing)
	//
	//	b1 := g.AddBox(ctx, evt.OStepAndOffsetFilter(0, 2), evt.ODistStep(23))
	//	b2 := g.AddBox(ctx, evt.OStepAndOffsetFilter(1, 2), evt.ODistStep(23))
	//	ctx.WRng(beat.Rng0To(25), func(ctx context.Context) {
	//		b1.AddEvent(ctx, opt.Ordinal(15, 45), fx.EaseNoneLinear)
	//		b2.AddEvent(ctx, opt.Ordinal(45, 15), fx.EaseNoneLinear)
	//	})
	//})

	ctx.WSeq(beat.Seq(292), func(ctx context.Context) {
		spotlightGhost(ctx, beat.Seq0,
			beat.RngStep(0, 1.2, 4),
			beat.RngStep(0, 1.2, 10),
			-80, 40, 0, 120,
			thesecond.SpotlightLeft, thesecond.SpotlightRight,
			fx.OBrightness(0.0, 1.0, 0, 2.4, ease.OutCirc),
			opt.T(evt.White, evt.Red),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(2.0)),
			evt.CW,
		)
		randFill(ctx, beat.Seq0,
			beat.RngStep(0, 0.8, 2),
			thesecond.Runway, thesecond.BigRing,
			fx.OBrightness(0, 1, 0, 2.4, ease.InOutCirc),
			opt.T(evt.White, evt.Red),
			evt.OBeatDistWave(1.7),
		)
		randFill(ctx, beat.Seq0,
			beat.RngStep(0, 0.8, 2),
			thesecond.SmallRing,
			fx.OBrightness(0, 1, 0, 1.0, ease.InOutCirc),
			opt.T(evt.Red, evt.White),
			evt.OBeatDistWave(1.7),
		)
	})

	fx.Off(ctx, 294, thesecond.Spotlight)

	//smolEscSeq := beat.Seq(294, 296, 298, 300)

	ctx.WSeq(beat.Seq(294, 296, 298, 300), func(ctx context.Context) {
		spotlightGhost(ctx, beat.Seq0,
			beat.RngStep(0, 1.4, 3),
			beat.RngStep(0, 1.4, 10),
			-20+15*ctx.SeqOrdinalF(), 20+15*ctx.SeqOrdinalF(), 15, -15-27*ctx.SeqOrdinalF(),
			thesecond.Spotlight,
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

		smolEscalation3(ctx, beat.Seq0,
			beat.RngStep(0, 1, 10),
			scale.FromUnitClamp(1.2, 1.8),
			scale.FromUnitClamp(15, 85),
			scale.FromUnitClamp(0, 90),
			scale.FromUnitClamp(60, 420),
			ease.InCirc,
		)

		ctx.WOpt(
			opt.IfSeqFirst[evt.ColorEventOption](
				fx.OBrightness(0.0, 1.0, 2.4, 0, ease.OutCirc),
				opt.Combine(
					fx.OBrightness(0.0, 0.5, 0, 3.2, ease.OutCirc),
					fx.OBrightness(0.5, 1.0, 3.2, 0, ease.OutCirc),
				),
			),
			opt.ColorBoxOnly(fx.InstantTransit),
		).Do(func(ctx context.Context) {
			stdColor(ctx, beat.Seq0,
				beat.RngStep(0, 1, 10),
				thesecond.SmallRing,
				opt.SeqOrdinal(
					opt.T(evt.Red, evt.White),
					opt.T(evt.White, evt.Blue),
					opt.T(evt.Blue, evt.Red),
				),
				evt.OBeatDistStep(0.04),
			)

			stdColor(ctx, beat.Seq0,
				beat.RngStep(0, 0.9, 10),
				thesecond.BigRing,
				opt.T(evt.White, evt.Red),
				evt.OBeatDistWave(2),
			)

			stdRotation(ctx, beat.Seq0,
				beat.Rng1(0),
				thesecond.BigRing,
				evt.OBeatDistWave(1.8),
				evt.ODistStep(-3*ctx.SeqOrdinalF()),
				evt.ODistAffectFirst(true),
				evt.ORotation(15*ctx.SeqOrdinalF()),
			)

			stdColor(ctx, beat.Seq0,
				beat.RngStep(0, 1, 10),
				thesecond.Runway,
				opt.SeqOrdinal(evt.Blue, evt.Red),
				evt.OBeatDistWave(1.4),
			)
		})
	})

	// kasmashii - 302, 303, 304, 305
	ctx.WSeq(beat.Seq(302, 303, 304, 305), func(ctx context.Context) {

		smolReduction2(ctx, beat.Seq0,
			beat.RngStep(0, 0.8, 10),
			scale.FromUnitClamp(0.6, 0.6),
			scale.FromUnitClamp(17, 111),
			scale.FromUnitClamp(420, 90),
			scale.FromUnitClamp(90, 420),
			ease.OutCirc,
			thesecond.BigRing,
		)

		ctx.WSeq(beat.Seq0, func(ctx context.Context) {
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

		stdColor(ctx, beat.Seq0,
			beat.RngStep(0, 1, 10),
			thesecond.BigRing,
			evt.OStepAndOffsetFilter(0, 1, evt.OReverse(true)),
			opt.T(evt.White, evt.Red, evt.Blue, evt.Blue),
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

		stdColor(ctx, beat.Seq0,
			beat.RngStep(0, 1, 10),
			thesecond.Runway,
			evt.OStepAndOffsetFilter(0, 1, evt.OReverse(true)),
			opt.SeqOrdinal(
				opt.SeqOrdinal(evt.White, evt.Red),
				opt.SeqOrdinal(evt.Blue, evt.White),
			),
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
	})

	horobiSeq := beat.Seq(307, 308, 309)
	stdColor(ctx,
		horobiSeq,
		beat.RngStep(0, 0.8, 3),
		thesecond.Runway,
		//evt.OBeatDistWave(1.0),
		evt.OSectionFilter(0, 0, evt.OReverse(true)),
		opt.Ordinal(evt.Blue, evt.White, evt.Blue),
		opt.Ordinal(evt.OBrightness(0), evt.OBrightness(0.2), evt.OBrightness(0.1)),
		fx.ExtendTransit,
	)

	incantation2(ctx, 310, 0, 5.5, 2.4, 1, 90, 180, 3, 45,
		thesecond.TopLasersRightTop, thesecond.BottomLasersRightBottom, opt.T(evt.Red, evt.Blue))
	incantation2(ctx, 311, 0, 6.5, 2.4, 1, 90, 180, 3, 45,
		thesecond.TopLasersLeftTop, thesecond.BottomLasersLeftTop, opt.T(evt.White, evt.Red))
	incantation2(ctx, 312, 0, 4.5, 2.4, 1, 135, 270, 3, 45,
		thesecond.TopLasersLeftBottom, thesecond.BottomLasersLeftBottom, opt.T(evt.White, evt.Blue))
	incantation2(ctx, 313, 0, 3.5, 2.4, 1, 135, 270, 3, 45,
		thesecond.TopLasersRightBottom, thesecond.BottomLasersRightBottom, opt.T(evt.Blue, evt.Red))
	incantation2(ctx, 314, 0, 2.5, 2.4, 1, 180, 260, 1, 30,
		thesecond.TopLasers, opt.T(evt.Red, evt.White))
	incantation2(ctx, 315, 0, 1.5, 2.4, 1, 180, 260, 1, 30,
		thesecond.BottomLasers, opt.T(evt.White, evt.Blue))
	incantation2(ctx, 316, 0, 0.5, 2.4, 1, 180, 270, 0.4, 90,
		thesecond.Lasers, opt.T(evt.White, evt.Red))

	fx.RotHold(ctx, 317-0.1, thesecond.Lasers)
	//cRng := beat.RngStep(0, 1.2, 2)
	//incantation(ctx,
	//	beat.Seq(310),
	//	cRng,
	//	beat.RngStep(0, 6.5, 30),
	//	opt.Ordinal(evt.Red, evt.White),
	//	opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
	//	thesecond.TopLasersLeftTop,
	//	opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
	//	fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	//)
	//incantation(ctx,
	//	beat.Seq(311),
	//	cRng,
	//	beat.RngStep(0, 5.5, 30),
	//	opt.Ordinal(evt.Red, evt.Blue),
	//	opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
	//	thesecond.TopLasersRightTop,
	//	opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
	//	fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	//)
	//incantation(ctx,
	//	beat.Seq(312),
	//	cRng,
	//	beat.RngStep(0, 4.5, 30),
	//	opt.Ordinal(evt.Red, evt.White),
	//	opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
	//	thesecond.TopLasersLeftBottom,
	//	opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(36)),
	//	fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	//)
	//incantation(ctx,
	//	beat.Seq(313),
	//	cRng,
	//	beat.RngStep(0, 3.5, 30),
	//	opt.Ordinal(evt.Blue, evt.Red),
	//	opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
	//	thesecond.TopLasersRightBottom,
	//	opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(36)),
	//	fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	//)
	//incantation(ctx,
	//	beat.Seq(314),
	//	cRng,
	//	beat.RngStep(0, 2.5, 30),
	//	opt.Ordinal(evt.Red, evt.White),
	//	opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
	//	thesecond.BottomLasersLeftBottom,
	//	opt.RotationBoxOnly(evt.OBeatDistWave(1), evt.ODistWave(121)),
	//	fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	//)
	//incantation(ctx,
	//	beat.Seq(315),
	//	cRng,
	//	beat.RngStep(0, 1.5, 30),
	//	opt.Ordinal(evt.Red, evt.Blue),
	//	opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
	//	thesecond.BottomLasersRightBottom,
	//	opt.RotationBoxOnly(evt.OBeatDistWave(1), evt.ODistWave(121)),
	//	fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	//)
	//incantation(ctx,
	//	beat.Seq(316),
	//	cRng,
	//	beat.RngStep(0, 0.5, 30),
	//	opt.Ordinal(evt.White, evt.Red),
	//	opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
	//	thesecond.BottomLasersTop,
	//	opt.RotationBoxOnly(evt.OBeatDistWave(0.4), evt.ODistWave(36)),
	//	fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	//)

	ctx.WSeq(beat.Seq(317), func(ctx context.Context) {
		_, cb := evt.ColorGroupWithBox(ctx, thesecond.TopLasers, thesecond.BottomLasers, thesecond.SmallRing)
		//_, rcb := evt.ColorGroupWithBox(ctx, thesecond.SmallRing, evt.OBeatDistWave(0.6))
		ctx.WRng(beat.RngStep(0, 0.4, 2), func(ctx context.Context) {
			cb.AddEvent(ctx,
				opt.Ordinal(evt.Red, evt.White),
				opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
				fx.InstantTransit,
			)
			//rcb.AddEvent(ctx,
			//	opt.Ordinal(evt.Red, evt.White),
			//	opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
			//	fx.InstantTransit,
			//)
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

	fx.Off(ctx, 317, thesecond.Runway)
}

func Bridge1(ctx context.Context) {
	// 318
	starfield(ctx,
		beat.SeqInterval(318, 382, 0.125),
		beat.RngStep(0.0, 1.5, 10),
		beat.RngStep(0.3, 1.8, 10),
		0.6,
		evt.OBeatDistWave(8),
	)
	starfieldSpin(ctx, 318, 382)

	// 320, 322, 324
	// 326, 328, 330, 332
	smolEscalation(ctx,
		beat.Seq(318, 321, 324, 325),
		beat.RngStep(0, 0.8, 10),
	)
	stdColor(ctx,
		beat.Seq(318, 321),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Blue),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	stdColor(ctx,
		beat.Seq(324),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OStepAndOffsetFilter(0, 2),
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Blue, evt.White, evt.Blue),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	stdColor(ctx,
		beat.Seq(325),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OStepAndOffsetFilter(0, 2, evt.OReverse(true)),
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(0.8), evt.OBrightness(0)),
	)

	incantation(ctx,
		beat.Seq(323),
		beat.RngStep(0, 8, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Blue, evt.White),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.TopLasersLeftTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(323.5),
		beat.RngStep(0, 8, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.TopLasersRightTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	ctx.WSeq(beat.Seq(323, 323.5), func(ctx context.Context) {
		light := opt.SeqOrdinal(thesecond.SpotlightLeft, thesecond.SpotlightRight)(ctx)
		spotlightGhost(ctx,
			beat.Seq(0),
			beat.RngStep(0, 1.6, 4),
			beat.RngStep(0, 2.4, 10),
			20, -20, 0, -45,
			light,
			fx.OBrightness(0.0, 0.5, 0, 2.4, ease.InOutCirc),
			fx.OBrightness(0.5, 1.0, 2.4, 0, ease.InOutCirc),
			opt.Ordinal(evt.Blue, evt.Blue, evt.White, evt.White),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
			evt.CCW,
		)
	})

	// seida chishki
	smolSimple(ctx,
		beat.Seq(326),
		beat.RngStep(0, 0.8, 10),
		evt.OBeatDistWave(1.1),
		evt.ODistWave(-72),
		fx.ORotation(0, 1, 0, 315, ease.OutCirc),
	)
	stdColor(ctx,
		beat.Seq(326),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.2), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	smolSimple(ctx,
		beat.Seq(328),
		beat.RngStep(0, 0.8, 10),
		evt.OBeatDistWave(1.1),
		evt.ODistWave(72),
		fx.ORotation(0, 1, 0, 45, ease.OutCirc),
	)
	stdColor(ctx,
		beat.Seq(328),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.2), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	smolSimple(ctx,
		beat.Seq(330),
		beat.RngStep(0, 0.8, 10),
		evt.OBeatDistWave(1.1),
		evt.ODistWave(-72),
		fx.ORotation(0, 1, 0, 315, ease.OutCirc),
	)
	stdColor(ctx,
		beat.Seq(330),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.2), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	smolSimple(ctx,
		beat.Seq(332),
		beat.RngStep(0, 0.8, 10),
		evt.OBeatDistWave(1.1),
		evt.ODistWave(72),
		fx.ORotation(0, 1, 0, 45, ease.OutCirc),
	)
	stdColor(ctx,
		beat.Seq(332),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.2), evt.OBrightness(0.8), evt.OBrightness(0)),
	)

	incantation(ctx,
		beat.Seq(326),
		beat.RngStep(0, 3, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Blue, evt.White),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersLeftTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(328),
		beat.RngStep(0, 3, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersRightTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(330),
		beat.RngStep(0, 3, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersLeftBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(332),
		beat.RngStep(0, 1.8, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersRightBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 91, 32, ease.OutCirc),
	)

	// 334
	sparkleFadeReset(ctx, 334-0.1)
	sparkleFade(ctx,
		beat.Seq(334),
		beat.RngStep(0, 6, 30),
		thesecond.Lasers,
		evt.OBeatDistStep(1),
		opt.T(evt.Blue, evt.White, evt.Red, evt.Blue, evt.White, evt.Blue),
		darkPeak(0.2, 0.5, 0.8, ease.OutCirc, ease.InCirc),
	)
	sparkleFadeMotion(ctx,
		beat.Seq(334),
		beat.RngStep(0, 6, 30),
		thesecond.Lasers,
	)

	// 341, 342
	ctx.WSeq(beat.Seq(341, 342), func(ctx context.Context) {
		light := opt.SeqOrdinal(thesecond.SpotlightLeft, thesecond.SpotlightRight)(ctx)
		spotlightGhost(ctx,
			beat.Seq(0),
			beat.RngStep(0, 1.6, 4),
			beat.RngStep(0, 2.4, 10),
			20, -20, 0, -45,
			light,
			fx.OBrightness(0.0, 0.5, 0, 1.4, ease.InOutCirc),
			fx.OBrightness(0.5, 1.0, 1.4, 0, ease.InOutCirc),
			opt.Ordinal(evt.Blue, evt.Blue, evt.White, evt.White),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
			evt.CCW,
		)
	})

	// piano 342
	pianoSparkleFadeOpts := opt.Combine(
		evt.OBeatDistStep(0.75),
		darkPeak(0.1, 0.5, 0.8, ease.OutCirc, ease.InCirc),
	)
	rotHold(ctx, 342-0.15, thesecond.TopLasers)
	sparkleFade(ctx,
		beat.Seq(342),
		beat.RngStep(0, 4, 30),
		thesecond.TopLasersLeft, evt.Blue,
		pianoSparkleFadeOpts,
	)
	sparkleFade(ctx,
		beat.Seq(342),
		beat.RngStep(0, 4, 30),
		thesecond.TopLasersRight, evt.White,
		pianoSparkleFadeOpts,
	)
	clock(ctx, 342, 1, 0.6, 90, 100, 5, evt.CW,
		thesecond.TopLasers)
	rotHold(ctx, 346-0.15, thesecond.BottomLasers)
	sparkleFade(ctx,
		beat.Seq(346),
		beat.RngStep(0, 4, 30),
		thesecond.BottomLasersLeft, evt.White,
		pianoSparkleFadeOpts,
	)
	sparkleFade(ctx,
		beat.Seq(346),
		beat.RngStep(0, 4, 30),
		thesecond.BottomLasersRight, evt.Blue,
		pianoSparkleFadeOpts,
	)
	clock(ctx, 346, 1, 0.6, 90, 100, 5, evt.CW,
		thesecond.BottomLasers)

	smolEscalation(ctx,
		beat.Seq(350, 352, 354, 356),
		beat.RngStep(0, 0.8, 10),
	)
	stdColor(ctx,
		beat.Seq(350, 352),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Blue),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	stdColor(ctx,
		beat.Seq(354),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OStepAndOffsetFilter(0, 2),
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Blue, evt.White, evt.Blue),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	stdColor(ctx,
		beat.Seq(356),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OStepAndOffsetFilter(0, 2, evt.OReverse(true)),
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(0.8), evt.OBrightness(0)),
	)

	// issaigasai
	smolSimple(ctx,
		beat.Seq(358),
		beat.RngStep(0, 0.8, 10),
		evt.OBeatDistWave(1.1),
		evt.ODistWave(-72),
		fx.ORotation(0, 1, 0, 315, ease.OutCirc),
	)
	stdColor(ctx,
		beat.Seq(358),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.2), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	smolSimple(ctx,
		beat.Seq(360),
		beat.RngStep(0, 0.8, 10),
		evt.OBeatDistWave(1.1),
		evt.ODistWave(72),
		fx.ORotation(0, 1, 0, 45, ease.OutCirc),
	)
	stdColor(ctx,
		beat.Seq(360),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.2), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	smolSimple(ctx,
		beat.Seq(362),
		beat.RngStep(0, 0.8, 10),
		evt.OBeatDistWave(1.1),
		evt.ODistWave(-72),
		fx.ORotation(0, 1, 0, 315, ease.OutCirc),
	)
	stdColor(ctx,
		beat.Seq(362),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.2), evt.OBrightness(0.8), evt.OBrightness(0)),
	)
	smolSimple(ctx,
		beat.Seq(364),
		beat.RngStep(0, 0.8, 10),
		evt.OBeatDistWave(1.1),
		evt.ODistWave(72),
		fx.ORotation(0, 1, 0, 45, ease.OutCirc),
	)
	stdColor(ctx,
		beat.Seq(364),
		beat.RngStep(0, 1, 3),
		thesecond.SmallRing,
		evt.OBeatDistWave(1.75),
		opt.Ordinal(evt.Red, evt.White, evt.Red),
		fx.InstantTransit,
		opt.Ordinal(evt.OBrightness(1.2), evt.OBrightness(0.8), evt.OBrightness(0)),
	)

	incantation(ctx,
		beat.Seq(358),
		beat.RngStep(0, 3, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Blue, evt.White),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersLeftTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(360),
		beat.RngStep(0, 3, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersRightTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(362),
		beat.RngStep(0, 3, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersLeftBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(364),
		beat.RngStep(0, 1.8, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersRightBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 91, 32, ease.OutCirc),
	)

	// 334
	sparkleFadeReset(ctx, 366-0.1)
	sparkleFade(ctx,
		beat.Seq(366),
		beat.RngStep(0, 6, 30),
		thesecond.Lasers,
		evt.OBeatDistStep(1),
		opt.T(evt.Blue, evt.White, evt.Red, evt.Blue, evt.White, evt.Blue),
		darkPeak(0.2, 0.5, 0.8, ease.OutCirc, ease.InCirc),
	)
	sparkleFadeMotion(ctx,
		beat.Seq(366),
		beat.RngStep(0, 6, 30),
		thesecond.Lasers,
	)

	// 374 ga
	incantation(ctx,
		beat.Seq(374),
		beat.RngStep(0, 8, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Blue, evt.White),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(377),
		beat.RngStep(0, 8, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.TopLasersLeftTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(377.5),
		beat.RngStep(0, 8, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.TopLasersRightTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(380),
		beat.RngStep(0, 8, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.BottomLasersBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(381),
		beat.RngStep(0, 8, 2),
		beat.RngStep(0, 8.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
		thesecond.TopLasersBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(51)),
		fx.ORotation(0, 1, 31, 92, ease.OutCirc),
	)
}

func Bridge2(ctx context.Context) {
	starfield(ctx,
		beat.SeqInterval(382, 446, 0.125),
		beat.RngStep(0.0, 1.5, 10),
		beat.RngStep(0.3, 1.8, 10),
		1.0,
		evt.OBeatDistWave(8),
	)

	victorySpin(ctx, 382)

	pianoSparkleFadeOpts := opt.Combine(
		evt.OBeatDistStep(0.75),
		darkPeak(0.1, 0.5, 0.8, ease.OutCirc, ease.InCirc),
	)
	rotHold(ctx, 390-0.15, thesecond.TopLasers)
	sparkleFade(ctx,
		beat.Seq(390),
		beat.RngStep(0, 4, 30),
		thesecond.TopLasersLeft, evt.Blue,
		pianoSparkleFadeOpts,
	)
	sparkleFade(ctx,
		beat.Seq(390),
		beat.RngStep(0, 4, 30),
		thesecond.TopLasersRight, evt.White,
		pianoSparkleFadeOpts,
	)
	clock(ctx, 390, 1, 0.6, 90, 100, 5, evt.CW,
		thesecond.TopLasers)
	rotHold(ctx, 394-0.15, thesecond.BottomLasers)
	sparkleFade(ctx,
		beat.Seq(394),
		beat.RngStep(0, 4, 30),
		thesecond.BottomLasersLeft, evt.White,
		pianoSparkleFadeOpts,
	)
	sparkleFade(ctx,
		beat.Seq(394),
		beat.RngStep(0, 4, 30),
		thesecond.BottomLasersRight, evt.Blue,
		pianoSparkleFadeOpts,
	)
	clock(ctx, 394, 1, 0.6, 90, 100, 5, evt.CW,
		thesecond.BottomLasers)

	rotHold(ctx, 398-0.15, thesecond.TopLasers)
	sparkleFade(ctx,
		beat.Seq(398),
		beat.RngStep(0, 4, 30),
		thesecond.TopLasersLeft, evt.Blue,
		pianoSparkleFadeOpts,
	)
	sparkleFade(ctx,
		beat.Seq(398),
		beat.RngStep(0, 4, 30),
		thesecond.TopLasersRight, evt.White,
		pianoSparkleFadeOpts,
	)
	clock(ctx, 398, 1, 0.6, 180, 190, 5, evt.CW,
		thesecond.TopLasers)

	rotHold(ctx, 402-0.15, thesecond.BottomLasers)
	sparkleFade(ctx,
		beat.Seq(402),
		beat.RngStep(0, 4, 30),
		thesecond.BottomLasersLeft, evt.White,
		pianoSparkleFadeOpts,
	)
	sparkleFade(ctx,
		beat.Seq(402),
		beat.RngStep(0, 4, 30),
		thesecond.BottomLasersRight, evt.Blue,
		pianoSparkleFadeOpts,
	)
	clock(ctx, 402, 1, 0.6, 180, 190, 5, evt.CW,
		thesecond.BottomLasers)

	ctx.WSeq(beat.Seq(390, 394, 396, 398, 402, 404, 406, 410), func(ctx context.Context) {
		spotlightGhost(ctx,
			beat.Seq(0),
			beat.RngStep(0, 1.6, 4),
			beat.RngStep(0, 2.4, 10),
			20, -20, 0, -45,
			thesecond.Spotlight,
			fx.OBrightness(0.0, 0.5, 0, 2.4, ease.InOutCirc),
			fx.OBrightness(0.5, 1.0, 2.4, 0, ease.InOutCirc),
			opt.Ordinal(evt.Red, evt.Red, evt.White, evt.White),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
			evt.CCW,
		)
	})
	victorySpin(ctx, 414)

	ctx.WSeq(beat.Seq(414, 418, 419, 420, 421, 422, 426, 427, 428, 430, 434, 436, 438, 442), func(ctx context.Context) {
		spotlightGhost(ctx,
			beat.Seq(0),
			beat.RngStep(0, 1.6, 4),
			beat.RngStep(0, 2.4, 10),
			20, -20, 0, -45,
			thesecond.Spotlight,
			fx.OBrightness(0.0, 0.5, 0, 2.4, ease.InOutCirc),
			fx.OBrightness(0.5, 1.0, 2.4, 0, ease.InOutCirc),
			opt.Ordinal(evt.Red, evt.Red, evt.White, evt.White),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
			evt.CCW,
		)
	})
}

func SilentBridge(ctx context.Context) {
	sukitoSeq := beat.Seq(443, 444, 445, 446, 450, 452, 454, 456, 458, 460, 462, 464, 466, 468, 470)

	stdColor(ctx,
		sukitoSeq,
		beat.RngStep(0, 2, 20),
		opt.SeqOrdinal(thesecond.TopLasersLeft, thesecond.BottomLasersRightBottom, thesecond.TopLasersLeftBottom, thesecond.BottomLasersRightBottom),
		evt.White, fx.InstantTransit,
		evt.OBeatDistWave(1.5),
		fx.OBrightness(0, 2, 1.1, 0, ease.OutCirc),
	)

	ctx.WSeq(beat.Seq(443), func(ctx context.Context) {
		_, b := evt.RotationGroupWithBox(ctx,
			thesecond.TopLasers, thesecond.BottomLasers,
			evt.ODistWave(17), evt.OBeatDistWave(2.4))
		ctx.WRng(beat.RngStep(0, 32, 20), func(ctx context.Context) {
			b.AddEvent(ctx, fx.ORotation(0, 1, 17, 180, ease.OutCirc))
		})
	})
}

func Chorus3(ctx context.Context) {
	spotlightGhost(ctx,
		beat.Seq(476),
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

	ctx.WSeq(beat.Seq(476), func(ctx context.Context) {
		stdColor(ctx, beat.Seq(0), beat.RngStep(0, 1, 2),
			thesecond.Runway, thesecond.BigRing,
			opt.Ordinal(evt.White, evt.Red),
			opt.Ordinal(evt.OBrightness(0), evt.OBrightness(2.8)),
			fx.InstantTransit,
			evt.OBeatDistStep(0.08),
			evt.OStepAndOffsetFilter(0, 1, evt.OOrder(evt.FilterOrderRandom, 123)),
		)
	})
	ctx.WSeq(beat.Seq(476), func(ctx context.Context) {
		stdColor(ctx, beat.Seq(0), beat.RngStep(0, 1, 2),
			thesecond.SmallRing,
			opt.Ordinal(evt.Red, evt.White),
			opt.Ordinal(evt.OBrightness(0), evt.OBrightness(1.0)),
			fx.InstantTransit,
			evt.OBeatDistStep(0.04),
			evt.OStepAndOffsetFilter(0, 1, evt.OOrder(evt.FilterOrderRandom, 123)),
		)
	})

	fx.Off(ctx, 478, thesecond.Spotlight)

	smolEscSeq := beat.Seq(478, 480, 482, 484)

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
	fusawashiSeq := beat.Seq(486, 487, 488, 489)

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
	horobiSeq := beat.Seq(491, 492, 493)
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
		beat.Seq(494),
		cRng,
		beat.RngStep(0, 6.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersLeftTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
		fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(495),
		cRng,
		beat.RngStep(0, 5.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersRightTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
		fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(496),
		cRng,
		beat.RngStep(0, 4.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersLeftBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(497),
		cRng,
		beat.RngStep(0, 3.5, 30),
		opt.Ordinal(evt.Blue, evt.Red),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersRightBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(498),
		cRng,
		beat.RngStep(0, 2.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.BottomLasersLeftBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(1), evt.ODistWave(121)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(499),
		cRng,
		beat.RngStep(0, 1.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.BottomLasersRightBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(1), evt.ODistWave(121)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(500),
		cRng,
		beat.RngStep(0, 0.5, 30),
		opt.Ordinal(evt.White, evt.Red),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.BottomLasersTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(0.4), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)

	ctx.WSeq(beat.Seq(501), func(ctx context.Context) {
		_, cb := evt.ColorGroupWithBox(ctx, thesecond.TopLasers, thesecond.BottomLasers)
		//_, rcb := evt.ColorGroupWithBox(ctx, thesecond.SmallRing, evt.OBeatDistWave(0.6))
		ctx.WRng(beat.RngStep(0, 0.4, 2), func(ctx context.Context) {
			cb.AddEvent(ctx,
				opt.Ordinal(evt.Red, evt.White),
				opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
				fx.InstantTransit,
			)
			//rcb.AddEvent(ctx,
			//	opt.Ordinal(evt.Red, evt.White),
			//	opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
			//	fx.InstantTransit,
			//)
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

	fx.Off(ctx, 501, thesecond.Runway)
}

func Chorus4(ctx context.Context) {
	//jinseiSeq := beat.Seq(504, 506, 508)
	// 318
	starfield(ctx,
		beat.SeqInterval(510, 510+16, 0.25),
		beat.RngStep(0.0, 1.5, 10),
		beat.RngStep(0.3, 1.8, 10),
		0.6,
		evt.OBeatDistWave(8),
	)
	starfieldSpin(ctx, 510, 510+16)

	ctx.WSeq(beat.Seq(504, 505, 506), func(ctx context.Context) {
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
	cRng := beat.RngStep(0, 1.2, 2)
	incantation(ctx,
		beat.Seq(504),
		cRng,
		beat.RngStep(0, 6.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersLeftTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
		fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(506),
		cRng,
		beat.RngStep(0, 5.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersRightTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
		fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(508),
		cRng,
		beat.RngStep(0, 4.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersLeftBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)

	escSeq := beat.Seq(510, 518, 526, 528)
	smolEscalation(ctx,
		escSeq,
		beat.RngStep(0, 1, 10),
	)
	// 534, 536, 538, 540
	stdColor(ctx,
		escSeq,
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

	ctx.WSeq(beat.Seq(510, 518), func(ctx context.Context) {
		spotlightGhost(ctx,
			beat.Seq(0),
			beat.RngStep(0, 1.6, 4),
			beat.RngStep(0, 2.4, 10),
			20, -20, 0, -45,
			thesecond.Spotlight,
			fx.OBrightness(0.0, 0.5, 0, 2.4, ease.InOutCirc),
			fx.OBrightness(0.5, 1.0, 2.4, 0, ease.InOutCirc),
			opt.Ordinal(evt.Red, evt.Red, evt.White, evt.White),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
			evt.CCW,
		)
	})

	ctx.WSeq(beat.Seq(534, 536, 538, 540), func(ctx context.Context) {
		spotlightGhost(ctx,
			beat.Seq(0),
			beat.RngStep(0, 1.6, 4),
			beat.RngStep(0, 2.4, 10),
			20, -20, 0, -45,
			thesecond.Spotlight,
			fx.OBrightness(0.0, 0.5, 0, 2.4, ease.InOutCirc),
			fx.OBrightness(0.5, 1.0, 2.4, 0, ease.InOutCirc),
			opt.Ordinal(evt.Red, evt.Red, evt.White, evt.White),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
			evt.CCW,
		)
	})

	ctx.WSeq(beat.Seq(542, 544, 546, 548), func(ctx context.Context) {
		spotlightGhost(ctx,
			beat.Seq(0),
			beat.RngStep(0, 1.6, 4),
			beat.RngStep(0, 2.4, 10),
			20, -20, 0, -45,
			thesecond.Spotlight,
			fx.OBrightness(0.0, 0.5, 0, 2.4, ease.InOutCirc),
			fx.OBrightness(0.5, 1.0, 2.4, 0, ease.InOutCirc),
			opt.Ordinal(evt.Red, evt.Red, evt.White, evt.White),
			opt.ColorBoxOnly(evt.OBeatDistWave(1.9)),
			opt.RotationBoxOnly(evt.OBeatDistWave(1.4)),
			evt.CCW,
		)
	})

	incantation(ctx,
		beat.Seq(558),
		cRng,
		beat.RngStep(0, 6.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersLeftTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
		fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(559),
		cRng,
		beat.RngStep(0, 5.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersRightTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(64)),
		fx.ORotation(0, 1, 0, 31, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(560),
		cRng,
		beat.RngStep(0, 4.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersLeftBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(561),
		cRng,
		beat.RngStep(0, 3.5, 30),
		opt.Ordinal(evt.Blue, evt.Red),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.TopLasersRightBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(3), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(562),
		cRng,
		beat.RngStep(0, 2.5, 30),
		opt.Ordinal(evt.Red, evt.White),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.BottomLasersLeftBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(1), evt.ODistWave(121)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(563),
		cRng,
		beat.RngStep(0, 1.5, 30),
		opt.Ordinal(evt.Red, evt.Blue),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.BottomLasersRightBottom,
		opt.RotationBoxOnly(evt.OBeatDistWave(1), evt.ODistWave(121)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)
	incantation(ctx,
		beat.Seq(564),
		cRng,
		beat.RngStep(0, 0.5, 30),
		opt.Ordinal(evt.White, evt.Red),
		opt.Ordinal(evt.OBrightness(2.4), evt.OBrightness(1)),
		thesecond.BottomLasersTop,
		opt.RotationBoxOnly(evt.OBeatDistWave(0.4), evt.ODistWave(36)),
		fx.ORotation(0, 1, 180, 210, ease.OutCirc),
	)

	ctx.WSeq(beat.Seq(565), func(ctx context.Context) {
		_, cb := evt.ColorGroupWithBox(ctx, thesecond.TopLasers, thesecond.BottomLasers)
		_, rcb := evt.ColorGroupWithBox(ctx, thesecond.SmallRing, evt.OBeatDistWave(0.6))
		ctx.WRng(beat.RngStep(0, 0.4, 2), func(ctx context.Context) {
			cb.AddEvent(ctx,
				opt.Ordinal(evt.Red, evt.White),
				opt.Ordinal(evt.OBrightness(3.6), evt.OBrightness(0)),
				fx.InstantTransit,
			)
			rcb.AddEvent(ctx,
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

	fx.Off(ctx, 565, thesecond.Runway)
}

func Outro(ctx context.Context) {
	pianoSeq := beat.Seq(574, 578, 581)

	stdColor(ctx,
		pianoSeq,
		beat.RngStep(0, 2, 20),
		opt.SeqOrdinal(thesecond.TopLasersLeft, thesecond.BottomLasersRightBottom, thesecond.TopLasersLeftBottom, thesecond.BottomLasersRightBottom),
		evt.White, fx.InstantTransit,
		evt.OBeatDistWave(1.5),
		fx.OBrightness(0, 2, 1.1, 0, ease.OutCirc),
	)

	ctx.WSeq(beat.Seq(574), func(ctx context.Context) {
		_, b := evt.RotationGroupWithBox(ctx,
			thesecond.TopLasers, thesecond.BottomLasers,
			evt.ODistWave(17), evt.OBeatDistWave(2.4))
		ctx.WRng(beat.RngStep(0, 32, 20), func(ctx context.Context) {
			b.AddEvent(ctx, fx.ORotation(0, 1, 17, 180, ease.OutCirc))
		})
	})
	// 585

	stdColor(ctx,
		beat.Seq(585),
		beat.RngStep(0, 2, 20),
		thesecond.Lasers,
		evt.White, fx.InstantTransit,
		evt.OBeatDistWave(1.5),
		fx.OBrightness(0, 2, 1.1, 0, ease.OutCirc),
	)
}
