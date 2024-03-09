package main

import (
	"fmt"
)

func main() {

	gameState, err := LoadGame()

	if gameState.Player.Level >= 50 {
		enemyTypes = append(tier1Enemies, tier2Enemies...)
	}
	if err != nil {

		var day_counter = 1

		fmt.Println("Выберите класс:")
		fmt.Println("1. Воин")
		fmt.Println("2. Маг")
		fmt.Println("3. Вор")
		fmt.Println("4. Бард")
		var classInput int
		fmt.Scanln(&classInput)

		playerClass := PlayerClass(classInput - 1)
		switch playerClass {
		case Воин:
			gameState.Player.Power = 9
			gameState.Player.Defense = 5
			gameState.Player.Gold = 0
		switch 
		}
	}
}
