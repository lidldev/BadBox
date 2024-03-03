package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name   string
	Health int
	Attack int
}

type Enemy struct {
	Name   string
	Health int
	Attack int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	player := Player{Name: "Kahraman", Health: 100, Attack: 20}

	fmt.Println("Hoş geldiniz,", player.Name, "!")
	fmt.Println("Bir düşman belirdi! Onu yenmeye çalışalım.")

	enemy := generateRandomEnemy()

	fmt.Println("Bir", enemy.Name, "ortaya çıktı!")

	for player.Health > 0 && enemy.Health > 0 {
		fmt.Println("Saldırı [s] veya Kaç [k] seçeneğini seçin:")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "s":
			playerAttack := rand.Intn(player.Attack)
			enemy.Health -= playerAttack
			fmt.Println(player.Name, "düşmana", playerAttack, "zarar verdi!")

			enemyAttack := rand.Intn(enemy.Attack)
			player.Health -= enemyAttack
			fmt.Println(enemy.Name, "sana", enemyAttack, "zarar verdi!")

			case "k":
				fmt.Println(player.Name, "kaçmaya çalışıyor...")
				break

				default:
					fmt.Println("Geçersiz seçenek! Saldırı [s] veya Kaç [k] seçeneğini seçin.")
		}

		if player.Health <= 0 {
			fmt.Println(player.Name, "öldü. Oyun bitti!")
		} else if enemy.Health <= 0 {
			fmt.Println(player.Name, "düşmanı yendi!")
		} else {
			fmt.Println(player.Name, "Sağlık:", player.Health, "Düşman Sağlık:", enemy.Health)
		}
	}
}

func generateRandomEnemy() Enemy {
	enemies := []Enemy{
		{Name: "Ork", Health: 50, Attack: 10},
		{Name: "Goblin", Health: 30, Attack: 5},
		{Name: "Troll", Health: 70, Attack: 15},
		}

		randomIndex := rand.Intn(len(enemies))
		return enemies[randomIndex]
}
