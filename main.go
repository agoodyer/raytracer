package main

import (
	. "raytracer/material"
	. "raytracer/objects"
	"raytracer/scenes"
)

func main() {

	var world Hittable_list
	var cam Camera

	cam.Aspect_ratio = 16.0 / 9.0
	cam.Image_width = 500

	// world = scenes.RandomSpheres()
	world = scenes.Texturedspheres()
	// world = scenes.Earth()

	bvh := NewBvh(world.Objects)

	cam.Render(&bvh)

}
