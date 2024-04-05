package common

type Hit_record struct {
	P          Point3
	Normal     Vec3
	T          float64
	Front_face bool
}

type Hittable interface {
	Hit(r Ray, ray_tmin float64, ray_tmax float64, rec *Hit_record) bool
}

func (h *Hit_record) Set_face_normal(r Ray, outward_normal Vec3) {
	h.Front_face = Dot(r.Direction, outward_normal) < 0
	if h.Front_face {
		h.Normal = outward_normal
	} else {
		h.Normal = NewVec3(0, 0, 0).Sub(outward_normal)
	}

	// logger := log.New(os.Stderr, "", 0)
	// logger.Print(h.normal)

}

type Hittable_list struct {
	objects []Hittable
}

func (l *Hittable_list) Add(h Hittable) {
	l.objects = append(l.objects, h)
}

func (l *Hittable_list) clear() {
	l.objects = l.objects[:0]
}

func (l *Hittable_list) Hit(r Ray, ray_tmin float64, ray_tmax float64, rec *Hit_record) bool {

	var temp_rec Hit_record
	hit_anything := false
	closest_so_far := ray_tmax

	for _, object := range l.objects {

		if object.Hit(r, ray_tmin, closest_so_far, &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.T
			*rec = temp_rec
		}

	}

	return hit_anything

}