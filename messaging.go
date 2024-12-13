package main

import (
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

// handleStream processes incoming messages.
func handleStream(stream network.Stream) {
	fmt.Println("New stream received!")
	defer stream.Close()

	buf := make([]byte, 256)
	n, err := stream.Read(buf)
	if err != nil {
		log.Printf("Error reading from stream: %s", err)
		return
	}

	message := string(buf[:n])
	fmt.Printf("Received message: %s\n", message)
}

// sendMessage sends a message to a peer.
func sendMessage(h host.Host, peerAddr string, message string) {
	peerMultiaddr, err := multiaddr.NewMultiaddr(peerAddr)
	if err != nil {
		log.Fatalf("Invalid peer address: %s", err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(peerMultiaddr)
	if err != nil {
		log.Fatalf("Failed to get peer info: %s", err)
	}

	stream, err := h.NewStream(context.Background(), peerInfo.ID, "/p2p/1.0.0")
	if err != nil {
		log.Fatalf("Failed to create stream: %s", err)
	}
	defer stream.Close()

	_, err = stream.Write([]byte(message))
	if err != nil {
		log.Printf("Failed to send message: %s", err)
	}
}
