package evt

import (
	"math/rand"

	"github.com/shasderias/iris/beatsaber"
)

type IndexFilter struct {
	IndexFilterType IndexFilterType
	ParamP, ParamT  int
	Reverse         bool
	Chunks          int
	Order           FilterOrder
	Seed            int
	Limit           float64
	LimitAffects    LimitAffects
}

type indexFilterFuncOption func(*IndexFilter)

func (fo indexFilterFuncOption) ApplyIndexFilter(f *IndexFilter) { fo(f) }

func OChunks(chunks int) IndexFilterOption {
	return indexFilterFuncOption(func(f *IndexFilter) { f.Chunks = chunks })
}

type IndexFilterType int

const (
	IndexFilterTypeSections      IndexFilterType = 1
	IndexFilterTypeStepAndOffset IndexFilterType = 2
)

type IndexFilterOption interface {
	ApplyIndexFilter(*IndexFilter)
}

func (f IndexFilter) ApplyRotationEventBox(b *RotationEventBox) { b.IndexFilter = f }
func (f IndexFilter) ApplyColorEventBox(b *ColorEventBox)       { b.IndexFilter = f }

func OSectionFilter(splitCount, targetSection int, options ...IndexFilterOption) IndexFilter {
	f := IndexFilter{
		IndexFilterType: IndexFilterTypeSections,
		ParamP:          splitCount,
		ParamT:          targetSection,
	}
	for _, opt := range options {
		opt.ApplyIndexFilter(&f)
	}
	return f
}

func OStepAndOffsetFilter(offset, step int, options ...IndexFilterOption) IndexFilter {
	f := IndexFilter{
		IndexFilterType: IndexFilterTypeStepAndOffset,
		ParamP:          offset,
		ParamT:          step,
	}
	for _, opt := range options {
		opt.ApplyIndexFilter(&f)
	}
	return f
}

func OReverse(reverse bool) IndexFilterOption {
	return indexFilterFuncOption(func(f *IndexFilter) { f.Reverse = reverse })
}

type FilterOrder int

const (
	FilterOrderInOrder             FilterOrder = 0
	FilterOrderRandom              FilterOrder = 2
	FilterOrderRandomStartingIndex FilterOrder = 3
)

const SeedRand = -1

func OOrder(order FilterOrder, seed int) IndexFilterOption {
	return IndexFilterOption(indexFilterFuncOption(func(f *IndexFilter) {
		f.Order = order
		f.Seed = seed
	}))
}

type LimitAffects int

const (
	LimitAffectsLightOnly               LimitAffects = 0
	LimitAffectsDuration                LimitAffects = 1
	LimitAffectsDistribution            LimitAffects = 2
	LimitAffectsDurationAndDistribution LimitAffects = 3
)

func OLimit(pct float64, affects LimitAffects) IndexFilterOption {
	return indexFilterFuncOption(func(f *IndexFilter) {
		f.Limit = pct
		f.LimitAffects = affects
	})
}

func (f IndexFilter) IndexFilterV300() beatsaber.IndexFilterV300 {
	return beatsaber.IndexFilterV300{
		FilterType: int(f.IndexFilterType),
		ParamP:     f.ParamP,
		ParamT:     f.ParamT,
		Reverse:    boolToIntBool(f.Reverse),
	}
}

func (f IndexFilter) IndexFilterV320() beatsaber.IndexFilterV320 {
	if f.Seed == SeedRand {
		f.Seed = rand.Intn(100000)
	}
	return beatsaber.IndexFilterV320{
		FilterType:   int(f.IndexFilterType),
		ParamP:       f.ParamP,
		ParamT:       f.ParamT,
		Reverse:      boolToIntBool(f.Reverse),
		Chunks:       f.Chunks,
		Order:        int(f.Order),
		Seed:         f.Seed,
		Limit:        f.Limit,
		LimitAffects: int(f.LimitAffects),
	}
}
