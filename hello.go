package main

import "fmt"
import "strings"

func main() {
	fplus := func(a, b int) int { return a + b }
	fmt.Println(fplus(2, 3))
}

func rune_f() {
	var s = "abc我是个中国人"
	var p = strings.Map(strRep, s)
	fmt.Println(s)
	fmt.Println(p)
}

func strRep(s rune) (p rune) {
	switch {
	case s >= 128:
		p = '?'
	default:
		p = s
	}
	return
}

func fibo_p() {
	for i := 0; i < 10; i++ {
		ret, pos := fibo(i)
		fmt.Printf("(%d, %d)", pos, ret)
	}
	fmt.Println()
}

func fibo(n int) (ret int, pos int) {
	switch n {
	case 0, 1:
		ret = 1
	default:
		ret1, _ := fibo(n - 1)
		ret2, _ := fibo(n - 2)
		ret = ret1 + ret2
	}
	pos = n
	return
}

func hello() {
	fmt.Println("Greetings, hello hello  gopher")
}
