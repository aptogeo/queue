package queue_test

import (
	"testing"

	"github.com/aptogeo/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := queue.NewQueue()
	q.Enqueue([]interface{}{12, 3, 5, 8, 16, 7})
	q.Enqueue([]interface{}{6, 2, 17, 4})
	assert.Equal(t, 12, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())
	assert.Equal(t, 5, q.Dequeue())
	assert.Equal(t, 8, q.Dequeue())
	assert.Equal(t, 16, q.Dequeue())
	assert.Equal(t, 7, q.Dequeue())
	assert.Equal(t, 6, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 17, q.Dequeue())
	assert.Equal(t, 4, q.Dequeue())
}

func TestSortedQueue(t *testing.T) {
	q := queue.NewQueue()
	q.SetMethod(queue.Sort)
	q.SetSortFn(func(left, right interface{}) bool {
		return left.(int) < right.(int)
	})
	q.Enqueue([]interface{}{12, 3, 5, 8, 16, 7})
	q.Enqueue([]interface{}{6, 2, 17, 4})
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())
	assert.Equal(t, 4, q.Dequeue())
	assert.Equal(t, 5, q.Dequeue())
	assert.Equal(t, 6, q.Dequeue())
	assert.Equal(t, 7, q.Dequeue())
	assert.Equal(t, 8, q.Dequeue())
	assert.Equal(t, 12, q.Dequeue())
	assert.Equal(t, 16, q.Dequeue())
	assert.Equal(t, 17, q.Dequeue())
}
