package main

import (
	. "raytracer/material"

	. "raytracer/objects"
	"raytracer/scenes"
)

func main() {

	var world Hittable_list
	var cam Camera

	// // world, cam = scenes.RandomSpheres()
	// // world, cam = scenes.Texturedspheres()
	// // world = scenes.Earth()

	// world, cam = scenes.Quads()

	// // world, cam = scenes.Boxes()

	world, cam = scenes.Meshes()

	bvh := NewBvh(world.Objects)

	cam.Render(&bvh)

}
