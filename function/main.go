package main

import (
	"fmt"
	"strings"
)

/**
 * Go언어에서는 함수에서 받는 인자값에 이것들이 무엇인지 알려주어야 한다. (타입 지정)
 * 또한 return에 있어서도 타입을 지정해주어야 한다. **/
func multiply(a, b int) int {
	return a * b
}

/**
 * multiple value를 반환하는 function을 만드는 방법입니다.
 */
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

/**
 * multiple value를 반환하는 function을 만드는 방법입니다.
 * "defer"은 해당 function이 끝나고 나서 코드를 실행시키는 코드입니다.
 */
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

/**
 * variable의 타입 앞에 ...을 찍게 된다면 arguments을 무제한으로 받을 수 있다.
 */
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	//fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("HyeonMun")
	fmt.Println(totalLength, upperName)

	/* Go가 자동으로 작동할 수 있게 처리를 해준다. 이 부분을 naked return이라고 명칭한다.
	 * return할 variable을 굳이 명시하지 않은 부분입니다.
	 */
	totalLength2, upperName2 := lenAndUpper2("HyeonMun")
	fmt.Println(totalLength2, upperName2)

	/* Go 언어에서는 value값을 무시하는 코드가 있는데
	 *underscore(_)는 무시된 value입니다. 컴파일러가 쳐다보지 않는 내용입니다.
	 */
	testLength, _ := lenAndUpper("HyeonMun")
	fmt.Println(testLength)

	repeatMe("mun", "lynn", "dal", "marl", "flynn")
}
