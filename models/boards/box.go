package boards

import (
	"district-21/models/tokens"
)

/**
 * Casilla del tablero.
 * Se identifica por una posicion en la coordeanada X y una posicion en la
 * coordenada Y
 * Está conectada a otras casillas.
 */ 
 type Box struct {

	// Coordenadas X e Y de la casilla
	Coord_x, Coord_y int

	//Token que ocupa la casilla.
	token tokens.Token

	//Lista de Casillas a las se puede llegar desde la posicion de la casilla
	//toBox []Box

	//Lista de Casillas desde las que podemos llegar 
	//fromBox []Box

}

//func (box *Box) loadToBoxes() {}

/**
 * Asocia una ficha a la casilla
 *
 */
func (box *Box) SetToken(item tokens.Token) {
	box.token = item
}

func (box *Box) IsOccuped() bool {
	//Por defecto no ocupada
	result := false
	//No se puede comparar un struct con un nil. Por eso comparamos con un valor
	//vacío
	if box.token != (tokens.Token{}) {
		result = true
	}
	return result
}