package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/websocket"
	websocketserver "github.com/polluxdev/go-serverx/websocket/server"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	// New router
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		conn, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}

			log.Printf("Received: %s", msg)
			if err := conn.WriteMessage(msgType, msg); err != nil {
				break
			}
		}
	})

	// New server
	server := websocketserver.New(mux, websocketserver.Port("8080"))

	// Start server
	server.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		log.Println("shutting down")
	case err := <-server.Notify():
		log.Fatal(err)
	}

	// Shutdown
	err := server.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
