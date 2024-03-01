package autoconfig

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/services/game/messages"
	"yamul-gateway/internal/services/login"
)

func TestModuleSetup(t *testing.T) {
	tests := []struct {
		name          string
		expectedError error
	}{
		{
			name: "ModuleSetup", expectedError: nil,
		},
	}
	for _, tt := range tests {
		loginModuleMock := &mockedModule{}
		login.Module = loginModuleMock
		t.Run(tt.name, func(t *testing.T) {
			var assertions = assert.New(t)
			err := Setup()
			assertions.Equal(tt.expectedError, err)
			assertions.Equal(loginModuleMock.setupInv, 1)
			assertions.Equal(loginModuleMock.closeInv, 0)
			Close()
			assertions.Equal(loginModuleMock.setupInv, 1)
			assertions.Equal(loginModuleMock.closeInv, 1)
		})
	}
}

type mockedModule struct {
	setupInv    int
	setupReturn error
	closeInv    int
}

func (m *mockedModule) Setup() error {
	m.setupInv++
	return m.setupReturn
}

func (m *mockedModule) Close() {
	m.closeInv++
}

type processorTest struct {
	msgName string
	msgType int32
}

func TestGameServiceProcessorsConfigured(t *testing.T) {
	testCases := make([]processorTest, len(services.MsgType_name))
	i := 0
	for msgType := range services.MsgType_name {
		testCases[i] = processorTest{
			msgName: services.MsgType_name[msgType],
			msgType: msgType,
		}
		i++
	}

	for _, tt := range testCases {
		loginModuleMock := &mockedModule{}
		login.Module = loginModuleMock
		t.Run(tt.msgName, func(t *testing.T) {
			var assertions = assert.New(t)
			err := Setup()
			assertions.NoError(err)
			_, isPresent := messages.Processors[services.MsgType(tt.msgType)]
			assertions.True(isPresent, "No processor present")
		})
	}
}
