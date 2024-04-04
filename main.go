package main

import (
	"fmt"
	"log"
	"os"
)

// func hit_sphere(center point3, radius float64, r ray) float64 {
// 	// oc := sub(r.origin, center)
// 	oc := r.origin.sub(center)

// 	a := r.direction.length_squared()
// 	half_b := dot(oc, r.direction)
// 	c := oc.length_squared() - radius*radius
// 	discriminant := half_b*half_b - a*c
// 	if discriminant < 0 {
// 		return -1.0
// 	} else {
// 		return (-half_b - math.Sqrt(discriminant)) / a
// 	}
// }

// func ray_color(r ray) color {

// 	t := hit_sphere(Point3(0, 0, -1), 0.5, r)

// 	if t > 0.0 {
// 		// N := unit_vector(sub(r.at(t), Vec3(0, 0, -1)))
// 		N := unit_vector(r.at(t).sub(Vec3(0, 0, -1)))

// 		return Color(N.x()+1, N.y()+1, N.z()+1).mult(0.5)
// 		// return mult(Color(N.x()+1, N.y()+1, N.z()+1), 0.5)
// 	}

// 	unit_direction := unit_vector(r.direction)
// 	a := 0.5 * (unit_direction.y() + 1.0)
// 	// return add(mult(Color(1.0, 1.0, 1.0), (1.0-a)), mult(Color(0.5, 0.7, 1.0), a))
// 	return Color(1.0, 1.0, 1.0).mult(1.0 - a).add(Color(0.5, 0.7, 1.0).mult(a))

// 	// return Color(0, 0, 0)
// }

func ray_color(r ray, world hittable) color {
	var rec hit_record
	if world.hit(r, 0, infinity, &rec) {

		// logger := log.New(os.Stderr, "", 0)
		// logger.Print(rec.normal)

		return rec.normal.add(Color(1, 1, 1)).mult(0.5)
	}
	unit_direction := unit_vector(r.direction)
	a := 0.5 * (unit_direction.y() + 1.0)
	return Color(1.0, 1.0, 1.0).mult(1.0 - a).add(Color(0.5, 0.7, 1.0).mult(a))
}

// func ray_color(r ray, s sphere) color {
// 	var rec hit_record
// 	if s.hit(r, 0, infinity, &rec) {

// 		logger := log.New(os.Stderr, "", 0)
// 		logger.Print(rec.normal)

// 		return rec.normal.add(Color(1, 1, 1)).mult(0.5)
// 	}
// 	unit_direction := unit_vector(r.direction)
// 	a := 0.5 * (unit_direction.y() + 1.0)
// 	return Color(1.0, 1.0, 1.0).mult(1.0 - a).add(Color(0.5, 0.7, 1.0).mult(a))
// }

// import "fmt"
func main() {

	var world hittable_list

	//Image sizing
	aspect_ratio := 16.0 / 9.0
	image_width := 400

	//Calculate image height
	image_height := int(float64(image_width) / aspect_ratio)
	if image_height < 1 {
		image_height = 1
	}

	// s := sphere{center: Point3(0, 0, -1), radius: 0.5}

	world.add(&sphere{center: Point3(0, 0, -1), radius: 0.5})
	world.add(&sphere{center: Point3(0, -100.5, -1), radius: 100})

	//Camera
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * float64(image_width) / float64(image_height)
	camera_center := Point3(0, 0, 0)

	viewport_u := Vec3(viewport_width, 0, 0)
	viewport_v := Vec3(0, -viewport_height, 0)

	// pixel_delta_u := div(viewport_u, float64(image_width))
	// pixel_delta_v := div(viewport_v, float64(image_height))

	pixel_delta_u := viewport_u.div(float64(image_width))
	pixel_delta_v := viewport_v.div(float64(image_height))

	// viewport_upper_left := sub(sub(sub(camera_center, Vec3(0, 0, focal_length)), div(viewport_u, 2)), div(viewport_v, 2))
	viewport_upper_left := camera_center.sub(Vec3(0, 0, focal_length)).sub(viewport_u.div(2)).sub(viewport_v.div(2))

	// pixel00_loc := add(viewport_upper_left, mult(add(pixel_delta_u, pixel_delta_v), 0.5))

	pixel00_loc := viewport_upper_left.add(pixel_delta_u.add(pixel_delta_v).mult(0.5))

	logger := log.New(os.Stderr, "", 0)

	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)

	for j := 0; j < image_height; j++ {

		logger.Printf("Scanlines remaining: %d", image_height-j)

		for i := 0; i < image_width; i++ {

			// pixel_center := add(pixel00_loc,
			// 	add(mult(pixel_delta_u, float64(i)), mult(pixel_delta_v, float64(j))),
			// )

			pixel_center := pixel00_loc.add(pixel_delta_u.mult(float64(i)).add(pixel_delta_v.mult(float64(j))))

			// ray_direction := sub(pixel_center, camera_center)
			ray_direction := pixel_center.sub(camera_center)

			r := Ray(camera_center, ray_direction)
			// pixel_color := ray_color(r, sphere{center: Point3(0, 0, -1), radius: 0.5})
			pixel_color := ray_color(r, &world)
			write_color(pixel_color)

		}

	}

}
