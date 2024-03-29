package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
)

type RoutingTable struct {
	Buckets []Bucket
	K       int
}

func makeRoutingTable() *RoutingTable {
	bucketAmount := int(math.Ceil(math.Log2(float64(keyspaceSize)))) + 1
	fmt.Println(bucketAmount)
	return &RoutingTable{
		Buckets: make([]Bucket, bucketAmount),
		K:       2,
	}
}

func (r *RoutingTable) addNode(node Node) {
	bucketIndex := r.getBucketIndexForId(node.LocalId)
	fmt.Println("adding node:", node.Id, "local:", node.LocalId, "to bucket", bucketIndex)
	r.addEntryToBucket(bucketIndex, Entry{node.LocalId, node.Ip, "Node"})
}

func (r *RoutingTable) getBucketIndexForId(id int) int {
	if id == 0 {
		return 0 // or handle this case according to your requirements
	}
	// TODO check buckets mathematically
	println("id ", float64(id))
	println("result ", math.Ceil(math.Log2(float64(id))+1))
	return int(math.Floor(math.Log2(float64(id)) + 1))
}

func (r *RoutingTable) addEntryToBucket(index int, e Entry) {
	if len(r.Buckets[index].Entries) == r.K {
		return
	}
	r.Buckets[index].Entries = append(r.Buckets[index].Entries, e)
}

func (r *RoutingTable) getLocalNode() Node {
	return localNode
}

func (r *RoutingTable) RegisterToServerList() {
	url := "http://176.230.36.90:5173/api/servers"

	type ServerEntry struct {
		Ip string `json:"ip"`
		Id int    `json:"id"`
	}

	ent := ServerEntry{
		Ip: getPublicIP(),
		Id: r.getLocalNode().Id,
	}

	jsonData, err := json.Marshal(ent)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	// Create a new POST request with the specified URL and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set Content-Type header to application/json since we are sending JSON data
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response Status:", resp.Status)

	// Read and print the response body
	// You may want to handle the response body according to your application's requirements
	// For example, you can parse it as JSON if the server sends JSON data back
	// Here, we simply print the response body as a string
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
}

func (r *RoutingTable) doesUserExist(username string) (bool, string) {
	// TODO: hash the username and check the node ID on the entire network i guess?
	id := calculateLocalId(hashUsername(username))
	for _, bucket := range r.Buckets {
		for _, entry := range bucket.Entries {
			fmt.Println("checking", entry.Id, "against", id)
			if entry.Id == id {
				return true, entry.Value
			}
		}
	}
	return false, ""
}
