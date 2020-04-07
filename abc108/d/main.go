package main

import "fmt"

func main() {
	var L int
	fmt.Scan(&L)

	a, b := calc(L)
	remain := L - 1 - a
	cnt := 0
	for remain > 0 {
		if remain&1 == 1 {
			cnt++
		}
		remain >>= 1
	}
	fmt.Printf("%d %d\n", b, 2*(b-1)+cnt)

	tmp := 1
	for i := 1; i < b; i++ {
		fmt.Printf("%d %d %d\n", i, i+1, 0)
		fmt.Printf("%d %d %d\n", i, i+1, tmp)
		tmp *= 2
	}

	remain = L - 1 - a
	t := 1
	t2 := 1
	t3 := 1
	for remain > 0 {
		if remain&1 == 1 {
			fmt.Printf("%d %d %d\n", t, b, a+t3)
			t3 += t2
		}
		remain >>= 1
		t++
		t2 <<= 1
	}
}

func calc(L int) (int, int) {
	ret1 := 1
	ret2 := 0
	for L > 0 {
		ret1 <<= 1
		ret2++
		L >>= 1
	}
	return ret1/2 - 1, ret2
}
