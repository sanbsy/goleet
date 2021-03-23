package leet

var maxD = [64]int{
	0, 1, 1, 2,
}

// 整数拆分
// https://leetcode-cn.com/problems/integer-break/
func integerBreak(n int) int {
	// 题设限制 n<=58
	if n > 59 {
		return 0
	}

	if maxD[n] > 0 {
		return maxD[n]
	}

	mm := 0
	for i := 1; i < n; i++ {
		num1, num2 := i, n-i
		if kk := integerBreak(num1); kk > num1 {
			num1 = kk
		}
		if kk := integerBreak(num2); kk > num2 {
			num2 = kk
		}
		if kk := num1 * num2; mm < kk {
			mm = kk
		}
	}

	maxD[n] = mm
	return mm
}
