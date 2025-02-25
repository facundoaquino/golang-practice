package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// juego de concurrencia, hay que arreglar el tema de los logs que se muestran todos juntos parece, aunque la logica no parece estar mal

type Person struct {
	Name   string
	Age    int
	Health int
	Weapon Weapon
}

type Weapon struct {
	Damage int
}

func (weapon Weapon) Shot() int {
	//returns damage
	return rand.Intn(10) + 1
}

type WeaponActions interface {
	Shot() int
}

func main() {
	player1 := Person{
		Name:   "Player 1",
		Age:    30,
		Health: 100,
		Weapon: Weapon{Damage: 10},
	}
	player2 := Person{
		Name:   "Player2",
		Age:    31,
		Health: 100,
		Weapon: Weapon{Damage: 10},
	}
	chann := make(chan *Person, 2)

	var wg sync.WaitGroup
	for {
		wg.Add(2)
		go playStart(player1, &player2, chann, &wg)
		go playStart(player2, &player1, chann, &wg)

		player := <-chann
		if player.Health <= 0 {
			fmt.Printf("Player %s is dead\n", player.Name)
			break
		}
		wg.Wait()
	}
}

func playStart(player Person, player2 *Person, chann chan<- *Person, wg *sync.WaitGroup) {
	wg.Done()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chanceToShot := r.Intn(10000)
	if chanceToShot == 5 {
		fmt.Printf("Player %s shot to player %s Health: %d \n", player.Name, player2.Name, player.Health)
		player2.Health -= player2.Weapon.Shot()
	}
	chann <- player2

}
