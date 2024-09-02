func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		start := i * n
		end := start + n
		result[i] = original[start:end]
	}
	return result    
}
