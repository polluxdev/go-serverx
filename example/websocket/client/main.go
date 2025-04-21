package main

import (
	"bufio"
	"log"
	"os"

	"github.com/gorilla/websocket"
	websocketclient "github.com/polluxdev/go-serverx/websocket/client"
)

func main() {
	// New client connection
	client, err := websocketclient.New(websocketclient.Target("ws://localhost:8080/ws"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Read message from server
	go func() {
		for {
			_, msg, err := client.Conn().ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}

			log.Printf("Received: %s\n", msg)
		}
	}()

	// Write message and send it to connection
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if err := client.Conn().WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
