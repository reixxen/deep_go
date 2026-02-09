package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type node struct {
	key   int
	value int
	left  *node
	right *node
}

type OrderedMap struct {
	root *node
	size int
}

// создать упорядоченный словарь
func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

// добавить элемент в словарь
func (m *OrderedMap) Insert(key, value int) {
	m.root = m.insert(m.root, key, value)
}

// удалить элемент из словари
func (m *OrderedMap) Erase(key int) {
	m.root = m.erase(m.root, key)
}

// проверить существование элемента в словаре
func (m *OrderedMap) Contains(key int) bool {
	current := m.root
	for current != nil {
		if key == current.key {
			return true
		}
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}

// получить количество элементов в словаре
func (m *OrderedMap) Size() int {
	return m.size
}

// применить функцию к каждому элементу словаря от меньшего к большему
func (m *OrderedMap) ForEach(action func(int, int)) {
	nodes := []*node{}
	current := m.root

	for current != nil || len(nodes) > 0 {
		for current != nil {
			nodes = append(nodes, current)
			current = current.left
		}
		current = nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		action(current.key, current.value)
		current = current.right
	}
}

func (m *OrderedMap) insert(n *node, key, value int) *node {
	if n == nil {
		n = &node{key: key, value: value}
		m.size++
		return n
	}

	if key < n.key {
		n.left = m.insert(n.left, key, value)
	} else if key > n.key {
		n.right = m.insert(n.right, key, value)
	} else {
		n.value = value
	}

	return n
}

func (m *OrderedMap) erase(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = m.erase(n.left, key)
	} else if key > n.key {
		n.right = m.erase(n.right, key)
	} else {
		m.size--
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}

		min := m.findMin(n.right)
		n.key = min.key
		n.value = min.value
		n.right = m.deleteMin(n.right)
	}

	return n
}

func (m *OrderedMap) findMin(n *node) *node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (m *OrderedMap) deleteMin(n *node) *node {
	if n.left == nil {
		return n.right
	}
	n.left = m.deleteMin(n.left)
	return n
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
