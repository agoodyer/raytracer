package common

type Interval struct {
	Min float64
	Max float64
}

var empty Interval = Interval{Min: Infinity, Max: -Infinity}
var universe Interval = Interval{Min: -Infinity, Max: Infinity}

func NewInterval(min float64, max float64) Interval {
	return Interval{Min: min, Max: max}
}

func (i *Interval) Contains(x float64) bool {
	return i.Min <= x && x <= i.Max
}

func (i *Interval) Surrounds(x float64) bool {
	return i.Min < x && x < i.Max
}
