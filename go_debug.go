package go_debug

import (
	"log"
	"flag"
)

var (
	debug = flag.Bool("debug", false, "Print debugging info")
)

func init() {
	flag.Parse()
}

func Debug(s string){
	if !*debug {
		return
	}
	log.Printf(s)
}