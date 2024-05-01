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

func (l *Lambertian) Emitted(u float64, v float64, p *Point3) Color {
	return NewColor(0, 0, 0)
}

type Diffuse_light struct {
	tex       Texture
	intensity float64
}

func (l *Diffuse_light) Emitted(u float64, v float64, p *Point3) Color {
	return l.tex.Value(u, v, p).Mult(l.intensity)
}
func (l *Diffuse_light) Scatter(r *Ray, rec *Hit_record, attenuation *Color, scattered *Ray) bool {
	return false
}

func NewDiffuse_light(c Color) Diffuse_light {
	// return Lambertian{Albedo: c}
	tex := NewSolid_color(&c)
	return Diffuse_light{tex: &tex, intensity: 1}
}

func NewTexturedDiffuse_Light(tex Texture, intensity float64) Diffuse_light {
	return Diffuse_light{tex: tex, intensity: intensity}
}
