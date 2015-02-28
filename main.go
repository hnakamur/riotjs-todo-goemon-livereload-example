package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var (
	assetsDirVar string
	hostVar      string
	portVar      int
)

func init() {
	flag.StringVar(&assetsDirVar, "a", "assets", "assets directory")
	flag.StringVar(&hostVar, "h", "localhost", "listening host")
	flag.IntVar(&portVar, "p", 3000, "listening port")
}

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(assetsDirVar))
	http.Handle("/", fs)

	hostAndPort := hostVar + ":" + strconv.Itoa(portVar)
	log.Printf("Listening at %s...\n", hostAndPort)
	http.ListenAndServe(hostAndPort, nil)
}
