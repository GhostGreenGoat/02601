package test

import (
	"fmt"
	"math"
	//"strconv"
	//"strings"
	//"testing"
)

//OrderedPair contains two float64 fields corresponding to
//the x and y coordinates of a point or vector in two-dimensional space.
type OrderedPair struct {
	x, y float64
}

//Boid represents our "bird" object. It contains two
//OrderedPair fields corresponding to its position, velocity, and acceleration.
type Boid struct {
	position, velocity, acceleration OrderedPair
}

//Sky represents a single time point of the simulation.
//It contains a width parameter indicating the boundary of the sky, as well as a slice of Boid objects.
//It also contains the system parameters.
type Sky struct {
	width                                             float64
	boids                                             []Boid
	maxBoidSpeed                                      float64 //fastest speed that a boid can fly
	proximity                                         float64 // used to determine if boids are close enough for forces to apply
	separationFactor, alignmentFactor, cohesionFactor float64 //multiply by each respective force
}

func updateAcceleration(currentsky Sky, b Boid) OrderedPair {
	var accel OrderedPair
	//compute net force vector acting on b
	force := computeNetForce(currentsky, b)
	//now, calculate acceleration (F = ma)
	accel.x = force.x
	accel.y = force.y
	return accel
}

//ComputeNetForce
//Input: A slice of boid objects and an individual boid
//Output: The net force vector (OrderedPair) acting on the given boid due to boids from all other boids in the sky
func computeNetForce(currentsky Sky, b Boid) OrderedPair {
	var force OrderedPair
	nearbyBoids := 0
	//range over all boids in the sky and add up the forces
	for i := range currentsky.boids {
		//don't add the force of the boid on itself
		//add the force within the proximity
		if currentsky.boids[i].position != b.position && Distance(currentsky.boids[i].position, b.position) <= currentsky.proximity {
			force1 := computeSeparation(b, currentsky.boids[i], currentsky.separationFactor)
			fmt.Println("seperation force on", i, "th boid to current boid:", force1)
			force2 := computeAlignment(currentsky.boids[i], b, currentsky.alignmentFactor)
			fmt.Println("alignment force on", i, "th boid to current boid:", force2)
			force3 := computeCohesion(b, currentsky.boids[i], currentsky.cohesionFactor)
			fmt.Println("alignment force on", i, "th boid to current boid:", force3)
			//add the three forces together
			//Seperation force has different direction than the other two forces
			force.x = force.x + force1.x + force2.x + force3.x
			force.y = force.x + force1.y + force2.y + force3.y
			nearbyBoids++
		}
	}
	if nearbyBoids != 0 {
		force.x = force.x / float64(nearbyBoids)
		force.y = force.y / float64(nearbyBoids)
	} else {
		force.x = 0
		force.y = 0
	}
	return force
}

//Compute separation force
//Input: Two boids
//Output: The force vector (OrderedPair) acting on the first boid due to the second boid
func computeSeparation(b1 Boid, b2 Boid, seperationFactor float64) OrderedPair {
	var force OrderedPair
	//compute the distance between the two boids
	dist := Distance(b1.position, b2.position)
	//compute the force vector
	force.x = seperationFactor * (b1.position.x - b2.position.x) / (dist * dist)
	force.y = seperationFactor * (b1.position.y - b2.position.y) / (dist * dist)
	//fmt.Println("seperation force: ", force)
	return force
}

//Compute alignment force
func computeAlignment(b1 Boid, b2 Boid, alignmentFactor float64) OrderedPair {
	var force OrderedPair
	//compute the distance between the two boids
	dist := Distance(b1.position, b2.position)
	//compute the force vector
	force.x = alignmentFactor * (b1.velocity.x) / (dist)
	force.y = alignmentFactor * (b1.velocity.y) / (dist)
	//fmt.Println("alignment force: ", force)
	return force
}

//Compute Cohesion force
func computeCohesion(b1 Boid, b2 Boid, cohesionFactor float64) OrderedPair {
	var force OrderedPair
	//compute the distance between the two boids
	dist := Distance(b1.position, b2.position)
	//compute the force vector
	force.x = cohesionFactor * (b2.position.x - b1.position.x) / (dist)
	force.y = cohesionFactor * (b2.position.y - b1.position.y) / (dist)
	//fmt.Println("cohesion force: ", force)
	return force
}

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

//test main function
func main() {
	//create a sky
	numBoids := 3
	var skySystem Sky
	skySystem.width = 400
	skySystem.boids = make([]Boid, numBoids)
	skySystem.maxBoidSpeed = 2
	skySystem.proximity = 200
	skySystem.separationFactor = 1
	skySystem.alignmentFactor = 1
	skySystem.cohesionFactor = 1
	skySystem.boids[0].position.x = 241.864115
	skySystem.boids[0].position.y = 376.203635
	skySystem.boids[0].velocity.x = -0.511419
	skySystem.boids[0].velocity.y = -0.859332
	skySystem.boids[0].acceleration.x = 0
	skySystem.boids[0].acceleration.y = 0
	skySystem.boids[1].position.x = 175.085675
	skySystem.boids[1].position.y = 169.854999
	skySystem.boids[1].velocity.x = -0.386609
	skySystem.boids[1].velocity.y = -0.922244
	skySystem.boids[1].acceleration.x = 0
	skySystem.boids[1].acceleration.y = 0
	skySystem.boids[2].position.x = 26.254808
	skySystem.boids[2].position.y = 62.607702
	skySystem.boids[2].velocity.x = 0.820062
	skySystem.boids[2].velocity.y = 0.572275
	skySystem.boids[2].acceleration.x = 0
	skySystem.boids[2].acceleration.y = 0
	for i := range skySystem.boids {
		netforce := computeNetForce(skySystem, skySystem.boids[i])
		fmt.Println("netforce on boid ", i, " is ", netforce)
	}
}
