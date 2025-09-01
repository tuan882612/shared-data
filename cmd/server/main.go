package main

import (
	"flag"
	"log"
	"os"
)

var (
	host string
	port int
)

func main() {
	flag.StringVar(&host, "host", "", "Server Host")
	flag.IntVar(&port, "port", 0, "Server Port")
	flag.Parse()

	if len(host) < 1 || port < 1 {
		log.Print("Invalid Host or Port")
		os.Exit(1)
	}
}
