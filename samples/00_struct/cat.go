package main

import "fmt"

type Cat struct {
	Name string
}

func main() {
	c := Cat{Name: "Kitty"}
	fmt.Println("Hello " + c.Name + "!")
}
