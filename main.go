package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/ce-admin/web_server"
)

func init() {
	log.SetPrefix("CE-Admin")
}

func main() {
	port := flag.Uint("port", 8080, "TCP port number for the admin app")
	flag.Parse()

	log.Println("Starting admin server...Listening on Port ", *port)

	ws := &web_server.WebServer{}
	go ws.Run()

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(*port)), nil))
}
