/// Link to problem : https://leetcode.com/problems/find-the-student-that-will-replace-the-chalk/
/// 

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	mod := k % sum
	if mod == 0 {
		return 0
	} else if len(chalk) == sum {
		return mod
	} else {
		rem := 0
		for i, v := range chalk {
			rem += v
			if rem >= mod {
				return i
			}
		}
	}
	return 0
}
