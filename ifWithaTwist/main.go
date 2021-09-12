package main

import "fmt"

func canIdrink(age int) bool {
	/** switch age {
	case 10:
		return false
	case 18:
		return true
	}

	// switch ~ case문에서도 if 바로 직후에 variable을 만들었던 것처럼 사용이 가능하다.
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	} **/

	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}

	return false
}

func canIDring(age int) bool {
	/* koreanAge := age + 2은 if koreanAge := age + 2;와 같은 내용입니다.
	 * Go언어에서는 if문에 변수 선언이 사용가능합니다. 다만 변수 선언이 세미콜론(;)은 꼭 기재되어야 합니다.
	 * variable expression으로 불린다.
	 */

	// if문에다가 작성한 변수는 "if-else의 조건에만 사용하기 위해 variable을 생성했구나"로 인식이 될 수 있다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

func main() {
	fmt.Println(canIDring(16))
	fmt.Println(canIdrink(50))
}
