package main

import "fmt"

/**
 * map은 key와 value가 존재한다.
 * Go 언어에서는 key의 타입과 value의 타입을 모두 지정해주어야 한다.
 */

func main() {
	nico := map[string]string{"name": "nico", "age": "12"}
	for _, value := range nico {
		fmt.Println(value)
	}
}
