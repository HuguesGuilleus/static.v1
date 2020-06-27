# static.v1

[![GoDoc](https://godoc.org/github.com/HuguesGuilleus/static.v1?status.svg)](https://godoc.org/github.com/HuguesGuilleus/static.v1)

[Use the new version](https://github.com/HuguesGuilleus/static.v2)

Create easily a `http.Handler` for CSS anf Js files to a Web server.

## Installation
```bash
go get github.com/HuguesGuilleus/static.v1
```

## Example
Create a directory `style` and a other `js` with some files and run the folowing
program.

```go
package main

import (
	"github.com/HuguesGuilleus/static.v1"
	"log"
	"net/http"
)

func main() {
	// static.Dev = false
	log.Println("go!")
	http.Handle("/style", static.Css("style/"))
	http.Handle("/js", static.Js("js/"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "text/html; charset=utf-8")
		w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<link rel="stylesheet" href="/style">
		<script src="/js" charset="utf-8"></script>
	</head>
	<body>
		<p>Hello World!</p>
	</body>
</html>`))
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
```
