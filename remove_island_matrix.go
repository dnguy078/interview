package main

// https://www.algoexpert.io/questions/Remove%20Islands

func RemoveIslands(matrix [][]int) [][]int {
	// Write your code here.
	connectedToBorder := make([][]bool, len(matrix))
	for i := range matrix {
		connectedToBorder[i] = make([]bool, len(matrix[0]))
	}
	// this makes a matrix of T/F for indexes on the border or touches the border
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			isBorder := row == 0 || col == 0 || row == len(matrix)-1 || col == len(matrix[row])-1
			if !isBorder {
				continue
			}
			if matrix[row][col] != 1 {
				continue
			}
			markConnectedToBorder(matrix, row, col, connectedToBorder)
		}
	}

	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[row])-1; col++ {
			if connectedToBorder[row][col] {
				continue
			}
			matrix[row][col] = 0
		}
	}

	return nil
}

func markConnectedToBorder(matrix [][]int, row, col int, connectedToBorder [][]bool) {
	stack := [][]int{{row, col}}
	var currentPosition []int
	for len(stack) > 0 {
		currentPosition, stack = stack[len(stack)-1], stack[:len(stack)-1]
		currentRow, currentCol := currentPosition[0], currentPosition[1]
		if connectedToBorder[currentRow][currentCol] {
			continue
		}

		connectedToBorder[currentRow][currentCol] = true
		neighbors := getNeighbors(matrix, currentRow, currentCol)
		for _, neighbor := range neighbors {
			row, col := neighbor[0], neighbor[1]
			if matrix[row][col] != 1 {
				continue
			}
			stack = append(stack, neighbor)
		}
	}
}

func getNeighbors(matrix [][]int, row, col int) [][]int {
	neighbors := make([][]int, 0)
	numRows := len(matrix)
	numCols := len(matrix[row])
	// up
	if row-1 >= 0 {
		neighbors = append(neighbors, []int{row - 1, col})
	}
	// down
	if row+1 < numRows {
		neighbors = append(neighbors, []int{row + 1, col})
	}
	//left
	if col-1 >= 0 {
		neighbors = append(neighbors, []int{row, col - 1})
	}
	//right
	if col+1 < numCols {
		neighbors = append(neighbors, []int{row, col + 1})
	}
	return neighbors
}
