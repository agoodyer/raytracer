package main

import (
	. "raytracer/common"
	. "raytracer/material"
	. "raytracer/objects"
)

func main() {

	var world Hittable_list

	material_ground := NewLambertian(NewColor(0.8, 0.8, 0.0))
	material_center := NewLambertian(NewColor(0.1, 0.2, 0.5))
	material_left := NewDielectric(1.5)
	material_right := NewMetal(NewColor(0.8, 0.6, 0.2), 0.0)

	world.Add(&Sphere{Center: NewPoint3(0, -100.5, -1), Radius: 100, Mat: &material_ground})
	world.Add(&Sphere{Center: NewPoint3(0, 0, -1), Radius: 0.5, Mat: &material_center})
	world.Add(&Sphere{Center: NewPoint3(-1, 0, -1), Radius: 0.5, Mat: &material_left})
	world.Add(&Sphere{Center: NewPoint3(-1, 0, -1), Radius: -0.3, Mat: &material_left})
	world.Add(&Sphere{Center: NewPoint3(1, -0, -1), Radius: 0.5, Mat: &material_right})

	var cam Camera

	cam.Aspect_ratio = 16.0 / 9.0
	cam.Image_width = 300
	cam.Render(&world)

}
