package argparser

import "flag"

type NetcatArgs struct {
	Listen *bool
	Help   *bool
	Port   *uint16
}

type _uint16 uint16

// Init the args struct
func Init() *NetcatArgs {
	args := NetcatArgs{}
	args.Help = flag.Bool("help", false, "Print help message.")
	args.Listen = flag.Bool("listen", false, "Whether to start a server to listen on incoming connection")
	// args.Port = flag.Uint()
	return &args
}
