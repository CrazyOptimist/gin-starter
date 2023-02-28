package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gin-starter/cmd/api/config"
)

func TestGenerateAccessToken(t *testing.T) {

	if err := config.LoadConfig("../../test.env"); err != nil {
		panic("Failed to load test env file")
	}

	t.Run("it should generate a valid access token", func(t *testing.T) {

		userId := uint(1)
		accessToken, err := GenerateAccessToken(userId)
		assert.NoError(t, err)

		isValid, _, _, _ := ValidateToken(accessToken)

		if isValid != true {
			t.Errorf("Expected a valid access token, but got an invalid one:\n %v ", accessToken)
		}
	})
}