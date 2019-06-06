package p0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n == 0 {
		return 0
	}
	m := len(obstacleGrid[0])
	count := make([][]int, n)
	for i := range obstacleGrid {
		count[i] = make([]int, m)
	}
	for i := range obstacleGrid[0] {
		if obstacleGrid[0][i] == 1 {
			break
		}
		count[0][i] = 1
	}
	for i := range obstacleGrid {
		if obstacleGrid[i][0] == 1 {
			break
		}
		count[i][0] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			count[i][j] = count[i-1][j] + count[i][j-1]
		}
	}
	return count[n-1][m-1]
}
