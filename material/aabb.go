package material

import (
	. "raytracer/common"
)

type Aabb struct {
	x, y, z Interval
}

func NewAabb(x Interval, y Interval, z Interval) Aabb {
	return Aabb{x: x, y: y, z: z}
}

func Merge(box0 Aabb, box1 Aabb) Aabb {
	x := MergeInterval(box0.x, box1.x)
	y := MergeInterval(box0.y, box1.y)
	z := MergeInterval(box0.z, box1.z)
	return Aabb{x: x, y: y, z: z}
}

func NewAabbFromPoints(a Point3, b Point3) Aabb {

	var x Interval
	var y Interval
	var z Interval

	if a.X() <= b.X() {
		x = NewInterval(a.X(), b.X())
	} else {
		x = NewInterval(b.X(), a.X())
	}

	if a.Y() <= b.Y() {
		y = NewInterval(a.Y(), b.Y())
	} else {
		y = NewInterval(b.Y(), a.Y())
	}

	if a.Z() <= b.Z() {
		z = NewInterval(a.Z(), b.Z())
	} else {
		z = NewInterval(b.Z(), a.Z())
	}

	// logger := log.New(os.Stderr, "", 0)

	// logger.Print(x, y, z)

	return Aabb{x: x, y: y, z: z}

}

func (a *Aabb) Axis_interval(n int) Interval {

	if n == 1 {
		return a.y
	}
	if n == 2 {
		return a.z
	}
	return a.x
}

func (a *Aabb) Hit(r *Ray, ray_t *Interval) bool {

	// logger := log.New(os.Stderr, "", 0)
	// logger.Print(r.Direction)

	ray_orig := r.Origin
	ray_dir := r.Direction

	for axis := 0; axis < 3; axis++ {

		ax := a.Axis_interval(axis)
		adinv := 1.0 / ray_dir.XYZ()[axis]

		t0 := (ax.Min - ray_orig.XYZ()[axis]) * adinv
		t1 := (ax.Max - ray_orig.XYZ()[axis]) * adinv

		if t0 < t1 {
			if t0 > ray_t.Min {
				ray_t.Min = t0
			}
			if t1 < ray_t.Max {
				ray_t.Max = t1
			}

		} else {
			if t1 > ray_t.Min {
				ray_t.Min = t1
			}
			if t0 < ray_t.Max {
				ray_t.Max = t0
			}

		}

		if ray_t.Max <= ray_t.Min {
			return false
		}

	}

	return true

}
