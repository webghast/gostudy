package main

import (
	"fmt"
)

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
}
