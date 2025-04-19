package 筛法

import "fmt"

// Eratosthenes 埃氏筛法
func Eratosthenes(n int) []int {
	isPrim := make([]bool, n+1)
	prims := make([]int, 0)

	isPrim[0], isPrim[1] = false, false
	for i := 2; i <= n; i++ {
		isPrim[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrim[i] {
			for j := i * i; j <= n; j += i {
				isPrim[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if isPrim[i] {
			prims = append(prims, i)
		}
	}

	return prims
}

func EraCaller(n int) {
	fmt.Println(Eratosthenes(n))
}
