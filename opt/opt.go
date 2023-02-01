package opt

import (
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/evt"
	"github.com/shasderias/iris/internal/calc"
	"github.com/shasderias/iris/scale"
)

type FuncContextOpt[T any] func(context.Context) T

func (fn FuncContextOpt[T]) ApplyBasicEvent(e *evt.BasicEvent) { panic("upgrade me") }
func (fn FuncContextOpt[T]) ApplyBasicEventContext(ctx context.Context, e *evt.BasicEvent) {
	if opt, ok := any(fn(ctx)).(evt.BasicEventOption); ok {
		evt.ApplyOptions(ctx, e, opt)
	}
}

func (fn FuncContextOpt[T]) ApplyColorEventGroup(e *evt.ColorEventGroup) { panic("upgrade me") }
func (fn FuncContextOpt[T]) ApplyColorEventGroupContext(ctx context.Context, e *evt.ColorEventGroup) {
	if opt, ok := any(fn(ctx)).(evt.ColorEventGroupOption); ok {
		evt.ApplyOptions(ctx, e, opt)
	}
}
func (fn FuncContextOpt[T]) ApplyColorEventBox(e *evt.ColorEventBox) { panic("upgrade me") }
func (fn FuncContextOpt[T]) ApplyColorEventBoxContext(ctx context.Context, e *evt.ColorEventBox) {
	if opt, ok := any(fn(ctx)).(evt.ColorEventBoxOption); ok {
		evt.ApplyOptions(ctx, e, opt)
	}
}
func (fn FuncContextOpt[T]) ApplyColorEvent(e *evt.ColorEvent) { panic("upgrade me") }
func (fn FuncContextOpt[T]) ApplyColorEventContext(ctx context.Context, e *evt.ColorEvent) {
	if opt, ok := any(fn(ctx)).(evt.ColorEventOption); ok {
		evt.ApplyOptions(ctx, e, opt)
	}
}

func (fn FuncContextOpt[T]) ApplyRotationEventGroup(e *evt.RotationEventGroup) { panic("upgrade me") }
func (fn FuncContextOpt[T]) ApplyRotationEventGroupContext(ctx context.Context, e *evt.RotationEventGroup) {
	if opt, ok := any(fn(ctx)).(evt.RotationEventGroupOption); ok {
		evt.ApplyOptions(ctx, e, opt)
	}
}
func (fn FuncContextOpt[T]) ApplyRotationEventBox(e *evt.RotationEventBox) { panic("upgrade me") }
func (fn FuncContextOpt[T]) ApplyRotationEventBoxContext(ctx context.Context, e *evt.RotationEventBox) {
	if opt, ok := any(fn(ctx)).(evt.RotationEventBoxOption); ok {
		evt.ApplyOptions(ctx, e, opt)
	}
}
func (fn FuncContextOpt[T]) ApplyRotationEvent(e *evt.RotationEvent) { panic("upgrade me") }
func (fn FuncContextOpt[T]) ApplyRotationEventContext(ctx context.Context, e *evt.RotationEvent) {
	if opt, ok := any(fn(ctx)).(evt.RotationEventOption); ok {
		evt.ApplyOptions(ctx, e, opt)
	}
}

func SeqOrdinal[T any](options ...T) FuncContextOpt[T] {
	return func(ctx context.Context) T {
		return calc.IdxWrap(options, ctx.SeqOrdinal())
	}
}

func Ordinal[T any](options ...T) FuncContextOpt[T] {
	return func(ctx context.Context) T {
		return calc.IdxWrap(options, ctx.Ordinal())
	}
}

func T[T any](options ...T) FuncContextOpt[T] {
	scaler := scale.FromUnitClamp(0, float64(len(options)-1))
	return func(ctx context.Context) T {
		return options[int(scaler(ctx.T()))]
	}
}

func IfFirst[T any](first, notFirst T) FuncContextOpt[T] {
	return func(ctx context.Context) T {
		if ctx.First() {
			return first
		}
		return notFirst
	}
}

func IfLast[T any](last, notLast T) FuncContextOpt[T] {
	return func(ctx context.Context) T {
		if ctx.Last() {
			return last
		}
		return notLast
	}
}

func IfSeqFirst[T any](first, notFirst T) FuncContextOpt[T] {
	return func(ctx context.Context) T {
		if ctx.SeqFirst() {
			return first
		}
		return notFirst
	}
}

func IfSeqLast[T any](last, notLast T) FuncContextOpt[T] {
	return func(ctx context.Context) T {
		if ctx.SeqLast() {
			return last
		}
		return notLast
	}
}

func Of[T any](options ...any) []T {
	ret := []T{}
	for _, opt := range options {
		if o, ok := opt.(T); ok {
			ret = append(ret, o)
		}
	}
	return ret
}

type Combined []any

func Combine(options ...any) Combined {
	return options
}

func (c Combined) ApplyBasicEvent(e *evt.BasicEvent) { panic("upgrade me") }
func (c Combined) ApplyBasicEventContext(ctx context.Context, e *evt.BasicEvent) {
	for _, opt := range Of[evt.BasicEventOption](c...) {
		if o, ok := opt.(evt.BasicEventContextOption); ok {
			o.ApplyBasicEventContext(ctx, e)
		} else {
			opt.ApplyBasicEvent(e)
		}
	}
}

func (c Combined) ApplyColorEventGroup(e *evt.ColorEventGroup) { panic("upgrade me") }
func (c Combined) ApplyColorEventGroupContext(ctx context.Context, e *evt.ColorEventGroup) {
	for _, opt := range Of[evt.ColorEventGroupOption](c...) {
		if o, ok := opt.(evt.ColorEventGroupContextOption); ok {
			o.ApplyColorEventGroupContext(ctx, e)
		} else {
			opt.ApplyColorEventGroup(e)
		}
	}
}

func (c Combined) ApplyColorEventBox(e *evt.ColorEventBox) { panic("upgrade me") }
func (c Combined) ApplyColorEventBoxContext(ctx context.Context, e *evt.ColorEventBox) {
	for _, opt := range Of[evt.ColorEventBoxOption](c...) {
		if o, ok := opt.(evt.ColorEventBoxContextOption); ok {
			o.ApplyColorEventBoxContext(ctx, e)
		} else {
			opt.ApplyColorEventBox(e)
		}
	}
}

func (c Combined) ApplyColorEvent(e *evt.ColorEvent) { panic("upgrade me") }
func (c Combined) ApplyColorEventContext(ctx context.Context, e *evt.ColorEvent) {
	for _, opt := range Of[evt.ColorEventOption](c...) {
		if o, ok := opt.(evt.ColorEventContextOption); ok {
			o.ApplyColorEventContext(ctx, e)
		} else {
			opt.ApplyColorEvent(e)
		}
	}
}

func (c Combined) ApplyRotationEventGroup(e *evt.RotationEventGroup) { panic("upgrade me") }
func (c Combined) ApplyRotationEventGroupContext(ctx context.Context, e *evt.RotationEventGroup) {
	for _, opt := range Of[evt.RotationEventGroupOption](c...) {
		if o, ok := opt.(evt.RotationEventGroupContextOption); ok {
			o.ApplyRotationEventGroupContext(ctx, e)
		} else {
			opt.ApplyRotationEventGroup(e)
		}
	}
}

func (c Combined) ApplyRotationEventBox(e *evt.RotationEventBox) { panic("upgrade me") }
func (c Combined) ApplyRotationEventBoxContext(ctx context.Context, e *evt.RotationEventBox) {
	for _, opt := range Of[evt.RotationEventBoxOption](c...) {
		if o, ok := opt.(evt.RotationEventBoxContextOption); ok {
			o.ApplyRotationEventBoxContext(ctx, e)
		} else {
			opt.ApplyRotationEventBox(e)
		}
	}
}

func (c Combined) ApplyRotationEvent(e *evt.RotationEvent) { panic("upgrade me") }
func (c Combined) ApplyRotationEventContext(ctx context.Context, e *evt.RotationEvent) {
	for _, opt := range Of[evt.RotationEventOption](c...) {
		if o, ok := opt.(evt.RotationEventContextOption); ok {
			o.ApplyRotationEventContext(ctx, e)
		} else {
			opt.ApplyRotationEvent(e)
		}
	}
}
