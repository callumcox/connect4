package game

import (
	"fmt"
	"io"

	"github.com/callumcox/connect4/game/board"
	"github.com/callumcox/connect4/game/player"
)

type game struct {
	board         board.BoardInterface
	player1       *player.Player
	player2       *player.Player
	currentPlayer *player.Player
	isFinished    bool
	outputWriter  io.Writer
}

func NewGame(player1, player2 *player.Player, board board.BoardInterface, outputWriter io.Writer) *game {
	game := game{
		board:         board,
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
		isFinished:    false,
		outputWriter:  outputWriter,
	}

	return &game
}

func (game *game) StartGame() {
	game.board.Print()

	for !game.isFinished {
		columnIndex := game.currentPlayer.GetMove(game.board)

		rowIndex := game.board.PlaceCounter(columnIndex, game.currentPlayer.Token)

		game.board.Print()

		if game.checkIsWinningMove(rowIndex, columnIndex) {
			game.writeMessage(fmt.Sprintf("Congratulations %s, you have won!\n", game.currentPlayer.Name))
			game.isFinished = true

		} else if game.checkIsDraw() {
			game.writeMessage("It's a draw!\n")
			game.isFinished = true

		} else if game.currentPlayer == game.player1 {
			game.currentPlayer = game.player2

		} else {
			game.currentPlayer = game.player1

		}
	}

	game.writeMessage("Game Over!\n")
}

func (game *game) checkIsDraw() bool {
	return game.board.IsBoardFull()
}

func (game *game) checkIsWinningMove(rowIndex, columnIndex int) bool {
	return game.board.IsWinningMove(rowIndex, columnIndex)
}

func (game *game) writeMessage(message string) {
	game.outputWriter.Write([]byte([]byte(message)))
}
