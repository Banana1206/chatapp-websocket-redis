package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	clients          = make(map[*websocket.Conn]bool)
	broadcastMsg     = make(chan []byte)
	broadcastMsgType = make(chan int)
)

func ReadCLientsAndMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, r.URL.Query())
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Err Read client: ", err)
		return
	}
	defer ws.Close()

	// Đăng ký client mới
	clients[ws] = true

	// In ra thông báo khi có client kết nối thành công
	log.Println("Client connected:", ws.RemoteAddr())

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, ws)
			break
		}
		fmt.Println("Received message:", string(message))
		// fmt.Println("Received message type:", messageType)
		broadcastMsg <- message
		broadcastMsgType <- messageType
	}

}

func WriteMessageThrowClients() {
	for {
		msg := <-broadcastMsg
		msgType := <-broadcastMsgType
		fmt.Println("msg: ", string(msg))
		for client := range clients {
			err := client.WriteMessage(msgType, msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", ReadCLientsAndMessage)
	go WriteMessageThrowClients()

	log.Println("http server started on :8085")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
