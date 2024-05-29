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

	fmt.Println(p1)
	mudeMe(&p1)

	fmt.Println(p1)

}
