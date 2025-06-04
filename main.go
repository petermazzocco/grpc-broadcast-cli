package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	mode := flag.String("mode", "server", "Mode to run: server or client")
	msg := flag.String("msg", "message", "The message you want to send to be broadcast on the server")
	flag.Parse()

	switch *mode {
	case "client":
		RunClient(msg)
	case "server":
		RunServer()
	default:
		log.Fatalf("Invalid mode: %s. Use 'server' or 'client'", *mode)
		os.Exit(1)
	}
}
