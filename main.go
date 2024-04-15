package main

import (
	"fmt"
    "math/rand/v2"
	// "text/template"
	// "net"
	"time"
)

const boardSize int = 30

var board [boardSize][boardSize]int
var tempBoard [boardSize][boardSize]int
var round = 0

func printTile(tileValue *int) {
    if *tileValue == 0 {
	    fmt.Printf("\033[48;5;255m  \033[0m")
    } else if *tileValue >= 244 {
        *tileValue = 1
	    fmt.Printf("\033[48;5;%dm  \033[0m", *tileValue)
    } else {
	    fmt.Printf("\033[48;5;%dm  \033[0m", *tileValue)
    }
}

func initPrintBoard() {
    fmt.Print("Board: Round 0\n")

    for row := 0; row < boardSize; row++ {
        for col := 0; col < boardSize; col++ {
            printTile(&board[col][row])
        }
        fmt.Print("\n")
    }
}

func printBoard() {
    fmt.Printf("\033[%dF2K\r", boardSize + 1)
    fmt.Printf("Board: Round %d\n", round)
    for row := 0; row < boardSize; row++ {
        for col := 0; col < boardSize; col++ {
            printTile(&board[col][row])
        }
        fmt.Print("\n")
    }
}

func updateBoardSpace(row int, col int) {
    liveAdjacentSpots := 0

    for rowOffset := -1; rowOffset < 2; rowOffset++ {
        currRowOffset := rowOffset + row
        if currRowOffset >= boardSize || currRowOffset < 0 {
           continue  
        }

        for colOffset := -1; colOffset < 2; colOffset++ {
            currColOffset := colOffset + col
            if currColOffset >= boardSize || currColOffset < 0 {
               continue  
            }
            
            if(board[currRowOffset][currColOffset] != 0 && 
                !(row == currRowOffset && col == currColOffset)) {
                liveAdjacentSpots++
            }
        }
    }

    if board[row][col] != 0 {
        if(liveAdjacentSpots == 2 || liveAdjacentSpots == 3) {
            tempBoard[row][col] = board[row][col] + 1
        } 
    } else {
        if liveAdjacentSpots == 3 {
            tempBoard[row][col] = 1
        }
    }
}

func updateBoard() {
    for row := 0; row < boardSize; row++ {
        for col := 0; col < boardSize; col++ {
            tempBoard[row][col] = 0
            updateBoardSpace(row, col)
        }
    }

    board = tempBoard
    round += 1
}

func checkBoard() bool {
    for row := 0; row < boardSize; row++ {
        for col := 0; col < boardSize; col++ {
            if board[row][col] != 0 {
                return false;
            }
        }
    }

    return true;
}

func gameMainLoop() {
    initPrintBoard()
    for { 
        updateBoard()
        time.Sleep(1000 * time.Millisecond)
        printBoard() 
        if checkBoard() {
            break
        }
    }
}

func initBoard() {
    for row := 0; row < boardSize; row++ {
        for col := 0; col < boardSize; col++ {
           if rand.IntN(100) > 70 {
                board[row][col] = 1
           } 
        }
    }
}

func main() {
    initBoard()
    gameMainLoop()
    
    for {
    }
}
