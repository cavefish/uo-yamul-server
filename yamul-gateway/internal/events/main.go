package events

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/events/clientEvents"
	"yamul-gateway/internal/events/gameEvents"
	"yamul-gateway/internal/services/game/messages"
	"yamul-gateway/internal/transport/multima/listeners"
)

func Setup() {
	listeners.Listeners.OnLoginRequest.SetListener(clientEvents.OnLoginRequest)
	listeners.Listeners.OnShardSelected.SetListener(clientEvents.OnShardSelected)
	listeners.Listeners.OnGameLoginRequest.SetListener(clientEvents.OnGameLoginRequest)
	listeners.Listeners.OnPreLogin.SetListener(clientEvents.OnCharacterPreLogin)

	messages.RegisterProcessor(services.MsgType_TypeUndefined, messages.UnimplementedProcessor)

	messages.RegisterProcessor(services.MsgType_TypeApplyWorldPatches, messages.UnimplementedProcessor)
	messages.RegisterProcessor(services.MsgType_TypeCharacterSelection, messages.UnimplementedProcessor)
	messages.RegisterProcessor(services.MsgType_TypeCreateCharacter, messages.UnimplementedProcessor)
	messages.RegisterProcessor(services.MsgType_TypeHealthBar, messages.UnimplementedProcessor)
	messages.RegisterProcessor(services.MsgType_TypeMapChange, messages.UnimplementedProcessor)
	messages.RegisterProcessor(services.MsgType_TypePlayMusic, gameEvents.OnPlayMusic)
	messages.RegisterProcessor(services.MsgType_TypeTeleportPlayer, gameEvents.OnTeleportPlayer)
}
