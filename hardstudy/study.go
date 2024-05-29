package main

import (
	"fmt"
	"sort"
)

type carro struct {
	Nome     string
	Consumo  int
	Potencia int
}

type ordernarPotencia []carro

func (x ordernarPotencia) Len() int      { return len(x) }
func (x ordernarPotencia) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x ordernarPotencia) Less(i, j int) bool {
	return !(x[i].Potencia < x[j].Potencia)
}

type ordenarConsumo []carro

func (x ordenarConsumo) Len() int      { return len(x) }
func (x ordenarConsumo) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x ordenarConsumo) Less(i, j int) bool {
	return !(x[i].Consumo < x[j].Consumo)
}

type ordenarLucroPosto []carro

func (x ordenarLucroPosto) Len() int      { return len(x) }
func (x ordenarLucroPosto) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x ordenarLucroPosto) Less(i, j int) bool {
	return x[i].Consumo < x[j].Consumo
}

func main() {
	carrosAnalise := []carro{
		{"civic", 15, 80},
		{"ford k", 25, 60},
		{"opala", 2, 70},
		{"ferrari", 5, 175},
	}
	sort.Sort(ordernarPotencia(carrosAnalise))
	fmt.Println(carrosAnalise)
	sort.Sort(ordenarConsumo(carrosAnalise))
	fmt.Println(carrosAnalise)
	sort.Sort(ordenarLucroPosto(carrosAnalise))
	fmt.Println(carrosAnalise)

	sort.Slice(carrosAnalise, func(i, j int) bool {
		return carrosAnalise[i].Consumo < carrosAnalise[j].Consumo
	})
	fmt.Println("por consumo:", carrosAnalise)

}
