package main

import (
	"math/bits"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
// go test -bench=. homework_test.go

func ToLittleEndian[T uint16 | uint32 | uint64](number T) T {
	var res T
	switch unsafe.Sizeof(number) {
	case 2:
		res = (number >> 8) | (number << 8)
	case 4:
		n := uint32(number)
		res = T(((n >> 24) & 0xFF) |
			((n >> 8) & 0xFF00) |
			((n << 8) & 0xFF0000) |
			((n << 24) & 0xFF000000))
	case 8:
		n := uint64(number)
		res = T(((n >> 56) & 0xFF) |
			((n >> 40) & 0xFF00) |
			((n >> 24) & 0xFF0000) |
			((n >> 8) & 0xFF000000) |
			((n << 8) & 0xFF00000000) |
			((n << 24) & 0xFF0000000000) |
			((n << 40) & 0xFF000000000000) |
			((n << 56) & 0xFF00000000000000))
	}

	return res
}

func ReverseBytes[T uint16 | uint32 | uint64](number T) T {
	var res T
	switch unsafe.Sizeof(number) {
	case 2:
		res = T(bits.ReverseBytes16(uint16(number)))
	case 4:
		res = T(bits.ReverseBytes32(uint32(number)))
	case 8:
		res = T(bits.ReverseBytes64(uint64(number)))
	}

	return res
}

func BenchmarkToLittleEndian(b *testing.B) {
	var v = []uint32{0x00000000, 0xFFFFFFFF, 0x00FF00FF, 0x0000FFFF, 0x01020304}

	for i := 0; i < b.N; i++ {
		_ = ToLittleEndian(v[i%len(v)])
	}
}

func BenchmarkReverseBytes(b *testing.B) {
	var v = []uint32{0x00000000, 0xFFFFFFFF, 0x00FF00FF, 0x0000FFFF, 0x01020304}

	for i := 0; i < b.N; i++ {
		_ = ReverseBytes(v[i%len(v)])
	}
}

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
