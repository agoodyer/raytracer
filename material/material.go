package material

import (
	. "raytracer/common"
)

type Material interface {
	Scatter(r *Ray, rec *Hit_record, attenuation *Color, scattered *Ray) bool
}
