package main

import (
	"io"
	"math/rand/v2"
	"net/http"
)

type Node struct {
	Id      int
	LocalId int
	Ip      string
}

type Bucket struct {
	Range   int
	Entries []Entry
}

type Entry struct {
	Id    int
	Value string
	Type  string
}

func generateID() int {
	return rand.IntN(32)
}

func getPublicIP() string {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(ip)
}

func calculateLocalId(id int) int {
	return id ^ routingTable.getLocalNode().Id
}
