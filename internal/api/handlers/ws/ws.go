package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

//global websockets handler variables
var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true //not verify origin
		},
	}
	Clients  map[*websocket.Conn]bool
	OutChann chan *Message
	InChann  chan *Message
)

//Message ws message structure
type Message struct {
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

//InitWebsocketsHandling initialize websockets handling
func InitWebsocketsHandling() error {
	OutChann = make(chan *Message)
	InChann = make(chan *Message)
	Clients = make(map[*websocket.Conn]bool)
	//go broadcastMessages()
	return nil
}

//ServeWs set set websocket connection
func ServeWs(w http.ResponseWriter, r *http.Request) {
	//upgrade
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Error().Msgf("websocket handshake error: %s", err.Error())
			//todo:
		}
		return
	}
	defer ws.Close()
	Clients[ws] = true
	ws.WriteMessage(websocket.TextMessage, []byte("websocket connection set"))
	for {
		msg := <-OutChann
		ws.WriteJSON(msg)
		if err != nil {
			log.Error().Msgf("[WS] send msg error: %s", err.Error())
			ws.Close()
			delete(Clients, ws)
		}
	}
}

func broadcastMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-OutChann
		// Send it out to every client that is currently connected
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Error().Msgf("[WS] send msg error: %s", err.Error())
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
