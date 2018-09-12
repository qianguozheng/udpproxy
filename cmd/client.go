package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", ":8900")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr fail.", err)
		os.Exit(1)
	}

	socket, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("net.DialUDP fail.", err)
		os.Exit(1)
	}
	defer socket.Close()
	r := bufio.NewReader(os.Stdin)
	for {
		switch line, ok := r.ReadString('\n'); true {
		case ok != nil:
			fmt.Printf("bye bye!\n")
			return
		default:
			socket.Write([]byte(line))
			data := make([]byte, 1024)
			_, remoteAddr, err := socket.ReadFromUDP(data)
			if err != nil {
				fmt.Println("error recv data")
				return
			}
			fmt.Printf("from %s:%s\n", remoteAddr.String(), string(data))
		}
	}
}
