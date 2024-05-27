package main

import (
	"fmt"
)

type lalo struct {
	lolo string
	lili int
	lulu bool
}

func main() {
	x := struct {
		m map[string]int
		s []string
	}{
		m: map[string]int{
			"jeff":  12,
			"mario": 15,
		},
		s: []string{"ola", "oi", "hello"},
	}
	fmt.Println(x)

	y := lalo{"ola", 12, true}
	fmt.Println(y)
}
