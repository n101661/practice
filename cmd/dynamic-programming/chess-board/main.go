// source: https://www.codewars.com/kata/5bc6f9110ca59325c1000254/csharp
package chessboard

func ChessBoard(board [][]int) map[int]int {
	cBoard := newSquareSizeCounterBoard(len(board), len(board[0]))

	for squareSize := 2; ; squareSize++ {
		cBoard.init(board, squareSize)

		for i := squareSize - 1; i < len(cBoard); i++ {
			checker := newChessChecker(board, squareSize, i-squareSize+2)
			for j := squareSize - 1; j < len(cBoard[i]); j++ {
				num := cBoard.Top(j, i, squareSize) + cBoard.Left(j, i, squareSize) - cBoard.TopLeft(j, i, squareSize)
				if !checker.CheckTopFirst(board, position{i: i + 1, j: j + 1}, squareSize) {
					num++
				}
				if num > 0 {
					cBoard[i][j][squareSize] = num
				}
			}
		}

		if cBoard[len(cBoard)-1][len(cBoard[0])-1][squareSize] == 0 ||
			squareSize == len(board) {
			break
		}
	}

	return cBoard[len(cBoard)-1][len(cBoard[0])-1]
}

type squareSizeCounterBoard [][]map[int]int

func newSquareSizeCounterBoard(iLen, jLen int) squareSizeCounterBoard {
	newBoard, cellNum := make(squareSizeCounterBoard, iLen-1), jLen-1
	for i := range newBoard {
		newBoard[i] = make([]map[int]int, cellNum)

		for j := range newBoard[i] {
			newBoard[i][j] = map[int]int{}
		}
	}
	return newBoard
}

func (board *squareSizeCounterBoard) init(chessBoard [][]int, squareSize int) {
	chessBoardStart := squareSize - 1
	boardStart := squareSize - 2

	// init row
	checker := newChessChecker(chessBoard, squareSize, 0)
	for x := chessBoardStart; x < len(chessBoard[chessBoardStart]); x++ {
		if x == chessBoardStart {
			if !checker.CheckTopFirst(chessBoard, position{i: chessBoardStart, j: chessBoardStart}, squareSize) {
				(*board)[boardStart][boardStart][squareSize] = 1
			}
			continue
		}

		num := (*board)[boardStart][x-2][squareSize]
		if !checker.CheckTopFirst(chessBoard, position{i: chessBoardStart, j: x}, squareSize) {
			num++
		}
		if num > 0 {
			(*board)[boardStart][x-1][squareSize] = num
		}
	}

	// init column
	for y := squareSize; y < len(chessBoard); y++ {
		num := (*board)[y-2][boardStart][squareSize]
		if !checker.CheckLeftFirst(chessBoard, position{i: y, j: chessBoardStart}, squareSize) {
			num++
		}
		if num > 0 {
			(*board)[y-1][boardStart][squareSize] = num
		}
	}
}

// TopLeft returns the number of the given square size.
func (board squareSizeCounterBoard) TopLeft(x, y int, squareSize int) int {
	if x == 0 || y == 0 || len(board) <= y || len(board[0]) <= x {
		return 0
	}

	return board[y-1][x-1][squareSize]
}

// Top returns the number of the given square size.
func (board squareSizeCounterBoard) Top(x, y int, squareSize int) int {
	if y == 0 || len(board) <= y {
		return 0
	}

	return board[y-1][x][squareSize]
}

// Left returns the number of the given square size.
func (board squareSizeCounterBoard) Left(x, y int, squareSize int) int {
	if x == 0 || len(board[0]) <= x {
		return 0
	}

	return board[y][x-1][squareSize]
}

type position struct{ i, j int }

type chessChecker struct {
	topCache  *position
	leftCache *position
}

func newChessChecker(board [][]int, checkRange, iShift int) *chessChecker {
	start := checkRange - 1

	checker := chessChecker{}

	// top first
topFirst:
	for j := start; j >= 0; j-- {
		for i := start + iShift; i >= iShift; i-- {
			if hasChess(board[i][j]) {
				checker.topCache = &position{
					i: i,
					j: j,
				}
				break topFirst
			}
		}
	}

	// left first
leftFirst:
	for i := start + iShift; i >= iShift; i-- {
		for j := start; j >= 0; j-- {
			if hasChess(board[i][j]) {
				checker.leftCache = &position{
					i: i,
					j: j,
				}
				break leftFirst
			}
		}
	}

	return &checker
}

// CheckTopFirst returns true if there is a chess at least.
// It checks top left direction.
//
// Example: There is a chess board as below:
//
//   [ 1, 0, 1, 1, 1]
//   [ 1, 1, 1, 1, 1]
//   [ 1, 1, 1, 1, 1]
//   [ 1, 1, 1, 1, 1]
//   [ 1, 1, 1, 1, 1]
//
// start=[2, 3], checkRange=3, it checks [2, 3], [1, 3], [0, 3],
// [2, 2], [1, 2], [0, 2], [2, 1], [1, 1], [0, 1] and finds the
// chess at [0, 1].
func (checker *chessChecker) CheckTopFirst(board [][]int, start position, checkRange int) bool {
	iEnd := start.i - checkRange + 1
	jEnd := start.j - checkRange + 1

	for i := start.i; i >= iEnd; i-- {
		if hasChess(board[i][start.j]) {
			if checker.topCache == nil {
				checker.topCache = &position{}
			}

			checker.topCache.i = i
			checker.topCache.j = start.j

			return true
		}
	}

	return checker.topCache != nil && iEnd <= checker.topCache.i && jEnd <= checker.topCache.j
}

// CheckLeftFirst is like CheckTopFirst but left first.
func (checker *chessChecker) CheckLeftFirst(board [][]int, start position, checkRange int) bool {
	iEnd := start.i - checkRange + 1
	jEnd := start.j - checkRange + 1

	for j := start.j; j >= jEnd; j-- {
		if hasChess(board[start.i][j]) {
			if checker.leftCache == nil {
				checker.leftCache = &position{}
			}

			checker.leftCache.i = start.i
			checker.leftCache.j = j

			return true
		}
	}

	return checker.leftCache != nil && iEnd <= checker.leftCache.i && jEnd <= checker.leftCache.j
}

func hasChess(i int) bool {
	return i == 0
}
