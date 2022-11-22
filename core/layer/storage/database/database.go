package database

import "io"

type Database interface {
	Reader
	Writer
	Stater
	io.Closer
}

type Reader interface {
	KeyValueReader
}

type Writer interface {
	KeyValueWriter
}

type Stater interface {
	KeyValueStater
}

type KeyValueReader interface {
	Has(key []byte) (bool, error)

	Get(key []byte) ([]byte, error)
}

type KeyValueWriter interface {
	Put(key []byte, value []byte) error

	Delete(key []byte) error
}

type KeyValueStater interface {
	Stat(property string) (string, error)
}
