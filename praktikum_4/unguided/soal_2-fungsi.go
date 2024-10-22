package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fogoh(a int) int {
	return f(g(h(a)))
}

func gohof(b int) int {
	return g(h(f(b)))
}

func hofog(c int) int {
	return h(f(g(c)))
}

func soal_2_fungsi() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Printf("\n(fogog)(%d) = %d\n", a, fogoh(a))
	fmt.Printf("(gohof)(%d) = %d\n", b, gohof(b))
	fmt.Printf("(hofog)(%d) = %d\n", c, hofog(c))
}
