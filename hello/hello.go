package main

import "fmt"

var g int = 10 // 패키지 전역 변수 선언

func main() {
	// 변수 선언의 다른 형태
	minimumWage := 10 // 선언 대입문 :=을 사용해서 var 키워드와 타입을 생략할 수 있습니다.
	workingHour := 10

	income := minimumWage * workingHour

	/* 변수 선언 후, 별도로 사용하지 않는 경우 에러로 간주합니다.
	 * 타입을 생략하면 우변의 타입으로 좌변(변수)의 타입이 저장됩니다.
	 */
	helloWorld := "Hello World"

	fmt.Println(minimumWage, workingHour, income, helloWorld)

	a := 3
	var b float64 = 3.5

	// var c int = b		Error - float64 변수를 int에 대입 불가
	// d := a * b			Error - 다른 타입인 int 변수와 float64 연산 불가

	var e int64 = 7
	// f := a * e			Error - 타입이 달라서 연산 불가
	// var g int = b * 3 	Error - 실수가 정수로 자동으로 바뀌지 않습니다.

	fmt.Println(a, b, e)

	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// 01. 타입 변환
	typeChange()
}
