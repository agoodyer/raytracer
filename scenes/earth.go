package scenes

import (
	. "raytracer/common"
	. "raytracer/material"
	. "raytracer/objects"
)

func Earth() (Hittable_list, Camera) {

	var world Hittable_list
	c := NewCamera()

	c.Aspect_ratio = 16.0 / 9.0
	c.Image_width = 1200

	c.Sample_per_pixel = 100 //250
	c.Max_depth = 33         //50

	c.Vfov = 68
	c.Look_from = NewPoint3(34, 42, -5)
	c.Look_at = NewPoint3(0, -2, 0)
	c.Vup = NewVec3(0, 1, 0)

	c.Defocus_angle = 0.0
	c.Focus_dist = 10.0

	c.Background = NewColor(0.7, 0.7, 0.7)

	c.Log_scanlines = true

	earth_tex := NewImage_texture("assets/earthmap.jpg")
	earth_surface := NewTexturedLambertian(&earth_tex)

	globe := NewSphere(NewPoint3(0, 0, 0), 2, &earth_surface)

	world.Add(&globe)

	// l := NewLambertian(NewColor(240.0/255, 230.0/255, 140.0/255))

	metal := NewMetal(NewColor(0.5, 0.5, 0.5), 0.3)
	gg := NewMeshFromFile("assets/vase.stl", &metal, 0.50)
	bvh := NewBvh(gg.Objects)

	abc := NewRotationY(bvh, 50)

	world.Add(&abc)

	sky_tex := NewImage_texture("assets/skysphere.jpg")
	sky_mat := NewTexturedDiffuse_Light(&sky_tex, 1.0)

	sky := NewSphere(NewPoint3(0, 0, 0), 100000, &sky_mat)

	world.Add(&sky)

	return world, c

}
