package objects

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "raytracer/common"
	. "raytracer/material"
)

func NewMeshFromFile(path string, material Material, scale float64) Hittable_list {

	var triangles Hittable_list

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error Loading Mesh from ", path)
		return triangles
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if fields[0] == "facet" {

			scanner.Scan()
			scanner.Scan()

			// read vertex 1 data
			line = scanner.Text()
			// fmt.Println(line)
			fields = strings.Fields(line)
			x, _ := strconv.ParseFloat(fields[1], 64)
			y, _ := strconv.ParseFloat(fields[2], 64)
			z, _ := strconv.ParseFloat(fields[3], 64)
			a := NewPoint3(x, y, z)
			// fmt.Println(a)

			scanner.Scan()
			// read vertex 2 data
			line = scanner.Text()
			// fmt.Println(line)
			fields = strings.Fields(line)
			x, _ = strconv.ParseFloat(fields[1], 64)
			y, _ = strconv.ParseFloat(fields[2], 64)
			z, _ = strconv.ParseFloat(fields[3], 64)
			b := NewPoint3(x, y, z)
			// fmt.Println(b)

			scanner.Scan()
			// read vertex 3 data
			line = scanner.Text()
			// fmt.Println(line)
			fields = strings.Fields(line)
			x, _ = strconv.ParseFloat(fields[1], 64)
			y, _ = strconv.ParseFloat(fields[2], 64)
			z, _ = strconv.ParseFloat(fields[3], 64)
			c := NewPoint3(x, y, z)
			// fmt.Println(c)

			// green := NewLambertian(NewColor(0.12, 0.45, 0.15))
			// red := NewLambertian(NewColor(0.72, 0.25, 0.15))
			tri := Triangle(a.Mult(scale), b.Mult(scale), c.Mult(scale), material)

			// v1 := NewSphere(a, 0.5, &green)
			// v2 := NewSphere(b, 0.5, &green)
			// v3 := NewSphere(c, 0.5, &green)

			triangles.Add(&tri)
			// triangles.Add(&v1)
			// triangles.Add(&v2)
			// triangles.Add(&v3)

		}

	}

	return triangles

}
