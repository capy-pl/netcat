package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func listen(port uint) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	log.Printf("Receive connection from %v.", conn.RemoteAddr())
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("Connection Closed by %v.\n", conn.RemoteAddr())
			} else {
				log.Println("Read Error.")
			}
			break
		}
		log.Printf("%v bytes received.\n", len(buf))
		fmt.Printf("%s", fmt.Sprintf("%s", buf))
	}

	// buf, err := ioutil.ReadAll(os.Stdin)

	// if err != nil {
	// 	log.Panicln(err)
	// }

	// if len(buf) > 0 {
	// 	conn.Write(buf)
	// }
}
