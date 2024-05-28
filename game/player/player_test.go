package player_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/callumcox/connect4/game/board"
	"github.com/callumcox/connect4/game/player"
)

func setup() (board.BoardInterface, *bytes.Buffer) {
	testOutputBuffer := &bytes.Buffer{}
	boardConfig := &board.BoardConfig{
		Rows:    6,
		Columns: 7,
	}

	testBoard := board.NewBoard(boardConfig, testOutputBuffer)
	return testBoard, testOutputBuffer
}

func TestGetMove(t *testing.T) {
	testBoard, testOutputBuffer := setup()

	testCases := []struct {
		name           string
		input          string
		expectedOutput string
		expected       int
	}{
		{
			name:           "Valid move returns column index",
			input:          "4\n",
			expectedOutput: "Alice, it's your turn, please enter a column: ",
			expected:       3,
		},
		{
			name:           "Initial move is invalid input",
			input:          "invalid\n2\n",
			expectedOutput: "Invalid input. Please enter a valid column number.",
			expected:       1,
		},
		{
			name:           "Initial move greater than max columns",
			input:          "8\n3\n",
			expectedOutput: "Invalid move, please try again.",
			expected:       2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testOutputBuffer.Reset()

			playerName := "Alice"

			testPlayer := player.NewPlayer(playerName, "X", strings.NewReader(tc.input), testOutputBuffer)

			move := testPlayer.GetMove(testBoard)

			outputString := testOutputBuffer.String()

			if !strings.Contains(outputString, tc.expectedOutput) {
				t.Errorf("Expected output to be %s, but got %s", tc.expectedOutput, outputString)
			}

			if move != tc.expected {
				t.Errorf("Expected move to be %d, but got %d", tc.expected, move)
			}
		})
	}
}
