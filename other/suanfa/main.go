package main

import (
	"fmt"
	"suanfa/public"
)

func main()  {
	arr :=[]int{1,2,3,-1,-5,11,-20,6}
fmt.Println(public.MaxSubArray(arr))
	fmt.Println(public.MaxSubArrayDT(arr))
}
