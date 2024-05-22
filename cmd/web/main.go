package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

// Define an application struct to hold the app-wide dependencies for the web app.
type application struct {
	logger *slog.Logger
	config config
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":8080", "HTTP Network Address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	app := &application{
		logger: logger,
		config: cfg,
	}

	logger.Info("Starting server", "addr", &cfg.addr)

	err := http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
}
