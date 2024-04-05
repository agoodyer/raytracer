package common

import (
	"fmt"
)

type Color = Vec3

func NewColor(x float64, y float64, z float64) Vec3 {
	return NewVec3(x, y, z)
}

func Write_color(pixel_color Color) {

	fmt.Printf("%d %d %d \n",
		int(255.999*pixel_color.X()),
		int(255.999*pixel_color.Y()),
		int(255.999*pixel_color.Z()),
	)

}
