package main

import "fmt"

// pizza interface
type pizza interface {
	price() int
}

// veggieMania exact pizza
type veggieMania struct {
}

func (v *veggieMania) price() int {
	return 10
}

// tomatoTopping pizza decorator
type tomatoTopping struct {
	pizza
}

func (t *tomatoTopping) price() int {
	return t.pizza.price() + 1
}

// cheeseTopping pizza decorator
type cheeseTopping struct {
	pizza
}

func (c *cheeseTopping) price() int {
	return c.pizza.price() + 2
}

// main
func main() {
	var p pizza
	p = &cheeseTopping{
		&tomatoTopping{
			&veggieMania{},
		},
	}

	fmt.Println(p.price())
}
