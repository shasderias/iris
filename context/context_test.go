package context_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/shasderias/iris/beat"
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/evt"
)

type contextValues struct {
	B       float64
	BOffset float64
	T       float64
	Ordinal int
	First   bool
	Last    bool

	SeqB           float64
	SeqT           float64
	SeqOrdinal     int
	SeqLen         int
	SeqNextB       float64
	SeqNextBOffset float64
	SeqPrevB       float64
	SeqPrevBOffset float64
	SeqFirst       bool
	SeqLast        bool
}

func valuesFromContext(ctx context.Context) contextValues {
	return contextValues{
		B:       ctx.B(),
		BOffset: ctx.BOffset(),
		T:       ctx.T(),
		Ordinal: ctx.Ordinal(),
		First:   ctx.First(),
		Last:    ctx.Last(),

		SeqB:           ctx.SeqB(),
		SeqT:           ctx.SeqT(),
		SeqOrdinal:     ctx.SeqOrdinal(),
		SeqLen:         ctx.SeqLen(),
		SeqNextB:       ctx.SeqNextB(),
		SeqNextBOffset: ctx.SeqNextBOffset(),
		SeqPrevB:       ctx.SeqPrevB(),
		SeqPrevBOffset: ctx.SeqPrevBOffset(),
		SeqFirst:       ctx.SeqFirst(),
		SeqLast:        ctx.SeqLast(),
	}
}

func TestSanity(t *testing.T) {
	ctx := context.NewBase()

	assertContextValues(t, ctx, contextValues{
		B:              0,
		BOffset:        0,
		T:              1,
		Ordinal:        0,
		First:          true,
		Last:           true,
		SeqB:           0,
		SeqT:           1,
		SeqOrdinal:     0,
		SeqLen:         0,
		SeqNextB:       0,
		SeqNextBOffset: 0,
		SeqPrevB:       0,
		SeqPrevBOffset: 0,
		SeqFirst:       true,
		SeqLast:        true,
	})

}

func TestSeq(t *testing.T) {
	ctx := context.NewBase()

	wantContextValues := []contextValues{
		{1, 1, 0.0 / 3.0, 0, true, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{2, 2, 1.0 / 3.0, 1, false, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{3, 3, 2.0 / 3.0, 2, false, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{4, 4, 3.0 / 3.0, 3, false, true, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
	}
	i := 0
	ctx.WSeq(beat.Seq(1, 2, 3, 4), func(ctx context.Context) {
		assertContextValues(t, ctx, wantContextValues[i])
		i++
	})
}

func TestRng(t *testing.T) {
	ctx := context.NewBase()

	wantContextValues := []contextValues{
		{1.0, 1.0, 0.0, 0, true, false, 0, 1, 0, 0, 0, 0, 0, 0, true, true},
		{1.5, 1.5, 1.0, 1, false, true, 0, 1, 0, 0, 0, 0, 0, 0, true, true},
	}
	i := 0
	ctx.WRng(beat.RngStep(1, 2, 2), func(ctx context.Context) {
		assertContextValues(t, ctx, wantContextValues[i])
		i++
	})
}

func TestSeqRng(t *testing.T) {
	ctx := context.NewBase()

	wantContextValues := []contextValues{
		{0.00, 1.00, 0.0 / 3.0, 0, true, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.25, 1.25, 1.0 / 3.0, 1, false, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.50, 1.50, 2.0 / 3.0, 2, false, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.75, 1.75, 3.0 / 3.0, 3, false, true, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.00, 2.00, 0.0 / 3.0, 0, true, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.25, 2.25, 1.0 / 3.0, 1, false, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.50, 2.50, 2.0 / 3.0, 2, false, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.75, 2.75, 3.0 / 3.0, 3, false, true, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.00, 3.00, 0.0 / 3.0, 0, true, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.25, 3.25, 1.0 / 3.0, 1, false, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.50, 3.50, 2.0 / 3.0, 2, false, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.75, 3.75, 3.0 / 3.0, 3, false, true, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.00, 4.00, 0.0 / 3.0, 0, true, false, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.25, 4.25, 1.0 / 3.0, 1, false, false, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.50, 4.50, 2.0 / 3.0, 2, false, false, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.75, 4.75, 3.0 / 3.0, 3, false, true, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
	}

	i := 0
	ctx.WSeq(beat.Seq(1, 2, 3, 4), func(ctx context.Context) {
		ctx.WRng(beat.RngStep(0, 1, 4), func(ctx context.Context) {
			assertContextValues(t, ctx, wantContextValues[i])
			i++
		})
	})
}

func TestSeqSeqRng(t *testing.T) {
	ctx := context.NewBase()

	wantContextValues := []contextValues{
		{0.00, 11.00, 0.0 / 3.0, 0, true, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.25, 11.25, 1.0 / 3.0, 1, false, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.50, 11.50, 2.0 / 3.0, 2, false, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.75, 11.75, 3.0 / 3.0, 3, false, true, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.00, 12.00, 0.0 / 3.0, 0, true, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.25, 12.25, 1.0 / 3.0, 1, false, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.50, 12.50, 2.0 / 3.0, 2, false, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.75, 12.75, 3.0 / 3.0, 3, false, true, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.00, 13.00, 0.0 / 3.0, 0, true, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.25, 13.25, 1.0 / 3.0, 1, false, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.50, 13.50, 2.0 / 3.0, 2, false, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.75, 13.75, 3.0 / 3.0, 3, false, true, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.00, 14.00, 0.0 / 3.0, 0, true, false, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.25, 14.25, 1.0 / 3.0, 1, false, false, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.50, 14.50, 2.0 / 3.0, 2, false, false, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.75, 14.75, 3.0 / 3.0, 3, false, true, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.00, 21.00, 0.0 / 3.0, 0, true, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.25, 21.25, 1.0 / 3.0, 1, false, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.50, 21.50, 2.0 / 3.0, 2, false, false, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.75, 21.75, 3.0 / 3.0, 3, false, true, 1, 0.0 / 3.0, 0, 4, 2, 1, 0, 1, true, false},
		{0.00, 22.00, 0.0 / 3.0, 0, true, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.25, 22.25, 1.0 / 3.0, 1, false, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.50, 22.50, 2.0 / 3.0, 2, false, false, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.75, 22.75, 3.0 / 3.0, 3, false, true, 2, 1.0 / 3.0, 1, 4, 3, 1, 1, 1, false, false},
		{0.00, 23.00, 0.0 / 3.0, 0, true, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.25, 23.25, 1.0 / 3.0, 1, false, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.50, 23.50, 2.0 / 3.0, 2, false, false, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.75, 23.75, 3.0 / 3.0, 3, false, true, 3, 2.0 / 3.0, 2, 4, 4, 1, 2, 1, false, false},
		{0.00, 24.00, 0.0 / 3.0, 0, true, false, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.25, 24.25, 1.0 / 3.0, 1, false, false, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.50, 24.50, 2.0 / 3.0, 2, false, false, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
		{0.75, 24.75, 3.0 / 3.0, 3, false, true, 4, 3.0 / 3.0, 3, 4, 5, 1, 3, 1, false, true},
	}

	i := 0
	ctx.WSeq(beat.Seq(10, 20), func(ctx context.Context) {
		ctx.WSeq(beat.Seq(1, 2, 3, 4), func(ctx context.Context) {
			ctx.WRng(beat.RngInterval(0, 1, 4), func(ctx context.Context) {
				assertContextValues(t, ctx, wantContextValues[i])
				i++
			})
		})
	})
}

func TestEvents(t *testing.T) {
	ctx := context.NewBase()

	ctx.WSeq(beat.Seq(1, 2, 3, 4), func(ctx context.Context) {
		ctx.WRng(beat.RngInterval(0, 1, 10), func(ctx context.Context) {
			evt.Basic(ctx)
		})
	})

	if len(*ctx.Events()) != 4*10 {
		t.Errorf("got %d events, want %d", len(*ctx.Events()), 4*10)
	}
}

func assertContextValues(t *testing.T, ctx context.Context, want contextValues) {
	t.Helper()
	got := valuesFromContext(ctx)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatal(diff)
	}
}
