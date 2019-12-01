package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func connectTo(arg1 string, arg2 string) {
	var port uint64
	var err error
	host := "127.0.0.1"

	if len(arg1) > 0 && len(arg2) > 0 {
		// host and port are provided.
		host = arg1
		port, err = strconv.ParseUint(arg2, 10, 32)

		log.Fatal(err)
	} else if len(arg1) > 0 && len(arg2) == 0 {
		// only port are provided
		port, err = strconv.ParseUint(arg1, 10, 32)
	} else {
		log.Fatalln("Host and port must be provided.")
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	// connW := bufio.NewWriter(conn)

	for {
		buf, err := reader.ReadBytes('\n')

		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("End of line.")
				break
			} else {
				log.Panicln(err)
			}
		}

		if len(buf) > 0 {
			conn.Write(buf)
			log.Printf("%d bytes send.", len(buf))
		}
	}
}
