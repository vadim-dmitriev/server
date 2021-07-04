package pkg

import (
	"log"
	"net"

	"google.golang.org/protobuf/proto"
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
	log.Print(request)

	bytes, err := proto.Marshal(request)
	if err != nil {
		panic(err)
	}
	tcps.conn.Write(bytes)

	return &Empty{}
}
