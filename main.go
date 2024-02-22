package main

import (
	"fmt"
)

type Cache interface {
	Get(k string) ([]byte, bool)
	Set(k string, v []byte)
}

// var _ Cache = (*cacheImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func newCache(lim int) Cache {
	return &cacheImpl{
		limit: lim,
	}
}

type node struct {
	key   string
	value []byte
	prev  *node
	next  *node
}

type cacheImpl struct {
	len   int
	tail  *node
	head  *node
	limit int
}

func (a *cacheImpl) Get(k string) ([]byte, bool) {
	// TODO implement me
	c := a.tail

	for c != nil {
		if c.key == k {
			return c.value, true
		}
		c = c.next
	}
	return nil, false
}

func (c *cacheImpl) Set(k string, v []byte) {
	// TODO implement me
	c.len++
	if c.len == c.limit {
		c.tail = c.tail.next
		c.len--
	}
	if c.head == nil {
		c.head = &node{
			key:   k,
			value: v,
			prev:  c.head,
		}
	} else {
		c.head.next = &node{
			key:   k,
			value: v,
			prev:  c.head,
		}
		c.head = c.head.next
	}
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string, v []byte) {
	v, ok := d.cache.Get(k)
	if ok {
		// return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
		return fmt.Printf("answer from cache: key: %s\n", k), ok
	}

	v, ok = d.dbs[k]
	// return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
	return fmt.Printf("answer from dbs: key: %s\n", k), ok
}

func main() {
	c := newCache()
	db := newDbImpl(c)
	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
}
