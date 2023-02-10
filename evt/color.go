package evt

import (
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/iter"
)

type ColorEventGroup struct {
	Beat  float64
	Group []int
	Boxes []*ColorEventBox
}

type ColorEventGroupOption interface {
	ApplyColorEventGroup(*ColorEventGroup)
}
type ColorEventGroupContextOption interface {
	ApplyColorEventGroupContext(ctx context.Context, e *ColorEventGroup)
}

func ColorGroup(ctx context.Context, options ...ColorEventGroupOption) *ColorEventGroup {
	g := &ColorEventGroup{Beat: ctx.BOffset()}

	ApplyOptions(ctx, g, getOptions(ctx, options)...)

	ctx.AddEvent(g)
	return g
}

func ColorGroupWithBox(ctx context.Context, options ...any) (*ColorEventGroup, *ColorEventBox) {
	g := ColorGroup(ctx, optionsOf[ColorEventGroupOption](options...)...)
	b := g.AddBox(ctx, optionsOf[ColorEventBoxOption](options...)...)
	return g, b
}

type ColorEventBoxOption interface {
	ApplyColorEventBox(*ColorEventBox)
}
type ColorEventBoxContextOption interface {
	ApplyColorEventBoxContext(ctx context.Context, e *ColorEventBox)
}

func (g *ColorEventGroup) AddBox(ctx context.Context, options ...ColorEventBoxOption) *ColorEventBox {
	b := &ColorEventBox{
		IndexFilter: IndexFilter{
			IndexFilterType: IndexFilterTypeSections,
		},
		BeatDistribution: BeatDistribution{
			Type: DistributionTypeWave,
		},
		BrightnessDistribution: BrightnessDistribution{
			Type: DistributionTypeWave,
		},
	}

	ApplyOptions(ctx, b, getOptions[ColorEventBoxOption](ctx, options)...)

	g.Boxes = append(g.Boxes, b)
	return b
}

type ColorEventBox struct {
	IndexFilter
	BeatDistribution
	BrightnessDistribution

	Events []*ColorEvent
}

type ColorEventOption interface {
	ApplyColorEvent(*ColorEvent)
}
type ColorEventContextOption interface {
	ApplyColorEventContext(ctx context.Context, e *ColorEvent)
}

type BrightnessDistribution struct {
	Type        DistributionType
	Param       float64
	AffectFirst bool
	Easing      Easing
}

func (b *ColorEventBox) AddEvent(ctx context.Context, options ...ColorEventOption) *ColorEvent {
	e := &ColorEvent{Beat: ctx.B(), Brightness: 1.0}

	ApplyOptions(ctx, e, getOptions(ctx, options)...)

	b.Events = append(b.Events, e)
	return e
}

type ColorEvent struct {
	Beat            float64
	TransitionType  ColorEventTransitionType
	Color           ColorEventColor
	Brightness      float64
	StrobeFrequency float64
}

type ColorEventTransitionType int

const (
	ColorEventTransitionInstant    ColorEventTransitionType = 0
	ColorEventTransitionTransition ColorEventTransitionType = 1
	ColorEventTransitionExtend     ColorEventTransitionType = 2
)

func (t ColorEventTransitionType) ApplyColorEvent(e *ColorEvent) {
	e.TransitionType = t
}

type ColorEventColor int

const (
	ColorEventRed   ColorEventColor = 0
	ColorEventBlue  ColorEventColor = 1
	ColorEventWhite ColorEventColor = 2

	Red   = ColorEventRed
	Blue  = ColorEventBlue
	White = ColorEventWhite
)

func (c ColorEventColor) ApplyColorEvent(e *ColorEvent) {
	e.Color = c
}

type colorEventGroupFuncOption func(*ColorEventGroup)

func (fo colorEventGroupFuncOption) ApplyColorEventGroup(g *ColorEventGroup) {
	fo(g)
}

func OColorEventGroup(g ...int) colorEventGroupFuncOption {
	return func(ge *ColorEventGroup) {
		for _, gid := range g {
			if !iter.SliceHas(ge.Group, gid) {
				ge.Group = append(ge.Group, gid)
			}
		}
	}
}

type colorEventFuncOption func(*ColorEvent)

func (fo colorEventFuncOption) ApplyColorEvent(e *ColorEvent) { fo(e) }

func OBrightness(b float64) colorEventFuncOption {
	return func(e *ColorEvent) { e.Brightness = b }
}

func OStrobe(f float64) colorEventFuncOption {
	return func(e *ColorEvent) { e.StrobeFrequency = f }
}

type colorEventContextFuncOption func(ctx context.Context, e *ColorEvent)

func (fo colorEventContextFuncOption) ApplyColorEvent(*ColorEvent) { panic("upgrade me") }
func (fo colorEventContextFuncOption) ApplyColorEventContext(ctx context.Context, e *ColorEvent) {
	fo(ctx, e)
}

func OBrightnessT(funcs ...func(float64) float64) colorEventContextFuncOption {
	return func(ctx context.Context, e *ColorEvent) {
		v := ctx.T()
		for _, fn := range funcs {
			v = fn(v)
		}
		e.Brightness = v
	}
}
