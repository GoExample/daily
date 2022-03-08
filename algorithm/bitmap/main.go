package main

import (
	"errors"
	"fmt"
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

func (b *Bitmap) coordinate(n uint64) (int, uint8) {
	//two ways of writing
	// first writing:
	//return int(n / 8), 1 << (n % 8)
	// second writing:
	return int(n / 8), uint8(n & (8 - 1))
}

func (b *Bitmap) Set(n uint64) error {
	index, position := b.coordinate(n)
	if index >= len(b.store) {
		return ErrOutOfRange
	}

	b.store[index] |= 1 << position
	return nil
}

func (b *Bitmap) Unset(n uint64) error {
	index, position := b.coordinate(n)
	if index >= len(b.store) {
		return ErrOutOfRange
	}
	b.store[index] &^= 1 << position
	return nil
}

func (b *Bitmap) Contain(n uint64) bool {
	index, position := b.coordinate(n)
	if index >= len(b.store) {
		return false
	}
	return b.store[index]&(1<<position) != 0
}

func (b *Bitmap) Print() {
	for i := 0; i < len(b.store); i++ {
		binStr := []rune(fmt.Sprintf("%b", b.store[i]))
		for index, value := range revertString(binStr) {
			if string(value) == "1" {
				fmt.Println(i*8 + index)
			}
		}
	}
}

func (b *Bitmap) Log() {
	for i := 0; i < len(b.store); i++ {
		for j := 0; j < 8; j++ {
			if b.store[i]>>j&1 == 1 {
				fmt.Println(i*8 + j)
			}
		}
	}
}

func revertString(from []rune) string {
	var result []rune
	for i := len(from) - 1; i >= 0; i-- {
		result = append(result, from[i])
	}
	return string(result)
}

func main() {
	bitOperation()
	fmt.Println("=========================================================")

	m := NewBitmap(3)
	var testData = [...]uint64{0, 5, 15, 23, 24}

	for _, data := range testData {
		fmt.Printf("插入前查询 %t\n", m.Contain(data))
		if err := m.Set(data); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("插入后查询 %t\n", m.Contain(data))

	}
	m.Print()
	for _, data := range testData {
		if err := m.Unset(data); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("删除后查询 %t\n", m.Contain(data))
		fmt.Println("=========================================================")
	}
	m.Log()
}

func bitOperation() {
	// https://go.dev/ref/spec#Arithmetic_operators
	var a uint8 = 2
	var b uint8 = 10
	var c uint8 = 8
	d := c

	and := a & b
	or := a | b
	xor := a ^ b
	d &^= 1 << 3

	fmt.Printf("a = %b, b = %b, and = %b (%d)\n", a, b, and, and)
	fmt.Printf("a = %b, b = %b, or = %b (%d)\n", a, b, or, or)
	fmt.Printf("a = %b, b = %b, xor = %b (%d)\n", a, b, xor, xor)
	fmt.Printf("c = %b, and not = %b (%d)\n", c, d, d)

	aLeftMove := a << 1
	fmt.Printf("a = %b, aLeftMove = %b (%d)\n", a, aLeftMove, aLeftMove)

	bLeftMove := b << 5
	fmt.Printf("b = %b, bLeftMove = %b (%d)\n", b, bLeftMove, bLeftMove)
}
