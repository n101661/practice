package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChessBoard(t *testing.T) {
	assert := assert.New(t)

	{
		board := [][]int{
			{1, 1},
			{1, 1},
		}
		assert.Equal(map[int]int{
			2: 1,
		}, ChessBoard(board))
	}
	{
		board := [][]int{
			{0, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{0, 1, 1, 0, 1},
			{1, 1, 1, 1, 1},
		}
		assert.Equal(map[int]int{
			2: 9,
			3: 2,
		}, ChessBoard(board))
	}
	{
		board := [][]int{
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
		}
		assert.Equal(map[int]int{
			2: 16,
			3: 9,
			4: 4,
			5: 1,
		}, ChessBoard(board))
	}
	{
		board := [][]int{
			{1, 1, 1, 1, 1, 1, 1},
			{1, 0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 0, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1, 0},
		}
		assert.Equal(map[int]int{
			2: 26,
			3: 11,
			4: 4,
		}, ChessBoard(board))
	}
	{
		board := [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 0, 1, 1},
		}
		assert.Equal(map[int]int{
			2: 1,
		}, ChessBoard(board))
	}
	{
		board := [][]int{
			{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
			{1, 1, 1, 0, 1, 1, 1, 1, 1, 1},
			{0, 0, 1, 0, 1, 1, 1, 1, 1, 1},
			{0, 1, 1, 0, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 1, 1, 1, 0, 1},
			{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0, 1, 1, 1, 1},
			{0, 1, 1, 0, 1, 1, 1, 1, 1, 1},
			{1, 1, 0, 1, 0, 1, 1, 1, 1, 1},
		}
		assert.Equal(map[int]int{
			2: 38,
			3: 12,
			4: 2,
		}, ChessBoard(board))
	}
}

func Test_squareSizeCounterBoard_init(t *testing.T) {
	assert := assert.New(t)

	{
		input := [][]int{
			{1, 1},
			{1, 1},
		}

		board := newSquareSizeCounterBoard(len(input), len(input[0]))
		board.init(input, 2)

		assert.Equal(squareSizeCounterBoard{
			{
				{2: 1},
			},
		}, board)
	}
	{
		input := [][]int{
			{1, 1},
			{0, 1},
		}

		board := newSquareSizeCounterBoard(len(input), len(input[0]))
		board.init(input, 2)

		assert.Equal(squareSizeCounterBoard{
			{{}},
		}, board)
	}
	{
		input := [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{0, 1, 1},
		}

		board := newSquareSizeCounterBoard(len(input), len(input[0]))
		board.init(input, 2)

		assert.Equal(squareSizeCounterBoard{
			{{2: 1}, {2: 2}},
			{{2: 1}, {}},
		}, board)
	}
	{
		input := [][]int{
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 0, 1},
			{0, 1, 1, 1},
		}

		board := newSquareSizeCounterBoard(len(input), len(input[0]))
		board.init(input, 2)

		assert.Equal(squareSizeCounterBoard{
			{{}, {2: 1}, {2: 2}},
			{{2: 1}, {}, {}},
			{{2: 1}, {}, {}},
		}, board)
	}
	{
		input := [][]int{
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
		}

		board := newSquareSizeCounterBoard(len(input), len(input[0]))
		board.init(input, 2)

		assert.Equal(squareSizeCounterBoard{
			{{2: 1}, {2: 2}, {2: 3}, {2: 4}, {2: 5}},
			{{2: 1}, {}, {}, {}, {}},
			{{2: 1}, {}, {}, {}, {}},
			{{2: 1}, {}, {}, {}, {}},
			{{2: 1}, {}, {}, {}, {}},
		}, board)
	}
	{
		input := [][]int{
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
		}

		board := newSquareSizeCounterBoard(len(input), len(input[0]))
		board.init(input, 3)

		assert.Equal(squareSizeCounterBoard{
			{{}, {}, {}, {}, {}},
			{{}, {}, {3: 1}, {3: 2}, {3: 3}},
			{{}, {}, {}, {}, {}},
			{{}, {3: 1}, {}, {}, {}},
			{{}, {3: 2}, {}, {}, {}},
		}, board)
	}
}

func Test_squareSizeCounterBoard(t *testing.T) {
	assert := assert.New(t)

	board := squareSizeCounterBoard{
		{{}, {2: 1}, {2: 1}},
		{{2: 1}, {}, {}},
		{{}, {}, {}},
	}

	assert.Equal(0, board.TopLeft(0, 0, 2))
	assert.Equal(0, board.TopLeft(1, 1, 2))
	assert.Equal(1, board.TopLeft(1, 2, 2))

	assert.Equal(0, board.Top(0, 0, 2))
	assert.Equal(1, board.Top(0, 2, 2))

	assert.Equal(0, board.Left(0, 0, 2))
	assert.Equal(1, board.Left(2, 0, 2))
}

func Test_chessChecker_CheckTopFirst(t *testing.T) {
	assert := assert.New(t)

	{
		board := [][]int{
			{1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
		}

		checker := newChessChecker(board, 3, 0, false)
		assert.Equal(&position{i: 0, j: 1}, checker.topCache)

		assert.True(checker.CheckTopFirst(board, position{i: 2, j: 2}, 3))
		assert.Equal(&position{i: 0, j: 1}, checker.topCache)

		assert.True(checker.CheckTopFirst(board, position{i: 2, j: 3}, 3))
		assert.Equal(&position{i: 0, j: 1}, checker.topCache)

		assert.False(checker.CheckTopFirst(board, position{i: 2, j: 4}, 3))
		assert.Equal(&position{i: 0, j: 1}, checker.topCache)
	}
	{
		board := [][]int{
			{0, 0, 1, 1, 1},
			{0, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
		}

		checker := newChessChecker(board, 2, 0, false)
		assert.Equal(&position{i: 0, j: 1}, checker.topCache)

		assert.True(checker.CheckTopFirst(board, position{i: 1, j: 1}, 2))
		assert.Equal(&position{i: 0, j: 1}, checker.topCache)

		assert.True(checker.CheckTopFirst(board, position{i: 1, j: 2}, 2))
		assert.Equal(&position{i: 0, j: 1}, checker.topCache)

		assert.False(checker.CheckTopFirst(board, position{i: 1, j: 3}, 2))
		assert.Equal(&position{i: 0, j: 1}, checker.topCache)
	}
}

func Test_chessChecker_CheckLeftFirst(t *testing.T) {
	assert := assert.New(t)

	{
		board := [][]int{
			{1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1},
		}

		checker := newChessChecker(board, 3, 0, true)
		assert.Equal(&position{i: 0, j: 1}, checker.leftCache)

		assert.True(checker.CheckLeftFirst(board, position{i: 2, j: 2}, 3))
		assert.Equal(&position{i: 0, j: 1}, checker.leftCache)

		assert.False(checker.CheckLeftFirst(board, position{i: 3, j: 2}, 3))
		assert.Equal(&position{i: 0, j: 1}, checker.leftCache)

		assert.True(checker.CheckLeftFirst(board, position{i: 4, j: 2}, 3))
		assert.Equal(&position{i: 4, j: 0}, checker.leftCache)
	}
	{
		board := [][]int{
			{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
			{1, 1, 1, 0, 1, 1, 1, 1, 1, 1},
			{0, 0, 1, 0, 1, 1, 1, 1, 1, 1},
			{0, 1, 1, 0, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 1, 1, 1, 0, 1},
			{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0, 1, 1, 1, 1},
			{0, 1, 1, 0, 1, 1, 1, 1, 1, 1},
			{1, 1, 0, 1, 0, 1, 1, 1, 1, 1},
		}

		checker := newChessChecker(board, 2, 0, true)
		assert.Nil(checker.leftCache)

		assert.False(checker.CheckLeftFirst(board, position{i: 1, j: 1}, 2))
		assert.Nil(checker.leftCache)

		assert.True(checker.CheckLeftFirst(board, position{i: 2, j: 1}, 2))
		assert.Equal(&position{i: 2, j: 1}, checker.leftCache)

		assert.True(checker.CheckLeftFirst(board, position{i: 3, j: 1}, 2))
		assert.Equal(&position{i: 3, j: 0}, checker.leftCache)

		assert.True(checker.CheckLeftFirst(board, position{i: 4, j: 1}, 2))
		assert.Equal(&position{i: 3, j: 0}, checker.leftCache)

		assert.False(checker.CheckLeftFirst(board, position{i: 5, j: 1}, 2))
		assert.Equal(&position{i: 3, j: 0}, checker.leftCache)

		assert.False(checker.CheckLeftFirst(board, position{i: 6, j: 1}, 2))
		assert.Equal(&position{i: 3, j: 0}, checker.leftCache)

		assert.True(checker.CheckLeftFirst(board, position{i: 7, j: 1}, 2))
		assert.Equal(&position{i: 7, j: 0}, checker.leftCache)

		assert.True(checker.CheckLeftFirst(board, position{i: 8, j: 1}, 2))
		assert.Equal(&position{i: 8, j: 0}, checker.leftCache)

		assert.True(checker.CheckLeftFirst(board, position{i: 9, j: 1}, 2))
		assert.Equal(&position{i: 8, j: 0}, checker.leftCache)
	}
}
