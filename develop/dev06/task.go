package main

import (
	"flag"
	"log"

	"github.com/antonovegorv/wbschool_exam_L2/develop/dev06/cut"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	f := flag.Int("f", 0, "choose a field (column)")
	d := flag.String("d", "\t", "choose a delimiter")
	s := flag.Bool("s", false, "lines with delimiter only")

	flag.Parse()

	filename := flag.Arg(0)

	c := cut.New(filename, *f, *d, *s)
	if err := c.Run(); err != nil {
		log.Fatalln(err)
	}
}
