package tcp

import (
	"log"
	"net"

	"github.com/pkg/errors"

	"github.com/vadim-dmitriev/server/internal/config"
	"github.com/vadim-dmitriev/server/internal/server"
)

const network = "tcp"

type TCPServer struct {
	listener net.Listener
}

func NewServer(config config.Configer) (server.IServer, error) {
	addr := "0.0.0.0:" + config.GetPort()

	listener, err := net.Listen(network, addr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create listener")
	}

	log.Printf("TCP server started on %s", addr)

	return &TCPServer{
		listener: listener,
	}, nil
}

func (tcps *TCPServer) Run() {
	for {
		newConn, err := tcps.listener.Accept()
		if err != nil {
			// TODO
		}

		go handleConn(newConn)
	}
}
