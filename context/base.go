package context

import (
	"github.com/shasderias/iris/beat"
)

type baseContext struct {
	events *[]any
}

func NewBase() Context {
	events := make([]any, 0)
	return baseContext{events: &events}
}

func (b baseContext) parent() Context { return nil }

func (b baseContext) AddEvent(event ...any) {
	*b.events = append(*b.events, event...)
}
func (b baseContext) Events() *[]any {
	return b.events
}
func (b baseContext) Options() []any {
	return []any{}
}

func (b baseContext) B() float64              { return 0 }
func (b baseContext) BOffset() float64        { return 0 }
func (b baseContext) T() float64              { return 1 }
func (b baseContext) Ordinal() int            { return 0 }
func (b baseContext) OrdinalF() float64       { return 0 }
func (b baseContext) First() bool             { return true }
func (b baseContext) Last() bool              { return true }
func (b baseContext) SeqB() float64           { return 0 }
func (b baseContext) SeqT() float64           { return 1 }
func (b baseContext) SeqOrdinal() int         { return 0 }
func (b baseContext) SeqOrdinalF() float64    { return 0 }
func (b baseContext) SeqLen() int             { return 0 }
func (b baseContext) SeqNextB() float64       { return 0 }
func (b baseContext) SeqNextBOffset() float64 { return 0 }
func (b baseContext) SeqPrevB() float64       { return 0 }
func (b baseContext) SeqPrevBOffset() float64 { return 0 }
func (b baseContext) SeqFirst() bool          { return true }
func (b baseContext) SeqLast() bool           { return true }

func (b baseContext) WSeq(seq beat.Sequence, callback func(ctx Context)) {
	WSeq(b, seq, callback)
}
func (b baseContext) WRng(rng beat.Range, callback func(ctx Context)) {
	WRng(b, rng, callback)
}
func (b baseContext) WOpt(options ...any) Doer {
	return doer{WOptions(b, options...)}
}
