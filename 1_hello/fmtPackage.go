package main

// fmt 패키지는 표준 입출력을 담당하는 패키지입니다.
import "fmt"

func fmtPackage() {
	var a int = 10
	var b int = 20
	var c int = 123456789
	var f float64 = 32799438743.8297

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f)
	fmt.Printf("a: %d b: %d f:%f\n", a, b, f)

	/** Print => 함수 입력값들을 출력합니다.
	 	Println => 함수 입력값들을 출력하고 개행합니다.
		Printf => 서식에 맞도록 입력값들을 출력합니다. **/

	/**
	 * Printf 서식
	 * %d => 10진수 정숫값으로 출력합니다.
	 * %f => 실숫값 그대로 출력 (123.456)
	 * %s => 문자열을 출력합니다.
	**/

	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d, %05d\n", a, b)
	fmt.Printf("%-05d, %-05d\n", a, b)

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)

	// 실수 소수점 이하 자리수
	var m = 324.13455
	var l = 3.14

	fmt.Printf("%08.2f\n", m)
	fmt.Printf("%08.2g\n", m)
	fmt.Printf("%8.5g\n", m)
	fmt.Printf("%f\n", l)
}
