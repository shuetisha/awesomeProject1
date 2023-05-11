package main

import (
	"awesomeProject1/config"
	"flag"
	"os"
)

func main() {
	println("test")
	var config config.Testing

	flag.StringVar(&config.Test, "testing test", "", "testing only")

	config.Test = os.Getenv("testingTest")
	println(config.Test)
}
