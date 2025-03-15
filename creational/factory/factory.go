package factory

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing Circle")
}

type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Drawing Square")
}

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing Rectangle")
}

// Factory function to create shapes
func GetShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	case "rectangle":
		return &Rectangle{}
	default:
		return nil
	}
}
