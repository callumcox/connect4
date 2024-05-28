package main

import (
	"os"

	"github.com/callumcox/connect4/constants"
	"github.com/callumcox/connect4/game"
	"github.com/callumcox/connect4/game/board"
	"github.com/callumcox/connect4/game/player"
)

func main() {

	in := os.Stdin
	out := os.Stdout

	config := &board.BoardConfig{
		Rows:    constants.BoardRows,
		Columns: constants.BoardColumns,
	}

	board := board.NewBoard(config, out)

	player1 := player.NewPlayer("Player 1", constants.Player1Token, in, out)
	player2 := player.NewPlayer("Player 2", constants.Player2Token, in, out)

	game := game.NewGame(player1, player2, board, out)

	game.StartGame()

}
