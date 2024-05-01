package material

import (
	"math"
	"math/rand"
	. "raytracer/common"
)

type Dielectric struct {
	ir float64
}

func NewDielectric(index_of_refraction float64) Dielectric {
	return Dielectric{ir: index_of_refraction}
}

func (d *Dielectric) Scatter(r *Ray, rec *Hit_record, attenuation *Color, scattered *Ray) bool {

	*attenuation = NewColor(1.0, 1.0, 1.0)
	var refraction_ratio float64

	if rec.Front_face {
		refraction_ratio = 1.0 / d.ir
	} else {
		refraction_ratio = d.ir
	}

	unit_direction := Unit_vector(r.Direction)

	cos_theta := math.Min(Dot(NullVector.Sub(unit_direction), rec.Normal), 1.0)
	sin_theta := math.Sqrt(1.0 - cos_theta*cos_theta)

	cannot_refract := refraction_ratio*sin_theta > 1.0

	var direction Vec3

	if cannot_refract || reflectance(cos_theta, refraction_ratio) > rand.Float64() {
		direction = Reflect(unit_direction, rec.Normal)
	} else {
		direction = Refract(unit_direction, rec.Normal, refraction_ratio)
	}

	// refracted := Refract(unit_direction, rec.Normal, refraction_ratio)

	*scattered = NewRay(rec.P, direction)
	return true

}

func reflectance(cosine float64, ref_idx float64) float64 {
	r0 := (1 - ref_idx) / (1 + ref_idx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

func (d *Dielectric) Emitted(u float64, v float64, p *Point3) Color {
	return NewColor(0, 0, 0)
}
