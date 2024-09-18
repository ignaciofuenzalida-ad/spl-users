package queue

import (
	"sync"
)

type MapQueue[K ~int | ~string] struct {
	lock   *sync.Mutex
	Values map[K]bool
}

func NewMapQueue[K ~int | ~string]() *MapQueue[K] {
	return &MapQueue[K]{&sync.Mutex{}, make(map[K]bool, 0)}
}

func (q *MapQueue[K]) PushMany(keys []K) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for _, key := range keys {
		_, exist := q.Values[key]

		if exist {
			continue
		}

		q.Values[key] = false
	}
}

func (q *MapQueue[K]) Push(key K) {
	q.lock.Lock()
	defer q.lock.Unlock()

	_, exist := q.Values[key]
	if exist {
		return
	}

	q.Values[key] = false
}

func (q *MapQueue[K]) PopMany(n int) *[]K {
	q.lock.Lock()
	defer q.lock.Unlock()

	keys := make([]K, 0)

	for key, value := range q.Values {
		if !value {
			q.Values[key] = true
			keys = append(keys, key)
		}

		if len(keys) == n {
			break
		}
	}

	return &keys
}

func (q *MapQueue[K]) Pop() *K {
	q.lock.Lock()
	defer q.lock.Unlock()

	for key, value := range q.Values {
		if !value {
			q.Values[key] = true
			return &key
		}
	}

	return nil
}

func (q *MapQueue[K]) Remove(key K) {
	q.lock.Lock()
	defer q.lock.Unlock()

	delete(q.Values, key)
}
