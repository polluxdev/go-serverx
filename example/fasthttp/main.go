package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/polluxdev/go-serverx/fasthttp"
)

func main() {
	// New app
	app := fiber.New()
	app.Get("/random", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World 👋!")
	})

	// New server
	server := fasthttp.New(app.Handler(), fasthttp.Port("8080"))

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
