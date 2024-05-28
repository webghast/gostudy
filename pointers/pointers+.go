package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func mudeMe(p *pessoa) {
	(*p).nome = "ricardo"
}

func main() {
	p1 := pessoa{"marquinhos", 16}
	x := &p1
	fmt.Println(x)
	mudeMe(*&x)

	fmt.Println(*&x)

}
