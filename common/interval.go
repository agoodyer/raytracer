package common

import "math"

type Interval struct {
	Min float64
	Max float64
}

var empty Interval = Interval{Min: Infinity, Max: -Infinity}
var universe Interval = Interval{Min: -Infinity, Max: Infinity}

func NewInterval(min float64, max float64) Interval {
	return Interval{Min: min, Max: max}
}

func MergeInterval(a Interval, b Interval) Interval {
	min := math.Min(a.Min, b.Min)
	max := math.Max(a.Max, b.Max)

	return Interval{Min: min, Max: max}

}

func (i *Interval) Contains(x float64) bool {
	return i.Min <= x && x <= i.Max
}

func (i *Interval) Surrounds(x float64) bool {
	return i.Min < x && x < i.Max
}

func (i *Interval) Clamp(x float64) float64 {
	if x < i.Min {
		return i.Min
	}
	if x > i.Max {
		return i.Max
	}
	return x
}

func (i *Interval) Expand(delta float64) Interval {
	padding := delta / 2
	return NewInterval(i.Min-padding, i.Max+padding)
}
