package main

type ray struct {
	origin    point3
	direction vec3
}

func Ray(origin point3, direction vec3) ray {
	return ray{origin: origin, direction: direction}
}

func (r *ray) at(t float64) point3 {
	return r.origin.add(r.direction.mult(t))
}
