package evt

import (
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/scale"
)

type RotationEventFuncOpt func(ctx context.Context, e *RotationEvent)

func (o RotationEventFuncOpt) ApplyRotationEvent(e *RotationEvent) { panic("upgrade me") }

func (o RotationEventFuncOpt) ApplyRotationEventContext(ctx context.Context, e *RotationEvent) {
	o(ctx, e)
}

func ORotate(r1, r2 float64, dir RotationDirection) RotationEventFuncOpt {
	rotScaler := scale.FromUnitClamp(r1, r2)
	return func(ctx context.Context, e *RotationEvent) {
		if ctx.First() {
			e.TransitionType = RotationTransitionTypeTransition
			e.Easing = EasingNone
		} else {
			e.TransitionType = RotationTransitionTypeTransition
			e.Easing = EasingLinear
			e.Direction = dir
		}
		e.Rotation = rotScaler(ctx.T())
	}
}
