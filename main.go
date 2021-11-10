package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	first := l.PushFront("first")
	last := l.PushBack("last")
	l.InsertBefore("mid",last) //在 last前边插入元素
	l.InsertAfter("second",first) // 在first后边插入一个元素
	//遍历一下所有的元素
	for i := l.Front();i != nil ;i = i.Next() {
		fmt.Println(i.Value)
	}
}
