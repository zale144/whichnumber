package types

const (
	StoredGameKeyPrefix = "StoredGame/value/"
)

// StoredGameKey returns the store key to retrieve a StoredGame from the index fields
func StoredGameKey(index string) []byte {
	var key []byte
	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)
	return key
}
