package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_get_color(t *testing.T) {
	player := Player{
		name:  "benauca",
		color: "blue",
	}
	assert.Equal(t, player.color, player.GetColor(), "The color should be the same.")
}

func Test_get_name(t *testing.T) {
	player := Player{
		name:  "benauca",
		color: "blue",
	}
	assert.Equal(t, player.name, player.GetName(), "The name should be the same.")
}

func Test_GetDirection(t *testing.T) {
	token := Token{
		direction: 1,
	}
	assert.Equal(t, token.GetDirection(), token.direction)
}

func Test_GetTokenType(t *testing.T) {
	token := Token{
		tokenType: 2,
	}
	assert.Equal(t, token.GetType(), token.tokenType)

}
