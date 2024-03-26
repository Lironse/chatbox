package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

/*
Console colors
*/
var green = "\033[32m"
var red = "\033[31m"
var magenta = "\033[35m"
var reset = "\033[0m"

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const keyspaceSize int = 256

var localNode = Node{generateID(), 0, getPublicIP()}
var routingTable = makeRoutingTable()
