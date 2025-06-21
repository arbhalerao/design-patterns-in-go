package main

import (
	"fmt"

	"github.com/arbhalerao/design-patterns-in-go/creational/prototype"
)

func main() {
	// Original DockerContainer
	baseContainer := prototype.NewDockerContainer()
	baseContainer.Describe()
	fmt.Println()

	// Clone 1: For Backend Service
	backendContainer := baseContainer.Clone().(*prototype.DockerContainer)
	backendContainer.CPU = 4
	backendContainer.Ports = append(backendContainer.Ports, 3000)
	backendContainer.EnvVars["SERVICE"] = "backend"

	backendContainer.Describe()
	fmt.Println()

	// Clone 2: For Database Service
	dbContainer := baseContainer.Clone().(*prototype.DockerContainer)
	dbContainer.Image = "postgres:latest"
	dbContainer.CPU = 8
	dbContainer.Memory = 8192
	dbContainer.EnvVars["DB_USER"] = "admin"

	dbContainer.Describe()
}
