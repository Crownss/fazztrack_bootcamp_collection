package middleware_test

import (
	"testing"

	"github.com/crownss/fazztrack_bootchamp/week_9/src/middleware"
	"github.com/stretchr/testify/assert"
)

func TestCheckToken(t *testing.T) {
	var payload_jwt struct {
		Username string
		Is_admin bool
		Exp int
	}
	payload_jwt.Username = "fantom"
	payload_jwt.Is_admin = true
	payload_jwt.Exp = 1664638427
	/*
	{
  		"Username": "fantom",
  		"Is_Admin": true,
  		"exp": 1664638427
	}
	*/
	result, err := middleware.CheckToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImZhbnRvbSIsIklzX0FkbWluIjp0cnVlLCJleHAiOjE2NjQ2Mzg0Mjd9.cO2Uy5JFRw493Q-WfKkAX6-1DTqkliHAkPjiEtCeJy4")
	// assert.Same(t, "fantom", result.Username, err.Error())
	assert.Equal(t, payload_jwt.Username, result.Username, err.Error())
}
