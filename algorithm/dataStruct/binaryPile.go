package main

import "fmt"

// BinaryPile 大顶堆
type BinaryPile struct {
	h []int
}

func NewBinaryPile(arr []int) *BinaryPile {
	return &BinaryPile{h: arr}
}

func (b *BinaryPile) Build() {
	for i := 1; i < len(b.h); i++ {
		b.up(i)
	}
	fmt.Println("构建成功")
}

func (b *BinaryPile) Get() int {
	return b.h[1]
}

func (b *BinaryPile) Add(val int) {
	b.h = append(b.h, val)
	b.up(len(b.h) - 1)
}

func (b *BinaryPile) Delete() {
	b.h[0], b.h[len(b.h)-1] = b.h[len(b.h)-1], b.h[0]
	b.h = b.h[:len(b.h)-1]
	b.down(0)
}

func (b *BinaryPile) up(x int) {
	for b.h[x] > b.h[x/2] && x > 1 {
		b.h[x], b.h[x/2] = b.h[x/2], b.h[x]
		x = x / 2
	}
}

func (b *BinaryPile) down(x int) {
	for x*2 <= len(b.h) {
		if b.h[x*2] > b.h[x] {
			b.h[x*2], b.h[x] = b.h[x], b.h[x*2]
			x = x * 2
		} else if b.h[x*2+1] > b.h[x] {
			b.h[x*2+1], b.h[x] = b.h[x], b.h[x*2+1]
			x = x*2 + 1
		}
	}
}

func main() {
	arr := []int{0, 1, 3, 5, 2}
	bt := NewBinaryPile(arr)
	fmt.Println("创建成功")
	bt.Build()

	bt.Add(6)
	fmt.Println(bt.h)
	fmt.Println(bt.Get())
}
