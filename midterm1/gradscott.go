package main

type Board [][]float64

func UpdateBoard(currentBoard Board, DiffusionRate float64, kernel [3][3]float64) Board {
	numRows := CountRows(currentBoard)
	numCols := CountCols(currentBoard)
	newBoard := InitializeBoard(numRows, numCols)
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			newBoard[row][col] = UpdateCell(currentBoard, row, col, DiffusionRate, kernel)
		}
	}
	return newBoard
}

func UpdateCell(currentBoard Board, row, col int, DiffusionRate float64, kernel [3][3]float64) float64 {
	currentCell := currentBoard[row][col]
	diffusionValues := ChangeDueToDiffusion(currentBoard, row, col, DiffusionRate, kernel)
	return SumCells(currentCell, diffusionValues)
}

func InitializeBoard(rows, cols int) Board {
	newBoard := make([][]float64, rows)
	for r := 0; r < rows; r++ {
		newBoard[r] = make([]float64, cols)
	}
	return newBoard
}

func SumCells(cells ...float64) float64 {
	var sum float64
	for _, val := range cells {
		sum = sum + val
	}
	return sum

}

func InField(currentBoard Board, row, col int) bool {
	numRows := CountRows(currentBoard)
	numCols := CountCols(currentBoard)
	if row < 0 || row > numRows-1 || col < 0 || col > numCols-1 {
		return false
	} else {
		return true
	}
}
func CountRows(currentBoard Board) int {
	return len(currentBoard)
}
func CountCols(currentBoard Board) int {
	return len(currentBoard[0])
}

func ChangeDueToDiffusion(currentBoard Board, row, col int, DiffusionRate float64, kernel [3][3]float64) float64 {
	var newCell float64
	newCell = 0
	lastrow := CountRows(currentBoard) - 1
	lastcol := CountCols(currentBoard) - 1

	if InField(currentBoard, row-1, col-1) {
		newCell += DiffusionRate * kernel[0][0] * currentBoard[row-1][col-1]
	} else if row-1 < 0 && col-1 < 0 {
		newCell += DiffusionRate * kernel[0][0] * currentBoard[lastrow][lastcol]
	} else if col-1 < 0 {
		newCell += DiffusionRate * kernel[0][0] * currentBoard[row-1][lastcol]
	} else if row-1 < 0 {
		newCell += DiffusionRate * kernel[0][0] * currentBoard[lastrow][col-1]
	}
	if InField(currentBoard, row-1, col) {
		newCell += DiffusionRate * kernel[0][1] * currentBoard[row-1][col]
	} else {
		newCell += DiffusionRate * kernel[0][1] * currentBoard[lastrow][col]
	}
	if InField(currentBoard, row-1, col+1) {
		newCell += DiffusionRate * kernel[0][2] * currentBoard[row-1][col+1]
	} else if row-1 < 0 && col+1 > lastcol {
		newCell += DiffusionRate * kernel[0][2] * currentBoard[lastrow][0]
	} else if row-1 < 0 {
		newCell += DiffusionRate * kernel[0][2] * currentBoard[lastrow][col+1]
	} else if col+1 > lastcol {
		newCell += DiffusionRate * kernel[0][2] * currentBoard[row-1][0]
	}
	if InField(currentBoard, row, col-1) {
		newCell += DiffusionRate * kernel[1][0] * currentBoard[row][col-1]
	} else {
		newCell += DiffusionRate * kernel[1][0] * currentBoard[row][lastcol]
	}
	if InField(currentBoard, row, col+1) {
		newCell += DiffusionRate * kernel[1][2] * currentBoard[row][col+1]
	} else {
		newCell += DiffusionRate * kernel[1][2] * currentBoard[row][0]
	}
	if InField(currentBoard, row+1, col-1) {
		newCell += DiffusionRate * kernel[2][0] * currentBoard[row+1][col-1]
	} else if row+1 > lastrow && col-1 < 0 {
		newCell += DiffusionRate * kernel[2][0] * currentBoard[0][lastcol]
	} else if row+1 > lastrow {
		newCell += DiffusionRate * kernel[2][0] * currentBoard[0][col-1]
	} else if col-1 < 0 {
		newCell += DiffusionRate * kernel[2][0] * currentBoard[row+1][lastcol]
	}
	if InField(currentBoard, row+1, col) {
		newCell += DiffusionRate * kernel[2][1] * currentBoard[row+1][col]
	} else {
		newCell += DiffusionRate * kernel[2][1] * currentBoard[0][col]
	}
	if InField(currentBoard, row+1, col+1) {
		newCell += DiffusionRate * kernel[2][2] * currentBoard[row+1][col+1]
	} else if row+1 > lastrow && col+1 > lastcol {
		newCell += DiffusionRate * kernel[2][2] * currentBoard[0][0]
	} else if row+1 > lastrow {
		newCell += DiffusionRate * kernel[2][2] * currentBoard[0][col+1]
	} else if col+1 > lastcol {
		newCell += DiffusionRate * kernel[2][2] * currentBoard[row+1][0]
	}
	newCell += DiffusionRate * kernel[1][1] * currentBoard[row][col]
	return newCell
}

/*
func DiffuseBoardOneParticleTorus(currentBoard [][]float64, diffusionRate float64, kernel [3][3]float64) [][]float64 {

}
*/
