package sorter

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMonths() map[string]int {
	return map[string]int{
		"january":   0,
		"february":  1,
		"march":     2,
		"april":     3,
		"may":       4,
		"june":      5,
		"july":      6,
		"august":    7,
		"september": 8,
		"october":   9,
		"november":  10,
		"december":  11,
	}
}

func isMonth(month string) bool {
	months := make([]string, 0, len(getMonths()))

	for m := range getMonths() {
		months = append(months, m)
	}

	for _, m := range months {
		if m == month {
			return true
		}
	}

	return false
}

// SortStrategy — strategy how we sort lines.
type SortStrategy interface {
	Sort(map[string][]string) []string
}

// DefaultSort — default sort strategy.
type DefaultSort struct{}

// Sort — sorts lines with default sort.
func (ds *DefaultSort) Sort(data map[string][]string) []string {
	keys := make([]string, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	sorted := make([]string, 0) // Could pass s.lines to find the capacity.

	for _, k := range keys {
		sorted = append(sorted, data[k]...)
	}

	return sorted
}

// NumericSort — sorts by numerics strategy.
type NumericSort struct{}

// Sort — sorts lines with numeric sort.
func (ns *NumericSort) Sort(data map[string][]string) []string {
	numKeys := make([]int, 0, len(data))
	strKeys := make([]string, 0, len(data))

	for k := range data {
		if num, err := strconv.Atoi(k); err == nil {
			numKeys = append(numKeys, num)
		} else {
			strKeys = append(strKeys, k)
		}
	}

	sort.Ints(numKeys)
	sort.Strings(strKeys)

	sorted := make([]string, 0)

	for _, k := range strKeys {
		sorted = append(sorted, data[k]...)
	}

	for _, k := range numKeys {
		sorted = append(sorted, data[strconv.Itoa(k)]...)
	}

	return sorted
}

// MonthSort — sort strategy by month.
type MonthSort struct{}

// Sort — sorts by months.
func (ms *MonthSort) Sort(data map[string][]string) []string {
	monthKeys := make([]string, 0, len(data))
	strKeys := make([]string, 0, len(data))

	for k := range data {
		if isMonth(strings.ToLower(k)) {
			monthKeys = append(monthKeys, k)
		} else {
			strKeys = append(strKeys, k)
		}
	}

	sort.Strings(strKeys)

	for i := 0; i < len(monthKeys); i++ {
		for j := 0; j < len(monthKeys)-i-1; j++ {
			m1 := strings.ToLower(monthKeys[j])
			m2 := strings.ToLower(monthKeys[j+1])

			if getMonths()[m1] > getMonths()[m2] {
				monthKeys[j], monthKeys[j+1] = monthKeys[j+1], monthKeys[j]
			}
		}
	}

	sorted := make([]string, 0)

	for _, k := range strKeys {
		sorted = append(sorted, data[k]...)
	}

	for _, k := range monthKeys {
		sorted = append(sorted, data[k]...)
	}

	return sorted
}

// NumericSuffixSort — sort strategy.
type NumericSuffixSort struct{}

// Sort — sorts by numeric and suffix.
func (nss *NumericSuffixSort) Sort(data map[string][]string) []string {
	return nil
}

// Sorter — main struct to store all fields and methods for sort util.
type Sorter struct {
	filename string
	options  map[string]bool
	k        int
	lines    []string
	strategy SortStrategy
}

// New — constructor for the Sorter struct.
func New(filename string, options map[string]bool, k int) *Sorter {
	return &Sorter{filename: filename, options: options, k: k}
}

// Run — main method to start util.
func (s *Sorter) Run() error {
	// Read lines from file.
	if err := s.readLines(); err != nil {
		return err
	}

	// Set sort strategy.
	if s.options["n"] {
		s.setStrategy(&NumericSort{})
	} else if s.options["M"] {
		s.setStrategy(&MonthSort{})
	} else if s.options["h"] {
		s.setStrategy(&NumericSuffixSort{})
	} else {
		s.setStrategy(&DefaultSort{})
	}

	// Sort lines.
	m := s.groupByColumn()
	result := s.strategy.Sort(m)

	// Reverse result if necessary.
	if s.options["r"] {
		s.reverse(result)
	}

	for _, l := range result {
		fmt.Println(l)
	}

	return nil
}

func (s *Sorter) setStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) groupByColumn() map[string][]string {
	m := make(map[string][]string)

	for _, l := range s.lines {
		columns := strings.Split(l, " ")

		sortColumn := columns[0]
		if s.k < len(columns) {
			sortColumn = columns[s.k]
		}

		sortColumn = strings.TrimSpace(sortColumn)

		m[sortColumn] = append(m[sortColumn], l)
	}

	return m
}

func (s *Sorter) readLines() error {
	content, err := os.ReadFile(s.filename)
	if err != nil {
		return err
	}

	if s.options["u"] {
		u := make(map[string]bool)

		for _, line := range bytes.Split(content, []byte{'\n'}) {
			u[string(line)] = true
		}

		for k := range u {
			if s.options["b"] {
				k = strings.TrimSpace(k)
			}
			s.lines = append(s.lines, k)
		}
	} else {
		for _, line := range bytes.Split(content, []byte{'\n'}) {
			if s.options["b"] {
				line = bytes.TrimSpace(line)
			}
			s.lines = append(s.lines, string(line))
		}
	}

	return nil
}

func (s *Sorter) reverse(data []string) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
