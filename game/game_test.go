package game_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/callumcox/connect4/game"
	"github.com/callumcox/connect4/game/player"
)

type mockBoard struct {
	WinningMoveCol int
	IsDraw         bool
}

func (b *mockBoard) Print()                                    {}
func (b *mockBoard) PlaceCounter(column int, token string) int { return 0 }
func (b *mockBoard) IsValidMove(column int) bool               { return true }
func (b *mockBoard) IsWinningMove(row, column int) bool        { return column == b.WinningMoveCol }
func (b *mockBoard) IsBoardFull() bool                         { return b.IsDraw }

func setup(winningMove int, isDraw bool) *mockBoard {

	brd := &mockBoard{
		WinningMoveCol: winningMove,
		IsDraw:         isDraw,
	}

	return brd
}

func TestStartGame(t *testing.T) {
	testOutput := &bytes.Buffer{}

	player1Name := "Alice"
	player1Move := 3

	player2Name := "Bob"
	player2Move := 2

	inputStr := "%d\n"

	gameOverStr := "Game Over!"
	winStr := "Congratulations %s, you have won!"
	drawStr := "It's a draw!"

	testCases := []struct {
		name           string
		winningMoveCol int
		isDraw         bool
		expectedOutput []string
	}{
		{
			name:           "Player 1 wins",
			winningMoveCol: player1Move,
			expectedOutput: []string{fmt.Sprintf(winStr, player1Name), gameOverStr},
			isDraw:         false,
		},
		{
			name:           "Player 2 wins",
			winningMoveCol: player2Move,
			expectedOutput: []string{fmt.Sprintf(winStr, player2Name), gameOverStr},
			isDraw:         false,
		},
		{
			name:           "Draw",
			winningMoveCol: 0,
			expectedOutput: []string{drawStr, gameOverStr},
			isDraw:         true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testOutput.Reset()

			mockBoard := setup(tc.winningMoveCol, tc.isDraw)

			player1 := player.NewPlayer(player1Name, "X", strings.NewReader(fmt.Sprintf(inputStr, player1Move+1)), testOutput)
			player2 := player.NewPlayer(player2Name, "O", strings.NewReader(fmt.Sprintf(inputStr, player2Move+1)), testOutput)

			testGame := game.NewGame(player1, player2, mockBoard, testOutput)

			testGame.StartGame()

			out := testOutput.String()

			for _, expectedOutput := range tc.expectedOutput {
				if !strings.Contains(out, expectedOutput) {
					t.Errorf("Expected output %s to contain %s", out, expectedOutput)
				}
			}
		})
	}
}
