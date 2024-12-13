package main

import (
	"fmt"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
)

func createNode(port int) host.Host {
	h, err := libp2p.New(libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)))
	if err != nil {

		log.Fatalf("Failed to create host :%s", err)
	}
	fmt.Printf("Node ID: %s\n", h.ID().String())
	fmt.Printf("Node Address : /ip4/127.0.0.1/tcp/%d/p2p/%s\n", port, h.ID())
	return h

}
