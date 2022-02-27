package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

// main to run the solution
func main() {

	// Create a square
	square, err := NewSquare(1, 1, 10)
	fmt.Printf("Square Before Move : %+v \n", square)

	if err != nil {
		log.Fatalf(" Error : Can not create a square")
	}

	square.Move(2, 3)

	fmt.Printf("Square After Move : %+v \n", square)
	fmt.Printf("Square Area Is  : %+v \n", square.Area())
}

// TODO Assignment : Define a struct square with a Center of type Point and Length of type int
// Define method Move and newSquare

type Square struct {
	Center Point
	Length int
}

func (square *Square) Move(dx int, dy int) {
	square.Center.Move(dx, dy)
}

func (square *Square) Area() int {
	return square.Length * square.Length
}

func NewSquare(x int, y int, length int) (*Square, error) {

	if length <= 0 {
		return nil, fmt.Errorf("Length must be greater than 0")
	}

	s := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return s, nil

}
