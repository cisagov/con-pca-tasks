package main

import (
	"flag"
	"fmt"
	"os"
)

var Version = "0.0.1"

func version() {
	v := flag.Bool("version", false, "prints current app version")
	flag.Parse()
	if *v {
		fmt.Println(string(Version))
		os.Exit(0)
	}
}
