// static.v1
// Copyright (c) 2019, HuguesGuilleus
// BSD 3-Clause License

package static

import (
	"io/ioutil"
	"log"
	"net/http"
)

type FileServer struct {
	data []byte
	mime string
}

// Css create a http.Handler that response with the file from path.
// The mime can be empty.
func File(path, mime string) (f *FileServer) {
	f = &FileServer{
		mime: mime,
	}
	go func() {
		for {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println("Error:", err)
			} else {
				f.data = data
			}
			sleep()
		}
	}()
	return f
}

func (f *FileServer) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if f.mime != "" {
		w.Header().Add("Content-Type", f.mime)
	}
	w.Write(f.data)
}
