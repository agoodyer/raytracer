package common

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type Color = Vec3

func NewColor(x float64, y float64, z float64) Vec3 {
	return NewVec3(x, y, z)
}

func Write_color(pixel_color Color, samples_per_pixel int, img *image.RGBA, i int, j int) {

	r := pixel_color.X()
	g := pixel_color.Y()
	b := pixel_color.Z()

	scale := 1.0 / float64(samples_per_pixel)

	r *= scale
	g *= scale
	b *= scale

	r = linear_to_gamma(r)
	g = linear_to_gamma(g)
	b = linear_to_gamma(b)

	intensity := NewInterval(0.000, 0.999)

	fmt.Printf("%d %d %d \n",
		int(256*intensity.Clamp(r)),
		int(256*intensity.Clamp(g)),
		int(256*intensity.Clamp(b)),
	)

	img.Set(i, j, color.RGBA{
		uint8(256 * intensity.Clamp(r)),
		uint8(256 * intensity.Clamp(g)),
		uint8(256 * intensity.Clamp(b)), 255})

}

func linear_to_gamma(linear_component float64) float64 {
	return math.Sqrt(linear_component)
}
