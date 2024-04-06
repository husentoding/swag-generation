package main

import (
	swaggeneration "github.com/husentoding/swag-generation"
)

func main() {
	
	builder := swaggeneration.Init("./base-template/", "./example/identity-service")
	builder.Generate("output.swagger.yaml")

}
