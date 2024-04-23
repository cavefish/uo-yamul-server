package clientEvents

import (
	"yamul-gateway/backend/services"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/services/login"
	"yamul-gateway/internal/transport/multima/listeners"
)

func OnCharacterPreLogin(event listeners.CommandEvent[commands.PreLogin]) {
	success, deniedReason := login.ValidateLogin(event.Client.GetLoginDetails().Username, event.Command.Password)
	if !success {
		login.DenyLogin(event.Client, deniedReason)
		return
	}

	event.Client.GetLoginDetails().CharacterSlot = event.Command.Slot

	err := event.Client.CreateGameConnection()
	if err != nil {
		event.Client.KillConnection(err)
		return
	}
	body := services.MsgCharacterSelection{
		Username: event.Client.GetLoginDetails().Username,
		Slot:     int32(event.Command.Slot),
	}
	selection := services.Message_CharacterSelection{CharacterSelection: &body}
	event.Client.GetGameService().Send(services.MsgType_TypeCharacterSelection, &services.Message{Msg: &selection})
}
