package main

import "fmt"

func logPacket(packet Packet) {
	fmt.Println("New packet received:")
	fmt.Println(" - type:", packet.Type)
	fmt.Println(" - action:", packet.Action)
	// fmt.Println(" - payload:", packet.Payload)
	fmt.Println(" - from:", packet.From)
	fmt.Println(" - to:", packet.To)
	fmt.Println()
}
