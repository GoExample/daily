package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		a uint
		b uint
	)

	a = 2
	b = 10

	and := a & b
	or := a | b
	xor := a ^ b
	fmt.Printf("a = %b, b = %b, and = %b (%d)\n", a, b, and, and)
	fmt.Printf("a = %b, b = %b, or = %b (%d)\n", a, b, or, or)
	fmt.Printf("a = %b, b = %b, xor = %b (%d)\n", a, b, xor, xor)

	aLeftMove := a << 1
	fmt.Printf("a = %b, aLeftMove = %b (%d)\n", a, aLeftMove, aLeftMove)

	bLeftMove := b << 5
	fmt.Printf("b = %b, bLeftMove = %b (%d)\n", b, bLeftMove, bLeftMove)

	fmt.Println("=========================================================")
	m := NewBitmap(3)
	// 越界
	if err := m.Add(16); err != nil {
		log.Println(err)
	}
	// 添加成功
	if err := m.Add(15); err != nil {
		log.Println(err)
	}
	fmt.Println(m.Contain(15))
	// 删除元素
	if err := m.Delete(15); err != nil {
		log.Println(err)
	}
	fmt.Println(m.Contain(15))

	m.Add(12)
	m.Add(15)
	m.Add(40)
	fmt.Println(len(m.store))
	fmt.Printf("%08b\n", m.store[0])
	fmt.Printf("%08b\n", m.store[1])
	fmt.Printf("%08b\n", m.store[2])
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", m.store)

	for i := range m.store {
		fmt.Println(i)
	}
}
