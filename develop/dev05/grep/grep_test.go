package grep

import (
	"strings"
	"testing"
)

func createOptionsInt(a, b, c int) map[string]*int {
	r := make(map[string]*int, 3)

	r["A"] = &a
	r["B"] = &b
	r["C"] = &c

	return r
}

func createOptionsBool(c, i, v, f, n bool) map[string]*bool {
	r := make(map[string]*bool, 5)

	r["c"] = &c
	r["i"] = &i
	r["v"] = &v
	r["F"] = &f
	r["n"] = &n

	return r
}

type grepTest struct {
	grep     Grep
	expected []string
}

var grepTests = []grepTest{
	{
		*New("vous", "grep_test.txt", createOptionsInt(0, 0, 0), createOptionsBool(false, false, false, false, false)),
		[]string{
			"Je vous aimais: et mon amour, peut-être,",
			"Mais plus sa peine en vous ne doit renaître.",
			"Je ne voudrais vous faire aucun chagrin.",
			"Je vous aimais sans bruit, sans rien attendre",
			"Je vous aimais d'un coeur si pur, si tendre.",
			"Qu'un autre, priez Dieu, vous aime autant.",
		},
	},
	{
		*New("vous", "grep_test.txt", createOptionsInt(0, 0, 0), createOptionsBool(false, false, true, false, true)),
		[]string{
			"1:N'est point au fond de l'âme encore éteint",
			"5:Jaloux et puis farouche en mon tourment,",
		},
	},
	{
		*New("voudrais", "grep_test.txt", createOptionsInt(1, 2, 0), createOptionsBool(false, false, false, false, false)),
		[]string{
			"N'est point au fond de l'âme encore éteint",
			"Mais plus sa peine en vous ne doit renaître.",
			"Je ne voudrais vous faire aucun chagrin.",
			"Je vous aimais sans bruit, sans rien attendre",
		},
	},
}

func TestGrep(t *testing.T) {
	for _, test := range grepTests {
		if err := test.grep.Start(); err != nil {
			t.Errorf("Error occurred while testing: %v", err)
		}

		if len(test.grep.result) != len(test.expected) {
			t.Errorf("Output %q not equal to expected %q", test.grep.result, test.expected)
		}

		for i := range test.grep.result {
			if strings.TrimSpace(test.grep.result[i]) != test.expected[i] {
				t.Errorf("Output %q not equal to expected %q", test.grep.result, test.expected)
			}
		}
	}
}
