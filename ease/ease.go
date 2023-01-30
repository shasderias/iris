// Package ease implements Robert Penner's easing functions.
//
// http://robertpenner.com/easing/
//
// See https://easings.net/ for visualizations of the easing functions.
package ease

import (
	"math"
)

const (
	pi = math.Pi
	// http: //void.heteml.jp/blog/archives/2014/05/easing_magicnumber.html
	c1 = 1.70158
	c2 = c1 * 1.525
	c3 = c1 + 1
	c4 = (2 * pi) / 3
	c5 = (2 * pi) / 4.5
)

var (
	pow  = math.Pow
	sqrt = math.Sqrt
	sin  = math.Sin
	cos  = math.Cos
)

type SetEasingable interface {
	SetEasing(easing Ing)
}
type R interface {
	Ease(x float64) float64
}

type Func func(x float64) float64

func bounceOut(x float64) float64 {
	const n1 = 7.5625
	const d1 = 2.75

	if x < 1/d1 {
		return n1 * x * x
	} else if x < 2/d1 {
		x -= 1.5 / d1
		return n1*x*x + 0.75
	} else if x < 2.5/d1 {
		x -= 2.25 / d1
		return n1*x*x + 0.9375
	} else {
		x -= 2.625 / d1
		return n1*x*x + 0.984375
	}
}

type Ing struct {
	Name string
	Func
}

var Nil = Ing{"", nil}

var (
	Linear = Ing{"easeLinear", func(x float64) float64 {
		return x
	}}
	Step = Ing{"easeStep", func(x float64) float64 {
		return math.Floor(x)
	}}
)

var (
	InQuad = Ing{"easeInQuad", func(x float64) float64 {
		return x * x
	}}
	OutQuad = Ing{"easeOutQuad", func(x float64) float64 {
		return 1 - (1-x)*(1-x)
	}}
	InOutQuad = Ing{"easeInOutQuad", func(x float64) float64 {
		if x < 0.5 {
			return 2 * x * x
		} else {
			return 1 - pow(-2*x+2, 2)/2
		}
	}}
)

var (
	InCubic = Ing{"easeInCubic", func(x float64) float64 {
		return x * x * x
	}}
	OutCubic = Ing{"easeOutCubic", func(x float64) float64 {
		return 1 - pow(1-x, 3)
	}}
	InOutCubic = Ing{"easeInOutCubic", func(x float64) float64 {
		if x < 0.5 {
			return 4 * x * x * x
		} else {
			return 1 - pow(-2*x+2, 3)/2
		}
	}}
)

var (
	InQuart = Ing{"easeInQuart", func(x float64) float64 {
		return x * x * x * x
	}}
	OutQuart = Ing{"easeOutQuart", func(x float64) float64 {
		return 1 - pow(1-x, 4)
	}}
	InOutQuart = Ing{"easeInOutQuart", func(x float64) float64 {
		if x < 0.5 {
			return 8 * x * x * x * x
		} else {
			return 1 - pow(-2*x+2, 4)/2
		}
	}}
)

var (
	InQuint = Ing{"easeInQuint", func(x float64) float64 {
		return x * x * x * x * x
	}}
	OutQuint = Ing{"easeOutQuint", func(x float64) float64 {
		return 1 - pow(1-x, 5)
	}}
	InOutQuint = Ing{"easeInOutQuint", func(x float64) float64 {
		if x < 0.5 {
			return 16 * x * x * x * x * x
		} else {
			return 1 - pow(-2*x+2, 5)/2
		}
	}}
)

var (
	InSin = Ing{"easeInSin", func(x float64) float64 {
		return 1 - cos((x*pi)/2)
	}}
	OutSin = Ing{"easeOutSin", func(x float64) float64 {
		return sin((x * pi) / 2)
	}}
	InOutSin = Ing{"easeInOutSin", func(x float64) float64 {
		return -(cos(pi*x) - 1) / 2
	}}
)

var (
	InExpo = Ing{"easeInExpo", func(x float64) float64 {
		if x == 0 {
			return 0
		} else {
			return pow(2, 10*x-10)
		}
	}}
	OutExpo = Ing{"easeOutExpo", func(x float64) float64 {
		if x == 1 {
			return 1
		} else {
			return 1 - pow(2, -10*x)
		}
	}}
	InOutExpo = Ing{"easeInOutExpo", func(x float64) float64 {
		switch {
		case x == 0:
			return 0
		case x == 1:
			return 1
		case x < 0.5:
			return pow(2, 20*x-10) / 2
		default:
			return (2 - pow(2, -20*x+10)) / 2
		}
	}}
)

var (
	InCirc = Ing{"easeInCirc", func(x float64) float64 {
		return 1 - sqrt(1-pow(x, 2))
	}}
	OutCirc = Ing{"easeOutCirc", func(x float64) float64 {
		return sqrt(1 - pow(x-1, 2))
	}}
	InOutCirc = Ing{"easeInOutCirc", func(x float64) float64 {
		if x < 0.5 {
			return (1 - sqrt(1-pow(2*x, 2))) / 2
		} else {
			return (sqrt(1-pow(-2*x+2, 2)) + 1) / 2
		}
	}}
)

var (
	InBack = Ing{"easeInBack", func(x float64) float64 {
		return c3*x*x*x - c1*x*x
	}}
	OutBack = Ing{"easeOutBack", func(x float64) float64 {
		return 1 + c3*pow(x-1, 3) + c1*pow(x-1, 2)
	}}
	InOutBack = Ing{"easeInOutBack", func(x float64) float64 {
		if x < 0.5 {
			return (pow(2*x, 2) * ((c2+1)*2*x - c2)) / 2
		} else {
			return (pow(2*x-2, 2)*((c2+1)*(x*2-2)+c2) + 2) / 2
		}
	}}
)

var (
	InElastic = Ing{"easeInElastic", func(x float64) float64 {
		switch {
		case x == 0:
			return 0
		case x == 1:
			return 1
		default:
			return -pow(2, 10*x-10) * sin((x*10-10.75)*c4)
		}
	}}
	OutElastic = Ing{"easeOutElastic", func(x float64) float64 {
		switch {
		case x == 0:
			return 0
		case x == 1:
			return 1
		default:
			return pow(2, -10*x)*sin((x*10-0.75)*c4) + 1
		}
	}}
	InOutElastic = Ing{"easeInOutElastic", func(x float64) float64 {
		switch {
		case x == 0:
			return 0
		case x == 1:
			return 1
		case x < 0.5:
			return -(pow(2, 20*x-10) * sin((20*x-11.125)*c5)) / 2
		default:
			return pow(2, -20*x+10)*sin((20*x-11.125)*c5)/2 + 1
		}
	}}
)

var (
	InBounce = Ing{"easeInBounce", func(x float64) float64 {
		return 1 - bounceOut(1-x)
	}}
	OutBounce = Ing{"easeOutBounce", func(x float64) float64 {
		return bounceOut(x)
	}}
	InOutBounce = Ing{"easeInOutBounce", func(x float64) float64 {
		if x < 0.5 {
			return (1 - bounceOut(1-2*x)) / 2
		} else {
			return (1 + bounceOut(2*x-1)) / 2
		}
	}}
)
