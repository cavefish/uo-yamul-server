package events

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/events/clientEvents"
	"yamul-gateway/internal/events/gameEvents"
	"yamul-gateway/internal/services/game/messages"
	"yamul-gateway/internal/transport/multima/listeners"
)

func Setup() {
	listeners.OnClientMoveAck.SetListener(clientEvents.OnClientMoveAck)
	listeners.OnClientDoubleClick.SetListener(clientEvents.OnClientDoubleClick)
	listeners.OnClientMoveRequest.SetListener(clientEvents.OnClientMoveRequest)
	listeners.OnGameLoginRequest.SetListener(clientEvents.OnGameLoginRequest)
	listeners.OnLoginRequest.SetListener(clientEvents.OnLoginRequest)
	listeners.OnOpenChatWindow.SetListener(clientEvents.OnOpenChatWindow)
	listeners.OnPreLogin.SetListener(clientEvents.OnCharacterPreLogin)
	listeners.OnShardSelected.SetListener(clientEvents.OnShardSelected)
	listeners.OnUnicodeSpeechRequest.SetListener(clientEvents.OnUnicodeSpeechRequest)

	messages.RegisterProcessor(services.MsgType_TypeUndefined, messages.UnimplementedProcessor)

	messages.RegisterProcessor(services.MsgType_TypeApplyWorldPatches, gameEvents.ApplyWorldPatches)
	messages.RegisterProcessor(services.MsgType_TypeCharacterSelection, messages.UnimplementedProcessor) // TODO
	messages.RegisterProcessor(services.MsgType_TypeExtendedStats, gameEvents.ExtendedStats)
	messages.RegisterProcessor(services.MsgType_TypeGeneralLightLevel, gameEvents.GeneralLightLevel)
	messages.RegisterProcessor(services.MsgType_TypeHealthBar, gameEvents.OnHealthBarUpdate)
	messages.RegisterProcessor(services.MsgType_TypeLoginComplete, gameEvents.OnLoginComplete)
	messages.RegisterProcessor(services.MsgType_TypeMapChange, gameEvents.OnMapChange)
	messages.RegisterProcessor(services.MsgType_TypeMoveAck, gameEvents.ServerMoveAck)
	messages.RegisterProcessor(services.MsgType_TypeSkillUpdateServer, gameEvents.SkillUpdateServer)
	messages.RegisterProcessor(services.MsgType_TypeOpenPaperDoll, gameEvents.OnOpenPaperDoll)
	messages.RegisterProcessor(services.MsgType_TypePlayerStartConfirmation, gameEvents.OnPlayerStartConfirmation)
	messages.RegisterProcessor(services.MsgType_TypePlayMusic, gameEvents.OnPlayMusic)
	messages.RegisterProcessor(services.MsgType_TypeStatWindowInfo, gameEvents.StatWindowInfo)
	messages.RegisterProcessor(services.MsgType_TypeSystemSendText, gameEvents.SystemSendText)
	messages.RegisterProcessor(services.MsgType_TypeTeleportPlayer, gameEvents.OnTeleportPlayer)
	messages.RegisterProcessor(services.MsgType_TypeUpdateObject, gameEvents.OnUpdateObject)
	messages.RegisterProcessor(services.MsgType_TypeWarmode, gameEvents.WarMode)
}
