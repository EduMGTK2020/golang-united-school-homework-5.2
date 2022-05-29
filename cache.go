package cache

import "time"

var always time.Time = time.Time{}

type valueStorage struct {
	value    string
	deadline time.Time
}

type Cache struct {
	storage map[string]valueStorage
}

func NewCache() Cache {
	return Cache{storage: make(map[string]valueStorage)}
}

func (c Cache) Get(key string) (string, bool) {
	if v, inMap := c.storage[key]; !inMap {
		return "", false
	} else {
		if v.deadline != always && time.Now().After(v.deadline) {
			delete(c.storage, key)
			return "", false
		} else {
			return v.value, true
		}
	}
}

func (c Cache) Put(key, value string) {
	c.storage[key] = valueStorage{value: value, deadline: always}
}

func (c Cache) Keys() []string {
	var keys []string
	for kv := range c.storage {
		if _, ok := c.Get(kv); ok {
			keys = append(keys, kv)
		}
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.storage[key] = valueStorage{value: value, deadline: deadline}
}
