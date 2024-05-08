package main

import (
	"fmt"
	"singleton/pkg"
)

var singleton *pkg.Singleton

func init() {
	fmt.Println("Init project")
}

func main() {
	singleton = pkg.NewSingleton(singleton, "First")
	singleton.Print()

	singleton = pkg.NewSingleton(singleton, "Second")
	singleton.Print()
}
