package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand/v2"
)

type KeySpace [256]string

// Node represents a node in the Kademlia DHT.
type Node struct {
	ID   int
	Data KeySpace // Stores usernames and public keys
}

// NewNode creates a new node with a random ID.
func NewNode() *Node {
	id := generateID()
	var ks KeySpace
	return &Node{
		ID:   id,
		Data: ks,
	}
}

// Store stores the username and public key in the DHT.
func (n *Node) Store(username, publicKey string) {
	hash := sha256.Sum256([]byte(username))

	// Convert the hash to uint8
	var uint8Hash uint8
	for _, b := range hash {
		uint8Hash += uint8(b)
	}

	n.Data[uint8Hash] = publicKey
	for key, value := range n.Data {
		fmt.Println("Key:", key, "Value:", value)
	}
}

// Lookup searches for a username in the DHT.
// func (n *Node) Lookup(username string) (string, error) {
// 	publicKey, ok := n.Data[username]
// 	if !ok {
// 		return "", fmt.Errorf("username not found")
// 	}
// 	return publicKey, nil
// }

// Ping pings another node to update contact information.
// func (n *Node) Ping(other *Node) {
// 	// This is where you would update contact information.
// 	fmt.Printf("Pinging node %s from node %s\n", other.ID, n.ID)
// }

// generateID generates a random ID for the node.
func generateID() int {
	return rand.IntN(256)
}

func (n *Node) init() {
	n.Data[n.ID] = "node-" + fmt.Sprint(n.ID)
	for key, value := range n.Data {
		fmt.Println("Key:", key, "Value:", value)
	}
}

// func main() {
// 	node1 := NewNode()
// 	node2 := NewNode()

// 	node1.Store("user1", "publicKey1")
// 	node1.Store("user2", "publicKey2")

// 	node2.Ping(node1)

// 	// Perform a lookup
// 	publicKey, err := node2.Lookup("user1")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Public key for user1:", publicKey)
// 	}
// }
