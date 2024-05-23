package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
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

	dsn := "host=db user=postgres password=postgres dbname=postgres_db port=5432 sslmode=disable"

	db, err := openDB(dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	query, _ := db.Query(context.Background(), "select * from snippets")
	logger.Info("query %v", query)

	app := &application{
		logger: logger,
		config: cfg,
	}

	logger.Info("Starting server", "addr", &cfg.addr)

	err = http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection
// pool for a given DSN.
func openDB(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		dbPool.Close()
		return nil, err
	}

	return dbPool, nil
}
