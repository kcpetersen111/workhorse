package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func StreamFile(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
	}

	res := "{'success':'true'}"
	conn.WriteJSON(res)
	// w.WriteHeader(201)
	// w.Write([]byte(res))
}

func recvFile(conn *websocket.Conn) (*os.File, error) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		return nil, fmt.Errorf("Error creating file: %v", err)
	}
	for {
		conn.ReadMessage()
		typ, msg, err := conn.ReadMessage()
		if typ == websocket.BinaryMessage {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Error reading json: %v", err)
		}
		sb := strings.Builder{}
		sb.Write(msg)
		_, err = f.WriteString(sb.String())
		if err != nil {
			return nil, fmt.Errorf("Error writing to the temp file: %v", err)
		}
	}
	return f, nil
}
