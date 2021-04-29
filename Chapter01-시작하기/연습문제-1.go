package main

import "fmt"

func main(){

	var N int
	fmt.Scanln(&N)

	for i := 1; i <= N; i++{
		fmt.Printf("타잔이 %d원짜리 팬티를 입고, %d원짜리 칼을 차고 노래를 한다\n",i*10,(i+1)*10)
	}
}