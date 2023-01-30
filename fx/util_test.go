package fx_test

import (
	"testing"

	"github.com/shasderias/iris/fx"
)

func TestCurve(t *testing.T) {
	testCases := []struct {
		name       string
		points     []fx.Point
		testPoints []struct {
			x, want float64
		}
	}{
		{
			name: "Sanity",
			points: []fx.Point{
				{Y: 0}, {Y: 1},
			},
			testPoints: []struct{ x, want float64 }{
				{0, 0}, {0.5, 0.5}, {1, 1},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			curve := fx.Curve(tt.points...)

			for _, tp := range tt.testPoints {
				got := curve(tp.x)
				if got != tp.want {
					t.Errorf("curve(%f) = %f, want %f", tp.x, got, tp.want)
				}
			}
		})
	}
}
