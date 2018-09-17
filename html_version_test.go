package HTMLVersion

import (
	"testing"
)

var htmlDocument = `
		<!DOCTYPE html>
		<html>
			<head></head>
			<body>
				<h1>Title</h1>
			</body>
		</html>
	`
var htmlURL = "https://golang.org/"

func AssertOutput(t *testing.T, got string, expected string) {
	t.Helper()

	if expected != got {
		t.Errorf("got %s want %s", got, expected)
	}
}

func TestDetectFromString(t *testing.T) {

	t.Run("should return html version from string", func(t *testing.T) {
		expected := "HTML 5"
		got := DetectFromString(htmlDocument)

		AssertOutput(t, got, expected)

	})

}

func TestDetectFromURL(t *testing.T) {

	t.Run("should return html version from URL", func(t *testing.T) {
		expected := "HTML 5"
		got, err := DetectFromURL(htmlURL)

		if err != nil {
			t.Error(err)
		}

		AssertOutput(t, got, expected)

	})

}

func TestDetectFromFile(t *testing.T) {

	t.Run("should return html version from File", func(t *testing.T) {
		expected := "HTML 4.01 Strict"
		got, err := DetectFromFile("testdata/test.html")

		if err != nil {
			t.Error(err)
		}

		AssertOutput(t, got, expected)

	})

}
