package socket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type M map[string]interface{}

var CurrentConnection WebSocketConnection

var connections = make([]*WebSocketConnection, 0)

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From   string
	Code   int
	Status string
	Domain string
	Prefix string
}

type WebSocketConnection struct {
	*websocket.Conn
	Username string
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Handler WebSocket
func Wshandler(w http.ResponseWriter, r *http.Request) {
	currentGorillaConn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	username := r.URL.Query().Get("username")
	currentConn := WebSocketConnection{Conn: currentGorillaConn, Username: username}
	CurrentConnection = currentConn
	connections = append(connections, &currentConn)
}

//func HandleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
//	defer func() {
//		if r := recover(); r != nil {
//			log.Println("ERROR", fmt.Sprintf("%v", r))
//		}
//	}()
//
//	//BrodacastMessage(currentConn, MESSAGE_NEW_USER, "")
//
//	for {
//		payload := SocketPayload{}
//		err := currentConn.ReadJSON(&payload)
//		if err != nil {
//			if strings.Contains(err.Error(), "websocket: close") {
//				//BrodacastMessage(currentConn, MESSAGE_LEAVE, "")
//				//ejectConnection(currentConn)
//				return
//			}
//
//			log.Println("ERROR", err.Error())
//			continue
//		}
//
//		//BrodacastMessage(currentConn, MESSAGE_CHAT, payload.Message)
//	}
//}

// BrodacastMessage to send data to client
func BrodacastMessage(currentConn *WebSocketConnection, code int, status, domain, prefix string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(SocketResponse{
			From:   currentConn.Username,
			Code:   code,
			Status: status,
			Domain: domain,
			Prefix: prefix,
		})
	}
}

// GetCurrentConnection get current connection detail
func GetCurrentConnection() *WebSocketConnection {
	return &CurrentConnection
}
