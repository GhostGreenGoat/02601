package main

//place your functions from the assignment here.
func SumCells(cells ...Cell) Cell {
    var sum Cell
    for _,val:=range cells{
        sum[0]=sum[0]+val[0]
        sum[1]=sum[1]+val[1]
    }
    return sum

}
func ChangeDueToReactions(currentCell Cell, feedRate, killRate float64) Cell {
    var newCell Cell
    A:=currentCell[0]
    B:=currentCell[1]
    newA:=feedRate*(1-A)-A*B*B
    newB:=-killRate*B+A*B*B
    newCell[0]=newA
    newCell[1]=newB
    return newCell
}

func InField(currentBoard Board, row, col int) bool{
    numRows:=CountRows(currentBoard)
    numCols:=CountCols(currentBoard)
    if row<0||row>numRows-1||col<0||col>numCols-1{
        return false} else {return true}
}
func CountRows(currentBoard Board)int{
    return len(currentBoard)
}
func CountCols(currentBoard Board)int{
    return len(currentBoard[0])
}

func ChangeDueToDiffusion(currentBoard Board, row, col int, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell {
    var newCell Cell
    newCell[0]=0
    newCell[1]=0

    if InField(currentBoard,row-1,col-1){
            newCell[0]+=preyDiffusionRate*kernel[0][0]*currentBoard[row-1][col-1][0]
            newCell[1]+=predatorDiffusionRate*kernel[0][0]*currentBoard[row-1][col-1][1]
            }
    if InField(currentBoard,row-1,col){
            newCell[0]+=preyDiffusionRate*kernel[0][1]*currentBoard[row-1][col][0]
            newCell[1]+=predatorDiffusionRate*kernel[0][1]*currentBoard[row-1][col][1]
            }
    if InField(currentBoard,row-1,col+1){
            newCell[0]+=preyDiffusionRate*kernel[0][2]*currentBoard[row-1][col+1][0]
            newCell[1]+=predatorDiffusionRate*kernel[0][2]*currentBoard[row-1][col+1][1]
            }
    if InField(currentBoard,row,col-1){
            newCell[0]+=preyDiffusionRate*kernel[1][0]*currentBoard[row][col-1][0]
            newCell[1]+=predatorDiffusionRate*kernel[1][0]*currentBoard[row][col-1][1]
            }
    if InField(currentBoard,row,col+1){
            newCell[0]+=preyDiffusionRate*kernel[1][2]*currentBoard[row][col+1][0]
            newCell[1]+=predatorDiffusionRate*kernel[1][2]*currentBoard[row][col+1][1]
            }
    if InField(currentBoard,row+1,col-1){
            newCell[0]+=preyDiffusionRate*kernel[2][0]*currentBoard[row+1][col-1][0]
            newCell[1]+=predatorDiffusionRate*kernel[2][0]*currentBoard[row+1][col-1][1]
            }
    if InField(currentBoard,row+1,col){
            newCell[0]+=preyDiffusionRate*kernel[2][1]*currentBoard[row+1][col][0]
            newCell[1]+=predatorDiffusionRate*kernel[2][1]*currentBoard[row+1][col][1]
            }
    if InField(currentBoard,row+1,col+1){
            newCell[0]+=preyDiffusionRate*kernel[2][2]*currentBoard[row+1][col+1][0]
            newCell[1]+=predatorDiffusionRate*kernel[2][2]*currentBoard[row+1][col+1][1]
            }
    newCell[0]+=preyDiffusionRate*kernel[1][1]*currentBoard[row][col][0]
    newCell[1]+=predatorDiffusionRate*kernel[1][1]*currentBoard[row][col][1]
    return newCell
}

func UpdateCell(currentBoard Board, row,col int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell{
    currentCell:=currentBoard[row][col]
    diffusionValues:=ChangeDueToDiffusion(currentBoard, row, col, preyDiffusionRate, predatorDiffusionRate, kernel)
    reactionValues:=ChangeDueToReactions(currentCell, feedRate, killRate)
    return SumCells(currentCell, diffusionValues, reactionValues)
}
func UpdateBoard(currentBoard Board, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Board {
    numRows:=CountRows(currentBoard)
    numCols:=CountCols(currentBoard)
    newBoard:=InitializeBoard(numRows, numCols)
    for row:=0;row<numRows;row++{
        for col:=0;col<numCols;col++{
            newBoard[row][col]=UpdateCell(currentBoard, row, col, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
        }
    }
    return newBoard
}

func InitializeBoard(rows,cols int) Board{
    newBoard:=make([][]Cell,rows)
    for r:=0;r<rows;r++{
        newBoard[r]=make([]Cell,cols)
    }
    return newBoard
}

func SimulateGrayScott(initialBoard Board, numGens int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) []Board {
    boards:=make([]Board,numGens+1)
    boards[0]=initialBoard
    for i:=1;i<=numGens;i++{
        boards[i]=UpdateBoard(boards[i-1], feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
    }
    return boards
}
