package main

import (
	"fmt"

	"github.com/arbhalerao/design-patterns-in-go/creational/builder"
)

func main() {
	// Builder
	req := builder.NewRequestBuilder().
		SetMethod("GET").
		SetURL("https://example.com").
		SetHeader("Accept", "application/json").
		SetQueryParam("page", "1").
		Build()

	req.Describe()

	fmt.Println()

	// Step Builder
	stepReq := builder.NewRequestStepBuilder().
		SetMethod("POST").
		SetURL("https://example.com").
		SetHeader("Content-Type", "application/json").
		SetQueryParam("id", "123").
		SetBody("{\"name\":\"test\"}").
		Build()

	stepReq.Describe()
}
