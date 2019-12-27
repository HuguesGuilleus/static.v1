// static.v1
// Copyright (c) 2019, HuguesGuilleus
// BSD 3-Clause License

package static

import (
	"io/ioutil"
	"log"
	"text/template"
)

// Read target file and return the template, regularly it will be update.
func Templ(target string) (t *template.Template) {
	t = template.New("")
	go func() {
		for {
			f, err := ioutil.ReadFile(target)
			if err != nil {
				log.Printf("%s, %v", target, err)
			} else {
				t.Parse(string(f))
			}
			sleep()
		}
	}()
	return t
}
