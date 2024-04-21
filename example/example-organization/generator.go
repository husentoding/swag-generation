package main

import (
	swaggeneration "github.com/husentoding/swag-generation"
)

// path are configured from example-organization as the workdir
func main() {
	builder := swaggeneration.Init("../../base-template/",
		"./array-service",
		"./array-service/output.swagger.yaml")
	builder.Generate()
}
