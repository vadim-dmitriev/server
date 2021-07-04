package tcp

import (
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	log.Printf("new conn '%s'", conn.RemoteAddr().String())

	for {

	}
}
