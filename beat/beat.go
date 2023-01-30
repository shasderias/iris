package beat

import (
	"math"

	"github.com/shasderias/iris/iter"
)

type Sequence struct {
	Beats     []float64
	GhostBeat float64
}

func Seq(b ...float64) Sequence {
	switch len(b) {
	case 0:
		return Sequence{[]float64{}, 0}
	case 1:
		return Sequence{[]float64{b[0]}, b[0] + b[0]}
	default:
		return Sequence{b, b[len(b)-1] + (b[len(b)-1] - b[len(b)-2])}
	}
}

func SeqInterval(start, end float64, divisor float64) Sequence {
	count := int(math.Ceil((end - start) * divisor))
	beats := make([]float64, count)
	for i := 0; i < count; i++ {
		beats[i] = start + float64(i)*(1.0/float64(divisor))
	}
	return Sequence{beats, beats[len(beats)-1] + (1.0 / float64(divisor))}
}

func (s Sequence) Iterate() iter.Iter[*SeqIter] {
	return &SeqIter{s, -1}
}

type SeqIter struct {
	Sequence
	ordinal int
}

func (s *SeqIter) Next() (*SeqIter, bool) {
	s.ordinal++
	if s.ordinal == s.SeqLen() {
		return s, false
	}
	return s, true
}

func (s *SeqIter) SeqB() float64 { return s.Beats[s.ordinal] }
func (s *SeqIter) SeqT() float64 {
	if s.SeqLen() == 1 {
		return 1
	} else {
		return float64(s.ordinal) / float64(s.SeqLen()-1)
	}
}
func (s *SeqIter) SeqOrdinal() int { return s.ordinal }
func (s *SeqIter) SeqLen() int     { return len(s.Beats) }
func (s *SeqIter) SeqNextB() float64 {
	if s.SeqLast() {
		return s.GhostBeat
	} else {
		return s.Beats[s.ordinal+1]
	}
}
func (s *SeqIter) SeqNextBOffset() float64 { return s.SeqNextB() - s.Beats[s.ordinal] }
func (s *SeqIter) SeqPrevB() float64 {
	if s.SeqFirst() {
		return 0
	} else {
		return s.Beats[s.ordinal-1]
	}
}
func (s *SeqIter) SeqPrevBOffset() float64 { return s.Beats[s.ordinal] - s.SeqPrevB() }
func (s *SeqIter) SeqFirst() bool          { return s.ordinal == 0 }
func (s *SeqIter) SeqLast() bool           { return s.ordinal == s.SeqLen()-1 }

type Range struct {
	start, end float64
	step       float64
}

func RngStep(start, end float64, steps float64) Range {
	if steps == 1 {
		return Range{start, start, 1}
	}
	if steps == 2 {
		return Range{start, end, end - start}
	}
	return Range{start, end, (end - start) / (steps - 1)}
}

func RngInterval(start, end float64, divisor int) Range {
	return Range{start, end, 1.0 / float64(divisor)}
}

func (r Range) Iterate() iter.Iter[*RngIter] {
	return &RngIter{r, -1}
}

type RngIter struct {
	Range
	ordinal int
}

func (r *RngIter) Next() (*RngIter, bool) {
	r.ordinal++
	if r.ordinal == r.RngLen() {
		return r, false
	}
	return r, true
}

func (r *RngIter) RngB() float64 { return r.start + float64(r.ordinal)*r.step }
func (r *RngIter) RngT() float64 {
	if r.RngLen() == 1 {
		return 1
	} else {
		return float64(r.ordinal) / float64(r.RngLen()-1)
	}
}
func (r *RngIter) RngOrdinal() int { return r.ordinal }
func (r *RngIter) RngLen() int {
	l := int(math.Floor((r.end-r.start)/r.step)) + 1
	return l
}
func (r *RngIter) First() bool { return r.ordinal == 0 }
func (r *RngIter) Last() bool  { return r.ordinal == r.RngLen()-1 }
