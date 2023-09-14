package main

import (
	"fmt"
	"net/http"
	"strings"

	"log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type wsConn struct {
	c *websocket.Conn
}

var connPool = make(map[string]wsConn)

func main() {
	http.HandleFunc("/eventSub", func(w http.ResponseWriter, r *http.Request) {
		log.Println("")
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		for {
			// Read message from browser
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if strings.HasPrefix(string(msg), "subID") {
				connPool[string(msg)] = wsConn{
					c: conn,
				}
				break
			}
		}
	})
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn := connPool[r.FormValue("id")]
		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.c.RemoteAddr(), r.FormValue("id"))

		// Write message back to browser
		if err := conn.c.WriteMessage(websocket.TextMessage, []byte(r.FormValue("id"))); err != nil {
			return
		}
	})

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "websockets.html")
	// })
	log.Println("start listen ...")
	http.ListenAndServe(":8080", nil)
}
