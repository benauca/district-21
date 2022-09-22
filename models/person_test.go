package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/** Test SetName */
func Test_set_name(t *testing.T) {
	person := Person {}
	//TODO: Create name as random string
	person.SetName("benauca")
	assert.NotNil(t, person.Name, ":( The name mustn't be null...")
}

/** Test Set UUID */
func Test_set_uuid(t *testing.T) {
	person := Person {}
	person.SetUUID()
	assert.NotNil(t, person.UUID, ":( The uuid must not be null...")
}

/** Test GetUUID */
func Test_get_uuid(t *testing.T) {
	person := Person {}
	person.SetUUID()
	assert.Equal(t, person.GetUUID(), person.UUID, ":( The uuid should be the same.")
}

/** Test Set nickname*/
func Test_set_nickname(t *testing.T) {
	person := Person {}
	person.SetNickname("benauca")
	assert.NotNil(t, person.nickname, ":( The nickname must not be null...")
}

/** Test GetNickname */
func Test_get_nickname(t *testing.T) {
	person := Person {}
	person.SetNickname("benauca")
	assert.Equal(t, person.GetNickname(), person.nickname, ":( The nickname should be the same.")
}

/** Test Set email*/
func Test_set_email(t *testing.T) {
	person := Person {}
	person.SetNickname("myname@domain.com")
	assert.NotNil(t, person.email, ":( The email must not be null...")
}

/** Test GetNickname */
func Test_get_email(t *testing.T) {
	person := Person {}
	person.SetEmail("myname@domain.com")
	assert.Equal(t, person.GetEmail(), person.email, ":( The email should be the same.")
}