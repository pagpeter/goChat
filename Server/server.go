package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn, connections *[]net.Conn) {
	// handle a connection
	for {
		// get message
		encoded, _ := bufio.NewReader(conn).ReadString('\n')
		m, success := decodeMessage(encoded)
		m.Timestamp = time.Now().Unix() // add timestamp

		// message is a valid message
		if success {
			// logging messages
			t := time.Unix(m.Timestamp, 0)
			fmt.Printf("%s [%s] - %s\n", t.Format("15:04:05"), m.Author, m.Text)

			// send the message to all connections
			for _, conn := range *connections {
				conn.Write([]byte(m.Encode()))
			}
		} else {
			// invalid message = closing connection
			fmt.Println("Closing connection")
			conn.Close()
			return
		}

	}
}

func server(addr string) {
	var connections []net.Conn

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[+] Started server on %s\n", addr)

	for {
		conn, err := ln.Accept()
		connections = append(connections, conn)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("New user joined!")
		go handleConnection(conn, &connections)
	}

}
