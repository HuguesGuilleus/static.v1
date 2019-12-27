// static.v1
// Copyright (c) 2019, HuguesGuilleus
// BSD 3-Clause License

package static

import (
	"bytes"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// An regexp to identify Js file.
var suffixJs = regexp.MustCompile(`\.js`)

// A HTTP server for Js.
type JsServer []byte

func (c JsServer) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-type", "application/javascript")
	w.Write(c)
}

// Css create a http.Handler who response with all the Js file from directory.
// Regularly the server are update.
func Js(directory string) *JsServer {
	serv := &JsServer{}
	go func() {
		for {
			serv.update(directory)
			sleep()
		}
	}()
	return serv
}

// Read and minify all file from dir to updage the JsServer.
func (c *JsServer) update(dir string) {
	defer reco()
	data := make([]byte, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		panicing(err)
		if info.IsDir() || !suffixJs.MatchString(path) {
			return nil
		}
		f, err := ioutil.ReadFile(path)
		panicing(err)
		data = append(data, f...)
		return nil
	})
	*c = jsMinify(data)
}

func jsMinify(data []byte) JsServer {
	if !Dev {
		in := bytes.NewBuffer(data)
		out := &bytes.Buffer{}
		m := minify.New()
		m.AddFunc("application/javascript", js.Minify)
		panicing(m.Minify("application/javascript", out, in))
		return JsServer(out.Bytes())
	} else {
		return JsServer(data)
	}
}
