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
	flag.BoolVar(&production, "production", false, "")
	flag.IntVar(&port, "port", 8080, "")
	flag.Parse()

	server, err := server.NewServer()
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	if err := server.Start(port); err != nil {
		log.Panic(err)
		panic(err)
	}
}
