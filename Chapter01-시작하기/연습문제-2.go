package main

import "fmt"

func Hanoi(N int){
	fmt.Println("Number of disks:3")
	hanoi(N, 1,3,2)
}

func hanoi(n, start, work, target int){
	if n == 1{
		fmt.Println(start, "→", target)
	}else{
		hanoi(n-1, start, target, work)
		fmt.Println(start, "→", target)
		hanoi(n-1, work, start, target)
	}
}

func main(){	
	Hanoi(3)
}