package material

import (
	. "raytracer/common"
)

type Metal struct {
	Albedo Color
}

func NewMetal(c Color) Metal {
	return Metal{Albedo: c}
}

func (m *Metal) Scatter(r *Ray, rec *Hit_record, attenuation *Color, scattered *Ray) bool {
	reflected := Reflect(Unit_vector(r.Direction), rec.Normal)
	*scattered = NewRay(rec.P, reflected)
	*attenuation = m.Albedo
	return true

}
