package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/polluxdev/go-serverx/http"
)

func main() {
	// New router
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// New server
	server := http.New(router, http.Port("8080"))

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
