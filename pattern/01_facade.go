package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

func main() {
	h := House{}
	h.turnOnSystems()
	h.shutDown()
}

// Plumbing System
type PlumbingSystem struct {
}

func (ps *PlumbingSystem) setPressure(v int) {
	fmt.Printf("Setting pressure to %v \n", v)
	_ = v
}

func (ps *PlumbingSystem) turnOn() {
	fmt.Println("Turning On Plumbing System")
}

func (ps *PlumbingSystem) turnOff() {
	fmt.Println("Turning Off Plumbing System")
}

// Electrical System
type ElectricalSystem struct {
}

func (es *ElectricalSystem) setVoltage(v int) {
	fmt.Printf("Setting voltage to %v \n", v)
	_ = v
}

func (es *ElectricalSystem) turnOn() {
	fmt.Println("Turning On Electrical System")
}

func (es *ElectricalSystem) turnOff() {
	fmt.Println("Turning Off Electrical System")
}

// House (Facade)
type House struct {
	ps PlumbingSystem
	es ElectricalSystem
}

func (h *House) turnOnSystems() {
	h.es.setVoltage(220)
	h.es.turnOn()
	h.ps.setPressure(500)
	h.ps.turnOn()
}

func (h *House) shutDown() {
	h.ps.turnOff()
	h.es.turnOff()
}
