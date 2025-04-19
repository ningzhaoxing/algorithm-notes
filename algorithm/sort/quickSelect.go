package main

func quickSelect(l, r, k int, a []int) int {
	if l == r {
		return a[l]
	}

	i, j := partition(a, l, r)
	lessCount := i - l + 1
	equalCount := j - i - 1

	if k <= lessCount { // k在小于区域
		return quickSelect(l, i, k, a)
	} else if k > lessCount+equalCount { // k在大于区域
		return quickSelect(j, r, k-(lessCount+equalCount), a)
	}
	// k在等于区域
	return a[i+1]
}

func quickSelectCaller(nums []int, k int) int {
	return quickSelect(0, len(nums)-1, len(nums)-k+1, nums)
}
