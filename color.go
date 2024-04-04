package main

import (
	"fmt"
)

type color = vec3

func Color(x float64, y float64, z float64) vec3 {
	return Vec3(x, y, z)
}

func write_color(pixel_color color) {

	fmt.Printf("%d %d %d \n",
		int(255.999*pixel_color.x()),
		int(255.999*pixel_color.y()),
		int(255.999*pixel_color.z()),
	)

}
