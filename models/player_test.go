package models

import (
	"fmt"
	"testing"
	"district-21/models/tokens"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

/** Test SetName */
func Test_set_name_player(t *testing.T) {
	player := Player {}
	player.SetName("nameplayer")
	asString := fmt.Sprint(player)
	assert.Contains(t, asString, "{{nameplayer 00000000-0000-0000-0000-000000000000  }", ":( The name mustn't be null...")
}

/** Test SetName */
func Test_definition_player(t *testing.T) {
	player := Player {}
	player.SetName("nameplayer")
	player.SetRealm("blue")
	asString := fmt.Sprint(player)
	assert.Contains(t, asString, "{{nameplayer 00000000-0000-0000-0000-000000000000  } blue", ":( The name mustn't be null...")
}

/** Test SetChief */
func Test_setChief(t *testing.T) {
	player := Player {}
	player.SetName("nameplayer")
	player.SetChief(tokens.Token{TypeToken: 1, UUID: uuid.New()})
	assert.Equal(t, player.GetChief().TypeToken,1, ":( GetToken dont return the typetoken rightly...")
}

/**Test IsLive without token*/
func Test_isLive_without_token(t *testing.T) {
	player := Player {}
	player.SetName("nameplayer")
	assert.Equal(t, player.IsLive(), false, ":( User player is kied...")
}

/** Test IsLive with token */
func Test_isLive_with_token(t *testing.T) {
	player := Player {}
	player.SetName("nameplayer")
	player.SetChief(tokens.Token{TypeToken: 1, UUID: uuid.New()})
	assert.Equal(t, player.IsLive(), true, ":( User player is live...")
}

/** Test Get Detectives */
func Test_get_detectives(t *testing.T)  {
	player := Player{}
	detectives := make([]tokens.Token, 0)
	detectives = append(
		detectives,
		tokens.Token {
			TypeToken: 2,
			UUID: uuid.New(),
		},
	)
	player.SetDetectives(detectives)
	assert.NotNil(t, player.GetDetectives(), "[:(] Detectives is null")
	assert.Equal(t, len(player.GetDetectives()), 1, "[:(] The size detectives is no the same...")
}

/** Test Get Policies */
func Test_get_policies(t *testing.T)  {
	player := Player{}
	policies := make([]tokens.Token, 0)
	policies = append(
		policies,
		tokens.Token {
			TypeToken: 2,
			UUID: uuid.New(),
		},
	)
	player.SetPolicies(policies)
	assert.NotNil(t, player.GetPolicies(), "[:(] Policies is null")
	assert.Equal(t, len(player.GetPolicies()), 1, "[:(] The size policies is no the same...")
}