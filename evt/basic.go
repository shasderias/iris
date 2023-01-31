package evt

import (
	"github.com/shasderias/iris/context"
)

type BasicEvent struct {
	Beat       float64
	Type       int
	Value      int
	FloatValue float64
}

type BasicEventOption interface {
	ApplyBasicEvent(e *BasicEvent)
}
type BasicEventContextOption interface {
	ApplyBasicEventContext(ctx context.Context, e *BasicEvent)
}

func Basic(ctx context.Context, options ...BasicEventOption) *BasicEvent {
	be := &BasicEvent{
		Beat:       ctx.BOffset(),
		Type:       0,
		Value:      0,
		FloatValue: 0,
	}

	ApplyOptions(ctx, be, getOptions(ctx, options)...)

	ctx.AddEvent(be)
	return be
}

func OType(t int) basicEventFuncOption {
	return func(e *BasicEvent) { e.Type = t }
}
func OValue(v int) basicEventFuncOption {
	return func(e *BasicEvent) { e.Value = v }
}
func OFloatValue(v float64) basicEventFuncOption {
	return func(e *BasicEvent) { e.FloatValue = v }
}

type basicEventFuncOption func(*BasicEvent)

func (fo basicEventFuncOption) ApplyBasicEvent(e *BasicEvent) { fo(e) }

var LightOff = basicEventFuncOption(func(e *BasicEvent) { e.Value = 0 })
