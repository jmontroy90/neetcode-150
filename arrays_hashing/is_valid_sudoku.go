package arrays_hashing

func IsValidSudoku(board [][]int) bool {
	for i1 := 0; i1 < 9; i1++ {
		var rowBM, colBM, boxBM [10]byte
		for i2 := 0; i2 < 9; i2++ {
			if v := board[i1][i2]; v != 0 && rowBM[v] != 0 {
				return false
			} else {
				rowBM[v]++
			}
			if v := board[i2][i1]; v != 0 && colBM[v] != 0 {
				return false
			} else {
				colBM[v]++
			}
			// boxes; this can be thought of just a slow index and a fast index slowly cascading through.
			row := i2%3 + (i1%3)*3
			col := i2/3 + (i1/3)*3
			if v := board[row][col]; v != 0 && boxBM[v] != 0 {
				return false
			} else {
				boxBM[v]++
			}
		}
	}
	return true
}
