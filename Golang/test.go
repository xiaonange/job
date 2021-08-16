package main

import (
	"fmt"
	"regexp"
	"sync"
)

func main() {
	str := "8765213"
	matched, err := regexp.MatchString(`^([1-9])\d{7}$`, str)
	if err !=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(matched)

}

func printNum(group *sync.WaitGroup,i,num int,ch chan int )  {
	select {
	case <-ch:
		for n:=num; n > 0;n-- {
			i++
			fmt.Println(i)
		}


	default:
		break

}
}

func printLetter( group *sync.WaitGroup,num int,ch chan string )  {

}