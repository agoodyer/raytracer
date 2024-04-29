package material

import (
	. "raytracer/common"
)

type Lambertian struct {
	// Albedo Color
	tex Texture
}

func NewLambertian(c Color) Lambertian {
	// return Lambertian{Albedo: c}
	tex := NewSolid_color(&c)
	return Lambertian{tex: &tex}
}

func NewTexturedLambertian(tex Texture) Lambertian {
	return Lambertian{tex: tex}
}

func (l *Lambertian) Scatter(r *Ray, rec *Hit_record, attenuation *Color, scattered *Ray) bool {
	scatter_direction := rec.Normal.Add(Random_unit_vector())

	if scatter_direction.Near_zero() {
		scatter_direction = rec.Normal
	}

	*scattered = NewRay(rec.P, scatter_direction)

	// *attenuation = l.Albedo

	*attenuation = l.tex.Value(rec.U, rec.V, &rec.P)

	return true

}
