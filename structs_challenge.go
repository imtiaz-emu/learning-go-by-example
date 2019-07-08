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
	p.X = dx
	p.Y = dy
}

type Square struct {
	Length int
	Center Point
}

func NewSquare(length int, x int, y int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Length of a square cannot less than 0")
	}
	s := &Square{
		Length: length,
		Center: Point{x, y},
	}

	return s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	square, err := NewSquare(-3, 1, 1)
	if err == nil {
		fmt.Println("Area of the square = ", square.Area())
		square.Move(3, 2)
		fmt.Println("Square points after move = ", square.Center)
	} else {
		log.Fatalf("Cannot continue ", err)
	}
}
