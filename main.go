package main

import (
	"distributed-file-storage/p2p"
	"distributed-file-storage/server"
	"distributed-file-storage/store"
	"log"
	"time"
)

func main() {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandskaeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := server.FileServerOpts{

		StorageRoot:       "3000_network",
		PathTransformFunc: store.CASPathTransformFunc,
		Transport:         tcpTransport,
	}
	s := server.NewFileServer(fileServerOpts)

	go func() {
		time.Sleep(time.Second * 3)
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
