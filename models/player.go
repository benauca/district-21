package models

import (
	"district-21/models/tokens"
)
type Player struct {
	
	//Hacemos uso de la herencia Player -> Person
	Person
	//El jugador selecciona el reino que lidera.
	realm string
	//Chief
	chief tokens.Token
	//Lista de detectives
	detectives []tokens.Token
	//Lista de policias
	policies []tokens.Token
}

/**
 * Método que nos dice si el jugador sigue en el juego o ha sido derrotado 
 */
func (player *Player) IsLive() bool {
	result := false
	if player.chief != (tokens.Token{}) {
		result = true
	}
	return result
}

/**
 * Set Realm
 */
 func (player *Player) SetRealm(realm string)  {
	player.realm = realm	
 }

/**
 * Return Color o reino con el que juega
 */
func (player *Player) GetRealm() string {
	return player.realm
}

/**
 * Set Chief
 * [ES] Asigna el rey
 */
 func (player *Player) SetChief(token tokens.Token) {
	player.chief = token
 }

 /** Get chief */
 func (player *Player) GetChief() tokens.Token {
	return player.chief
 }

 /**
  * Set Detectives tokens
  * [ES] Asigna las fichas de caballería
  */
 func (player *Player) SetDetectives(tokens []tokens.Token){
	player.detectives = tokens
 }

/** Get Detectives */
func (player *Player) GetDetectives() [] tokens.Token {
	return player.detectives
}

 /**
  * Set polices tokens
  * [ES] Asigna las fichas de infantería
  */
 func (player *Player) SetPolicies(tokens []tokens.Token){
	player.policies = tokens
 }

 /** Get Detectives */
func (player *Player) GetPolicies() [] tokens.Token {
	return player.policies
}
