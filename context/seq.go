package context

import (
	"github.com/shasderias/iris/beat"
)

type seqContext struct {
	p Context
	*beat.SeqIter
}

func (s seqContext) parent() Context { return s.p }

func (s seqContext) AddEvent(event ...any) {
	s.p.AddEvent(event...)
}
func (s seqContext) Events() *[]any {
	return s.p.Events()
}
func (s seqContext) Options() []any {
	return s.p.Options()
}

func (s seqContext) B() float64              { return s.SeqIter.SeqB() }
func (s seqContext) BOffset() float64        { return s.p.BOffset() + s.SeqIter.SeqB() }
func (s seqContext) T() float64              { return s.SeqIter.SeqT() }
func (s seqContext) Ordinal() int            { return s.SeqIter.SeqOrdinal() }
func (s seqContext) OrdinalF() float64       { return float64(s.SeqIter.SeqOrdinal()) }
func (s seqContext) First() bool             { return s.SeqIter.SeqFirst() }
func (s seqContext) Last() bool              { return s.SeqIter.SeqLast() }
func (s seqContext) SeqT() float64           { return s.SeqIter.SeqT() }
func (s seqContext) SeqOrdinal() int         { return s.SeqIter.SeqOrdinal() }
func (s seqContext) SeqOrdinalF() float64    { return float64(s.SeqIter.SeqOrdinal()) }
func (s seqContext) SeqLen() int             { return s.SeqIter.SeqLen() }
func (s seqContext) SeqNextB() float64       { return s.SeqIter.SeqNextB() }
func (s seqContext) SeqNextBOffset() float64 { return s.SeqIter.SeqNextBOffset() }
func (s seqContext) SeqPrevB() float64       { return s.SeqIter.SeqPrevB() }
func (s seqContext) SeqPrevBOffset() float64 { return s.SeqIter.SeqPrevBOffset() }
func (s seqContext) SeqFirst() bool          { return s.SeqIter.SeqFirst() }
func (s seqContext) SeqLast() bool           { return s.SeqIter.SeqLast() }

func (s seqContext) WSeq(seq beat.Sequence, callback func(ctx Context)) {
	WSeq(s, seq, callback)
}
func (s seqContext) WRng(rng beat.Range, callback func(ctx Context)) {
	WRng(s, rng, callback)
}
func (s seqContext) WOpt(options ...any) Doer {
	return doer{WOptions(s, options...)}
}
