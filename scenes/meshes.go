package scenes

import (
	. "raytracer/common"
	. "raytracer/material"
	. "raytracer/objects"
)

func Meshes() (Hittable_list, Camera) {

	var world Hittable_list
	cam := NewCamera()

	cam.Aspect_ratio = 16.0 / 9.0
	cam.Image_width = 300
	cam.Sample_per_pixel = 15
	cam.Max_depth = 6

	cam.Vfov = 50
	cam.Look_from = NewPoint3(40, 42, 80)
	cam.Look_at = NewPoint3(0, 0, 0)
	cam.Vup = NewVec3(0, 1, 0)
	cam.Log_scanlines = true

	cam.Defocus_angle = 0

	cam.Background = NewColor(0.05, 0.05, 0.2)

	c1 := NewColor(0.2, 0.3, 0.1)
	c2 := NewColor(0.9, 0.9, 0.9)
	checker := NewChecker_texture(7.32, &c1, &c2)
	checkerLambertian := NewTexturedLambertian(&checker)

	// green := NewLambertian(NewColor(0.9, 0.9, 0.8))
	gg := NewMeshFromFile("assets/vase.stl", &checkerLambertian, 1.6)

	bvh := NewBvh(gg.Objects)

	world.Add(&bvh)

	sun_surface := NewDiffuse_light(NewColor(6, 6, 6))
	s3 := NewSphere(NewPoint3(-25, 77, 4), 40, &sun_surface)

	s4 := NewSphere(NewPoint3(150, 107, 4), 40, &sun_surface)

	world.Add(&s3)
	world.Add(&s4)

	// s := NewSphere(NewPoint3(3, 3, 3), 1, &sun_surface)

	// world

	return world, cam

}
