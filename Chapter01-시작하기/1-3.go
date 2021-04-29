package main

import "fmt"

func main(){
    fmt.Println("Hello, playground") //세미콜론이 없이 보이지만 붙어서 컴파일러에게 전달됨
	
	var x int // "변수 x는 int형이다", 파스칼, 델파이도 이와 같이 사용
	var arr [5]int
	var p *int
	var x int = 10

	//GO는 정적으로 자료형을 사용할 수 있음 (파이썬처럼 동적으로 사용되는 것처럼 보이나 아님)
	var i = 10	//C++ auto i = 10;와 비슷
	var p = &i  //C++ auto *p = &i;

	i := 10
	p := &i
}