package main

import (
	"errors"
	"fmt"
)

var bitsmap = make([]byte, 2, 2)

func getbit(bm []byte, i int) bool {
	i -= 1
	p := i / 8
	l := i % 8
	bytem := bm[p]            //这就取出8位了
	return bytem > bytem^1<<l //如果原来那一位为1 再一次异或 一定为 0 那么相比原来就变小了
}

func setbit(bm []byte, i int) {
	i -= 1     //数组索引是从0 开始的 传入是从1开始的
	p := i / 8 //计算商表示在第几个字节 等价也可以成 p := i >> 3
	l := i % 8 //计算余数 表示位
	fmt.Printf("i = %d, l = %d\n", i, l)
	bytem := bitsmap[p] //这就取出8位了
	fmt.Printf("改变前状态:%08b 第 %d 位要置为1 \n", bytem, l+1)
	bytem = bytem ^ 1<<l //取异或操作
	fmt.Printf("改变后    :%08b\n", bytem)
	bm[p] = bytem //保存更改
}

func iter(bm []byte) []int {
	list := make([]int, 0, len(bm))
	//2个字节 总共可以表达 16的大小 遍历16位
	for i := 1; i <= len(bm)*8; i++ { //这里字节 要乘以 8=16bit
		val := getbit(bm, i)
		if val {
			list = append(list, i)
		}
	}
	return list
}

func UnSet(bm []byte, i int) {
	fmt.Println("删除数字:", i)
	val := getbit(bm, i) //判断存在吗
	if val {
		setbit(bm, i) //再一次插入就是删除了
	}
}

func getPosition(n int) (index, position uint8) {
	return uint8(n / 8), uint8(n & (8 - 1))
}

func setValue(n int) error {
	index, position := getPosition(n)
	a[index] |= 1 << (position - 1)
	return nil
}

var a = make([]uint8, 5)

func main() {
	fmt.Println(iter(bitsmap))
	setbit(bitsmap, 8)
	setbit(bitsmap, 5)
	setbit(bitsmap, 16)
	fmt.Println(iter(bitsmap))
	fmt.Println("=============")
	fmt.Println(getbit(bitsmap, 8))
	fmt.Println(getbit(bitsmap, 2))
	fmt.Println(iter(bitsmap))
	UnSet(bitsmap, 8) //删除
	fmt.Println(iter(bitsmap))

	a[4] = 5
	//for index, value := range a {
	//	fmt.Println(index, value)
	//	fmt.Printf("index = %d, value = %d (%08b)\n", index, value, value)
	//	fmt.Println(reflect.TypeOf(value))
	//}
	fmt.Println("==========================")
	err := setValue(22)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("==========================")
	for index, value := range a {
		fmt.Printf("index = %d, value = %d (%08b)\n", index, value, value)
		//fmt.Println(reflect.TypeOf(value))
	}
	fmt.Println(iter(a))

	/*
		intA := 3
		intB := 5
		c := intA &^ intB
		fmt.Printf("intA = %d %08b, intB = %d %08b, result = %d %08b\n", intA, intA, intB, intB, c, c)

		d := ^intB
		fmt.Printf("%d %08b\n", d, d)
		e := intA & d
		fmt.Printf("%d %08b\n", e, e)

		hd, _ := hamming([]byte{0}, []byte{255})
		fmt.Println(hd)
	*/
}

func hamming(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a b are not the same length")
	}

	diff := 0
	for i := 0; i < len(a); i++ {
		b1 := a[i]
		b2 := b[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (b1 & mask) != (b2 & mask) {
				diff++
			}
		}
	}
	return diff, nil
}
