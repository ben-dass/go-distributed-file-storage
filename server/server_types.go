package server

import (
	"distributed-file-storage/p2p"
	"distributed-file-storage/store"
)

type FileServerOpts struct {
	StorageRoot       string
	PathTransformFunc store.PathTransformFunc
	Transport         p2p.Transport
}

type FileServer struct {
	FileServerOpts
	store  *store.Store
	quitCh chan struct{}
}
