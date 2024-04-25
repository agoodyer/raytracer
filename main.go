package main

import (
	"fmt"
	. "raytracer/material"
	. "raytracer/objects"
	"raytracer/scenes"
)

func main() {

	var world Hittable_list
	var cam Camera

	cam.Aspect_ratio = 16.0 / 9.0
	cam.Image_width = 400

	world = scenes.RandomSpheres()

	bvh := NewBvh(world.Objects)
	fmt.Print(bvh)

	cam.Render(&bvh)

}
