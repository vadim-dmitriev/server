package pkg

import (
	"fmt"
	"net"
)

type Server interface {
	SendButtonPressed(request *KeyPressedRequest) *Empty
}

func NewTCPServer(conn net.Conn) Server {
	return &TCPServer{
		conn: conn,
	}
}

type TCPServer struct {
	conn net.Conn
}

func (tcps *TCPServer) SendButtonPressed(request *KeyPressedRequest) *Empty {
	fmt.Println(request.Key)

	return &Empty{}
}
