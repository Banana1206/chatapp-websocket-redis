package websocket

import (
	"chatapp-websoket-redis/model"
	"chatapp-websoket-redis/server/redisrepo"
	"encoding/json"
	// "log"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Username string
}

type Message struct {
	Type string     `json:"type"`
	User string     `json:"user,omitempty"`
	Chat model.Chat `json:"chat,omitempty"`
}

var clients = make(map[*Client]bool)
var broadcast = make(chan *model.Chat)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}


// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host, r.URL.Query())

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	client := &Client{Conn: ws}
	// register client
	clients[client] = true
	log.Println("clients", len(clients), clients, ws.RemoteAddr())

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	receiver(client)

	log.Println("exiting", ws.RemoteAddr().String())
	delete(clients, client)
}

// define a receiver which will listen for
// new messages being sent to our WebSocket
// endpoint
func receiver(client *Client) {
	for {
		// read in a message
		// readMessage returns messageType, message, err
		// messageType: 1-> Text Message, 2 -> Binary Message
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		msgFrame := &Message{}
		// log.Println("====================")
		// log.Println(message)
		// log.Println(msgFrame)

		err = json.Unmarshal(message, msgFrame)
		if err != nil {
			log.Println("error while unmarshaling chat", err)
			continue
		}

		log.Println("host", client.Conn.RemoteAddr())
		if msgFrame.Type == "bootup" {
			// do mapping on bootup
			client.Username = msgFrame.User
			log.Println("client successfully mapped", &client, client, client.Username)
		} else {
			log.Println("received message", msgFrame.Type, msgFrame.Chat)
			c := msgFrame.Chat
			c.Timestamp = time.Now().Unix()

			// save in redis
			id, err := redisrepo.CreateChat(&c)
			if err != nil {
				log.Println("error while saving chat in redis", err)
				return
			}
			log.Printf("Create chat#%d on redis success", c.Timestamp)

			c.ID = id
			broadcast <- &c
		}
	}
}

func broadcaster() {
	for {
		message := <-broadcast
		// send to every client that is currently connected
		log.Println("new message", message)

		for client := range clients {
			// send message only to involved users
			log.Println("username:", client.Username,
				"from:", message.From,
				"to:", message.To)

			if client.Username == message.From || client.Username == message.To {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					log.Printf("Websocket error: %s", err)
					client.Conn.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Panicln(w, "Simple Server")
	})
	// map our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func StartWebsocketServer() {
	redisClient := redisrepo.InitialiseRedis()
	defer redisClient.Close()

	go broadcaster()
	setupRoutes()
	http.ListenAndServe(":8081", nil)
}

