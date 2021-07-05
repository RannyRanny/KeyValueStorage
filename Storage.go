package KeyValueStorage

import "sync"

// Storage Thread safe key-value storage
// new instance should be created by NewStorage()
type Storage struct {
	storage map[interface{}]interface{}
	mutex   *sync.Mutex
}

// NewStorage Creates new key-value storage
func NewStorage() *Storage {
	s := new(Storage)
	s.storage = make(map[interface{}]interface{})
	var mutex sync.Mutex
	s.mutex = &mutex
	return s
}

// AddValue Adds value associated with key in storage
func (s Storage) AddValue(key interface{}, value interface{}) interface{} {
	s.mutex.Lock()
	if s.storage[key] != nil {
		s.mutex.Unlock()
		return nil
	}
	s.storage[key] = value
	s.mutex.Unlock()
	return value
}

// ReadValue Reads value by key from storage
func (s Storage) ReadValue(key interface{}) interface{} {
	s.mutex.Lock()
	var value = s.storage[key]
	s.mutex.Unlock()
	return value
}

// UpdateValue Updates value associated with key in storage
func (s Storage) UpdateValue(key interface{}, value interface{}) interface{} {
	s.mutex.Lock()
	if s.storage[key] == nil {
		s.mutex.Unlock()
		return nil
	}
	s.storage[key] = value
	s.mutex.Unlock()
	return value
}

// DeleteValue Deletes key and associated value from storage
func (s Storage) DeleteValue(key interface{}) {
	s.mutex.Lock()
	if s.storage[key] != nil {
		delete(s.storage, key)
	}
	s.mutex.Unlock()
}
