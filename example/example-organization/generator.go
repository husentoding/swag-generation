package main

import (
	swaggeneration "github.com/husentoding/swag-generation"
)

// path are configured from root swag-generation repository
func main() {
	builder := swaggeneration.Init("../../base-template/",
		"./identity-service",
		"./identity-service/output.swagger.yaml")
	builder.Generate()
}
