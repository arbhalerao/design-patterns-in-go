package main

import "github.com/abha1erao/design-patterns-in-go/creational/builder"

func main() {
	request := builder.NewRequestBuilder().
		SetMethod("POST").
		SetURL("https://api.app.com/v1/users").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer XYZ123").
		SetQueryParam("limit", "50").
		SetQueryParam("offset", "10").
		SetBody(`{"username": "aditya", "email": "abha@gmail.com"}`).
		Build()

	request.Describe()
}
