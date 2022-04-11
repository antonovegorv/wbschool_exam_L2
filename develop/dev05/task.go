package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/antonovegorv/wbschool_exam_L2/develop/dev05/grep"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	optionsInt := make(map[string]*int)
	optionsBool := make(map[string]*bool)

	optionsInt["A"] = flag.Int("A", 0, "print NUM lines of trailing context after matching lines")
	optionsInt["B"] = flag.Int("B", 0, "print NUM lines of leading context before matching lines")
	optionsInt["C"] = flag.Int("C", 0, "print NUM lines of output context")

	optionsBool["c"] = flag.Bool("c", false, "output amount of lines")
	optionsBool["i"] = flag.Bool("i", false, "ignore case distinctions in patterns and input data, so that characters that differ only in case  match each other")
	optionsBool["v"] = flag.Bool("v", false, "invert the sense of matching, to select non-matching lines")
	optionsBool["F"] = flag.Bool("F", false, "interpret PATTERNS as fixed strings, not regular expressions")
	optionsBool["n"] = flag.Bool("n", false, "prefix each line of output with the 1-based line number within its input file")

	flag.Parse()

	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	if pattern == "" || filename == "" {
		fmt.Println("Please, be sure to provide both PATTERN and FILENAME!")
		return
	}

	g := grep.New(pattern, filename, optionsInt, optionsBool)
	if err := g.Start(); err != nil {
		log.Fatalln(err)
	}
}
