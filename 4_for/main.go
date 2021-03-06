package main

import "fmt"

/**
 * Go언어에서는 반복문이 for문 밖에 없다.
 * range : array에 loop 동작을 적용할 수 있도록 해줍니다.
 * range는 for문을 돌리면서 index를 던져줍니다.
 */

func superAdd(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
