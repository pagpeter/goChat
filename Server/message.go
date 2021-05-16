package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

type Message struct {
	Text      string `json:"text"`
	Author    string `json:"author"`
	Timestamp int64  `json:"timestamp"`
}

func (m Message) Encode() string {
	// struct to JSON
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	// JSON to base64
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(j))
	return base64Encoded + seperator
}

func decodeMessage(raw string) (Message, bool) {
	// base64 decode
	decoded, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		//fmt.Println("Error decoding message")
		return Message{}, false
	}

	// JSON to struct
	var m Message
	if json.Unmarshal(decoded, &m) != nil {
		//fmt.Println("Error unmarshalling message")
		return Message{}, false
	}
	return m, true
}
