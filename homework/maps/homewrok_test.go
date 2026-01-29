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
	return OrderedMap{
		root: nil,
		size: 0,
	}
}

// добавить элемент в словарь
func (m *OrderedMap) Insert(key, value int) {
	if m.root == nil {
		m.root = &node{key: key, value: value}
		m.size++
		return
	}

	current := m.root
	for current != nil {
		if key < current.key {
			if current.left == nil {
				current.left = &node{key: key, value: value}
				m.size++
				return
			}
			current = current.left
		} else if key > current.key {
			if current.right == nil {
				current.right = &node{key: key, value: value}
				m.size++
				return
			}
			current = current.right
		} else {
			current.value = value
			return
		}
	}
}

// удалить элемент из словари
func (m *OrderedMap) Erase(key int) {
	if m.root == nil {
		return
	}

	parent := &node{}
	current := m.root

	// find a node and its parent
	for current != nil && current.key != key {
		parent = current
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	if current == nil {
		return
	}

	// node has 2 childs
	if current.left != nil && current.right != nil {
		minRightParent := current
		minRight := current.right

		for minRight.left != nil {
			minRightParent = minRight
			minRight = minRight.left
		}

		current.key = minRight.key
		current.value = minRight.value

		current = minRight
		parent = minRightParent
	}

	// node has 0 or 1 child
	child := &node{}
	if current.left != nil {
		child = current.left
	} else {
		child = current.right
	}

	// if we delete the root
	if parent == nil {
		m.root = child
	} else if parent.left == current {
		parent.left = child
	} else {
		parent.right = child
	}

	m.size--
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
