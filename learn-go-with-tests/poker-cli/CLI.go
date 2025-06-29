package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}

}

func (cli *CLI) Playpoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))

}
func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(s string) string {
	return strings.Replace(s, " wins", "", 1)

}
