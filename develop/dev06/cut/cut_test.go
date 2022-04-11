package cut

import "testing"

type cutTest struct {
	str          string
	f            int
	d            string
	s            bool
	expectedStr  string
	expectedBool bool
}

var cutTests = []cutTest{
	{"Hello\tWorld!", 0, "\t", false, "Hello", true},
	{"", 1, "\t", false, "", false},
	{"Hello World!", 1, " ", false, "World!", true},
	{"Hello, World!", 1, "\t", true, "", false},
}

func TestCut(t *testing.T) {
	for _, test := range cutTests {
		if res, ok := cut(test.str, test.f, test.d, test.s); res != test.expectedStr || ok != test.expectedBool {
			t.Errorf("Output %v %v not equal to expected %v %v", res, ok, test.expectedStr, test.expectedBool)
		}
	}
}
