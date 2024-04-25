package scenes

import (
	"math/rand"
	. "raytracer/common"
	. "raytracer/material"
	. "raytracer/objects"
)

func RandomSpheres() Hittable_list {

	var world Hittable_list

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {

			choose_mat := rand.Float64()
			center := NewPoint3(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if center.Sub(NewPoint3(4, 0.2, 0)).Length() > 0.9 {
				// var sphere_material Material

				if choose_mat < 0.8 {
					//diffuse
					var albedo Color = ComponentMultiply(RandomClampedVector(), RandomClampedVector())
					sphere_material := NewLambertian(albedo)

					sp := NewSphere(center, 0.2, &sphere_material)
					world.Add(&sp)
					// world.Add(&Sphere{Center: center, Radius: 0.2, Mat: &sphere_material})
				} else if choose_mat < 0.95 {
					var albedo Color = NewColor(0.5, 0.5, 0.5).Add(RandomClampedVector().Mult(0.499))
					fuzz := Random_float(0.5, 1)
					sphere_material := NewMetal(albedo, fuzz)
					sp := NewSphere(center, 0.2, &sphere_material)
					world.Add(&sp)

					// world.Add(&Sphere{Center: center, Radius: 0.2, Mat: &sphere_material})

				} else {
					sphere_material := NewDielectric(1.5)
					// world.Add(&Sphere{Center: center, Radius: 0.2, Mat: &sphere_material})
					sp := NewSphere(center, 0.2, &sphere_material)
					world.Add(&sp)

				}
			}

		}
	}

	material1 := NewDielectric(1.5)
	material2 := NewLambertian(NewColor(0.4, 0.2, 0.1))
	material3 := NewMetal(NewColor(0.7, 0.6, 0.5), 0.0)

	s1 := NewSphere(NewPoint3(0, 1, 0), 1.0, &material1)
	s2 := NewSphere(NewPoint3(-4, 1, 0), 1.0, &material2)
	s3 := NewSphere(NewPoint3(4, 1, 0), 1.0, &material3)

	world.Add(&s1)
	world.Add(&s2)
	world.Add(&s3)

	ground_material := NewLambertian(NewColor(0.5, 0.5, 0.5))
	ground := NewSphere(NewPoint3(0, -1000, 0), 1000, &ground_material)
	world.Add(&ground)

	return world

}