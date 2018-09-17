// package HTMLVersion provides methods for detecting version
// of a HTML Document from a string, url or filename
package HTMLVersion

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

var doctypes = make(map[string]string)

func init() {
	doctypes["HTML 4.01 Strict"] = `"-//W3C//DTD HTML 4.01//EN"`
	doctypes["HTML 4.01 Transitional"] = `"-//W3C//DTD HTML 4.01 Transitional//EN"`
	doctypes["HTML 4.01 Frameset"] = `"-//W3C//DTD HTML 4.01 Frameset//EN"`
	doctypes["XHTML 1.0 Strict"] = `"-//W3C//DTD XHTML 1.0 Strict//EN"`
	doctypes["XHTML 1.0 Transitional"] = `"-//W3C//DTD XHTML 1.0 Transitional//EN"`
	doctypes["XHTML 1.0 Frameset"] = `"-//W3C//DTD XHTML 1.0 Frameset//EN"`
	doctypes["XHTML 1.1"] = `"-//W3C//DTD XHTML 1.1//EN"`
	doctypes["HTML 5"] = `<!DOCTYPE html>`
}

func checkDoctype(html string) string {
	var version = "UNKNOWN"

	for doctype, matcher := range doctypes {
		match := strings.Contains(html, matcher)

		if match == true {
			version = doctype
		}
	}

	return version
}

// Detect HTML Document Version From a string
func DetectFromString(htmlDoc string) string {

	htmlVersion := checkDoctype(htmlDoc)

	return htmlVersion

}

// Detect HTML Version From a http URL
func DetectFromURL(url string) (string, error) {

	var htmlVersion string

	resp, err := http.Get(url)
	if err != nil {
		return htmlVersion, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return htmlVersion, errors.New("Error Retrieving Document")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return htmlVersion, err
	}

	html, err := doc.Html()
	if err != nil {
		return htmlVersion, err
	}

	htmlVersion = checkDoctype(html)

	return htmlVersion, nil

}

// Detect HTML Document Version from a filename
func DetectFromFile(filename string) (string, error) {

	var htmlVersion string

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return htmlVersion, err
	}

	htmlVersion = checkDoctype(string(content[:]))

	return htmlVersion, nil

}
