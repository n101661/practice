function count(chessBoard) {
    let cBoard = new squareSizeCounterBoard(chessBoard.length, chessBoard[0].length)

    for (let squareSize = 2; ; squareSize++) {
        cBoard.init(chessBoard, squareSize)

        for (let i = squareSize - 1; i < cBoard.board.length; i++) {
            let checker = new chessChecker(chessBoard, squareSize, i - squareSize + 2, false)
            for (let j = squareSize - 1; j < cBoard.board[i].length; j++) {
                let num = cBoard.top(j, i, squareSize) + cBoard.left(j, i, squareSize) - cBoard.topLeft(j, i, squareSize)
                if (!checker.checkTopFirst(chessBoard, { i: i + 1, j: j + 1 }, squareSize)) {
                    num++
                }
                if (num > 0) {
                    cBoard.board[i][j][squareSize] = num
                }
            }
        }

        if (cBoard.board[cBoard.board.length - 1][cBoard.board[0].length - 1][squareSize] == 0 || squareSize == chessBoard.length) {
            break
        }
    }

    return cBoard.board[cBoard.board.length - 1][cBoard.board[0].length - 1]
}

function squareSizeCounterBoard(iLen, jLen) {
    this.board = []
    this.board.length = iLen - 1

    let cellNum = jLen - 1

    for (let i = 0; i < iLen - 1; i++) {
        this.board[i] = []
        this.board[i].length = cellNum

        for (let j = 0; j < cellNum; j++) {
            this.board[i][j] = {}
        }
    }

    this.init = function (chessBoard, squareSize) {
        let chessBoardStart = squareSize - 1
        let boardStart = squareSize - 2

        // init row
        let checker = new chessChecker(chessBoard, squareSize, 0, true)
        for (let x = chessBoardStart; x < chessBoard[chessBoardStart].length; x++) {
            if (x == chessBoardStart) {
                if (!checker.checkTopFirst(chessBoard, { i: chessBoardStart, j: chessBoardStart }, squareSize)) {
                    this.board[boardStart][boardStart][squareSize] = 1
                }
                continue
            }

            let num = this.board[boardStart][x - 2][squareSize] || 0
            if (!checker.checkTopFirst(chessBoard, { i: chessBoardStart, j: x }, squareSize)) {
                num++
            }
            if (num > 0) {
                this.board[boardStart][x - 1][squareSize] = num
            }
        }

        // init column
        for (let y = squareSize; y < chessBoard.length; y++) {
            let num = this.board[y - 2][boardStart][squareSize] || 0
            if (!checker.checkLeftFirst(chessBoard, { i: y, j: chessBoardStart }, squareSize)) {
                num++
            }
            if (num > 0) {
                this.board[y - 1][boardStart][squareSize] = num
            }
        }
    }
    this.topLeft = function (x, y, squareSize) {
        if (x == 0 || y == 0 || this.board.length <= y || this.board[0].length <= x) {
            return 0
        }

        if (this.board[y - 1][x - 1][squareSize] === undefined) {
            return 0
        }

        return this.board[y - 1][x - 1][squareSize]
    }
    this.top = function (x, y, squareSize) {
        if (y == 0 || this.board.length <= y) {
            return 0
        }

        if (this.board[y - 1][x][squareSize] === undefined) {
            return 0
        }

        return this.board[y - 1][x][squareSize]
    }
    this.left = function (x, y, squareSize) {
        if (x == 0 || this.board[0].length <= x) {
            return 0
        }

        if (this.board[y][x - 1][squareSize] === undefined) {
            return 0
        }

        return this.board[y][x - 1][squareSize]
    }
}

function chessChecker(board, checkRange, iShift, genLeft) {
    this.topCache = null;
    this.leftCache = null;

    let start = checkRange - 1

    // top first
    let quit = false
    for (let j = start; j >= 0; j--) {
        for (let i = start + iShift; i >= iShift; i--) {
            if (hasChess(board[i][j])) {
                this.topCache = { i: i, j: j }
                quit = true

                break
            }
        }

        if (quit) {
            break
        }
    }

    if (genLeft) {
        // left first
        quit = false
        for (let i = start + iShift; i >= iShift; i--) {
            for (let j = start; j >= 0; j--) {
                if (hasChess(board[i][j])) {
                    this.leftCache = { i: i, j: j }
                    quit = true

                    break
                }
            }

            if (quit) {
                break
            }
        }
    }

    this.checkTopFirst = function (board, start, checkRange) {
        let iEnd = start.i - checkRange + 1
        let jEnd = start.j - checkRange + 1

        for (let i = start.i; i >= iEnd; i--) {
            if (hasChess(board[i][start.j])) {
                if (this.topCache === null) {
                    this.topCache = {}
                }

                this.topCache.i = i
                this.topCache.j = start.j

                return true
            }
        }

        return this.topCache != null && iEnd <= this.topCache.i && jEnd <= this.topCache.j
    }
    this.checkLeftFirst = function (board, start, checkRange) {
        let iEnd = start.i - checkRange + 1
        let jEnd = start.j - checkRange + 1

        for (let j = start.j; j >= jEnd; j--) {
            if (hasChess(board[start.i][j])) {
                if (this.leftCache === null) {
                    this.leftCache = {}
                }

                this.leftCache.i = start.i
                this.leftCache.j = j

                return true
            }
        }

        return this.leftCache != null && iEnd <= this.leftCache.i && jEnd <= this.leftCache.j
    }
}

function hasChess(i) { return i === 0 }
