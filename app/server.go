package main

import (
	"bufio"
	"fmt"
	"log"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection: ", err.Error())

		}
		go func(conn net.Conn) {
			defer conn.Close()
			s := bufio.NewScanner(conn)
			for s.Scan() {
				if s.Err() != nil {
					log.Fatalln("Error scan: ", s.Err())
				}

				if s.Text() == "ping" {
					buf := bufio.NewWriter(conn)
					if _, err = buf.WriteString("+PONG\r\n"); err != nil {
						log.Fatalln("Error write to buf ", err.Error())
					}

					if err := buf.Flush(); err != nil {
						log.Fatalln("Error reply: ", err.Error())
					}
				}
			}

		}(conn)
	}
}
