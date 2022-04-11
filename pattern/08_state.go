package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

func main() {
	dm := &dayMeals{s: &breakfast{}}

	dm.request()
	dm.request()
	dm.request()
	dm.request()
}

type state interface {
	getMeal(*dayMeals)
}

type dayMeals struct {
	s state
}

func (dm *dayMeals) setState(s state) {
	dm.s = s
}

func (dm *dayMeals) request() {
	dm.s.getMeal(dm)
}

type breakfast struct{}

func (b *breakfast) getMeal(dm *dayMeals) {
	fmt.Println("Breakfast: Here's your scrambled eggs and fried bacon for breakfast!")
	dm.setState(&dinner{})
}

type dinner struct{}

func (d *dinner) getMeal(dm *dayMeals) {
	fmt.Println("Dinner: Your magnificent borscht with green onions is already on the table!")
	dm.setState(&supper{})
}

type supper struct{}

func (s *supper) getMeal(dm *dayMeals) {
	fmt.Println("Supper: The beef stroganoff fried potatoes are getting cold. Hurry up!")
	dm.setState(&breakfast{})
}
