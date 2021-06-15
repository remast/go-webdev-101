package main

import (
	"errors"
	"fmt"
)

type Cat struct {
	Name string
}

func (c Cat) feed(food string) error {
	if c.Name == "Kitty" && food != "Steak Tartare" {
		return errors.New("Won't eat!")
	}
	return nil
}

func main() {
	c := Cat{Name: "Kitty"}

	// Handle error
	err := c.feed("Caviar")
	if err != nil {
		fmt.Printf("%v won't eat it.", c.Name)
	}
}
