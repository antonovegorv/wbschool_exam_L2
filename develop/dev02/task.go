package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const escapedRune = 92

func main() {
	fmt.Println(unpack(`\5`))
}

// unpack — function to unpack string.
func unpack(s string) (string, error) {
	// Create slices to store the result and the number.
	var result, number []rune

	// Explicitly convert the string to a slice of runes.
	runes := []rune(s)

	// Check the string for potential errors before we start processing it.
	if err := validateString(runes); err != nil {
		return "", err
	}

	// Create a variable in which we will store the value of the previous character.
	var previous rune
	for i := 0; i < len(runes); i++ {
		// Create an intermediate variable that is the current symbol for convenience.
		current := runes[i]

		// We check the current symbol for three conditions:
		// 1. Is the current character a digit;
		// 2. Whether the current character is an escaped character;
		// 3. Everything else.
		if unicode.IsDigit(current) {
			number = append(number, current)
		} else if current == escapedRune {
			i++
			previous = runes[i]
			result = append(result, runes[i])
		} else {
			if err := cloneRune(&result, &number, previous); err != nil {
				return "", err
			}

			result = append(result, current)
			previous = current
		}
	}

	if err := cloneRune(&result, &number, previous); err != nil {
		return "", err
	}

	return string(result), nil
}

// cloneRune — clones "previous" rune "number" number of time and adds the result to the "result".
func cloneRune(result, number *[]rune, previous rune) error {
	if len(*number) != 0 {
		// Convert string presentation of a string to int.
		count, err := strconv.Atoi(string(*number))

		// Check for error.
		if err != nil {
			return err
		}

		// If the count is zero, then we need to remove the last character from the resulting
		// string, because we have already added it. Otherwise, in the loop, add the rune to the
		// result.
		if count != 0 {
			for j := 0; j < count-1; j++ {
				*result = append(*result, previous)
			}
		} else {
			*result = (*result)[:len(*result)-1]
		}

		// Zeroing out the number.
		*number = nil
	}

	return nil
}

// validateString — function to check if a string is correct.
func validateString(r []rune) error {
	// Check for the first character not to be a digit. And last character not to be an escaped one.
	if len(r) > 0 {
		if unicode.IsDigit(r[0]) {
			return fmt.Errorf("error: string starts with a number")
		} else if r[len(r)-1] == 92 {
			return fmt.Errorf("error: last symbol is escaped symbol")
		}
	}

	return nil
}
