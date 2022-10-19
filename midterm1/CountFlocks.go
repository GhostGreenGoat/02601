package main

import "math"

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
//It corresponds to a slice of Boid objects
type Sky []*Boid

func find(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	return find(parent, parent[i])
}
func Union(parent []int, rank []int, x int, y int) {
	xroot := find(parent, x)
	yroot := find(parent, y)
	if rank[xroot] < rank[yroot] {
		parent[xroot] = yroot
	} else if rank[xroot] > rank[yroot] {
		parent[yroot] = xroot
	} else {
		parent[yroot] = xroot
		rank[xroot]++
	}
}

func initializeParent(s Sky) []int {
	parent := make([]int, len(s))
	for i := range s {
		parent[i] = i
	}
	return parent
}

func Count(parent []int) int {
	count := 0
	for i := range parent {
		if parent[i] == i {
			count++
		}
	}
	return count
}

//Insert your CountFlocks() function here, along with any subroutines and type declarations that you need.
func (s Sky) CountFlocks(flockDistance float64) int {
	parent := initializeParent(s)
	rank := make([]int, len(s))

	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if Distance(s[i].position, s[j].position) < flockDistance {
				Union(parent, rank, i, j)
			}
		}
	}
	Flocks := Count(parent)
	return Flocks
}

func Distance(a, b OrderedPair) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}
