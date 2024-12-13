package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// fmt.Print("hello") test
	port := 1000 + rand.Intn(1000)
	fmt.Print(port, "\n")
	// Step 1: Create a new node (port is randomized)
	node := createNode(port)

	// Step 2: Set a stream handler for incoming messages
	node.SetStreamHandler("/p2p/1.0.0", handleStream)

	// Step 3: Example to connect and send a message (optional)
	peerAddr := "/ip4/127.0.0.1/tcp/1947/p2p/12D3KooWCMZgfrJ1Ze1QwjBcnESUdyfD45Rii4dAG6MrhdPsTKEb"
	connectToPeer(node, peerAddr)
	sendMessage(node, peerAddr, "Hello, P2P World!")

	// Keep the node running
	select {}

}
