package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var seperator = "\n\r\n\r"
var addr = ":1000"

func printMessages(c chan Message) {
	// print all incomming messages for a channel
	for {
		m := <-c
		t := time.Unix(m.Timestamp, 0)
		fmt.Printf("%s [%s] - %s\n", t.Format("15:04:05"), m.Author, m.Text)
	}
}

func main() {
	name := os.Args[1]
	incomming, outgoing := client(addr)

	go printMessages(incomming)

	for {
		// get input
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		// send the message
		m := Message{Author: name, Text: strings.TrimSpace(text)}
		outgoing <- m
	}
}
