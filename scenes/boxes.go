package scenes

import (
	// "fmt"

	. "raytracer/common"
	. "raytracer/material"
	. "raytracer/objects"
)

func Boxes() (Hittable_list, Camera) {

	var world Hittable_list

	cam := NewCamera()

	cam.Aspect_ratio = 1.0
	cam.Image_width = 400
	cam.Sample_per_pixel = 50
	cam.Max_depth = 50
	cam.Background = NewColor(0.0, 0.0, 0.0)

	cam.Vfov = 40
	cam.Look_from = NewPoint3(278, 278, -800)
	cam.Look_at = NewPoint3(278, 278, 0)
	cam.Vup = NewVec3(0, 1, 0)

	cam.Log_scanlines = true

	cam.Defocus_angle = 0

	red := NewLambertian(NewColor(0.65, 0.05, 0.05))
	white := NewLambertian(NewColor(0.73, 0.73, 0.73))
	green := NewLambertian(NewColor(0.12, 0.45, 0.15))
	light_mat := NewDiffuse_light(NewColor(15, 15, 15))

	left := NewQuad(NewPoint3(555, 0, 0), NewVec3(0, 555, 0), NewVec3(0, 0, 555), &green)
	right := NewQuad(NewPoint3(0, 0, 0), NewVec3(0, 555, 0), NewVec3(0, 0, 555), &red)
	top := NewQuad(NewPoint3(555, 555, 555), NewVec3(-555, 0, 0), NewVec3(0, 0, -555), &white)
	back := NewQuad(NewPoint3(0, 0, 555), NewVec3(555, 0, 0), NewVec3(0, 555, 0), &white)
	bottom := NewQuad(NewPoint3(0, 0, 0), NewVec3(555, 0, 0), NewVec3(0, 0, 555), &white)
	light := NewQuad(NewPoint3(343, 554, 332), NewVec3(-130, 0, 0), NewVec3(0, 0, -105), &light_mat)

	box1 := NewBox(NewPoint3(130, 0, 65), NewPoint3(295, 165, 230), &white)
	box2 := NewBox(NewPoint3(265, 0, 295), NewPoint3(430, 330, 460), &white)

	world.Add(&box1)
	world.Add(&box2)
	world.Add(&left)
	world.Add(&right)
	world.Add(&top)
	world.Add(&bottom)
	world.Add(&back)
	world.Add(&light)

	// fmt.Print(world)

	return world, cam

}
