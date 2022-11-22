package database

type Snapshot interface {
	// Has retrieves if a key is present in the snapshot backing by a key-value
	// data store.
	Has(key []byte) (bool, error)

	// Get retrieves the given key if it's present in the snapshot backing by
	// key-value data store.
	Get(key []byte) ([]byte, error)

	// Release releases associated resources. Release should always succeed and can
	// be called multiple times without causing error.
	Release()
}
