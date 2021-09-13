package main

import "fmt"

/**
 * array는 변수명 := [배열 개수] 타입 {배열값}
 * slices는 기존 배열 선언에서 [배열 개수] 부분을 []로 지정해주면 됩니다.
 */

func main() {
	/** names := [5]string{"hyeonmun", "test", "nico", "lynn", "dal"}
	fmt.Println(names) **/

	// javascript나 python과 다르게 Go 언어에서는 append를 하게되면 array가 수정되지않고 새로운 slice를 return합니다.
	names := []string{"noci", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}
