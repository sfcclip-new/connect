package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		production bool
		port       int
	)
	flag.BoolVar(&production, "production", false, "")
	flag.IntVar(&port, "port", 8080, "")
	flag.Parse()

	log.Info("Connecting to the database...")

	db, err := NewDatabase(production, "connect.sfcclip.net")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Info("Starting the server...")

	portStr := fmt.Sprintf(":%d", port)
	server := NewServer(portStr, db, production)

	log.Info("The server is started")

	go func() {
		if err := server.Start(); err != nil {
			log.Error(err)
			log.Info("Shutting down...")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Info("The server is shut down")

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

	log.Info("The database is disconnected")
}
