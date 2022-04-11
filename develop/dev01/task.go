package main

import (
	"fmt"
	"os"

	"github.com/antonovegorv/wbschool_exam_L2/develop/dev01/datetime"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

const layoutRFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
const exitFailure = 1

const hostname = "1.ru.pool.ntp.org"

func main() {
	// Create new datetime instance.
	dt := datetime.New(hostname)

	lt := dt.GetLocal()      // Get local time.
	ct, err := dt.GetExact() // Get exact time using ntp.

	// Check for error.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v; \n", err)
		os.Exit(exitFailure)
	}

	fmt.Printf("Local time: %v; \n", lt.Format(layoutRFC1123))
	fmt.Printf("Current time: %v; \n", ct.Format(layoutRFC1123))
}
