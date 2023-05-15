package main

import (
	"awesomeProject1/config"
	"os"
)

var (
	version = "dev"
)

func main() {
	println("test")
	var config config.Testing

	config.Test = os.Getenv("TESTING_TEST")

	println(config.Test)

	println("app version is: %s", version)
}
