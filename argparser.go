package main

import (
	"flag"
	"strconv"
)

type _uint16 uint16

type NetcatArgs struct {
	Help *bool

	// Args required in listening mode.
	Listen *bool
	Port   *uint

	// Use Ipv4 or IPv6 to connect.
	IPv4 *bool
	IPv6 *bool
}

func (num *_uint16) Set(value string) error {
	parsedNum, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return err
	}
	*num = _uint16(parsedNum)
	return nil
}

func (num *_uint16) String() string {
	return strconv.FormatUint(uint64(*num), 10)
}

// Init the args struct
func Init() *NetcatArgs {
	args := NetcatArgs{}
	args.Help = flag.Bool("help", false, "Print help message.")
	args.Listen = flag.Bool("l", false, "Whether to start a server to listen on incoming connection")
	args.Port = flag.Uint("p", 0, "The listening port.")
	args.IPv4 = flag.Bool("4", true, "Use IPv4 address only.")
	args.IPv6 = flag.Bool("6", false, "Use IPv6 address only.")
	return &args
}
