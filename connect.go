package main

import (
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func connectToPeer(h host.Host, peerAddress string) {
	peerMultiAddress, err := multiaddr.NewMultiaddr(peerAddress)
	if err != nil {
		log.Fatalf("Invalid peer address: %s", err)
	}
	peerInfo, err := peer.AddrInfoFromP2pAddr(peerMultiAddress)
	if err != nil {
		log.Fatalf("Failed to get peer info: %s", err)
	}
	if err := h.Connect(context.Background(), *peerInfo); err != nil {
		log.Printf("Failed to connect to peer: %s", err)
	} else {
		fmt.Printf("Connected to peer: %s\n", peerInfo.ID.String())
	}

}
