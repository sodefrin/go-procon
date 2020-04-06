package main

import (
	"fmt"
	"math"
)

const Mod = 1000000007

func main() {
	var N, M int

	fmt.Scan(&N, &M)

	p := Prime(M)
	n := make([]int, len(p))

	for i, v := range p {
		for M%v == 0 {
			n[i]++
			M /= v
		}
	}

	if M != 1 {
		p = append(p, M)
		n = append(n, 1)
	}

	var ret int64 = 1
	for _, v := range n {
		if v > 0 {
			ret = ret * C(int64(v+N-1), int64(v)) % Mod
		}
	}

	fmt.Println(ret)
}

func Prime(M int) []int {
	max := int(math.Pow(float64(M), 0.5)) + 1
	tmp := make([]bool, max+1)

	for i := 2; i < max+1; i++ {
		tmp[i] = true
	}

	ret := []int{}
	for i := 2; i < max; i++ {
		if tmp[i] {
			for j := i * i; j < max+1; j += i {
				tmp[j] = false
			}
		}
	}

	for i, v := range tmp {
		if v {
			ret = append(ret, i)
		}
	}

	return ret
}

func C(a, b int64) int64 {
	var ret int64 = 1
	var i int64
	for i = 1; i <= b; i++ {
		ret = (ret * (a - int64(i) + 1)) % Mod
		ret = (ret * Inv(i)) % Mod
	}

	return ret
}

func Inv(b int64) int64 {
	base := Mod - 2
	ret := int64(1)
	for base > 0 {
		if base&1 == 1 {
			ret = ret * b % Mod
		}
		b = (b * b) % Mod
		base >>= 1
	}
	return ret
}
