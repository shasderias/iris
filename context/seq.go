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
func (s seqContext) SeqT() float64           { return parentPassthrough(s).SeqT() }
func (s seqContext) SeqOrdinal() int         { return parentPassthrough(s).SeqOrdinal() }
func (s seqContext) SeqOrdinalF() float64    { return float64(parentPassthrough(s).SeqOrdinal()) }
func (s seqContext) SeqLen() int             { return parentPassthrough(s).SeqLen() }
func (s seqContext) SeqNextB() float64       { return parentPassthrough(s).SeqNextB() }
func (s seqContext) SeqNextBOffset() float64 { return parentPassthrough(s).SeqNextBOffset() }
func (s seqContext) SeqPrevB() float64       { return parentPassthrough(s).SeqPrevB() }
func (s seqContext) SeqPrevBOffset() float64 { return parentPassthrough(s).SeqPrevBOffset() }
func (s seqContext) SeqFirst() bool          { return parentPassthrough(s).SeqFirst() }
func (s seqContext) SeqLast() bool           { return parentPassthrough(s).SeqLast() }

func (s seqContext) WSeq(seq beat.Sequence, callback func(ctx Context)) {
	WSeq(s, seq, callback)
}
func (s seqContext) WRng(rng beat.Range, callback func(ctx Context)) {
	WRng(s, rng, callback)
}
func (s seqContext) WOpt(options ...any) Doer {
	return doer{WOptions(s, options...)}
}

func parentPassthrough(s seqContext) seqPassthrough {
	if len(s.Sequence.Beats) == 1 && s.Sequence.Beats[0] <= 0 {
		return s.parent()
	}
	return s.SeqIter
}

type seqPassthrough interface {
	SeqB() float64
	SeqT() float64
	SeqOrdinal() int
	SeqLen() int
	SeqNextB() float64
	SeqNextBOffset() float64
	SeqPrevB() float64
	SeqPrevBOffset() float64
	SeqFirst() bool
	SeqLast() bool
}
