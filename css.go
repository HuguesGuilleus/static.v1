// static.v1
// Copyright (c) 2019, HuguesGuilleus
// BSD 3-Clause License

package static

import (
	"bytes"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// An regexp to identify CSS file.
var suffixCss = regexp.MustCompile(`\.css`)

// A HTTP server for CSS.
type CssServer []byte

func (c CssServer) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-type", "text/css; charset=utf-8")
	w.Write(c)
}

// Css create a http.Handler who response with all the CSS file from directory.
// Regularly the server are update.
func Css(directory string) *CssServer {
	serv := &CssServer{}
	go func() {
		for {
			serv.update(directory)
			sleep()
		}
	}()
	return serv
}

// Read and minify all file from dir to updage the CssServer.
func (c *CssServer) update(dir string) {
	defer reco()
	data := make([]byte, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		panicing(err)
		if info.IsDir() || !suffixCss.MatchString(path) {
			return nil
		}
		f, err := ioutil.ReadFile(path)
		panicing(err)
		data = append(data, f...)
		return nil
	})
	*c = cssMinify(data)
}

func cssMinify(data []byte) CssServer {
	if !Dev {
		in := bytes.NewBuffer(data)
		out := &bytes.Buffer{}
		m := minify.New()
		m.AddFunc("text/css", css.Minify)
		panicing(m.Minify("text/css", out, in))
		return CssServer(out.Bytes())
	} else {
		return CssServer(data)
	}
}
