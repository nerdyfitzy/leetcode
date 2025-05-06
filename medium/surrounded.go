package main

//dfs with preset directions
func solveR(board [][]byte, r, c int, directions [][]int) {
	//everything we reach will be equal to . so its easy to keep track of
	board[r][c] = '.'

	//go through each direction and get the new coords
	for _, v := range directions {
		newCoord := []int{r + v[0], c + v[1]}
		//outside bounds == skip
		if newCoord[0] <= 0 || newCoord[0] >= len(board)-1 || newCoord[1] <= 0 || newCoord[1] >= len(board[0])-1 {
			continue
		}
		//if its an o, run dfs from that
		if board[newCoord[0]][newCoord[1]] == 'O' {
			solveR(board, newCoord[0], newCoord[1], directions)
		}
	}

	return
}

func solve(board [][]byte) {
	x, y := len(board), len(board[0])
	directions := [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	//run a dfs from all the borders, if we can reach an O then its safe (simple)
	for row := 0; row < x; row++ {
		if board[row][0] == 'O' {
			solveR(board, row, 0, directions)
		}
		if board[row][y-1] == 'O' {
			solveR(board, row, y-1, directions)
		}
	}

	for col := 0; col < y; col++ {
		if board[0][col] == 'O' {
			solveR(board, 0, col, directions)
		}
		if board[x-1][col] == 'O' {
			solveR(board, x-1, col, directions)
		}
	}

	//weve changed all the reached ones to ., so now just make the . = O and anything else = X
	for i, v := range board {
		for j, _ := range v {
			if board[i][j] == '.' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}
