package main

import (
	"ChatService/internal/service"
	"ChatService/internal/wsClient"
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	hub := wsClient.NewHub()
	go hub.Run()
	//http.HandleFunc("/", service.ServeHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		service.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}