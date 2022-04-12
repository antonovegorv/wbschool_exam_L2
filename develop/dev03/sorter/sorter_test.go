package sorter

import "testing"

type defaultSortTest struct {
	lines    []string
	expected []string
}

var defaultSortTests = []defaultSortTest{
	{
		[]string{"hello", "gopher", "world", "very", "good"},
		[]string{"good", "gopher", "hello", "very", "world"},
	},
}

type numericSortTest struct {
	lines    []string
	expected []string
}

var numericSortTests = []numericSortTest{
	{
		[]string{"10", "2", "21", "15", "10"},
		[]string{"2", "10", "10", "15", "21"},
	},
}

type monthsSortTest struct {
	lines    []string
	expected []string
}

var monthsSortTests = []monthsSortTest{
	{
		[]string{"December", "May", "June", "March", "November", "January"},
		[]string{"January", "March", "May", "June", "November", "December"},
	},
}

func TestDefaultSort(t *testing.T) {
	for _, test := range defaultSortTests {
		s := Sorter{lines: test.lines}
		m := s.groupByColumn()
		s.setStrategy(&DefaultSort{})
		res := s.strategy.Sort(m)

		for i := range res {
			if res[i] != test.expected[i] {
				t.Errorf("Output %v; Expected %v;", res, test.expected)
			}
		}
	}
}

func TestNumericSort(t *testing.T) {
	for _, test := range numericSortTests {
		s := Sorter{lines: test.lines}
		m := s.groupByColumn()
		s.setStrategy(&NumericSort{})
		res := s.strategy.Sort(m)

		for i := range res {
			if res[i] != test.expected[i] {
				t.Errorf("Output %v; Expected %v;", res, test.expected)
			}
		}
	}
}

func TestMonthSort(t *testing.T) {
	for _, test := range monthsSortTests {
		s := Sorter{lines: test.lines}
		m := s.groupByColumn()
		s.setStrategy(&MonthSort{})
		res := s.strategy.Sort(m)

		for i := range res {
			if res[i] != test.expected[i] {
				t.Errorf("Output %v; Expected %v;", res, test.expected)
			}
		}
	}
}
