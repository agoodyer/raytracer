package objects

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	. "raytracer/common"
	. "raytracer/material"
	"sync"
	"time"
)

type Camera struct {
	Aspect_ratio     float64
	Image_width      int
	Sample_per_pixel int
	image_height     int
	center           Point3
	pixel00_loc      Point3
	pixel_delta_u    Vec3
	pixel_delta_v    Vec3
	max_depth        int
	Vfov             float64
	Look_from        Point3
	Look_at          Point3
	Vup              Point3
	u, v, w          Vec3
	Defocus_angle    float64
	Focus_dist       float64
	defocus_disk_u   Vec3
	defocus_disk_v   Vec3
}

func (c *Camera) initialize() {

	// c.Aspect_ratio = 1.0
	// c.Image_width = 100

	c.Sample_per_pixel = 10 //250
	c.max_depth = 4         //50

	c.Vfov = 20
	c.Look_from = NewPoint3(13, 2, 3)
	c.Look_at = NewPoint3(0, 0, 0)
	c.Vup = NewVec3(0, 1, 0)

	c.Defocus_angle = 0.6
	c.Focus_dist = 10.0

	//Calculate image height
	c.image_height = int(float64(c.Image_width) / c.Aspect_ratio)
	if c.image_height < 1 {
		c.image_height = 1
	}

	c.center = c.Look_from

	//Camera
	// focal_length := (c.Look_from.Sub(c.Look_at).Length())
	theta := Degrees_to_radians(c.Vfov)
	h := math.Tan(theta / 2)

	viewport_height := 2.0 * h * c.Focus_dist
	viewport_width := viewport_height * float64(c.Image_width) / float64(c.image_height)

	c.w = Unit_vector(c.Look_from.Sub(c.Look_at))
	c.u = Unit_vector(Cross(c.Vup, c.w))
	c.v = Cross(c.w, c.u)

	viewport_u := c.u.Mult(viewport_width)
	viewport_v := c.v.Mult(-viewport_height)

	c.pixel_delta_u = viewport_u.Div(float64(c.Image_width))
	c.pixel_delta_v = viewport_v.Div(float64(c.image_height))

	viewport_upper_left := c.center.Sub(c.w.Mult(c.Focus_dist)).Sub(viewport_u.Div(2)).Sub(viewport_v.Div(2))

	c.pixel00_loc = viewport_upper_left.Add(c.pixel_delta_u.Add(c.pixel_delta_v).Mult(0.5))

	defocus_radius := c.Focus_dist * math.Tan(Degrees_to_radians(c.Defocus_angle/2))
	c.defocus_disk_u = c.u.Mult(defocus_radius)
	c.defocus_disk_u = c.v.Mult(defocus_radius)

}

func (c *Camera) Render(world Hittable) {

	start := time.Now() //time execution

	c.initialize()

	img := image.NewRGBA(image.Rect(0, 0, c.Image_width, c.image_height))

	for j := 0; j < c.image_height; j++ {

		fmt.Printf("Scanlines remaining: %d\n", c.image_height-j)

		for i := 0; i < c.Image_width; i++ {

			pixel_color := NewColor(0, 0, 0)

			for sample := 0; sample < c.Sample_per_pixel; sample++ {
				r := c.get_ray(i, j)
				pixel_color = pixel_color.Add(ray_color(&r, c.max_depth, world))
			}

			Write_color(pixel_color, c.Sample_per_pixel, img, i, j)

		}

	}

	//THREAD TESTING
	// var wg sync.WaitGroup

	// numWorkers := 2

	// wg.Add(numWorkers)

	// go c.renderBlock(0, c.image_height/4, world, img, &wg)

	// go c.renderBlock(c.image_height/4, c.image_height/4, world, img, &wg)

	// wg.Wait()

	//END TEST

	// Create a file to save the image
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the image to PNG and save to the file
	if err := png.Encode(file, img); err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	fmt.Print("~~~~~~~~~~~~~~~~~~~~~~~~~~\nElapsed Time: ", elapsed, "\n~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

}

func (c *Camera) renderBlock(startLine int, numLines int, world Hittable, img *image.RGBA, wg *sync.WaitGroup) {

	defer wg.Done()

	for j := startLine; j < startLine+numLines; j++ {

		for i := 0; i < c.Image_width; i++ {

			pixel_color := NewColor(0, 0, 0)

			for sample := 0; sample < c.Sample_per_pixel; sample++ {
				r := c.get_ray(i, j)
				pixel_color = pixel_color.Add(ray_color(&r, c.max_depth, world))
				// logger.Print(pixel_color)
			}

			// pixel_center := c.pixel00_loc.Add(c.pixel_delta_u.Mult(float64(i)).Add(c.pixel_delta_v.Mult(float64(j))))
			// Ray_direction := pixel_center.Sub(c.center)
			// r := NewRay(c.center, Ray_direction)
			// // logger.Print(c.pixel_delta_u)
			// pixel_color := ray_color(r, world)

			Write_color(pixel_color, c.Sample_per_pixel, img, i, j)

		}

	}

}

func ray_color(r *Ray, depth int, world Hittable) Color {
	var rec Hit_record

	if depth <= 0 {
		return NewColor(0, 0, 0)
	}

	if world.Hit(r, NewInterval(0.001, Infinity), &rec) {

		// direction := rec.Normal.Add(Random_unit_vector())

		// return ray_color(NewRay(rec.P, direction), depth-1, world).Mult(0.50)

		var scattered Ray
		var attenuation Color

		if rec.Mat.Scatter(r, &rec, &attenuation, &scattered) {
			return ComponentMultiply(attenuation, ray_color(&scattered, depth-1, world))
		}

	}
	unit_direction := Unit_vector(r.Direction)
	a := 0.5 * (unit_direction.Y() + 1.0)
	return NewColor(1.0, 1.0, 1.0).Mult(1.0 - a).Add(NewColor(0.5, 0.7, 1.0).Mult(a))
}

func (c *Camera) get_ray(i int, j int) Ray {

	pixel_center := c.pixel00_loc.Add(c.pixel_delta_u.Mult(float64(i)).Add(c.pixel_delta_v.Mult(float64(j))))
	pixel_sample := pixel_center.Add(c.pixel_sample_square())

	ray_origin := c.center

	if c.Defocus_angle > 0 {
		ray_origin = c.defocus_disk_sample()
	}

	ray_direction := pixel_sample.Sub(ray_origin)

	return NewRay(ray_origin, ray_direction)
}

func (c *Camera) defocus_disk_sample() Point3 {
	p := Random_in_unit_disk()
	return c.center.Add(c.defocus_disk_u.Mult(p.X())).Add(c.defocus_disk_v.Mult(p.Y()))
}

func (c *Camera) pixel_sample_square() Vec3 {
	px := -0.5 + rand.Float64()
	py := -0.5 + rand.Float64()
	return c.pixel_delta_u.Mult(px).Add(c.pixel_delta_v.Mult(py))
}
