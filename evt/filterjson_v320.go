package evt

import (
	"math/rand"

	"github.com/shasderias/iris/beatsaber"
)

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
		Limit:        beatsaber.Float64(f.Limit),
		LimitAffects: int(f.LimitAffects),
	}
}
