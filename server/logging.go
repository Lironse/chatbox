package main

import (
	"fmt"
)

func logPacket(packet Packet) {
	fmt.Print(magenta + "[PKT] ")
	fmt.Println("New packet recieved:")
	fmt.Println("  ➜  Action:", packet.Action)
	fmt.Println("  ➜  Payload:", packet.Payload)
	fmt.Println("  ➜  From:", packet.From)
	fmt.Println("  ➜  To:", packet.To)
	fmt.Println(reset)
}

func logError(err error) {
	fmt.Print(red + "[ERR] ")
	fmt.Println(err.Error())
	fmt.Println(reset)
}

func logInfo(args ...any) {
	fmt.Print(green + "[LOG] ")
	fmt.Println(args...)
	fmt.Println(reset)
}

func logClients() {
	fmt.Print(green + "[LOG] ")
	fmt.Println("Online clients:", len(clients))
	for client := range clients {
		fmt.Println("  ➜  Name:", client.Name)
	}
	fmt.Print(reset)
}
