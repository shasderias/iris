package evt

import (
	"github.com/shasderias/iris/context"
)

type RotationEventGroup struct {
	Beat  float64
	Group []int
	Boxes []*RotationEventBox
}

type RotationEventGroupOption interface {
	ApplyRotationEventGroup(*RotationEventGroup)
}
type RotationEventGroupContextOption interface {
	ApplyRotationEventGroupContext(ctx context.Context, e *RotationEventGroup)
}

func RotationGroup(ctx context.Context, options ...RotationEventGroupOption) *RotationEventGroup {
	g := &RotationEventGroup{Beat: ctx.BOffset()}

	ApplyOptions(ctx, g, getOptions(ctx, options)...)

	ctx.AddEvent(g)
	return g
}

func RotationGroupWithBox(ctx context.Context, options ...any) (*RotationEventGroup, *RotationEventBox) {
	g := RotationGroup(ctx, optionsOf[RotationEventGroupOption](options...)...)
	b := g.AddBox(ctx, optionsOf[RotationEventBoxOption](options...)...)
	return g, b
}

type RotationEventBoxOption interface {
	ApplyRotationEventBox(*RotationEventBox)
}
type RotationEventBoxContextOption interface {
	ApplyRotationEventBoxContext(ctx context.Context, b *RotationEventBox)
}

func (g *RotationEventGroup) AddBox(ctx context.Context, options ...RotationEventBoxOption) *RotationEventBox {
	b := &RotationEventBox{
		IndexFilter: IndexFilter{
			IndexFilterType: IndexFilterTypeSections,
		},
		BeatDistribution: BeatDistribution{
			Type: DistributionTypeWave,
		},
		RotationDistribution: RotationDistribution{
			Type: DistributionTypeWave,
		},
	}

	ApplyOptions(ctx, b, getOptions(ctx, options)...)

	g.Boxes = append(g.Boxes, b)
	return b
}

type RotationEventBox struct {
	IndexFilter
	BeatDistribution
	RotationDistribution

	Events []*RotationEvent
}

type RotationDistribution struct {
	Type        DistributionType
	Param       float64
	AffectFirst bool
	Flip        bool
	Axis        Axis
	Easing      Easing
}

type RotationEventOption interface {
	ApplyRotationEvent(*RotationEvent)
}
type RotationEventContextOption interface {
	ApplyRotationEventContext(ctx context.Context, e *RotationEvent)
}

func (b *RotationEventBox) AddEvent(ctx context.Context, options ...RotationEventOption) *RotationEvent {
	e := &RotationEvent{
		Beat:           ctx.B(),
		TransitionType: RotationTransitionTypeTransition,
		Easing:         EasingLinear,
		Direction:      RotationDirectionAutomatic,
	}

	ApplyOptions(ctx, e, getOptions(ctx, options)...)

	b.Events = append(b.Events, e)
	return e
}

type RotationEvent struct {
	Beat           float64
	TransitionType RotationEventTransitionType
	Easing         Easing
	LoopsCount     int
	Rotation       float64
	Direction      RotationDirection
}

type RotationEventTransitionType int

const (
	RotationTransitionTypeTransition RotationEventTransitionType = 0
	RotationTransitionTypeExtend     RotationEventTransitionType = 1
)

func (t RotationEventTransitionType) ApplyRotationEvent(e *RotationEvent) {
	e.TransitionType = t
}

type RotationDirection int

const (
	RotationDirectionAutomatic        RotationDirection = 0
	RotationDirectionClockwise        RotationDirection = 1
	RotationDirectionCounterClockwise RotationDirection = 2
)

const (
	RotAuto = RotationDirectionAutomatic
	CW      = RotationDirectionClockwise
	CCW     = RotationDirectionCounterClockwise
)

func (d *RotationDirection) Flip() {
	if *d == CW {
		*d = CCW
	} else if *d == CCW {
		*d = CW
	}
}

func (d RotationDirection) ApplyRotationEvent(e *RotationEvent) {
	e.Direction = d
}

type rotationBoxFuncOption func(*RotationEventBox)

func (fo rotationBoxFuncOption) ApplyRotationEventBox(b *RotationEventBox) { fo(b) }

type rotationEventFuncOption func(*RotationEvent)

func (fo rotationEventFuncOption) ApplyRotationEvent(e *RotationEvent) { fo(e) }

func OLoop(n int) rotationEventFuncOption {
	return func(e *RotationEvent) { e.LoopsCount = n }
}

func ORotation(r float64) rotationEventFuncOption {
	return func(e *RotationEvent) { e.Rotation = r }
}
