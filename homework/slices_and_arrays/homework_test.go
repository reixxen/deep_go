package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// Предположим, что эта очередь будет оперировать только положительными
// числами (отрицательные числа ей никогда не поступят на вход)
type CircularQueue struct {
	values []int
	front  int
	rear   int
	isFull bool
}

// создать очередь с определенным размером буффера
func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		front:  0,
		rear:   0,
		isFull: false,
	}
}

// добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	q.values[q.rear] = value
	q.rear = (q.rear + 1) % len(q.values)
	if q.rear == q.front {
		q.isFull = true
	}
	return true
}

// удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.front = (q.front + 1) % len(q.values)
	q.isFull = false
	return true
}

// получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

// получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	lastIdx := (q.rear - 1 + len(q.values)) % len(q.values)
	return q.values[lastIdx]
}

// проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return !q.isFull && (q.front == q.rear)
}

// проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	return q.isFull
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
