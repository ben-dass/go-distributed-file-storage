package p2p

import (
	"encoding/gob"
	"io"
)

func (dec GOBDecoder) Decode(r io.Reader, msg *RPC) error {
	return gob.NewDecoder(r).Decode(msg)
}

func (dec DefaultDecoder) Decode(r io.Reader, msg *RPC) error {
	buff := make([]byte, 1028)

	n, err := r.Read(buff)
	if err != nil {
		return err
	}

	msg.Payload = buff[:n]

	return nil
}
