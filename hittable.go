package main

type hit_record struct {
	p          point3
	normal     vec3
	t          float64
	front_face bool
}

type hittable interface {
	hit(r ray, ray_tmin float64, ray_tmax float64, rec *hit_record) bool
}

func (h *hit_record) set_face_normal(r ray, outward_normal vec3) {
	h.front_face = dot(r.direction, outward_normal) < 0
	if h.front_face {
		h.normal = outward_normal
	} else {
		h.normal = Vec3(0, 0, 0).sub(outward_normal)
	}

	// logger := log.New(os.Stderr, "", 0)
	// logger.Print(h.normal)

}

type hittable_list struct {
	objects []hittable
}

func (l *hittable_list) add(h hittable) {
	l.objects = append(l.objects, h)
}

func (l *hittable_list) clear() {
	l.objects = l.objects[:0]
}

func (l *hittable_list) hit(r ray, ray_tmin float64, ray_tmax float64, rec *hit_record) bool {

	var temp_rec hit_record
	hit_anything := false
	closest_so_far := ray_tmax

	for _, object := range l.objects {

		if object.hit(r, ray_tmin, closest_so_far, &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.t
			*rec = temp_rec
		}

	}

	return hit_anything

}
