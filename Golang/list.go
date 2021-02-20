package main

import (
	"container/list"
	"fmt"
	"sort"
) ///kənˈteɪnə(r)

func main() {
	//队列:是限制在两端进行插入操作和删除操作的线性表。插入在队尾，删除在对头；特点：先进先出
	queue := list.New()
	queue.PushBack(123)
	queue.PushBack("1111")
	print(queue.Front().Value)
	queue.Remove(queue.Front())
	queue.Remove(queue.Front())
	print(queue.Len())
	queue.PushBack(123)
	print(queue.Len())
	//推栈:必须按"后进先出"的规则进行操作
	stack := list.New()
	stack.PushBack(123)
	stack.PushBack("1111")
	print(stack.Front().Value)
	queue.Remove(stack.Front())
	queue.Remove(stack.Front())
	print(stack.Len())
	stack.PushBack(123)
	print(stack.Len())
	//哈希
	//HashMap之所以高效，是因为其结合了顺序存储(数组)和链式存储(链表)两种存储结构

	//快速排序
	fmt.Print("快速排序")
	TestSort()
}

//快速排序
func TestSort()  {
	a :=[]int{1,2,3,4,5,6,7,8,9,20,10,11}
	sort.Ints(a)
	fmt.Println(a)
}