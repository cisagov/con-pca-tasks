package main

import (
	"flag"
	"log"
	"os"
)

var Version = "0.0.1"

func version() {
	v := flag.Bool("version", false, "prints current app version")
	flag.Parse()
	if *v {
		log.SetFlags(0)
		log.Println(string(Version))
		os.Exit(0)
	}
}
