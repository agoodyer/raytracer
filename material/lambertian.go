package material

import (
	. "raytracer/common"
)

type Lambertian struct {
	Albedo Color
}

func NewLambertian(c Color) Lambertian {
	return Lambertian{Albedo: c}
}

func (l *Lambertian) Scatter(r *Ray, rec *Hit_record, attenuation *Color, scattered *Ray) bool {
	scatter_direction := rec.Normal.Add(Random_unit_vector())

	if scatter_direction.Near_zero() {
		scatter_direction = rec.Normal
	}

	*scattered = NewRay(rec.P, scatter_direction)

	*attenuation = l.Albedo
	return true

}
