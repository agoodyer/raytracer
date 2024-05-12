package scenes

import (
	. "raytracer/common"
	. "raytracer/material"
	. "raytracer/objects"
)

func Texturedspheres() (Hittable_list, Camera) {

	var world Hittable_list

	c := NewCamera()
	c.Aspect_ratio = 16.0 / 9.0
	c.Image_width = 1920

	c.Sample_per_pixel = 250 //250
	c.Max_depth = 30         //50

	c.Vfov = 68
	c.Look_from = NewPoint3(14, 12, -5)
	c.Look_at = NewPoint3(0, -2, 0)
	c.Vup = NewVec3(0, 1, 0)

	c.Defocus_angle = 0.0
	c.Focus_dist = 10.0

	c.Background = NewColor(0.0, 0.0, 0.0085)

	c.Log_scanlines = true

	// c1 := NewColor(0.2, 0.3, 0.1)
	// c2 := NewColor(0.9, 0.9, 0.9)
	// checker := NewChecker_texture(0.32, &c1, &c2)

	// checkerLambertian := NewTexturedLambertian(&checker)

	// material3 := NewMetal(NewColor(0.7, 0.6, 0.5), 0.0)

	earth_tex := NewImage_texture("assets/earthmap.jpg")
	earth_surface := NewTexturedLambertian(&earth_tex)

	moon_tex := NewImage_texture("assets/moon.jpg")
	moon_surface := NewTexturedLambertian(&moon_tex)

	// sky_tex := NewImage_texture("assets/expsky2.png")
	// sky_surface := NewTexturedDiffuse_Light(&sky_tex, 0.4)

	sun_surface := NewDiffuse_light(NewColor(18, 18, 18))

	s1 := NewSphere(NewPoint3(0, -10, 0), 8.0, &earth_surface)
	s2 := NewSphere(NewPoint3(0, 6, -5), 2.0, &moon_surface)

	s3 := NewSphere(NewPoint3(1400, 1600, 3400), 2000, &sun_surface)

	// q1 := NewQuad(NewPoint3(-3, -2, 5), NewVec3(0, 0, -4), NewVec3(0, 4, 0), &earth_surface)

	// sky := NewSphere(NewPoint3(10000, 10000, -100000), 3000000, &sky_surface)

	world.Add(&s1)
	world.Add(&s2)
	world.Add(&s3)
	// world.Add(&sky)
	// world.Add(&q1)

	return world, c

}
