package models

import (
	"district-21/models/tokens"
	"github.com/google/uuid"
)
type Player struct {
	
	//Hacemos uso de la herencia Player -> Person
	Person
	//UUID que identifica el jugador
	UUID uuid.UUID
	//El jugador selecciona el reino que lidera.
	Realm string
	//Chief
	Chief tokens.Token
	//Lista de detectives
	Detectives []tokens.Token
	//Lista de policias
	Policies []tokens.Token
}

func (player *Player) IsLive() bool {
	result := false
	if player.Chief != (tokens.Token{}) {
		result = true
	}
	return result
}
func (player *Player) GetName() string {
	return player.Name
}

/**
 * Return Color
 */
func (player *Player) GetRealm() string {
	return player.Realm
}

/**
 * Crea un uuid a la instancia del Jugador que nos servirá para definir el
 * objeto en los eventos
 */
func (player *Player) SetUUID() {
	player.UUID = uuid.New()
}

/**
 * Return Chief
 
 func (player *Player) GetChief() tokens.Chief {
	return player.chief
}
*/
/**
 * Set Chief
 */
 func (player *Player) SetChief(token tokens.Token) {
	player.Chief = token
 }

 func (player *Player) SetDetectives(tokens []tokens.Token){
	player.Detectives = tokens
 }

 func (player *Player) SetPolicies(tokens []tokens.Token){
	player.Policies = tokens
 }