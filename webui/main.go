package main

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var (
	client   *redis.Client
	upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins
		},
	}
	templates = template.Must(template.ParseFiles("index.html"))
)

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)

	go subscribeToRedis()

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.Execute(w, nil)
}

var clients = make(map[*websocket.Conn]bool)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			break
		}
	}
}

func subscribeToRedis() {
	ctx := context.Background()

	pubsub := client.PSubscribe(ctx, "__*:*")
	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println("Message received:", msg.Channel, msg.Payload)
		notifyClients(msg.Channel, msg.Payload)
	}
}

func notifyClients(channel, message string) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(channel+": "+message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			client.Close()
			delete(clients, client)
		}
	}
}
