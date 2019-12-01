package main

import (
	"flag"
)

type netcatArgs struct {
	Help *bool

	// Args required in listening mode.
	Listen *bool
	Port   *uint

	// Use Ipv4 or IPv6 to connect.
	IPv4 *bool
	IPv6 *bool
}

// Init the args struct
func configArgs() *netcatArgs {
	args := netcatArgs{}
	args.Help = flag.Bool("help", false, "Print help message.")
	args.Listen = flag.Bool("l", false, "Whether to start a server to listen on incoming connection")
	args.Port = flag.Uint("p", 0, "The listening port.")
	args.IPv4 = flag.Bool("4", true, "Use IPv4 address only.")
	args.IPv6 = flag.Bool("6", false, "Use IPv6 address only.")
	return &args
}
