package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logError(ErrUpgradeFailure)
		return // dismiss client
	}

	logInfo("Connection from: ", conn.RemoteAddr())

	for {
		var packet Packet
		err = conn.ReadJSON(&packet)
		if err != nil {
			logError(ErrReadPacketFailure)
			err = removeFromConnectedClients(conn)
			if err != nil {
				logError(err)
			}
			return // dismiss client
		}

		logPacket(packet)
		err = handlePacket(packet, conn)
		if err != nil {
			logError(err)
			continue // dismiss packet
		}
	}
}

var clients = make(map[Client]*websocket.Conn)

func handlePacket(packet Packet, conn *websocket.Conn) error {
	switch packet.Action {
	case "register":
		client := Client{
			Name:      packet.From,
			PublicKey: packet.Payload,
			Conn:      conn,
		}

		// this should be replaced with a db
		// node.Store(client.Name, client.PublicKey)

		addToConnectedClients(client)

	case "connect":
		client := Client{
			Name:      packet.From,
			PublicKey: packet.Payload,
			Conn:      conn,
		}
		addToConnectedClients(client)

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

func addToConnectedClients(client Client) error {
	_, found := findClientByName(client.Name)
	if found {
		return ErrClientAlreadyExists
	}

	clients[client] = client.Conn

	return nil
}

func removeFromConnectedClients(conn *websocket.Conn) error {
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
