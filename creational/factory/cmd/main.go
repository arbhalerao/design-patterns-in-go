package main

import (
	"github.com/arbhalerao/design-patterns-in-go/creational/factory"
)

func main() {
	circle := factory.GetShape("circle")
	square := factory.GetShape("square")
	rectangle := factory.GetShape("rectangle")

	circle.Draw()
	square.Draw()
	rectangle.Draw()
}
