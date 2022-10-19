package main

import (
	"math"
)

//place your non-drawing functions here.

//simulates Boids over a series of snap shots separated by equal unit time.
//Input: initial sky object, number of generations, and time parameter (in seconds).
//Output: a slice of generations+1 sky objects
func SimulateBoids(initialsky Sky, numGens int, time float64) []Sky {
	timePoints := make([]Sky, numGens+1)
	timePoints[0] = initialsky
	//now range over the number of generations and update the sky each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = updateSky(timePoints[i-1], time)
	}
	return timePoints
}

//UpdateSky updates a given sky over a specified time interval (in seconds).
//Input: A sky object and a float time.
//Output: A sky object corresponding to simulating boids over time seconds, assuming that acceleration is constant over this time.
func updateSky(currentsky Sky, time float64) Sky {
	newsky := copySky(currentsky)
	//range over all boids in the sky and update their acceleration, velocity, and position
	for i := range newsky.boids {
		newsky.boids[i].acceleration = updateAcceleration(currentsky, newsky.boids[i])
		newsky.boids[i].velocity = updateVelocity(newsky.boids[i], time, currentsky.maxBoidSpeed)
		newsky.boids[i].position = updatePosition(newsky.boids[i], time, currentsky.width)
	}
	return newsky
}

//UpdateAcceleration
//Input: sky object and a boid B
//Output: The net acceleration on B due to boids calculated by every boid in the sky
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
			force2 := computeAlignment(currentsky.boids[i], b, currentsky.alignmentFactor)
			force3 := computeCohesion(b, currentsky.boids[i], currentsky.cohesionFactor)
			//add the three forces together
			//Seperation force has different direction than the other two forces
			force.x = force.x + force1.x + force2.x + force3.x
			force.y = force.y + force1.y + force2.y + force3.y
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

//CopySky makes a copy of a sky object
func copySky(sky Sky) Sky {
	var newSky Sky
	newSky.width = sky.width
	newSky.maxBoidSpeed = sky.maxBoidSpeed
	newSky.proximity = sky.proximity
	newSky.separationFactor = sky.separationFactor
	newSky.alignmentFactor = sky.alignmentFactor
	newSky.cohesionFactor = sky.cohesionFactor
	newSky.boids = make([]Boid, len(sky.boids))
	for i := range sky.boids {
		newSky.boids[i].acceleration.x = sky.boids[i].acceleration.x
		newSky.boids[i].acceleration.y = sky.boids[i].acceleration.y
		newSky.boids[i].velocity.x = sky.boids[i].velocity.x
		newSky.boids[i].velocity.y = sky.boids[i].velocity.y
		newSky.boids[i].position.x = sky.boids[i].position.x
		newSky.boids[i].position.y = sky.boids[i].position.y
	}
	return newSky
}

//UpdateVelocity updates the velocity of a boid over a specified time interval (in seconds).
//Input: A boid object and a float time.
//Output: A boid object corresponding to simulating boid over time seconds, assuming that acceleration is constant over this time.
func updateVelocity(b Boid, time float64, maxBoidSpeed float64) OrderedPair {
	var newVelocity OrderedPair
	newVelocity.x = b.velocity.x + b.acceleration.x*time
	newVelocity.y = b.velocity.y + b.acceleration.y*time
	//limit the speed of the boid to maxBoidSpeed
	if absoluteVelocity(newVelocity) > maxBoidSpeed {
		tempx := maxBoidSpeed * newVelocity.x / absoluteVelocity(newVelocity)
		tempy := maxBoidSpeed * newVelocity.y / absoluteVelocity(newVelocity)
		newVelocity.x = tempx
		newVelocity.y = tempy
	}
	return newVelocity
}

//calculate the absolute velocity of a boid
func absoluteVelocity(velocity OrderedPair) float64 {
	return math.Sqrt(velocity.x*velocity.x + velocity.y*velocity.y)
}

//UpdatePosition updates the position of a boid over a specified time interval (in seconds).
//Input: A boid object and a float time.
//Output: A boid object corresponding to simulating boid over time seconds, assuming that acceleration is constant over this time.
func updatePosition(b Boid, time float64, width float64) OrderedPair {
	var newPosition OrderedPair
	newPosition.x = b.position.x + b.velocity.x*time + 0.5*b.acceleration.x*time*time
	newPosition.y = b.position.y + b.velocity.y*time + 0.5*b.acceleration.y*time*time
	//check if the boid is out of bounds
	//make the boids in the sky
	if newPosition.x > width {
		newPosition.x = newPosition.x - width
	}
	if newPosition.x < 0 {
		newPosition.x = newPosition.x + width
	}
	if newPosition.y > width {
		newPosition.y = newPosition.y - width
	}
	if newPosition.y < 0 {
		newPosition.y = newPosition.y + width
	}
	return newPosition
}
