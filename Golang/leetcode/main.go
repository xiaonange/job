package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	/*arr := &ListNode{Val:1,Next:&ListNode{Val:2,Next:&ListNode{Val:3,Next:&ListNode{Val:4,Next:&ListNode{}}}}}
	result := reverseList(arr)
	fmt.Print(result.Val)*/
	var a  uint = 1
	var b uint =2
	fmt.Print(a-b)
	var c  int = 1
	var d int =2
	fmt.Print(c-d)
}
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for head.Next != nil {
		t := head.Next.Next
		head.Next.Next = cur // 反转原指针方向
		cur = head.Next // 将新头节点移到下一位
		head.Next = t // 连接回断开的地方，继续重复上面操作
	}

	return cur
}