package main

import (
	"fmt"
)

func logPacket(packet Packet) {
	fmt.Print(magenta + "[PKT] ")
	fmt.Println("New packet recieved:")
	fmt.Println("> type:", packet.Type)
	fmt.Println("> action:", packet.Action)
	fmt.Println("> payload:", packet.Payload)
	fmt.Println("> from:", packet.From)
	fmt.Println("> to:", packet.To)
	fmt.Println(reset)
}

func logError(err error) {
	fmt.Print(red + "[ERR] ")
	fmt.Print(err.Error())
	fmt.Println(reset)
}

func logInfo(args ...any) {
	fmt.Print(green + "[LOG] ")
	fmt.Print(args...)
	fmt.Println(reset)
}
