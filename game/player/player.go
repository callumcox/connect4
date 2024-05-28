package player

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/callumcox/connect4/game/board"
)

type Player struct {
	Name    string
	Token   string
	writer  io.Writer
	scanner *bufio.Scanner
}

func NewPlayer(name, token string, reader io.Reader, writer io.Writer) *Player {
	scanner := bufio.NewScanner(reader)
	player := Player{Name: name, Token: token, writer: writer, scanner: scanner}
	return &player
}

func (player *Player) GetMove(board board.BoardInterface) int {

	for {
		player.writeMessage(fmt.Sprintf("%s, it's your turn, please enter a column: ", player.Name))

		if !player.scanner.Scan() {
			player.writeMessage("Failed to read input. Please try again.\n")
			continue
		}

		input := player.scanner.Text()
		column, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			player.writeMessage("Invalid input. Please enter a valid column number.\n")
			continue
		}

		columnIndex := column - 1

		if board.IsValidMove(columnIndex) {
			return columnIndex
		}

		player.writeMessage("Invalid move, please try again.\n")
	}
}

func (player *Player) writeMessage(message string) {
	player.writer.Write([]byte(message))
}
