package connection

import (
	"testing"
	"yamul-gateway/internal/dtos"
	"yamul-gateway/utils/tests/assertions"
)

func TestClientConnectionInterface(t *testing.T) {
	t.Run("Implements interface", func(t *testing.T) {
		sut := &clientConnection{
			loginDetails: dtos.LoginDetails{
				Username:      "username",
				Password:      "password",
				CharacterSlot: 0,
			},
		}
		assert := assertions.For(t)
		assert.Equals(&sut.loginDetails, sut.GetLoginDetails())
	})
}
