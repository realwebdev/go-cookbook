package main

import "fmt"

type Player struct {
	HP int
}

func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage. New HP ->", player.HP)
}

// 8 byte long int point to memmory address/ a space a slot in memory
func main() {
	player := &Player{
		HP: 100,
	}
	TakeDamage(player, 10)
	fmt.Println(player)
}
