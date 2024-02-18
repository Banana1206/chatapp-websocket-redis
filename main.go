package main

import (
	"chatapp-websoket-redis/server/httpserver"
	"chatapp-websoket-redis/server/websocket"
	"flag"
	"log"
)



func main() {
	server := flag.String("server", "", "http,websocket")
	flag.Parse()

	if *server == "http" {
		log.Println("http server is starting on :8082")
		httpserver.StartHTTPServer()
	} else if *server == "websocket" {
		log.Println("websocket server is starting on :8081")
		websocket.StartWebsocketServer()
	} else {
		log.Println("invalid server. Available server: http or websocket")
	}
}
