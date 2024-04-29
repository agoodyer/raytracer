package scenes

import (
	. "raytracer/common"
	. "raytracer/material"
	. "raytracer/objects"
)

func Texturedspheres() Hittable_list {

	var world Hittable_list

	// c1 := NewColor(0.2, 0.3, 0.1)
	// c2 := NewColor(0.9, 0.9, 0.9)
	// checker := NewChecker_texture(0.32, &c1, &c2)

	// checkerLambertian := NewTexturedLambertian(&checker)

	// material3 := NewMetal(NewColor(0.7, 0.6, 0.5), 0.0)

	earth_tex := NewImage_texture("assets/earthmap.jpg")
	earth_surface := NewTexturedLambertian(&earth_tex)

	moon_tex := NewImage_texture("assets/moon.jpg")
	moon_surface := NewTexturedLambertian(&moon_tex)

	s1 := NewSphere(NewPoint3(0, -10, 0), 10.0, &earth_surface)
	s2 := NewSphere(NewPoint3(0, 6, 0), 4.0, &moon_surface)

	world.Add(&s1)
	world.Add(&s2)

	return world

}
