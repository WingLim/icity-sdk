package icity_sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_GetWorld(t *testing.T) {
	user := login()

	diaries := user.GetWorld()
	assert.Len(t, diaries, 25)
}
