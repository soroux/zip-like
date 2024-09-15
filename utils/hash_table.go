package utils

type HashTable struct {
	table map[string][]byte
}

func NewHashTable() *HashTable {
	return &HashTable{
		table: make(map[string][]byte),
	}
}

func (ht *HashTable) Set(key string, value []byte) {
	ht.table[key] = value
}

func (ht *HashTable) Get(key string) ([]byte, bool) {
	value, exists := ht.table[key]
	return value, exists
}
