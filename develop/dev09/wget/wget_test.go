package wget

import (
	"net/http"
	"os"
	"testing"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type getLinksTest struct {
	hostname string
	expected []string
}

var getLinksTests = []getLinksTest{
	{
		"https://vaggroupsyzon.herokuapp.com/about",
		[]string{
			`https://vaggroupsyzon.herokuapp.com/about/`,
			`https://vaggroupsyzon.herokuapp.com/about/favicon.ico`,
			`https://vaggroupsyzon.herokuapp.com/about/_app/assets/pages/__layout.svelte-b2d4ec74.css`,
			`https://vaggroupsyzon.herokuapp.com/about/_app/start-13e59186.js`,
			`https://kit.svelte.dev/`,
			`https://vaggroupsyzon.herokuapp.com/about/images/VolkswagenTitle.svg`,
			`https://vaggroupsyzon.herokuapp.com/about/brands`,
			`https://tailwindcss.com/`,
			`https://vwgroup.ru/`,
			`https://www.volkswagenag.com/en.html`,
			`https://carobka.ru/`,
			`https://vaggroupsyzon.herokuapp.com/about/_app/pages/__layout.svelte-0a20b182.js`,
			`https://vaggroupsyzon.herokuapp.com/about/_app/pages/about.svelte-a0a0aa68.js`,
			`https://vaggroupsyzon.herokuapp.com/about/about`,
			`https://vaggroupsyzon.herokuapp.com/about/_app/chunks/vendor-cb4cbf22.js`,
			`https://developer.mozilla.org/en-US/docs/Learn`,
			`https://mai.ru/`,
		},
	},
	{
		"https://cars-disney-wiki.herokuapp.com/characters/1",
		[]string{
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/slides/slide-08.webp`,
			`https://github.com/antonovegorv/disney-cars-wiki`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/_app/pages/__layout.svelte-63413514.js`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/avatar.webp`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/slides/slide-02.webp`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/slides/slide-07.webp`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/slides/slide-09.webp`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/_app/start-63bb7f0c.js`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/_app/chunks/characters-0b1172fc.js`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/about`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/slides/slide-03.webp`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/slides/slide-05.webp`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/_app/chunks/vendor-f8de5bfc.js`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/_app/pages/characters/[id].svelte-df45f3a7.js`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/_app/assets/start-d5b4de3e.css`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/_app/assets/pages/__layout.svelte-4a34f711.css`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/characters`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/slides/slide-01.webp`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/slides/slide-04.webp`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/images/characters/lightning-mcqueen/slides/slide-06.webp`,
			`https://cars-disney-wiki.herokuapp.com/characters/1/favicon.png`,
		},
	},
}

type runTest struct {
	hostname string
	dir      string
	expected []string
}

var runTests = []runTest{
	{
		"https://web.ics.purdue.edu/~gchopra/class/public/pages/webdesign/05_simple.html",
		"wget_test",
		[]string{
			"05_simple.html",
			"index.html",
			"www.yahoo.com",
		},
	},
}

func TestGetLinks(t *testing.T) {
	for _, test := range getLinksTests {
		wg := New("", test.hostname)
		resp, err := http.Get(wg.hostname)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()
		res, _ := wg.getLinks(resp.Body)

		if len(res) != len(test.expected) {
			t.Errorf("Output: %v. Expected: %v", res, test.expected)
		}

		for _, v := range res {
			if !contains(test.expected, v) {
				t.Errorf("Output: %v. Expected: %v", res, test.expected)
			}
		}
	}
}

func TestRun(t *testing.T) {
	for _, test := range runTests {
		wg := New(test.dir, test.hostname)
		if err := wg.Run(); err != nil {
			t.Error(err)
		}

		files, err := os.ReadDir(test.dir)
		if err != nil {
			t.Error(err)
		}

		for _, f := range files {
			if !contains(test.expected, f.Name()) {
				t.Errorf("Output: %v. Expected: %v", files, test.expected)
			}
		}
	}
}
