package connection

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"yamul-gateway/internal/dtos"
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
		assertions := assert.New(t)
		assertions.Equal(&sut.loginDetails, sut.GetLoginDetails())
	})
}
