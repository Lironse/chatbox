package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func handlePacket(packet Packet, conn *websocket.Conn) error {
	// marshal the json struct into a string
	packetString, err := json.Marshal(packet)
	if err != nil {
		return ErrPacketMarshalFailure
	}

	// Follow packet.Action
	// Accepted actions:
	// register: write the received username and pubkey into the DHT
	// open: do nothing
	// passOffer: send the recieved offer SDP to the intended recipient specified in packet.To
	// passAnswer: send the recieved answer SDP to the intended recipient specified in packet.To
	switch packet.Action {
	case "open", "register":
		addClient(packet.From, conn)

	case "passOffer", "passAnswer":
		passPacket([]byte(packetString), packet.From, packet.To)
	}

	return nil
}

var clients = make(map[string]*websocket.Conn)

// var clientsMutex sync.Mutex

func passPacket(message []byte, sender string, to string) {
	_, ok := clients[to]
	if ok {
		err := clients[to].WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("Error writing message:", err)
			removeClient(sender)
		}
	} else {
		fmt.Println("User", to, "is not connected!")
	}
}

func addClient(name string, conn *websocket.Conn) {
	clients[name] = conn
}

func removeClient(name string) {
	delete(clients, name)
}
