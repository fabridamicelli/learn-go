package poker

import (
	"encoding/json"
	"fmt"
	"os"
)

type League []Player

func NewLeague(file *os.File) ([]Player, error) {
	var l []Player
	err := json.NewDecoder(file).Decode(&l)
	if err != nil {
		return nil, fmt.Errorf("Problem parsing league: %v", err)
	}
	return l, nil

}
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
