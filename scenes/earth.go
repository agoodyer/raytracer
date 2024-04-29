package scenes

import (
	. "raytracer/common"
	. "raytracer/material"
	. "raytracer/objects"
)

func Earth() Hittable_list {

	var world Hittable_list
	earth_tex := NewImage_texture("earthmap.jpg")
	earth_surface := NewTexturedLambertian(&earth_tex)

	globe := NewSphere(NewPoint3(0, 0, 0), 2, &earth_surface)

	world.Add(&globe)

	return world

}
