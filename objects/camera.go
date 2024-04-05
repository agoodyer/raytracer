package objects

import (
	"fmt"
	"log"
	"os"
	. "raytracer/common"
)

type Camera struct {
	Aspect_ratio float64
	Image_width  int
	// Sample_per_pixel int
	image_height  int
	center        Point3
	pixel00_loc   Point3
	pixel_delta_u Vec3
	pixel_delta_v Vec3
}

func (c *Camera) initialize() {

	// c.Aspect_ratio = 1.0
	// c.Image_width = 100

	//Calculate image height
	c.image_height = int(float64(c.Image_width) / c.Aspect_ratio)
	if c.image_height < 1 {
		c.image_height = 1
	}

	//Camera
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * float64(c.Image_width) / float64(c.image_height)
	c.center = NewPoint3(0, 0, 0)

	viewport_u := NewVec3(viewport_width, 0, 0)
	viewport_v := NewVec3(0, -viewport_height, 0)

	c.pixel_delta_u = viewport_u.Div(float64(c.Image_width))
	c.pixel_delta_v = viewport_v.Div(float64(c.image_height))

	viewport_upper_left := c.center.Sub(NewVec3(0, 0, focal_length)).Sub(viewport_u.Div(2)).Sub(viewport_v.Div(2))

	c.pixel00_loc = viewport_upper_left.Add(c.pixel_delta_u.Add(c.pixel_delta_v).Mult(0.5))

}

func (c *Camera) Render(world Hittable) {

	c.initialize()

	logger := log.New(os.Stderr, "", 0)

	logger.Print(world)

	fmt.Printf("P3\n%d %d\n255\n", c.Image_width, c.image_height)

	for j := 0; j < c.image_height; j++ {

		// logger.Printf("Scanlines remaining: %d", c.image_height-j)

		for i := 0; i < c.Image_width; i++ {

			pixel_center := c.pixel00_loc.Add(c.pixel_delta_u.Mult(float64(i)).Add(c.pixel_delta_v.Mult(float64(j))))
			Ray_direction := pixel_center.Sub(c.center)
			r := NewRay(c.center, Ray_direction)
			// logger.Print(c.pixel_delta_u)
			pixel_color := ray_color(r, world)

			Write_color(pixel_color)

		}

	}

}

func ray_color(r Ray, world Hittable) Color {
	var rec Hit_record

	if world.Hit(r, NewInterval(0, Infinity), &rec) {
		return rec.Normal.Add(NewColor(1, 1, 1)).Mult(0.5)
	}
	unit_direction := Unit_vector(r.Direction)
	a := 0.5 * (unit_direction.Y() + 1.0)
	return NewColor(1.0, 1.0, 1.0).Mult(1.0 - a).Add(NewColor(0.5, 0.7, 1.0).Mult(a))
}