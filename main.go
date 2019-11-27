package main

import (
	"flag"
	"os"

	"github.com/capy0812/netcat/argparser"
)

func main() {
	args := argparser.Init()
	flag.Parse()
	if *args.Help == true {
		flag.PrintDefaults()
		os.Exit(0)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

}
