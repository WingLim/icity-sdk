package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetWorld(t *testing.T) {
	user := login()

	diaries := user.GetWorld()
	assert.Len(t, diaries, 25)
}
