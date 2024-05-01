package material

import (
	. "raytracer/common"
)

type Metal struct {
	Albedo Color
	Fuzz   float64
}

func NewMetal(c Color, f float64) Metal {
	return Metal{Albedo: c, Fuzz: f}
}

func (m *Metal) Scatter(r *Ray, rec *Hit_record, attenuation *Color, scattered *Ray) bool {
	reflected := Reflect(Unit_vector(r.Direction), rec.Normal)
	*scattered = NewRay(rec.P, reflected.Add(Random_unit_vector().Mult(m.Fuzz)))
	*attenuation = m.Albedo
	return true

}

func (m *Metal) Emitted(u float64, v float64, p *Point3) Color {
	return NewColor(0, 0, 0)
}
