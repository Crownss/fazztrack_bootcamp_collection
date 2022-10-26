package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func HelloName(s string) string {
	return "hello " + s
}

func TestHelloName(t *testing.T) {
	result := HelloName("jesen")
	assert.Equal(t, "hello jesen", result, "return must be hello jesen")
}

func TestSubHelloName(t *testing.T){
	t.Run("test pertama", func(t *testing.T) {
		result := HelloName("fantom")
		assert.Equal(t, "hello fantom", result, "return must be hello fantom")
	})

	t.Run("test kedua", func(t *testing.T) {
		result := HelloName("jesen")
		assert.Equal(t, "hello jesen", result, "return must be hello jesen")
	})

	t.Run("test ketiga", func(t *testing.T) {
		result := HelloName("soehar-")
		assert.Equal(t, "hello soehar-", result, "return must be hello soehar-")
	})
}
