package main

import (
	"math"
)

type vec3 struct {
	e [3]float64
}

type point3 = vec3

func (v vec3) x() float64 {
	return v.e[0]
}

func (v vec3) y() float64 {
	return v.e[1]
}

func (v vec3) z() float64 {
	return v.e[2]
}

func Vec3(x float64, y float64, z float64) vec3 {
	return vec3{e: [3]float64{x, y, z}}
}

func Point3(x float64, y float64, z float64) vec3 {
	return point3{e: [3]float64{x, y, z}}
}

func (v1 vec3) add(v2 vec3) vec3 {
	return Vec3(v1.e[0]+v2.e[0], v1.e[1]+v2.e[1], v1.e[2]+v2.e[2])
}

func (v1 vec3) sub(v2 vec3) vec3 {
	return Vec3(v1.e[0]-v2.e[0], v1.e[1]-v2.e[1], v1.e[2]-v2.e[2])
}

// func add(v1 vec3, v2 vec3) vec3 {

// 	v := v1.clone()
// 	v.add(v2)
// 	return v
// }
// func sub(v1 vec3, v2 vec3) vec3 {
// 	v := v1.clone()
// 	v.sub(v2)
// 	return v
// }

// func mult(v vec3, t float64) vec3 {
// 	v2 := v.clone()
// 	v2.mult(t)
// 	return v2
// }

// func div(v vec3, t float64) vec3 {
// 	return mult(v, 1/t)
// }

func (v vec3) mult(t float64) vec3 {
	return Vec3(v.e[0]*t, v.e[1]*t, v.e[2]*t)
}

func (v vec3) div(t float64) vec3 {
	return v.mult(1 / t)
}

func (v vec3) length() float64 {
	return math.Sqrt(v.length_squared())
}

func (v vec3) length_squared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func dot(v1 vec3, v2 vec3) float64 {
	return (v1.e[0]*v2.e[0] +
		v1.e[1]*v2.e[1] +
		v1.e[2]*v2.e[2])
}

func cross(v1 vec3, v2 vec3) vec3 {
	return Vec3(
		v1.e[1]*v2.e[2]-v1.e[2]*v2.e[1],
		v1.e[2]*v2.e[0]-v1.e[0]*v2.e[2],
		v1.e[0]*v2.e[1]-v1.e[1]*v2.e[0],
	)
}

func (v vec3) clone() vec3 {
	return Vec3(v.x(), v.y(), v.z())
}

func unit_vector(v vec3) vec3 {
	v2 := v.clone()
	v2.div(v.length())
	return v2
}
