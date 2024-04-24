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

// func NewBvh(l Hittable_list, start int, end int) Bvh {

// }

func NewBvh2(objects []Hittable, start int, end int) Bvh {

	bvh := Bvh{}

	axis := Random_int(0, 2)

	var comparator func(a Hittable, b Hittable) bool

	// logger := log.New(os.Stderr, "", 0)
	// logger.Print(objects)

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

		// logger.Print(objects)
		sort.Slice(objects, func(i, j int) bool {
			return comparator(objects[i], objects[j])
		})

		// logger.Print(objects)

		mid := start + object_span/2
		l := NewBvh2(objects, start, mid)
		r := NewBvh2(objects, mid, end)
		bvh.left = &l
		bvh.right = &r

	}

	// logger.Print("left: ", reflect.TypeOf(bvh.left), bvh.left, "\nright: ", reflect.TypeOf(bvh.right), bvh.right, "\n ~~~~~~~~~~~~~~~~~~")

	bvh.bbox = Merge(bvh.left.Bounding_box(), bvh.right.Bounding_box())
	return bvh

}

func (bvh *Bvh) Bounding_box() Aabb {
	return bvh.bbox
}

func (bvh *Bvh) Hit(r Ray, ray_t Interval, rec *Hit_record) bool {
	// logger := log.New(os.Stderr, "", 0)

	// if bvh.bbox.Hit(r, &ray_t) {
	// 	logger.Print("!")
	// }

	if !bvh.bbox.Hit(&r, &ray_t) {
		return false
	}

	hit_left := bvh.left.Hit(r, ray_t, rec)

	max := rec.T
	if !hit_left {
		max = ray_t.Max
	}

	hit_right := bvh.right.Hit(r, NewInterval(ray_t.Min, max), rec)

	// logger.Print(hit_left || hit_right)

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
