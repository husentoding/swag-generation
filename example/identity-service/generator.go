package main

import (
	swaggeneration "github.com/husentoding/swag-generation"
)

// path are configured from root swag-generation repository
func main() {
	builder := swaggeneration.Init("./base-template/", "./example/identity-service")
	builder.Generate("./example/identity-service/output.swagger.yaml")
}
