package udpproxy

import (
	"log"
	"net"
)

func UDPServer(addr string) {
	l, err := newListener(addr)
	defer l.Close()
	if err != nil {
		log.Printf("newListener error %s", err.Error())
	}

	buf := make([]byte, 40960)
	var id int = 0

	for {
		rlen, remote, err := l.ReadFromUDP(buf)
		if err == nil {
			id++
			log.Println("connected from " + remote.String())
			//go handle(l, remote, id, buf[:rlen])
			go echoHandler(l, remote, id, buf[:rlen])
		}
	}
}

func echoHandler(conn *net.UDPConn, remote *net.UDPAddr, id int, buf []byte) {
	conn.WriteToUDP(buf, remote)
	log.Printf("id: %d, data:%s\n", id, string(buf))
}
