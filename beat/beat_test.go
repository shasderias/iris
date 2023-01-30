package beat

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/shasderias/iris/iter"
)

func TestSeq(t *testing.T) {
	testCases := []struct {
		name               string
		seqFn              func() Sequence
		wantSeqB           []float64
		wantSeqT           []float64
		wantSeqOrdinal     []int
		wantSeqNextB       []float64
		wantSeqNextBOffset []float64
		wantSeqPrevB       []float64
		wantSeqPrevBOffset []float64
		wantSeqFirst       []bool
		wantSeqLast        []bool
	}{
		{
			name:               "Empty",
			seqFn:              func() Sequence { return Seq() },
			wantSeqB:           []float64{},
			wantSeqT:           []float64{},
			wantSeqOrdinal:     []int{},
			wantSeqNextB:       []float64{},
			wantSeqNextBOffset: []float64{},
			wantSeqPrevB:       []float64{},
			wantSeqPrevBOffset: []float64{},
			wantSeqFirst:       []bool{},
			wantSeqLast:        []bool{},
		},
		{
			name:               "Seq",
			seqFn:              func() Sequence { return Seq(1, 2, 3, 4) },
			wantSeqB:           []float64{1, 2, 3, 4},
			wantSeqT:           []float64{0, 1.0 / 3.0, 2.0 / 3.0, 1},
			wantSeqOrdinal:     []int{0, 1, 2, 3},
			wantSeqNextB:       []float64{2, 3, 4, 5},
			wantSeqNextBOffset: []float64{1, 1, 1, 1},
			wantSeqPrevB:       []float64{0, 1, 2, 3},
			wantSeqPrevBOffset: []float64{1, 1, 1, 1},
			wantSeqFirst:       []bool{true, false, false, false},
			wantSeqLast:        []bool{false, false, false, true},
		},
		{
			name:               "SeqInterval",
			seqFn:              func() Sequence { return SeqInterval(0, 2, 4) },
			wantSeqB:           []float64{0, 0.25, 0.50, 0.75, 1.0, 1.25, 1.50, 1.75},
			wantSeqT:           []float64{0, 1.0 / 7.0, 2.0 / 7.0, 3.0 / 7.0, 4.0 / 7.0, 5.0 / 7.0, 6.0 / 7.0, 1},
			wantSeqOrdinal:     []int{0, 1, 2, 3, 4, 5, 6, 7},
			wantSeqNextB:       []float64{0.25, 0.50, 0.75, 1.0, 1.25, 1.50, 1.75, 2.0},
			wantSeqNextBOffset: []float64{0.25, 0.25, 0.25, 0.25, 0.25, 0.25, 0.25, 0.25},
			wantSeqPrevB:       []float64{0, 0, 0.25, 0.50, 0.75, 1.0, 1.25, 1.50},
			wantSeqPrevBOffset: []float64{0, 0.25, 0.25, 0.25, 0.25, 0.25, 0.25, 0.25},
			wantSeqFirst:       []bool{true, false, false, false, false, false, false, false},
			wantSeqLast:        []bool{false, false, false, false, false, false, false, true},
		},
		{
			name:               "SeqInterval",
			seqFn:              func() Sequence { return SeqInterval(1, 2, 4) },
			wantSeqB:           []float64{1.0, 1.25, 1.50, 1.75},
			wantSeqT:           []float64{0, 1.0 / 3.0, 2.0 / 3.0, 1},
			wantSeqOrdinal:     []int{0, 1, 2, 3},
			wantSeqNextB:       []float64{1.25, 1.50, 1.75, 2.0},
			wantSeqNextBOffset: []float64{0.25, 0.25, 0.25, 0.25},
			wantSeqPrevB:       []float64{0, 1.0, 1.25, 1.50},
			wantSeqPrevBOffset: []float64{1.0, 0.25, 0.25, 0.25},
			wantSeqFirst:       []bool{true, false, false, false},
			wantSeqLast:        []bool{false, false, false, true},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			seq := tt.seqFn()

			gotSeqB := iter.Map(seq.Iterate(), func(s *SeqIter) float64 { return s.SeqB() })
			if diff := cmp.Diff(gotSeqB, tt.wantSeqB); diff != "" {
				t.Fatal(diff)
			}
			gotSeqT := iter.Map(seq.Iterate(), func(s *SeqIter) float64 { return s.SeqT() })
			if diff := cmp.Diff(gotSeqT, tt.wantSeqT); diff != "" {
				t.Fatal(diff)
			}
			gotSeqOrdinal := iter.Map(seq.Iterate(), func(s *SeqIter) int { return s.SeqOrdinal() })
			if diff := cmp.Diff(gotSeqOrdinal, tt.wantSeqOrdinal); diff != "" {
				t.Fatal(diff)
			}
			gotSeqNextB := iter.Map(seq.Iterate(), func(s *SeqIter) float64 { return s.SeqNextB() })
			if diff := cmp.Diff(gotSeqNextB, tt.wantSeqNextB); diff != "" {
				t.Fatal(diff)
			}
			gotSeqNextBOffset := iter.Map(seq.Iterate(), func(s *SeqIter) float64 { return s.SeqNextBOffset() })
			if diff := cmp.Diff(gotSeqNextBOffset, tt.wantSeqNextBOffset); diff != "" {
				t.Fatal(diff)
			}
			gotSeqPrevB := iter.Map(seq.Iterate(), func(s *SeqIter) float64 { return s.SeqPrevB() })
			if diff := cmp.Diff(gotSeqPrevB, tt.wantSeqPrevB); diff != "" {
				t.Fatal(diff)
			}
			gotSeqPrevBOffset := iter.Map(seq.Iterate(), func(s *SeqIter) float64 { return s.SeqPrevBOffset() })
			if diff := cmp.Diff(gotSeqPrevBOffset, tt.wantSeqPrevBOffset); diff != "" {
				t.Fatal(diff)
			}
			gotSeqFirst := iter.Map(seq.Iterate(), func(s *SeqIter) bool { return s.SeqFirst() })
			if diff := cmp.Diff(gotSeqFirst, tt.wantSeqFirst); diff != "" {
				t.Fatal(diff)
			}
			gotSeqLast := iter.Map(seq.Iterate(), func(s *SeqIter) bool { return s.SeqLast() })
			if diff := cmp.Diff(gotSeqLast, tt.wantSeqLast); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestRng(t *testing.T) {
	testCases := []struct {
		name           string
		rngFn          func() Range
		wantRngB       []float64
		wantRngT       []float64
		wantRngOrdinal []int
	}{
		{
			name:           "RngInterval",
			rngFn:          func() Range { return RngInterval(0, 2, 4) },
			wantRngB:       []float64{0, 0.25, 0.50, 0.75, 1.0, 1.25, 1.50, 1.75, 2.00},
			wantRngT:       []float64{0, 1.0 / 8.0, 2.0 / 8.0, 3.0 / 8.0, 4.0 / 8.0, 5.0 / 8.0, 6.0 / 8.0, 7.0 / 8.0, 1},
			wantRngOrdinal: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:           "Fractional",
			rngFn:          func() Range { return RngInterval(0.1, 1.3, 4) },
			wantRngB:       []float64{0.1, 0.35, 0.60, 0.85, 1.1},
			wantRngT:       []float64{0, 1.0 / 4.0, 2.0 / 4.0, 3.0 / 4.0, 1},
			wantRngOrdinal: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			rng := tt.rngFn()

			gotRngB := iter.Map(rng.Iterate(), func(r *RngIter) float64 { return r.RngB() })
			if diff := cmp.Diff(gotRngB, tt.wantRngB); diff != "" {
				t.Fatal(diff)
			}
			gotRngT := iter.Map(rng.Iterate(), func(r *RngIter) float64 { return r.RngT() })
			if diff := cmp.Diff(gotRngT, tt.wantRngT); diff != "" {
				t.Fatal(diff)
			}
			gotRngOrdinal := iter.Map(rng.Iterate(), func(r *RngIter) int { return r.RngOrdinal() })
			if diff := cmp.Diff(gotRngOrdinal, tt.wantRngOrdinal); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
