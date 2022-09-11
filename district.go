package main

import (
	"fmt"
)

func main() {

	var name string
	var color string
	fmt.Println("Wellcome District 21 Game")
	fmt.Println("Introduce tu nombre: ")
	fmt.Scanln(&name)
	fmt.Println("Selecciona un color: \n\t [0] Rojo \n\t [1] Azul")
	fmt.Scanln(&color)
	player1 := Player{
		name:  name,
		color: color,
	}
	fmt.Printf("You are wellcome %s\n", player1)
}

// Representa el jugador de la partida
type Player struct {
	name  string
	color string
}

func (player *Player) GetName() string {
	return player.name
}

/**
 * Return Color
 */
func (player *Player) GetColor() string {
	return player.color
}

// Representa el Juego
type Game struct {
}

// Representa el tablero
type Board struct {
}

// Representa una ficha en el tablero, typo, position, movimiento
type Token struct {
	//Coordenadas que ocupa la ficha en el tablero
	coord_x int
	coord_y int
	//Direccion por la que se puede mover la ficha
	direction int
	//Tipo de la ficha. 0 - KING, 1 - KNIGHT, 2 - INFANTRY
	tokenType int
}

func (token *Token) GetDirection() int {
	return token.direction
}

func (token *Token) GetCoordX() int {
	return token.coord_x
}

func (token *Token) GetCoordY() int {
	return token.coord_y
}

func (token *Token) GetAsCord() string {
	return fmt.Sprintf("(%d,%d)", token.coord_x, token.coord_y)
}

func (token *Token) GetType() int {
	return token.tokenType
}
