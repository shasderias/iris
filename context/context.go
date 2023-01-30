package context

import (
	"github.com/shasderias/iris/beat"
	"github.com/shasderias/iris/iter"
)

type Context interface {
	parent() Context

	AddEvent(event ...any)
	Events() *[]any
	Options() []any

	B() float64
	BOffset() float64
	T() float64
	Ordinal() int
	OrdinalF() float64
	First() bool
	Last() bool

	SeqB() float64
	SeqT() float64
	SeqOrdinal() int
	SeqOrdinalF() float64
	SeqLen() int
	SeqNextB() float64
	SeqNextBOffset() float64
	SeqPrevB() float64
	SeqPrevBOffset() float64
	SeqFirst() bool
	SeqLast() bool

	WSeq(seq beat.Sequence, callback func(ctx Context))
	WRng(rng beat.Range, callback func(ctx Context))
	WOpt(options ...any) Doer
}

type Doer interface {
	Do(callback func(ctx Context))
}

type Iteratable[T any] interface {
	Iterate() iter.Iter[T]
}

func WRng(ctx Context, rng beat.Range, callback func(ctx Context)) {
	iter.ForEach(rng.Iterate(), func(r *beat.RngIter) {
		callback(rngContext{ctx, r})
	})
}
func WSeq(ctx Context, seq beat.Sequence, callback func(ct Context)) {
	iter.ForEach(seq.Iterate(), func(s *beat.SeqIter) {
		callback(seqContext{ctx, s})
	})
}
