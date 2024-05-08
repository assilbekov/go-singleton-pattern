package main

import (
	"fmt"
	"singleton/pkg"
)

var singleton *pkg.Singleton

func init() {
	fmt.Println("Init project")
	singleton = pkg.NewSingleton(singleton, "Init")
}

func main() {
	singleton = pkg.NewSingleton(singleton, "First")
	singleton.Print()

	singleton = pkg.NewSingleton(singleton, "Second")
	singleton.Print()
}
