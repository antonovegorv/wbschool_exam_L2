package wget

import (
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// Wget — struct for all fields and methods of wget util.
type Wget struct {
	dir      string
	hostname string
}

// New — constructor of the Wget struct.
func New(dir, hostname string) *Wget {
	return &Wget{dir: dir, hostname: hostname}
}

// Run — main method to start util.
func (wg *Wget) Run() error {
	// Get main page.
	resp, err := http.Get(wg.hostname)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Don't forget to download main page.
	wg.download(wg.hostname)

	// Find all links.
	links, err := wg.getLinks(resp.Body)
	if err != nil {
		return err
	}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(links))

	// Download all sublinks.
	for _, l := range links {
		go func(l string) {
			defer waitGroup.Done()
			wg.download(l)
		}(l)
	}

	waitGroup.Wait()

	return nil
}

// Gets all sublinks of src and href attributes.
func (wg *Wget) getLinks(r io.Reader) ([]string, error) {
	tokenizer := html.NewTokenizer(r)
	u := make(map[string]bool)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}

			return nil, err
		}

		token := tokenizer.Token()
		for _, attr := range token.Attr {
			if attr.Key == "href" || attr.Key == "src" {
				if strings.HasPrefix(attr.Val, "/") {
					u[wg.hostname+attr.Val] = true
				} else if strings.HasPrefix(attr.Val, ".") {
					u[wg.hostname+strings.TrimLeft(attr.Val, ".")] = true
				} else {
					u[attr.Val] = true
				}
			}
		}
	}

	links := make([]string, 0, len(u))
	for k := range u {
		links = append(links, k)
	}

	return links, nil
}

// Downloads the file.
func (wg *Wget) download(url string) error {
	log.Printf("Downloading %v\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filename, err := wg.getFilename(url, resp.Header)
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(wg.dir, filename))
	if err != nil {
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

// Gets the name of the file.
func (wg *Wget) getFilename(url string, header http.Header) (string, error) {
	ct := header.Get("Content-Type")

	mt, _, err := mime.ParseMediaType(ct)
	if err != nil {
		return "", err
	}

	ext := mt
	if strings.Contains(mt, "/") {
		ext = strings.Split(mt, "/")[1]
	}

	bn := path.Base(url)
	if strings.Contains(bn, "?") {
		bn = strings.Split(bn, "?")[0]
	}

	if path.Ext(bn) == "" && ext != "" {
		return bn + "." + ext, nil
	}

	return bn, nil
}
