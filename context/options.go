package context

import (
	"github.com/shasderias/iris/beat"
)

type optContext struct {
	p       Context
	options []any
}

func (c optContext) parent() Context { return c.p }

func (c optContext) AddEvent(event ...any) {
	c.p.AddEvent(event...)
}

func (c optContext) Options() []any {
	return append(c.options, c.p.Options()...)
}

func (c optContext) Events() *[]any {
	return c.p.Events()
}

func (c optContext) B() float64              { return c.p.B() }
func (c optContext) BOffset() float64        { return c.p.BOffset() }
func (c optContext) T() float64              { return c.p.T() }
func (c optContext) Ordinal() int            { return c.p.Ordinal() }
func (c optContext) OrdinalF() float64       { return float64(c.p.Ordinal()) }
func (c optContext) First() bool             { return c.p.First() }
func (c optContext) Last() bool              { return c.p.Last() }
func (c optContext) SeqB() float64           { return c.p.SeqB() }
func (c optContext) SeqT() float64           { return c.p.SeqT() }
func (c optContext) SeqOrdinal() int         { return c.p.SeqOrdinal() }
func (c optContext) SeqOrdinalF() float64    { return float64(c.p.SeqOrdinal()) }
func (c optContext) SeqLen() int             { return c.p.SeqLen() }
func (c optContext) SeqNextB() float64       { return c.p.SeqNextB() }
func (c optContext) SeqNextBOffset() float64 { return c.p.SeqNextBOffset() }
func (c optContext) SeqPrevB() float64       { return c.p.SeqPrevB() }
func (c optContext) SeqPrevBOffset() float64 { return c.p.SeqPrevBOffset() }
func (c optContext) SeqFirst() bool          { return c.p.SeqFirst() }
func (c optContext) SeqLast() bool           { return c.p.SeqLast() }

func (c optContext) WSeq(seq beat.Sequence, callback func(ctx Context)) {
	WSeq(c, seq, callback)
}
func (c optContext) WRng(rng beat.Range, callback func(ctx Context)) {
	WRng(c, rng, callback)
}
func (c optContext) WOpt(options ...any) Doer {
	return doer{WOptions(c, options...)}
}

func WOptions(ctx Context, options ...any) Context {
	return optContext{p: ctx, options: options}
}

type doer struct {
	Context
}

func (d doer) Do(callback func(ctx Context)) {
	callback(d)
}
