package main

import (
	"canvas"
	"image"
)

//place your drawing code here.
// let's place our drawing functions here.

//AnimateSystem takes a slice of sky objects along with a canvas width
//parameter and a frequency parameter. It generates a slice of images corresponding to drawing each sky whose index is divisible by the frequency parameter
//on a canvasWidth x canvasWidth canvas

func AnimateSystem(timePoints []Sky, canvasWidth, drawingFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		if i%drawingFrequency == 0 { //only draw if current index of universe
			//is divisible by some parameter frequency
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

//DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
//object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(s Sky, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the bodies and draw them.
	for _, b := range s.boids {
		c.SetFillColor(canvas.MakeColor(255, 255, 255))
		cx := (b.position.x / s.width) * float64(canvasWidth)
		cy := (b.position.y / s.width) * float64(canvasWidth)
		r := (5 / s.width) * float64(canvasWidth)
		c.Circle(cx, cy, r)
		c.Fill()

	}
	return c.GetImage()
}
