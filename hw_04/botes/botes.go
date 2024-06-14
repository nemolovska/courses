package botes

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Bote struct {
	ID   int
	Name string
}

func (b Bote) Print() {
	fmt.Printf("ID: %03d has %s\n", b.ID, b.Name)
}

func CreateBotes(num int) []Bote {

	var botepake []Bote

	rand.Seed(time.Now().UnixNano())

	existingIDs := make(map[int]bool)

	for i := 1; i <= 10; i++ {
		var randomID int
		for {
			randomID = rand.Intn(100)
			if !existingIDs[randomID] {
				existingIDs[randomID] = true
				break
			}
		}
		bot := Bote{
			ID:   randomID,
			Name: fmt.Sprintf("Bot%d", i),
		}
		botepake = append(botepake, bot)
	}

	sort.Slice(botepake, func(i, j int) bool {
		return botepake[i].ID < botepake[j].ID
	})

	return botepake
}
