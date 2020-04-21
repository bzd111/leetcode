package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans += 1
				dfs(grid, m, n, i, j)
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, m, n, i, j int) {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, m, n, i+1, j)
	dfs(grid, m, n, i-1, j)
	dfs(grid, m, n, i, j+1)
	dfs(grid, m, n, i, j-1)
}

func main() {

	// Input:
	// 11110
	// 11010
	// 11000
	// 00000
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
	//g Input:
	//g 11000
	//g 11000
	//g 00100
	//g 00011
	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}
