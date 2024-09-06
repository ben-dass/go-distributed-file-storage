package store

type PathKey struct {
	Pathname string
	Filename string
}

type PathTransformFunc func(string) PathKey

type StoreOpts struct {
	// Root is where all data is stored.
	Root              string
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}
