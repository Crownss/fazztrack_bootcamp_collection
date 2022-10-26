package middleware_test

import (
	"testing"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/middleware"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCheckToken(t *testing.T) {
	var payload_jwt struct {
		Username string
		Is_admin bool
		Exp      int
	}
	payload_jwt.Username = "fantom"
	payload_jwt.Is_admin = true
	payload_jwt.Exp = 1664638427
	result, err := middleware.CheckTokenString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImZhbnRvbSIsIklzX0FkbWluIjp0cnVlLCJleHAiOjE2NjQ4NTY4NDJ9.j7WyMFfUzgQBmIgGiqmxI_84HIpMV-vsVpvRbqrNTfo")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	assert.Equal(t, payload_jwt.Username, result, "error: username must same")
}

func TestNewToken(t *testing.T) {
	test1 := middleware.NewToken("fantom", true)
	assert.Equal(t, "fantom", test1.Username, "error: username must same")
}

func TestCheckPassword(t *testing.T) {
	userpw := "password1"
	dbpw := "$2a$04$Wr4MPWtX32DwbMGftUi33.J/qEy91rvQ2YUJhfFHXUxpNo.LO4Tkm"
	decrypt := bcrypt.CompareHashAndPassword([]byte(dbpw), []byte(userpw))
	assert.Nil(t, decrypt, "password not same with encrypted password")
}
