package material

import (
	. "raytracer/common"
)

type Material interface {
	Scatter(r *Ray, rec *Hit_record, attenuation *Color, scattered *Ray) bool
	Emitted(u float64, v float64, p *Point3) Color
}
