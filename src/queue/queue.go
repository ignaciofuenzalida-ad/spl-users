package queue

import (
	"sync"
)

type Queue[T any] struct {
	lock   *sync.Mutex
	Values []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{&sync.Mutex{}, make([]T, 0)}
}

func (q *Queue[T]) PushMany(values []T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.Values = append(q.Values, values...)
}

func (q *Queue[T]) Push(value T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.Values = append(q.Values, value)
}

func (q *Queue[T]) PopMany(n int) *[]T {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.Values) < n {
		n = len(q.Values)
	}

	firstElements := q.Values[:n]
	q.Values = q.Values[n:]

	return &firstElements
}

func (q *Queue[T]) Pop() *T {
	q.lock.Lock()
	defer q.lock.Unlock()

	firstElements := q.Values[0]
	q.Values = q.Values[1:]

	return &firstElements
}
