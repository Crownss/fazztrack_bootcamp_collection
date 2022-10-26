package helpers_test

import (
	"testing"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/helpers"
	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	result, err := helpers.EncryptPassword("password")
	assert.Equal(t, "password", result, err.Error())
}
