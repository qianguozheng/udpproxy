package udpproxy

import (
	"log"
	"net"
)

func newListener(addr string) (*net.UDPConn, error) {
	udp, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	l, err := net.ListenUDP("udp", udp)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func UDPProxy(addr string) {
	l, err := newListener(addr)
	defer l.Close()
	if err != nil {
		log.Printf("newListener error %s", err.Error())
	}

	buf := make([]byte, 40960)
	var id int = 0

	for {
		// conn, err := l.Accept()
		// if err != nil {
		// 	log.Printf("Accept error %s", err.Error())
		// }
		//
		// go handle(conn)

		rlen, remote, err := l.ReadFromUDP(buf)
		if err == nil {
			id++
			log.Println("connected from " + remote.String())
			go handle(l, remote, id, buf[:rlen])
			//go echoHandler(l, remote, id, buf[:rlen])
		}
	}

}

func handle(conn *net.UDPConn, remote *net.UDPAddr, id int, buf []byte) {
	client, err := UDPClient(":9001")
	defer client.Close()
	if err != nil {
		log.Println("connect to udp server failed", err.Error())
		return
	}
	client.Write(buf)

	data := make([]byte, 1024)
	_, remoteAddr, err := client.ReadFromUDP(data)
	if err != nil {
		log.Println("read from udp server failed", err.Error())
		return
	}
	log.Println("remoteAddr:", remoteAddr.String(), string(data))
	conn.WriteToUDP(data, remote)
}

// func read2(conn *net.Conn) ([]byte, error) {
//   defer conn.Close()
//   var buf bytes.Buffer
//   _, err := io.Copy(&buf, conn)
//   if err != nil {
//     // Error handler
//     return nil, err
//   }
//   return buf, nil
// }

// func read3(conn net.Conn) ([]byte, error) {
// 	defer conn.Close()
// 	buf, err := ioutil.ReadAll(conn)
// 	if err != nil {
// 		// Error Handler
// 		return nil, err
// 	}
// 	// use buf...
// 	return buf, nil
// }
