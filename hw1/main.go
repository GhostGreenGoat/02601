package main

import (
	"fmt"
	"gifhelper"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	//let's take command line arguments from the user
	//Go has built-in array called os.Args that is an array of strings.
	//os.Args[0] is program name ("jupiter".)
	//os.Args[1] is first command line argument.
	//os.Args[i] - the i-th command line argument.
	//etc.

	numBoids, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		fmt.Println("Error parsing number of boids")
	}
	//this would be bad: numGens := os.Args[1]. Must parse
	skyWidth, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic(err2)
	}
	initialSpeed, err3 := strconv.ParseFloat(os.Args[3], 64)
	if err3 != nil {
		panic("Error parsing initial Speed.")
	}
	maxBoidSpeed, err4 := strconv.ParseFloat(os.Args[4], 64)
	if err4 != nil {
		panic("Error parsing maxBoidSpeed.")
	}
	numGens, err5 := strconv.Atoi(os.Args[5])
	if err5 != nil {
		panic("Error parsing numGens.")
	}
	if numGens < 0 {
		panic("Error: numGens must be positive.")
	}
	proximity, err6 := strconv.ParseFloat(os.Args[6], 64)
	if err6 != nil {
		panic("Error parsing proximity.")
	}
	separationFactor, err7 := strconv.ParseFloat(os.Args[7], 64)
	if err7 != nil {
		panic("Error parsing separationFactor.")
	}
	alignmentFactor, err8 := strconv.ParseFloat(os.Args[8], 64)
	if err8 != nil {
		panic("Error parsing alignmentFactor.")
	}
	cohesionFactor, err9 := strconv.ParseFloat(os.Args[9], 64)
	if err9 != nil {
		panic("Error parsing cohesionFactor.")
	}
	timeStep, err10 := strconv.ParseFloat(os.Args[10], 64)
	if err10 != nil {
		panic("Error parsing timeStep.")
	}

	canvasWidth, err11 := strconv.Atoi(os.Args[11])
	if err11 != nil {
		panic("Error parsing canvasWidth.")
	}
	imageFrequency, err12 := strconv.Atoi(os.Args[12])
	if err12 != nil {
		panic("Error parsing imageFrequency.")
	}

	fmt.Println("Command line arguments read successfully.")

	fmt.Println("Simulating system.")

	// declaring sky and setting its fields.
	var skySystem Sky
	skySystem.width = float64(skyWidth)
	skySystem.boids = make([]Boid, numBoids)
	skySystem.maxBoidSpeed = float64(maxBoidSpeed)
	skySystem.proximity = float64(proximity)
	skySystem.separationFactor = float64(separationFactor)
	skySystem.alignmentFactor = float64(alignmentFactor)
	skySystem.cohesionFactor = float64(cohesionFactor)
	//initializing boids
	for i := 0; i < numBoids; i++ {
		skySystem.boids[i].position.x = rand.Float64() * skySystem.width
		skySystem.boids[i].position.y = rand.Float64() * skySystem.width
		skySystem.boids[i].velocity.x = rand.Float64() * initialSpeed
		skySystem.boids[i].velocity.y = initialSpeed - math.Sqrt(initialSpeed*initialSpeed-skySystem.boids[i].velocity.x*skySystem.boids[i].velocity.x)
		//generate random initial speed in all directions
		if rand.Float64() < 0.25 {
			skySystem.boids[i].velocity.x = -skySystem.boids[i].velocity.x
		} else if rand.Float64() < 0.5 {
			skySystem.boids[i].velocity.y = -skySystem.boids[i].velocity.y
		} else if rand.Float64() < 0.75 {
			skySystem.boids[i].velocity.x = -skySystem.boids[i].velocity.x
			skySystem.boids[i].velocity.y = -skySystem.boids[i].velocity.y
		}
		skySystem.boids[i].acceleration.x = 0
		skySystem.boids[i].acceleration.y = 0
	}
	fmt.Println("Sky initialized.")
	/*
		//test output

		fmt.Println("Boids has been simulated!")
		fmt.Println("Ready to draw images.")
		fmt.Println("timePoints: ")
		fmt.Println(timePoints[0].width)
		fmt.Println(timePoints[0].maxBoidSpeed)
		fmt.Println(timePoints[0].proximity)
		fmt.Println(timePoints[0].separationFactor, timePoints[0].alignmentFactor, timePoints[0].cohesionFactor)
		for i := range timePoints[0].boids {
			fmt.Println(timePoints[1].boids[i].position.x, timePoints[1].boids[i].position.y)
			fmt.Println(timePoints[1].boids[i].velocity.x, timePoints[1].boids[i].velocity.y)
		}
	*/
	timePoints := SimulateBoids(skySystem, numGens, timeStep)
	//fmt.Println("timePoints", timePoints[1])
	images := AnimateSystem(timePoints, canvasWidth, imageFrequency)

	fmt.Println("Images drawn!")

	fmt.Println("Making GIF.")

	gifhelper.ImagesToGIF(images, "Boids.gif")

	fmt.Println("Animated GIF produced!")

	fmt.Println("Exiting normally.")

}
