package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	logInfo("Connection from: ", r.RemoteAddr)

	// upgrade the http connection to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logError(ErrUpgradeFailure)
		return // close server
	}

	// start reading and writing to clients
	for {
		// read the packet into a Packet struct
		var packet Packet
		err = conn.ReadJSON(&packet)
		if err != nil {
			logError(ErrReadPacketFailure)
			err = removeClient(conn)
			if err != nil {
				logError(err)
			}
			return // dismiss client
		}

		// log the full incoming packet
		logPacket(packet)
		err = handlePacket(packet, conn)
		if err != nil {
			logError(err)
			continue // dismiss packet
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[Client]*websocket.Conn)

func handlePacket(packet Packet, conn *websocket.Conn) error {
	// Follow packet.Action
	// Accepted actions:
	// register: write the received username and pubkey into the DHT
	// open: do nothing
	// passPacket: forward the recieved packet
	switch packet.Action {
	case "register", "connect":
		newClient := Client{
			Name:      packet.From,
			PublicKey: packet.Payload,
		}
		addClient(newClient, conn)

	case "passPacket":
		err := passPacket(packet)
		if err != nil {
			logError(err)
		}
	}

	return nil
}

func passPacket(packet Packet) error {
	sender, found := findClientByName(packet.From)
	if !found {
		// TODO: maybe disconnect the client?
		return ErrClientNotFound
	}

	receiver, found := findClientByName(packet.To)
	if !found {
		// if client is offline
		packet.Action = "alertError"
		packet.Payload = "User offline"
		packet.From = "server"
		packet.To = receiver.Name

		// return the packet to the sender
		receiver = sender
	}

	// marshal the json struct into a byte slice
	packetBytes, err := json.Marshal(packet)
	if err != nil {
		return ErrPacketMarshalFailure
	}

	// send the packet to the appropriate client
	err = clients[receiver].WriteMessage(websocket.TextMessage, []byte(packetBytes))
	if err != nil {
		return ErrSendPacketFailure
	}

	return nil
}

func findClientByName(name string) (Client, bool) {
	for client := range clients {
		if client.Name == name {
			return client, true
		}
	}

	return Client{}, false
}

func addClient(client Client, conn *websocket.Conn) error {
	_, found := findClientByName(client.Name)
	if found {
		return ErrClientAlreadyExists
	}

	clients[client] = conn

	return nil
}

func removeClient(conn *websocket.Conn) error {
	client, found := findClientByConn(conn)
	if !found {
		return ErrClientNotFound
	}

	delete(clients, client)

	return nil
}

func findClientByConn(conn *websocket.Conn) (Client, bool) {
	for client := range clients {
		if clients[client] == conn {
			return client, true
		}
	}

	return Client{}, false
}
