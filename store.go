package main

import (
	"sync"
)

// struct for the key-value pair data
type keyValueStore struct {
	data sync.Map
}

// a constructor function to return a ptr to the struct
func newKeyValueStore() *keyValueStore {
	return &keyValueStore{
		data: sync.Map{}, // data is an empty map
	}
}

//CRUD functions for the sync map

func (s* keyValueStore) Create (key string, value string) {
	s.data.Store(key, value);
}

func (c* keyValueStore) Fetch (key string) (string, bool) {
	value, ok := s.data.Load(key)
	if !ok {
		return "", false
	} else {
		return value.(string), true
	}
}

func (s* keyValueStore) Delete (key string) {
	s.data.Delete(key)
}




