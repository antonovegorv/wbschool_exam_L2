package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

func main() {
	level2 := &level2{}

	level1 := &level1{}
	level1.setNext(level2)

	level0 := &level0{}
	level0.setNext(level1)

	trainee1 := &trainee{name: "Egor Antonov"}
	trainee2 := &trainee{name: "Gopher Go", level0Done: true}

	level0.execute(trainee1)
	println()
	level0.execute(trainee2)
}

type trainee struct {
	name       string
	level0Done bool
	level1Done bool
	level2Done bool
}

type stage interface {
	execute(*trainee)
	setNext(stage)
}

// Level 0
type level0 struct {
	next stage
}

func (l0 *level0) execute(t *trainee) {
	if t.level0Done {
		fmt.Printf("%v has already completed L0! \n", t.name)
		l0.next.execute(t)
		return
	}

	fmt.Printf("%v has successfully completed L0! \n", t.name)
	t.level0Done = true
	l0.next.execute(t)
}

func (l0 *level0) setNext(next stage) {
	l0.next = next
}

// Level 1
type level1 struct {
	next stage
}

func (l1 *level1) execute(t *trainee) {
	if t.level1Done {
		fmt.Printf("%v has already completed L1! \n", t.name)
		l1.next.execute(t)
		return
	}

	fmt.Printf("%v has successfully completed L1! \n", t.name)
	t.level1Done = true
	l1.next.execute(t)
}

func (l1 *level1) setNext(next stage) {
	l1.next = next
}

type level2 struct {
	next stage
}

func (l2 *level2) execute(t *trainee) {
	if t.level2Done {
		fmt.Printf("%v has already completed L2! \n", t.name)
		return
	}

	fmt.Printf("%v has successfully completed L2! \n", t.name)
	t.level2Done = true
}

func (l2 *level2) setNext(next stage) {
	l2.next = next
}
