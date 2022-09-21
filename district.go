package main

import (
	"fmt"
	"district-21/models"
	"district-21/models/boards"
	"district-21/models/tokens"
	"github.com/google/uuid"
)

var colors = make(map[string]string)

func main() {

	colors["0"]= "red"
	colors["1"]= "blue"

	var name string
	var color string
	isRunning := true
	
	
	// Format Wellcome page
	fmt.Println("Wellcome District 21 Game")
	fmt.Print("Player one: ")
	fmt.Scanln(&name)
	fmt.Println("Reinos disponibles: %v", colors)
	fmt.Println("Selecciona un color: \n\t [0] Rojo \n\t [1] Azul")
	fmt.Scanln(&color)
	playerOne := initPlayerOne(name, color)
	delete(colors, playerOne.Realm)
	fmt.Println("Reinos disponibles: %v\n", colors)
	for key, color := range colors {
		fmt.Println("Key:", key)
		fmt.Println("Realm:", color)
	}
	initPlayerTwo()
	initGame()
	for isRunning==true {
				
	}
}

/**
 * Inicializa el juego Distrito 21
 * 	- Inicializa el tablero
 * 	- Inicializa las fichas
 */
func initGame()  {
	initBoardGame()
}

/**
 * Inicializa el tablero
 */
func initBoardGame()  {
	fmt.Println("--Init boardgame")
	box00 := boards.Box { Coord_x: 0, Coord_y: 0,}
	fmt.Println("Box 0", box00)
	//Init Tokens...................................................
	}

/**
 * Inicializa las fichas en el tablero
 */
 func initPlayerOne(name string, realm string) models.Player  {
	fmt.Println("-- Init player one.......................")
	playerOne := models.Player{
		Realm: colors[realm],
	}
	playerOne.Name = name
	playerOne.SetUUID()
	fmt.Printf("You are wellcome \n\t[Nick]: %s \n\t[ID]: %s\n\t[Realm]: %s\n}\n", playerOne.Name, playerOne.UUID, playerOne.Realm)
	//Init tokens
	chief_1 := tokens.Token { TypeToken: 1, UUID: uuid.New()}
	playerOne.SetChief(chief_1)
	//Init Detectives
	var detectives []tokens.Token
	for i := 0; i < 7; i++ {
		detectives = append(detectives, tokens.Token { TypeToken: 2, UUID: uuid.New()})
	}
	
	playerOne.SetDetectives(detectives)

	//Init Policies
	var policies []tokens.Token 
	for i := 0; i < 8; i++ {
		policies = append(policies, tokens.Token { TypeToken: 3, UUID: uuid.New()})
	}
	playerOne.SetPolicies(policies)
	fmt.Println("player One", playerOne)
	return playerOne
}

/**
 * Inicializa las fichas en el tablero
 */
 func initPlayerTwo()  {
	fmt.Println("--Init player two.......................")
}

