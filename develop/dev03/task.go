package main

import (
	"flag"
	"log"

	"github.com/antonovegorv/wbschool_exam_L2/develop/dev03/sorter"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	k := flag.Int("k", 0, "column to sort")
	n := flag.Bool("n", false, "compare according to string numerical value")
	r := flag.Bool("r", false, "reverse the result of comparisons")
	u := flag.Bool("u", false, "output unique lines only")
	m := flag.Bool("M", false, "sort by monthes")
	b := flag.Bool("b", false, "trim right spaces")
	c := flag.Bool("c", false, "check for sorted input; do not sort")
	h := flag.Bool("h", false, "sort by numeric with suffix")
	flag.Parse()

	options := make(map[string]bool)
	options["n"] = *n
	options["r"] = *r
	options["u"] = *u
	options["M"] = *m
	options["b"] = *b
	options["c"] = *c
	options["h"] = *h

	if filename := flag.Arg(0); filename != "" {
		s := sorter.New(filename, options, *k)
		s.Run()
	} else {
		log.Fatalln("Please, provide a file to sort!")
	}
}
