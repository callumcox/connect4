package board

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/callumcox/connect4/constants"
)

var testBoardConfig = &BoardConfig{
	Rows:    6,
	Columns: 7,
}

func setup() (*board, *bytes.Buffer) {
	testOutputBuffer := &bytes.Buffer{}
	board := NewBoard(testBoardConfig, testOutputBuffer)
	return board, testOutputBuffer
}

var winTestCases = []struct {
	name     string
	moves    []string
	expected bool
}{
	{name: "Player Win", moves: []string{"X", "X", "X", "X"}, expected: true},
	{name: "Blocked by Player", moves: []string{"X", "O", "X", "X"}, expected: false},
	{name: "Not 4 in a row", moves: []string{"X", "X", "X"}, expected: false},
}

func TestNewBoard(t *testing.T) {
	board, _ := setup()

	if len(board.cells) != testBoardConfig.Rows {
		t.Errorf("Expected %d rows, but got %d", testBoardConfig.Rows, len(board.cells))
	}

	if len(board.cells[0]) != testBoardConfig.Columns {
		t.Errorf("Expected %d columns, but got %d", testBoardConfig.Columns, len(board.cells[0]))
	}

	if board.moves != 0 {
		t.Errorf("Expected moves to be 0, but got %d", board.moves)
	}

	for r := range board.cells {
		for c := range board.cells[r] {
			if board.cells[r][c] != constants.EmptyCell {
				t.Errorf("Expected cell at row %d, column %d to be empty, but got %s", r, c, board.cells[r][c])
			}
		}
	}

}

func TestIsValidMove(t *testing.T) {
	board, _ := setup()

	for r := range board.cells {
		board.cells[r][0] = "X"
	}

	testCases := []struct {
		name     string
		column   int
		expected bool
	}{
		{name: "Valid row column", column: 1, expected: true},
		{name: "Invalid column index, less than min index", column: -1, expected: false},
		{name: "Invalid column index, greater than max index", column: 7, expected: false},
		{name: "Invalid column index, column is full", column: 0, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if board.IsValidMove(tc.column) != tc.expected {
				t.Errorf("Expected move to be %t, but got %t", tc.expected, !tc.expected)
			}
		})
	}
}

func TestPlaceCounter(t *testing.T) {
	writer := &bytes.Buffer{}

	board := NewBoard(testBoardConfig, writer)

	moves := []string{"X", "O"}

	for idx, move := range moves {
		row := board.PlaceCounter(0, move)
		if row != testBoardConfig.Rows-1-idx {
			t.Errorf("Expected counter to be placed in row %d, but got row %d", testBoardConfig.Rows-1, row)
		}

		if board.cells[row][0] != move {
			t.Errorf("Expected counter at row %d, column 0 to be %s, but got %s", row, move, board.cells[row][0])
		}

		if board.moves != idx+1 {
			t.Errorf("Expected moves to be %d, but got %d", idx+1, board.moves)
		}
	}
}

func TestIsHorizontalWinningMove(t *testing.T) {

	for _, tc := range winTestCases {
		t.Run(tc.name, func(t *testing.T) {
			board, _ := setup()

			copy(board.cells[0], tc.moves)

			// Check all columns in the row to validate the winning horizontal move at each position
			for c := range tc.moves {
				if board.IsWinningMove(0, c) != tc.expected {
					t.Errorf("Expected IsWinningMove with row 0, column %d to be %t, but got %t", c, tc.expected, !tc.expected)
				}
			}
		})
	}
}

func TestIsVerticalWinningMove(t *testing.T) {
	for _, tc := range winTestCases {
		t.Run(tc.name, func(t *testing.T) {

			board, _ := setup()

			for r, m := range tc.moves {
				board.cells[r][0] = m
			}

			// Check the top-most cell in the column to validate the winning vertical move
			if board.IsWinningMove(0, 0) != tc.expected {
				t.Errorf("Expected IsWinningMove with row %d, column 0 to be %t, but got %t", len(winTestCases), tc.expected, !tc.expected)
			}

		})
	}
}

func TestIsDiagonalLeftWinningMove(t *testing.T) {
	for _, tc := range winTestCases {
		t.Run(tc.name, func(t *testing.T) {

			board, _ := setup()

			for r, m := range tc.moves {
				board.cells[r][3-r] = m
			}

			// Check all cells in the diagonal to validate the winning diagonal move at each position
			for r := range tc.moves {
				if board.IsWinningMove(r, 3-r) != tc.expected {
					t.Errorf("Expected IsWinningMove with row %d, column %d to be %t, but got %t", r, 3-r, tc.expected, !tc.expected)
				}
			}
		})
	}
}

func TestIsDiagonalRightWinningMove(t *testing.T) {
	for _, tc := range winTestCases {
		t.Run(tc.name, func(t *testing.T) {

			board, _ := setup()

			for r, move := range tc.moves {
				board.cells[r][r] = move
			}

			// Check all cells in the diagonal to validate the winning diagonal move at each position
			for r := range tc.moves {
				if board.IsWinningMove(r, r) != tc.expected {
					t.Errorf("Expected IsWinningMove with row %d, column %d to be %t, but got %t", r, r, tc.expected, !tc.expected)
				}
			}
		})
	}
}

func TestIsBoardFull(t *testing.T) {
	board, _ := setup()

	// Empty board
	if board.IsBoardFull() {
		t.Errorf("Expected IsBoardFull to be false, but got true")
	}

	board.moves = testBoardConfig.Rows * testBoardConfig.Columns

	if !board.IsBoardFull() {
		t.Errorf("Expected IsBoardFull to be true, but got false")
	}
}

func TestPrint(t *testing.T) {
	board, testBuf := setup()

	moves := []string{"X", "O", "X", "O", "X", "O", "X"}

	emptyRowStr := "\n| _  _  _  _  _  _  _ |"
	colNumStr := "\n  1  2  3  4  5  6  7  \n"
	fullRowStr := fmt.Sprintf("\n| %s |", strings.Join(moves, "  "))

	// Print empty board

	testCases := []struct {
		name     string
		moves    []string
		expected string
	}{
		{name: "Prints the empty board", moves: nil, expected: strings.Repeat(emptyRowStr, testBoardConfig.Rows) + colNumStr},
		{name: "Prints the board with tokens placed", moves: moves, expected: fullRowStr + strings.Repeat(emptyRowStr, testBoardConfig.Rows-1) + colNumStr},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testBuf.Reset()

			if tc.moves != nil {
				board.cells[0] = tc.moves
			}

			board.Print()

			out := testBuf.String()

			if out != tc.expected {
				t.Errorf("Expected output to be %s, but got %s", tc.expected, out)
			}
		})
	}
}
