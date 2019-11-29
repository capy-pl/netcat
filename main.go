package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := Init()
	flag.Parse()
	if *args.Help == true {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *args.Listen {
		if *args.Port != 0 {
			fmt.Printf("Start to listening on port %v.\n", *args.Port)
		} else {
			fmt.Fprintf(os.Stderr, "Error: Port not provided.\n")
			os.Exit(1)
		}
	}
	// host := flag.Arg(0)
	// port := flag.Arg(1)
}
