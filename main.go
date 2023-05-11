package main

import (
	"awesomeProject1/config"
	"os"
)

func main() {
	println("test")
	var config config.Testing
	config.Test = os.Getenv("testingTest")
	println(config.Test)
}
