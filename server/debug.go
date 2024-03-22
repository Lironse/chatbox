package main

import "fmt"

func input_commands() {
	var command string
	for {
		fmt.Scanln(&command)
		switch command {
		case "list":
			logClients()
		case "rt":
			logRoutingTable()
		}
	}
}
