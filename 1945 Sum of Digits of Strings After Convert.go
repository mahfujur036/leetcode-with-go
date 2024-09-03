// Problem 1945 https://leetcode.com/problems/sum-of-digits-of-string-after-convert
// Daily Challenge problem for September 3, 2024

func sumOfDigits(a int) int {
	sum := 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	return sum
}

func getLucky(s string, k int) int {
	digits := []int{}
	for _, v := range s {
		digits = append(digits, int(v-'a'+1))
	}
	sum := 0
	for i, v := range digits {
		if v > 9 {
			digits[i] = v/10 + v%10
		}
		sum += digits[i]
	}
	if k == 1 {
		return sum
	}
	for i := 1; i < k; i++ {
		sum = sumOfDigits(sum)
	}
	return sum
}
