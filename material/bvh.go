package material

import (
	. "raytracer/common"
	"sort"
)

type Bvh struct {
	left  Hittable
	right Hittable
	bbox  Aabb
}

func NewBvh(objects []Hittable) Bvh {
	return NewBvhNode(objects, 0, len(objects))
}

func NewBvhNode(objects []Hittable, start int, end int) Bvh {

	bvh := Bvh{}
	bvh.bbox = Aabb{}

	for object_index := start; object_index < end; object_index++ {
		bvh.bbox = Merge(bvh.bbox, objects[object_index].Bounding_box())
	}
	axis := bvh.bbox.Longest_axis()

	var comparator func(a Hittable, b Hittable) bool

	if axis == 0 {
		comparator = bvh.box_x_compare
	} else if axis == 1 {
		comparator = bvh.box_y_compare
	} else {
		comparator = bvh.box_z_compare
	}

	object_span := end - start

	if object_span == 1 {
		bvh.left = objects[start]
		bvh.right = objects[start]
	} else if object_span == 2 {
		bvh.left = objects[start]
		bvh.right = objects[start+1]
	} else {

		sort.Slice(objects[start:end], func(i, j int) bool {
			return comparator(objects[start+i], objects[start+j])
		})

		mid := start + object_span/2
		l := NewBvhNode(objects, start, mid)
		r := NewBvhNode(objects, mid, end)
		bvh.left = &l
		bvh.right = &r

	}

	// fmt.Print("\nleft: ", reflect.TypeOf(bvh.left), bvh.left, "\nright: ", reflect.TypeOf(bvh.right), bvh.right, "\n ~~~~~~~~~~~~~~~~~~\n\n")
	return bvh

}

func (bvh *Bvh) Bounding_box() Aabb {
	return bvh.bbox
}

func (bvh *Bvh) Hit(r *Ray, ray_t Interval, rec *Hit_record) bool {

	if !bvh.bbox.Hit(r, &ray_t) {
		return false
	}

	hit_left := bvh.left.Hit(r, ray_t, rec)

	max := rec.T
	if !hit_left {
		max = ray_t.Max
	}

	hit_right := bvh.right.Hit(r, NewInterval(ray_t.Min, max), rec)

	return hit_left || hit_right

}

func (bvh *Bvh) box_compare(a Hittable, b Hittable, axis_index int) bool {

	a_box := a.Bounding_box()
	b_box := b.Bounding_box()
	a_axis_interval := a_box.Axis_interval(axis_index)
	b_axis_interval := b_box.Axis_interval(axis_index)

	return a_axis_interval.Min < b_axis_interval.Max
}

func (bvh *Bvh) box_x_compare(a Hittable, b Hittable) bool {
	return bvh.box_compare(a, b, 0)
}

func (bvh *Bvh) box_y_compare(a Hittable, b Hittable) bool {
	return bvh.box_compare(a, b, 1)
}

func (bvh *Bvh) box_z_compare(a Hittable, b Hittable) bool {
	return bvh.box_compare(a, b, 2)
}
