package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var green = "\033[32m"
var red = "\033[31m"
var magenta = "\033[35m"
var reset = "\033[0m"

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var localNode = Node{generateID(), 0, getPublicIP()}
var routingTable = makeRoutingTable()
