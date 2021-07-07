package icity_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetWorld(t *testing.T) {
	user := login()

	diarys := user.GetWorld()
	assert.Len(t, diarys, 25)
}
