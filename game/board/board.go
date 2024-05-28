package board

import (
	"fmt"
	"io"
	"strings"

	"github.com/callumcox/connect4/constants"
)

type BoardConfig struct {
	Rows    int
	Columns int
}

type BoardInterface interface {
	IsValidMove(column int) bool
	PlaceCounter(column int, token string) int
	IsWinningMove(row, column int) bool
	IsBoardFull() bool
	Print()
}

type board struct {
	BoardConfig
	cells        [][]string
	outputWriter io.Writer
	moves        int
}

type directionVector struct {
	x, y int
}

// Direction vectors used to check for winning moves
var (
	Horizontal    = []directionVector{{0, -1}, {0, 1}}
	Vertical      = []directionVector{{1, 0}}
	DiagonalLeft  = []directionVector{{-1, -1}, {1, 1}}
	DiagonalRight = []directionVector{{-1, 1}, {1, -1}}
)

func NewBoard(config *BoardConfig, outputWriter io.Writer) *board {

	// Create a 2D slice of cells to represent the board and initialise all cells to be empty
	cells := make([][]string, config.Rows)

	for r := range cells {
		cells[r] = make([]string, config.Columns)
		for c := range cells[r] {
			cells[r][c] = constants.EmptyCell
		}
	}

	board := board{cells: cells, outputWriter: outputWriter, moves: 0, BoardConfig: *config}

	return &board
}

func (board *board) IsValidMove(column int) bool {
	if column < 0 || column >= board.Columns {
		return false
	}

	// We only need to check the top cell of the column to see if it's empty as that will be a valid move.
	return board.cells[0][column] == constants.EmptyCell
}

func (board *board) PlaceCounter(column int, token string) int {
	// Iterate backwards through the rows to find the first empty cell in the column
	for r := board.Rows - 1; r >= 0; r-- {
		if board.cells[r][column] == constants.EmptyCell {
			board.cells[r][column] = token

			// Increment the number of moves to prevent having to iterate through the entire row/board to check if it's full
			board.moves++
			// Return the row that the counter ended up in
			return r
		}
	}

	return -1
}

func (board *board) IsWinningMove(row, column int) bool {
	return board.checkHorizontal(row, column) || board.checkVertical(row, column) || board.checkDiagonalLeft(row, column) || board.checkDiagonalRight(row, column)
}

func (board *board) IsBoardFull() bool {
	// If the number of moves is equal to the number of cells on the board, then the board is full
	return board.moves == board.Rows*board.Columns
}

func (board *board) Print() {
	var builder strings.Builder

	builder.WriteString("\n")

	for r := range board.cells {
		builder.WriteString("|")
		for column := range board.cells[r] {
			builder.WriteString(fmt.Sprintf(" %s ", board.cells[r][column]))
		}
		builder.WriteString("|\n")
	}

	// Print column numbers
	builder.WriteString(" ")

	for i := 1; i <= board.Columns; i++ {
		builder.WriteString(fmt.Sprintf(" %d ", i))
	}

	builder.WriteString(" \n")

	board.outputWriter.Write([]byte(builder.String()))
}

// Convenience functions to check specific in specific directions
func (board *board) checkHorizontal(row, column int) bool {
	return board.checkDirection(row, column, Horizontal)
}

func (board *board) checkVertical(row, column int) bool {
	return board.checkDirection(row, column, Vertical)
}

func (board *board) checkDiagonalLeft(row, column int) bool {
	return board.checkDirection(row, column, DiagonalLeft)
}

func (board *board) checkDiagonalRight(row, column int) bool {
	return board.checkDirection(row, column, DiagonalRight)
}

// checkDirection checks for the player token sequence in given directions
func (board *board) checkDirection(row, column int, directionsToCheck []directionVector) bool {
	token := board.cells[row][column]
	count := 1

	// Check in both directions
	for _, dir := range directionsToCheck {
		// Starting from the current cell, modify the current array index by the direction vector until we either:
		// a) Reach the edge of the board
		// b) Encounter a cell that doesn't contain the token we're looking for
		for r, c := row+dir.x, column+dir.y; r >= 0 && r < board.Rows && c >= 0 && c < board.Columns; r, c = r+dir.x, c+dir.y {
			if board.cells[r][c] == token {
				count++
			} else {
				break
			}
		}
	}

	return count >= 4
}
