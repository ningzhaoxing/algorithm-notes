package 筛法

func getPrim(n int) []int {
	isPrim := make([]bool, n+1)
	prims := make([]int, 0)

	for i := 2; i <= n; i++ {
		if !isPrim[i] {
			prims = append(prims, i)
		}
		for _, p := range prims {
			if i*p > n {
				break
			}
			isPrim[i*p] = true
			if i%p == 0 {
				break
			}
		}
	}
	return prims
}
