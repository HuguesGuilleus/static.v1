// static.v1
// Copyright (c) 2019, HuguesGuilleus
// BSD 3-Clause License

package static

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/xml"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type HtmlServer struct {
	data []byte
}

func Html(path string) (f *HtmlServer) {
	f = &HtmlServer{}
	go func() {
		for {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println("Error:", err)
			} else {
				f.data = htmlMinify(data)
			}
			sleep()
		}
	}()
	return
}

func (h *HtmlServer) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Write(h.data)
}

func htmlMinify(in []byte) []byte {
	if Dev {
		return in
	}
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
		KeepEndTags:      true,
	})
	m.AddFunc("image/svg+xml", xml.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	out, _ := m.Bytes("text/html", in)
	return out
}
