# 3D Raytracer 
Created by Aidan Goodyer 


# About the Project

Raytracing is a computationally-expensive and math heavy task. My motivation to build this project was to create a program that is highly parallelizable and ripe for optimization, as to provide an environment to practice my **Go** skills. Additionally, this project supports the rendering of STL files, allowing users to preview their own 3D Meshes.



# Installation 

- After downloading, running **go build main.go** will produce the standalone binary required to run the raytracer
- Alternativel, run the project using **go run main.go** 



# Sample Renders 


![3D Mesh Scene](/sample_renders/mesh_scene.png)
![Earth and Moon Scene](/sample_renders/earth_scene.png)
![Random Sphere Scene](/sample_renders/sphere_scene.png)
![Cornell Box Scene](/sample_renders/cornell_box_scene.png)


# Performance 

- Supports both single-threaded and multi-threaded rendering modes
- Implements a Bounding Volume Hierarchy structure to improve performance on complex scenes

## CPU Profiling Info 

Observing the CPU Profiling info, we can see most of the render time is spent calculating Ray-Surface collisions, which is to be expected since this is fundamental to the raytracing process. We can also see that the Go Compiler has performed some clever inlining of common math functions in the project, including Dot Products, and Interval Checks. 


This Profiling information comes from the render of an STL File at high-resolution and high sample count. 

![CPU Profile for Mesh Render](/sample_renders/cpu_profiling.svg)

