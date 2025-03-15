package prototype

import "fmt"

// Prototype Interface
type Prototype interface {
	Clone() Prototype
}

type DockerContainer struct {
	Image   string
	CPU     int
	Memory  int
	Ports   []int
	EnvVars map[string]string
}

func NewDockerContainer() *DockerContainer {
	return &DockerContainer{
		Image:   "ubuntu:latest",
		CPU:     2,
		Memory:  4096,
		Ports:   []int{8080, 9000},
		EnvVars: map[string]string{"ENV": "production"},
	}
}

func (d *DockerContainer) Clone() Prototype {
	portsCopy := make([]int, len(d.Ports))
	copy(portsCopy, d.Ports)

	envVarsCopy := make(map[string]string)
	for key, value := range d.EnvVars {
		envVarsCopy[key] = value
	}

	return &DockerContainer{
		Image:   d.Image,
		CPU:     d.CPU,
		Memory:  d.Memory,
		Ports:   portsCopy,
		EnvVars: envVarsCopy,
	}
}

func (d *DockerContainer) Describe() {
	fmt.Printf("[Docker Container]\n")
	fmt.Printf("Image: %s\n", d.Image)
	fmt.Printf("CPU: %d\n", d.CPU)
	fmt.Printf("Memory: %d MB\n", d.Memory)
	fmt.Printf("Ports: %v\n", d.Ports)
	fmt.Printf("Env Vars: %v\n", d.EnvVars)
}
