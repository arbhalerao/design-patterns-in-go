package main

import (
	"github.com/abha1erao/design-patterns-in-go/creational/factory"
)

func main() {
	circle := factory.GetShape("circle")
	square := factory.GetShape("square")
	rectangle := factory.GetShape("rectangle")

	circle.Draw()
	square.Draw()
	rectangle.Draw()
}
