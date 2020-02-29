package main

import (
	"testing"
)

type TValue struct {
	Value int
}

func fillHashTable(hashTable *HashTable, testValues map[string]interface{}) {
	for key, value := range testValues {
		hashTable.Set(key, value)
	}
}
func TestHash(t *testing.T) {
	testValues := map[string]interface{}{
		"hello": "hello",
		"one":   1,
		"bool":  true,
	}
	hashTable := &HashTable{}
	hashTable.Init()
	fillHashTable(hashTable, testValues)
	for key, value := range testValues {
		if tableValue := hashTable.Get(key); tableValue != value {
			t.Errorf("Table value (%v) differente from test value (%v).", tableValue, value)
		}
	}
}
