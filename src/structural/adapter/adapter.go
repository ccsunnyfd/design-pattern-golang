package main

import (
	"fmt"
	"math"
)

// roundHole
type roundHole struct {
	radius float64
}

func (rHole *roundHole) Fits(peg iRoundPeg) bool {
	return rHole.radius >= peg.getRadius()
}

// iRoundPeg interface
type iRoundPeg interface {
	getRadius() float64
}

// roundPeg
type roundPeg struct {
	radius float64
}

func (rPeg *roundPeg) getRadius() float64 {
	return rPeg.radius
}

// squarePeg
type squarePeg struct {
	width float64
}

func (sPeg *squarePeg) getWidth() float64 {
	return sPeg.width
}

// roundPegAdapter
type squarePegAdapter struct {
	sPeg *squarePeg
}

func (sPegAdapter *squarePegAdapter) getRadius() float64 {
	return sPegAdapter.sPeg.getWidth() / 2 * math.Sqrt(2)
}

// main
func main() {
	rHole := &roundHole{
		1.5,
	}
	rPeg := &roundPeg{
		2,
	}
	sPeg := &squarePeg{
		1.2,
	}
	squarePegAdapter := &squarePegAdapter{
		sPeg,
	}
	fmt.Println(rHole.Fits(rPeg))
	fmt.Println(rHole.Fits(squarePegAdapter))
}
