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

	messages.RegisterProcessor(services.MsgType_TypeApplyWorldPatches, gameEvents.ApplyWorldPatches)
	messages.RegisterProcessor(services.MsgType_TypeCharacterSelection, messages.UnimplementedProcessor) // TODO
	messages.RegisterProcessor(services.MsgType_TypeExtendedStats, gameEvents.ExtendedStats)
	messages.RegisterProcessor(services.MsgType_TypeGeneralLightLevel, gameEvents.GeneralLightLevel)
	messages.RegisterProcessor(services.MsgType_TypeHealthBar, gameEvents.OnHealthBarUpdate)
	messages.RegisterProcessor(services.MsgType_TypeMapChange, gameEvents.OnMapChange)
	messages.RegisterProcessor(services.MsgType_TypePlayerStartConfirmation, gameEvents.OnPlayerStartConfirmation)
	messages.RegisterProcessor(services.MsgType_TypePlayMusic, gameEvents.OnPlayMusic)
	messages.RegisterProcessor(services.MsgType_TypeStatWindowInfo, gameEvents.StatWindowInfo)
	messages.RegisterProcessor(services.MsgType_TypeTeleportPlayer, gameEvents.OnTeleportPlayer)
	messages.RegisterProcessor(services.MsgType_TypeUpdateObject, gameEvents.OnUpdateObject)
	messages.RegisterProcessor(services.MsgType_TypeWarmode, gameEvents.WarMode)
}
