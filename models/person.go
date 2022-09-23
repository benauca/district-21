package models

import (
	"github.com/google/uuid"
)

type Person struct {

	Name string
	//UUID que identifica el jugador
	UUID uuid.UUID
	email string
	nickname string

}

/**
 * Crea un uuid a la instancia del Jugador que nos servirá para definir el
 * objeto en los eventos
 */
 func (person *Person) SetUUID() {
	person.UUID = uuid.New()
}

/** Get UUID */
 func (person *Person) GetUUID() uuid.UUID {
	return person.UUID
}

 

/** Set email */
func (person *Person) SetEmail(email string) {
	person.email = email
}

/** Get email */
func (person *Person) GetEmail() string {
	return person.email
}

/** Set name */
func (person *Person) SetName(name string) {
	person.Name = name
}

/** Get name */
func (person *Person) GetName() string {
	return person.Name 
}
/** Set nickname */
func (person *Person) SetNickname(nickname string) {
	person.nickname = nickname
}

/** Get nickname */
func (person *Person) GetNickname() string {
	return person.nickname
}
