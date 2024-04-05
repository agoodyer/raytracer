package main

import (
	. "raytracer/common"
	. "raytracer/objects"
)

func main() {

	var world Hittable_list

	world.Add(&Sphere{Center: NewPoint3(0, 0, -1), Radius: 0.5})
	world.Add(&Sphere{Center: NewPoint3(0, -100.5, -1), Radius: 100})

	var cam Camera

	cam.Aspect_ratio = 16.0 / 9.0
	cam.Image_width = 400
	cam.Render(&world)

}
