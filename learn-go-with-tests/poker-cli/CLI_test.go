package poker_test

import (
	"strings"
	"testing"

	"poker-cli"
)

func TestCLI(t *testing.T) {
	t.Run("record chris from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.Playpoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.Playpoker()
		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

}
