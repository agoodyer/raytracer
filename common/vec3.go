package common

import (
	"math"
)

type Vec3 struct {
	e [3]float64
}

type Point3 = Vec3

func (v Vec3) X() float64 {
	return v.e[0]
}

func (v Vec3) Y() float64 {
	return v.e[1]
}

func (v Vec3) Z() float64 {
	return v.e[2]
}

func NewVec3(x float64, y float64, z float64) Vec3 {
	return Vec3{e: [3]float64{x, y, z}}
}

func NewPoint3(x float64, y float64, z float64) Vec3 {
	return Point3{e: [3]float64{x, y, z}}
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return NewVec3(v1.e[0]+v2.e[0], v1.e[1]+v2.e[1], v1.e[2]+v2.e[2])
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return NewVec3(v1.e[0]-v2.e[0], v1.e[1]-v2.e[1], v1.e[2]-v2.e[2])
}

func (v Vec3) Mult(t float64) Vec3 {
	return NewVec3(v.e[0]*t, v.e[1]*t, v.e[2]*t)
}

func (v Vec3) Div(t float64) Vec3 {
	return v.Mult(1 / t)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Length_squared())
}

func (v Vec3) Length_squared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func Dot(v1 Vec3, v2 Vec3) float64 {
	return (v1.e[0]*v2.e[0] +
		v1.e[1]*v2.e[1] +
		v1.e[2]*v2.e[2])
}

func Cross(v1 Vec3, v2 Vec3) Vec3 {
	return NewVec3(
		v1.e[1]*v2.e[2]-v1.e[2]*v2.e[1],
		v1.e[2]*v2.e[0]-v1.e[0]*v2.e[2],
		v1.e[0]*v2.e[1]-v1.e[1]*v2.e[0],
	)
}

func (v Vec3) Clone() Vec3 {
	return NewVec3(v.X(), v.Y(), v.Z())
}

func Unit_vector(v Vec3) Vec3 {
	v2 := v.Clone()
	v2.Div(v.Length())
	return v2
}
