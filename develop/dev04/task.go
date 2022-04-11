package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	w := []string{"австралопитек", "ватерполистка", "обезьянство", "светобоязнь", "импортер", "пирометр", "реимпорт"}
	a := findAnagrams(w)
	fmt.Println(a)
}

// findAnagrams — makes a map with anagrams.
func findAnagrams(w []string) map[string][]string {
	t := make(map[string][]string) // Initialize temporary map.
	r := make(map[string][]string) // Initialize result map.

	// We go through each word and sort its letters. Due to this, we organize a map from anagrams.
	for _, v := range w {
		word := strings.ToLower(v)
		wordSlice := strings.Split(word, "")
		sort.Strings(wordSlice)
		t[strings.Join(wordSlice, "")] = append(t[strings.Join(wordSlice, "")], word)
	}

	// Exclude duplicates from temporary map.
	for k := range t {
		set := set(t[k])

		if len(set) > 1 {
			r[t[k][0]] = set
		}
	}

	return r
}

// set — makes set from a slice.
func set(slice []string) []string {
	u := make(map[string]bool)

	for _, v := range slice {
		u[v] = true
	}

	keys := make([]string, 0, len(u))
	for k := range u {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
