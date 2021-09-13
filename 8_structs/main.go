package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

//  field가 기재되어있다면 field:value로 value만 적어줄 수는 없습니다.
func main() {
	favFood := []string{"kimchi", "ramen"}
	// nico := person{"nico", 18, favFood}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}
