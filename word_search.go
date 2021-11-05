package main

import "fmt"

func exist(board [][]byte, word string) bool {
	maxRow, maxCol := len(board), len(board[0])
	seen := make(map[coordinate]bool)
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if dfs(seen, board, word, maxRow, maxCol, r, c, 0) {
				return true
			}
		}
	}
	return false
}

type coordinate struct {
	r, c int
}

func dfs(path map[coordinate]bool, board [][]byte, word string, maxRow, maxCol, r, c, i int) bool {
	if i == len(word) {
		fmt.Println(path)
		return true
	}
	if r < 0 || c < 0 || r >= maxRow || c >= maxCol || word[i] != board[r][c] || path[coordinate{r, c}] {
		return false
	}
	path[coordinate{r, c}] = true
	result := dfs(path, board, word, maxRow, maxCol, r+1, c, i+1) ||
		dfs(path, board, word, maxRow, maxCol, r-1, c, i+1) ||
		dfs(path, board, word, maxRow, maxCol, r, c+1, i+1) ||
		dfs(path, board, word, maxRow, maxCol, r, c-1, i+1)

	delete(path, coordinate{r, c})
	return result
}
