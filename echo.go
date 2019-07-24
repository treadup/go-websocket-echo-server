package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error trying to upgrade to a websocket connection: ", err)
        return
    }

    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            fmt.Println(err)
            return
        }

        if err := conn.WriteMessage(messageType, message); err != nil {
            fmt.Println(err)
            return
        }
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/echo", websocketHandler)
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
