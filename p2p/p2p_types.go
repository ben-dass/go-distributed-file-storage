package p2p

import (
	"io"
	"net"
)

// RPC holds any arbitrary data being sent over each transport
// between two nodes.
type RPC struct {
	From    net.Addr
	Payload []byte
}

// Peer represents the remote node.
type Peer interface {
	Close() error
}

// Transport hanldes communication between nodes in the network.
// (TCP, UDP, websockets, etc.)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}

// TCPPeer The remote node over an establised TCP connection.
type TCPPeer struct {
	conn net.Conn

	// When dialing, the peer is considered Outbound. When connection
	// accepted it becomes an inbound connection (Outbound false).
	outbound bool
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC
}

// HandshakeFunc
type HandshakeFunc func(Peer) error

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{}

type DefaultDecoder struct{}
