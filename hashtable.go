package main

import (
	"fmt"
	"math"
	"math/rand"
)

// HashTable implementation with Pearson hashing
type HashTable struct {
	table    [][]keyValuePair
	lkpTable []int
}

type keyValuePair struct {
	key   string
	value interface{}
}

// Init starts hashtable
func (ht *HashTable) Init() {
	ht.table = make([][]keyValuePair, math.MaxUint8)
	ht.lkpTable = rand.Perm(256)
	fmt.Printf("%v", ht.lkpTable)
}

func (ht HashTable) hash(key string) uint8 {
g	var hash uint8 = uint8(len(key) % 256)
	for _, b := range []byte(key) {
		hash = uint8(ht.lkpTable[hash^uint8(b)])
	}
	if hash < 0 {
		panic(ht.lkpTable)
	}
	return hash
}

// Set value on table
func (ht *HashTable) Set(key string, value interface{}) {
	keyHash := ht.hash(key)
	newValue := keyValuePair{key, value}
	slot := ht.table[keyHash]
	index := -1
	for i, kvPair := range slot {
		if key == kvPair.key {
			index = i
			break
		}
	}
	if index < 0 {
		slot = append(slot, newValue)
	} else {
		slot[index] = newValue
	}
	ht.table[keyHash] = slot
}

// Get return value for key
func (ht *HashTable) Get(key string) interface{} {
	slot := ht.table[ht.hash(key)]
	for _, kvPair := range slot {
		if kvPair.key == key {
			return kvPair.value
		}
	}
	return nil
}

// Delete remove key from table
func (ht *HashTable) Delete(key string) {
	ht.Set(key, nil)
}
