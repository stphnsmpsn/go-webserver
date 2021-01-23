package websocket

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net/http"
)

func RegisterWebsocketHandlers() {
	http.HandleFunc("/websocket/echo", WsEcho)
}

func WsEcho(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		// handle error
	}
	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				// handle error
			}
			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				// handle error
			}
		}
	}()
}
