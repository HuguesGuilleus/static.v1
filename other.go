// static.v1
// Copyright (c) 2019, HuguesGuilleus
// BSD 3-Clause License

// Create easily a http.Handler for CSS anf Js files to a Web server.
//
// Usage:
// 	http.Handle("/style", static.Css("style/"))
// 	http.Handle("/js", static.Js("js/"))
package static

import (
	"log"
	"time"
)

var (
	// Dev disable the minifing and reduce the sleeping between two updates.
	Dev bool = false
	// The duration of sleeping in standard mode
	SleepDev time.Duration = 200 * time.Millisecond
	// The duration of sleeping in developpment mode
	SleepProd time.Duration = 10 * time.Minute
)

// Sleep according to the mode.
func sleep() {
	if Dev {
		time.Sleep(SleepDev)
	} else {
		time.Sleep(SleepProd)
	}
}

// Panic if the error is not nil.
func panicing(err error) {
	if err != nil {
		panic(err)
	}
}

// Reco recover and print the error if no nil.
func reco() {
	err := recover()
	if err != nil {
		log.Println("error:", err)
	}
}
