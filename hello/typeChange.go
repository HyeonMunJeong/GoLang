package main

import "fmt"

func typeChange() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)
	d := float64(a * c)

	// 3 * 3 = 9

	var e int64 = 7
	f := int64(d) * e

	// 9 * 7 = 63

	var g int = int(b * 3) // 3 * 3 = 9
	var h int = int(b) * 3 // 3 * 3 = 9

	n := 3.5
	var m int = int(n * 3)

	fmt.Println(g, h, f, m)

	var t int16 = 3456
	var l int8 = int8(t)

	fmt.Println(t, l)

	var ml float32 = 1234.523
	var mb float32 = 3456.123
	var mc float32 = ml * mb
	var md float32 = mc * 3

	fmt.Println(ml, mb, mc, md)
}
