package fx

import (
	"math"

	"github.com/shasderias/iris/beat"
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/ease"
	"github.com/shasderias/iris/evt"
	"github.com/shasderias/iris/opt"
	"github.com/shasderias/iris/scale"
)

var (
	ExtendTransit  = opt.IfFirst(evt.Extend, evt.Transition)
	InstantTransit = opt.IfFirst(evt.Instant, evt.Transition)
	EaseNoneLinear = opt.IfFirst(evt.EasingNone, evt.EasingLinear)
)

type Point struct {
	X, Y float64
	Ease ease.Ing
}

type Curve struct {
	Points []Point
}

func (c Curve) Lerp(x float64) float64 {
	points := c.Points
	for i, pt := range points {
		if x >= pt.X && x <= points[i+1].X {
			x1 := scale.ToUnitClamp(pt.X, points[i+1].X)(x)

			var easeFn ease.Func

			if pt.Ease.Name == "" {
				easeFn = ease.Linear.Func
			} else {
				easeFn = pt.Ease.Func
			}

			x2 := easeFn(x1)

			return scale.FromUnitClamp(pt.Y, points[i+1].Y)(x2)
		}
	}
	panic("unreachable code")
}

func NewCurve(values ...float64) Curve {
	points := []Point{}
	for i, val := range values {
		points = append(points, Point{
			X:    float64(i) / float64(len(values)-1),
			Y:    val,
			Ease: ease.Linear,
		})
	}
	return Curve{points}
}

func RotHold(ctx context.Context, b float64, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(beat.Seq(b), func(ctx context.Context) {
			_, b := evt.RotationGroupWithBox(ctx)
			ctx.WRng(beat.Rng1(0), func(ctx context.Context) {
				b.AddEvent(ctx, evt.Extend)
			})
		})
	})
}

func RotReset(ctx context.Context, b float64, options ...any) {
	ctx.WOpt(options...).Do(func(ctx context.Context) {
		ctx.WSeq(beat.Seq(b), func(ctx context.Context) {
			_, b := evt.RotationGroupWithBox(ctx)
			ctx.WRng(beat.Rng1(0), func(ctx context.Context) {
				b.AddEvent(ctx, evt.EasingNone)
			})
		})
	})
}

func Off(ctx context.Context, b float64, options ...any) {
	ctx.WSeq(beat.Seq(b), func(ctx context.Context) {
		_, box := evt.ColorGroupWithBox(ctx, options...)

		ctx.WRng(beat.RngStep(0, 1, 1), func(ctx context.Context) {
			box.AddEvent(ctx, evt.OBrightness(0), evt.Instant)
		})
	})
}

type OptBrightness struct {
	st, et, sv, ev float64
	easing         ease.Ing
}

func (o OptBrightness) ApplyColorEvent(e *evt.ColorEvent) { panic("upgrade me") }
func (o OptBrightness) ApplyColorEventContext(ctx context.Context, e *evt.ColorEvent) {
	if ctx.T() < o.st || ctx.T() > o.et {
		return
	}

	t := scale.ToUnitClamp(o.st, o.et)(ctx.T())

	e.Brightness = scale.FromUnitClamp(o.sv, o.ev)(o.easing.Func(t))
}

var OBri = OBrightness

func OBrightness(st, et, sv, ev float64, easing ease.Ing) evt.ColorEventOption {
	return OptBrightness{st, et, sv, ev, easing}
}

type OptRotation struct {
	st, et, sr, er float64
	easing         ease.Ing
}

func (o OptRotation) ApplyRotationEvent(e *evt.RotationEvent) { panic("upgrade me") }
func (o OptRotation) ApplyRotationEventContext(ctx context.Context, e *evt.RotationEvent) {
	if ctx.T() < o.st || ctx.T() > o.et {
		return
	}

	t := scale.ToUnitClamp(o.st, o.et)(ctx.T())

	e.Rotation = math.Mod(scale.FromUnitClamp(o.sr, o.er)(o.easing.Func(t)), 360)
}

func ORotation(st, et, sr, er float64, easing ease.Ing) evt.RotationEventOption {
	return OptRotation{st, et, sr, er, easing}
}
