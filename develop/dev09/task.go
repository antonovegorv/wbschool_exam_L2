package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/antonovegorv/wbschool_exam_L2/develop/dev09/wget"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// https://vaggroupsyzon.herokuapp.com/

func main() {
	flag.Parse()

	dir := strings.TrimSpace(flag.Arg(0))
	hostname := strings.TrimRight(strings.TrimSpace(flag.Arg(1)), "/")

	if dir == "" || hostname == "" {
		fmt.Println("Please, provide directory and a hostname!")
		return
	}

	wg := wget.New(dir, hostname)
	if err := wg.Run(); err != nil {
		log.Fatalln(err)
	}
}
