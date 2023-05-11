package main

import (
	"awesomeProject1/config"
	"flag"
)

func main() {
	println("test")
	var config config.Testing

	flag.StringVar(&config.Test, "TESTINGTEST", "", "testing only")

	println(config.Test)
}
