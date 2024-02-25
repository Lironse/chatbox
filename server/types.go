package main

import "github.com/gorilla/websocket"

type Packet struct {
	Type    string `json:"type"`
	Action  string `json:"action"`
	Payload string `json:"payload"`
	From    string `json:"from"`
	To      string `json:"to"`
}

type Client struct {
	Conn *websocket.Conn
	Name string
}
