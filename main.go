package main

import (
	"flag"

	"github.com/akkyie/connect.sfcclip.net/server"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		production bool
		port       int
	)
	flag.BoolVar(&production, "production", false, "switch development or production database")
	flag.IntVar(&port, "port", 8080, "port for the server to be binded")
	flag.Parse()

	server, err := server.NewServer(production)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err := server.Start(port); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
