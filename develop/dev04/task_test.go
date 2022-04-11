package main

import "testing"

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type findAnagramsTest struct {
	data     []string
	expected map[string][]string
}

var findAnagramsTests = []findAnagramsTest{
	{
		[]string{"Пятак", "пятка", "Тяпка", "листок", "слиток", "столиК", "привет"},
		map[string][]string{
			"пятак":  {"пятак", "пятка", "тяпка"},
			"листок": {"листок", "слиток", "столик"},
		},
	},
	{
		[]string{"австралопитек", "ватерполистка", "обезьянство", "светобоязнь", "импортер", "пирометр", "реимпорт"},
		map[string][]string{
			"австралопитек": {"австралопитек", "ватерполистка"},
			"импортер":      {"импортер", "пирометр", "реимпорт"},
			"обезьянство":   {"обезьянство", "светобоязнь"},
		},
	},
}

func TestGetAnagrams(t *testing.T) {
	for _, test := range findAnagramsTests {
		result := findAnagrams(test.data)
		if len(result) != len(test.expected) {
			t.Errorf("Incorrect result. Expect: %v, Got %v\n", test.expected, result)
		}
		for k, v := range result {
			if !Equal(v, test.expected[k]) {
				t.Errorf("Incorrect result. Expect: %v, Got %v\n", test.expected[k], v)
			}
		}
	}
}
