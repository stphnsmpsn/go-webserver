package main

import (
	"./api"
	"./file"
	"./websocket"
	"net/http"
	"os"
)


const IpIdx int = 1
const PortIdx int = 2
const DocRootIdx int = 3

func main() {
	// todo: add command line argument validation
	file.RegisterFileHandlers(os.Args[DocRootIdx])
	websocket.RegisterWebsocketHandlers()
	api.RegisterVehicleHandlers()
	http.ListenAndServe(os.Args[IpIdx]+":"+os.Args[PortIdx], nil)
}
