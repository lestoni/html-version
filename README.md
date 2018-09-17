# HTML Version

Detect HTML Version of document from string, url or file


## Installation

`go get github.com/lestoni/html-version`

## Usage

Detect from html document string

```
  htmlDocument := `<!DOCTYPE html>
    <html>
      <head></head>
      <body>
        <h1>Title</h1>
      </body>
    </html>`
  version := HTMLVersion.DetectFromString(htmlDocument)

```

Detect From URL

```
  version, err := HTMLVersion.DetectFromURL("http://golang.org/")

```

Detect From File

```
  version, err := HTMLVersion.DetectFromFile("testdata/test.html")

```

---

[![GoDoc](https://godoc.org/github.com/lestoni/html-version?status.svg)](https://godoc.org/github.com/lestoni/html-version)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)