package tokens

import (
	"github.com/google/uuid"
)

const POLICE int = 3
const DETECTIVE int = 2
const CHIEF int = 1

//Representa una ficha en el tablero
type Token struct {

	TypeToken int
	UUID uuid.UUID

	/**
	 * Moviemiento de una ficha en el tablero
	 */
	//movement(direction int, space int)
	
	/**
	 * TODO: Pendiente de revisar. El juego solo va a saber de coodenadas.
	 *
	 * Devuelve la posicion de una ficha en el tablero
	 * Coordenada X
	 * Coordenada Y
	 */
	//GetPosition() (int, int)

}

func (token *Token) GetTypeToken() int {
	return token.TypeToken	
}