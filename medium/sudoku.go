package main

func isValidSudoku(board [][]byte) bool {
	horizontalMap := make([]map[byte]bool, 9)
	verticalMap := make([]map[byte]bool, 9)
	squareMap := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		horizontalMap[i] = make(map[byte]bool)
		verticalMap[i] = make(map[byte]bool)
		squareMap[i] = make(map[byte]bool)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			char := board[row][col]
			if char == '.' {
				continue
			}

			squareId := ((row / 3) * 3) + (col / 3)

			if horizontalMap[row][char] || verticalMap[col][char] || squareMap[squareId][char] {
				return false
			} else {
				horizontalMap[row][char] = true
				verticalMap[col][char] = true
				squareMap[squareId][char] = true
			}
		}
	}

	return true
}
