package main

import "fmt"

var testName bool = false

// testName := 'asd' 축약형은 오로지 func안에서만 가능하고 변수에만 적용이 가능하다.
func variableConstants() {
	// const name string = "JeongHyeonMun"
	name := "hyeonMun"
	name = "lynn"
	// name = "Lynn"	상수 값이라 값 변경 불가 X

	fmt.Println(name, testName)
}
