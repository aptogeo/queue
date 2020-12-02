package queue

import (
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Method method type
type Method int

const (
	// Fifo method
	Fifo Method = iota
	// Random method
	Random
	// Sort method
	Sort
)

// SortFn type
type SortFn func(left, right interface{}) bool

// Queue strut
type Queue struct {
	slice  []interface{}
	lock   sync.Mutex
	rand   *rand.Rand
	sortFn SortFn
	method Method
}

// NewQueue creates new queue
func NewQueue() *Queue {
	return &Queue{
		slice:  make([]interface{}, 0),
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		method: Fifo,
	}
}

// Enqueue enqueues
func (q *Queue) Enqueue(values []interface{}) {
	defer q.lock.Unlock()
	q.lock.Lock()
	if q.method == Sort && q.sortFn != nil {
		sort.Slice(values, func(i, j int) bool {
			return q.sortFn(values[i], values[j])
		})
		len1 := len(q.slice)
		q.slice = append(q.slice, values...)
		a := 0
		b := len(q.slice)
		for i := len1; i < b; i++ {
			for j := i; j > a && q.sortFn(q.slice[j], q.slice[j-1]); j-- {
				q.slice[j], q.slice[j-1] = q.slice[j-1], q.slice[j]
			}
		}
	} else {
		q.slice = append(q.slice, values...)
	}
}

// Dequeue dequeues
func (q *Queue) Dequeue() interface{} {
	defer q.lock.Unlock()
	q.lock.Lock()
	l := len(q.slice)
	if l > 0 {
		if q.method == Random {
			i := rand.Intn(l)
			v := q.slice[i]
			l--
			q.slice[i] = q.slice[l]
			q.slice = q.slice[:l]
			return v
		}
		v := q.slice[0]
		q.slice = q.slice[1:]
		return v
	}
	return nil
}

// SetMethod sets method
func (q *Queue) SetMethod(method Method) {
	q.method = method
	if q.sortFn != nil {
		sort.Slice(q.slice, func(i, j int) bool {
			data1 := q.slice[i]
			data2 := q.slice[j]
			return q.sortFn(data1, data2)
		})
	}
}

// SetSortFn sets sort less function
func (q *Queue) SetSortFn(sortFn SortFn) {
	q.sortFn = sortFn
	if q.method == Sort {
		sort.Slice(q.slice, func(i, j int) bool {
			data1 := q.slice[i]
			data2 := q.slice[j]
			return q.sortFn(data1, data2)
		})
	}
}

// Len gets len
func (q *Queue) Len() int {
	defer q.lock.Unlock()
	q.lock.Lock()
	return len(q.slice)
}

// Reset resets
func (q *Queue) Reset() {
	defer q.lock.Unlock()
	q.lock.Lock()
	q.slice = q.slice[:0]
}
