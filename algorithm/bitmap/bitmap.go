package main

import (
	"errors"
)

var ErrOutOfRange = errors.New("out of range")

type Bitmap struct {
	store []uint8
}

func NewBitmap(n int) *Bitmap {
	return &Bitmap{
		store: make([]uint8, n),
	}
}

func (b *Bitmap) coordinate(n uint8) (int, uint8) {
	//two ways of writing
	// first writing:
	//return int(n / 8), 1 << (n % 8)
	// second writing:
	return int(n / 8), n & (8 - 1)
}

func (b *Bitmap) Add(n uint8) error {
	index, position := b.coordinate(n)
	if index >= len(b.store) {
		return ErrOutOfRange
	}

	b.store[index] |= position
	return nil
}

func (b *Bitmap) Contain(n uint8) bool {
	index, position := b.coordinate(n)
	if index > len(b.store) {
		return false
	}
	return b.store[index]&position == 1
}

func (b *Bitmap) Delete(n uint8) error {
	index, position := b.coordinate(n)
	if index > len(b.store) {
		return ErrOutOfRange
	}
	b.store[index] &= ^position
	return nil
}
