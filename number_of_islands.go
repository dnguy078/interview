package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	ans := 0
	dirs := []int{0, 1, 0, -1, 0}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				queue := [][]int{}
				queue = append(queue, []int{i, j})
				for len(queue) != 0 {
					pop := queue[0]
					queue = queue[1:]
					for k := 0; k < 4; k++ {
						x := pop[0] + dirs[k]
						y := pop[1] + dirs[k+1]
						if x >= 0 && y >= 0 && x < m && y < n && grid[x][y] == '1' {
							grid[x][y] = '0'
							queue = append(queue, []int{x, y})
						}
					}
				}
			}
		}
	}
	return ans
}
