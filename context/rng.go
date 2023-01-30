package context

import (
	"github.com/shasderias/iris/beat"
)

type rngContext struct {
	p Context
	*beat.RngIter
}

func (r rngContext) parent() Context { return r.p }

func (r rngContext) AddEvent(event ...any) {
	r.p.AddEvent(event...)
}
func (r rngContext) Events() *[]any {
	return r.p.Events()
}
func (r rngContext) Options() []any {
	return r.p.Options()
}

func (r rngContext) B() float64              { return r.RngIter.RngB() }
func (r rngContext) BOffset() float64        { return r.p.BOffset() + r.RngIter.RngB() }
func (r rngContext) T() float64              { return r.RngIter.RngT() }
func (r rngContext) Ordinal() int            { return r.RngIter.RngOrdinal() }
func (r rngContext) OrdinalF() float64       { return float64(r.RngIter.RngOrdinal()) }
func (r rngContext) First() bool             { return r.RngIter.First() }
func (r rngContext) Last() bool              { return r.RngIter.Last() }
func (r rngContext) SeqB() float64           { return r.p.SeqB() }
func (r rngContext) SeqT() float64           { return r.p.SeqT() }
func (r rngContext) SeqOrdinal() int         { return r.p.SeqOrdinal() }
func (r rngContext) SeqOrdinalF() float64    { return float64(r.p.SeqOrdinal()) }
func (r rngContext) SeqLen() int             { return r.p.SeqLen() }
func (r rngContext) SeqNextB() float64       { return r.p.SeqNextB() }
func (r rngContext) SeqNextBOffset() float64 { return r.p.SeqNextBOffset() }
func (r rngContext) SeqPrevB() float64       { return r.p.SeqPrevB() }
func (r rngContext) SeqPrevBOffset() float64 { return r.p.SeqPrevBOffset() }
func (r rngContext) SeqFirst() bool          { return r.p.SeqFirst() }
func (r rngContext) SeqLast() bool           { return r.p.SeqLast() }

func (r rngContext) WSeq(seq beat.Sequence, callback func(ctx Context)) {
	WSeq(r, seq, callback)
}
func (r rngContext) WRng(rng beat.Range, callback func(ctx Context)) {
	WRng(r, rng, callback)
}
func (r rngContext) WOpt(options ...any) Doer {
	return doer{WOptions(r, options...)}
}
