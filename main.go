package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	args := configArgs()
	flag.Parse()
	if *args.Help == true {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *args.Listen {
		if *args.Port != 0 {
			log.Printf("Start to listening on port %d\n", *args.Port)
			listen(*args.Port)
		} else {
			panic("Error: Port not provided.\n")
		}
	}

	connectTo(flag.Arg(0), flag.Arg(1))
}
