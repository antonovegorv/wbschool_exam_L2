package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

func main() {
	h := NewHotdog("gluten free").addKetchup().addMustard()
	fmt.Println(h)
}

type Hotdog struct {
	bread   string
	ketchup bool
	mustard bool
	kraut   bool
}

func NewHotdog(bread string) *Hotdog {
	return &Hotdog{bread: bread}
}

func (h *Hotdog) addKetchup() *Hotdog {
	h.ketchup = true
	return h
}

func (h *Hotdog) addMustard() *Hotdog {
	h.mustard = true
	return h
}

func (h *Hotdog) addKraut() *Hotdog {
	h.kraut = true
	return h
}
