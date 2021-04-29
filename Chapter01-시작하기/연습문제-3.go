package main

import "fmt"

func fibo(N int) int {
	p := 0
	q := 1
	for N - 1 > 0 {
		tmp := p
		p = q
		q = tmp + q
		N--
	}
	return q
}

func main(){
	var N int
	fmt.Scanln(&N)
	fmt.Println(fibo(N))
}