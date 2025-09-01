package main

import (
	"flag"
	"log/slog"
	"os"
	"shareddata/internal/config"
	"shareddata/internal/data"
	"shareddata/internal/server"
)

func main() {
	path := flag.String("config-path", "", "Server configuration path")
	flag.Parse()

	if len(*path) < 1 || path == nil {
		slog.Error("Path flag is empty")
		os.Exit(1)
	}

	conf, err := config.NewConfiguration(*path)
	if err != nil {
		os.Exit(1)
	}

	dataHandler := data.NewDataHandler()
	svr, err := server.NewServer(&conf.Config.Server, dataHandler.GetHandlerMap())
	if err != nil {
		os.Exit(1)
	}

	if err != svr.Start() {
		os.Exit(1)
	}
}
