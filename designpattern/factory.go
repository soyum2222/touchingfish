package main

import "fmt"

type shape interface {
	draw()
}

func GetShape(shape string) shape {

	switch shape {
	case "square":
		return &Square{}
	case "circle":
		return &Circle{}
	}
	return nil

}

type Square struct {
}

func (s *Square) draw() {
	fmt.Println("im square")
}

type Circle struct {
}

func (c *Circle) draw() {
	fmt.Println("im circle")
}

func main() {

	GetShape("circle").draw()
	GetShape("square").draw()

}
