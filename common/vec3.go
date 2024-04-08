package common

import (
	"math"
	"math/rand"
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

var NullVector Vec3 = NewVec3(0, 0, 0)

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

func ComponentMultiply(v1 Vec3, v2 Vec3) Vec3 {
	return NewVec3(v1.e[0]*v2.e[0], v1.e[1]*v2.e[1], v1.e[2]*v2.e[2])
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
	return v.Div(v.Length())
}

func RandomClampedVector() Vec3 {
	return NewVec3(rand.Float64(), rand.Float64(), rand.Float64())
}

func RandomVector(min float64, max float64) Vec3 {
	return NewVec3(random_float(min, max), random_float(min, max), random_float(min, max))
}

func Random_in_unit_sphere() Vec3 {
	for {
		p := RandomVector(-1, 1)
		if p.Length_squared() < 1 {
			return p
		}
	}
}

func Random_in_unit_disk() Vec3 {
	for {
		p := NewVec3(random_float(-1, 1), random_float(-1, 1), 0)
		if p.Length_squared() < 1 {
			return p
		}
	}
}

func Random_unit_vector() Vec3 {
	return Unit_vector(Random_in_unit_sphere())
}

func Random_on_hemisphere(normal Vec3) Vec3 {
	on_unit_sphere := Random_unit_vector()

	if Dot(on_unit_sphere, normal) > 0.0 {
		return on_unit_sphere
	} else {
		return NullVector.Sub(on_unit_sphere)
	}

}

func (v Vec3) Near_zero() bool {
	s := 1e-8
	return (math.Abs(v.e[0]) < s && math.Abs(v.e[1]) < s && math.Abs(v.e[2]) < s)
}

func Reflect(v Vec3, n Vec3) Vec3 {
	return v.Sub(n.Mult(Dot(v, n) * 2))
}

func Refract(uv Vec3, n Vec3, etai_over_etat float64) Vec3 {
	cos_theta := math.Min(Dot(NullVector.Sub(uv), n), 1.0)
	r_out_perp := uv.Add(n.Mult(cos_theta)).Mult(etai_over_etat)
	r_out_parallel := n.Mult(math.Sqrt(math.Abs(1.0-r_out_perp.Length_squared())) * -1)
	return r_out_parallel.Add(r_out_perp)
}
