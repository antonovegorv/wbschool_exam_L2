package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

func main() {
	n := &navigator{}

	n.setRouteStrategy(&roadStrategy{})
	n.getRoute()

	n.setRouteStrategy(&publicTransportStrategy{})
	n.getRoute()

	n.setRouteStrategy(&walkingStrategy{})
	n.getRoute()
}

// Route Strategy
type routeStrategy interface {
	buildRoute()
}

// 1st algo
type roadStrategy struct{}

func (rs *roadStrategy) buildRoute() {
	fmt.Println("Building route by road")
}

// 2nd algo
type publicTransportStrategy struct{}

func (pts *publicTransportStrategy) buildRoute() {
	fmt.Println("Building route by public transport")
}

// 3rd algo
type walkingStrategy struct{}

func (ws *walkingStrategy) buildRoute() {
	fmt.Println("Building route by walking")
}

// Context
type navigator struct {
	rs routeStrategy
}

func (n *navigator) setRouteStrategy(rs routeStrategy) {
	n.rs = rs
}

func (n *navigator) getRoute() {
	n.rs.buildRoute()
}
