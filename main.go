package main

import (
	"os"
	. "raytracer/material"
	"raytracer/scenes"
	"runtime/pprof"

	. "raytracer/objects"
)

func main() {

	// Start CPU profiling.
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}

	defer pprof.StopCPUProfile()

	//MAIN CODE

	var world Hittable_list
	var cam Camera

	// world, cam = scenes.RandomSpheres()
	// world, cam = scenes.Texturedspheres()
	// // world = scenes.Earth()

	// world, cam = scenes.Quads()

	// world, cam = scenes.Boxes()

	world, cam = scenes.Meshes()

	bvh := NewBvh(world.Objects)
	cam.RenderMultithreaded(&bvh)

	// END MAIN CODE

	// Generate heap profile
	h, err := os.Create("heap.prof")
	if err != nil {
		panic(err)
	}
	defer h.Close()

	if err := pprof.WriteHeapProfile(h); err != nil {
		panic(err)
	}

}
