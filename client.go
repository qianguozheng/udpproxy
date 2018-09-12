package udpproxy

import (
	"log"
	"net"
)

func UDPClient(address string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Printf("net.ResolveUDPAddr fail %s.", err.Error())
		return nil, err
	}

	socket, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Printf("net.DialUDP fail %s.", err.Error())
		return nil, err
	}
	return socket, nil
}
