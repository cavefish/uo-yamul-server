package handlers

import (
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
)

func SystemSendText(client interfaces.ClientConnection, body commands.SystemSendText) { // 0x1C
	client.StartPacket()
	defer client.EndPacket()

	client.WriteByte(0x1C)
	client.WriteUShort(uint16(44 + len(body.Body) + 1))
	client.WriteUInt(body.Serial)
	client.WriteUShort(body.Model)
	client.WriteByte(body.Type)
	client.WriteUShort(body.Hue)
	client.WriteUShort(body.Font)
	client.WriteFixedString(30, body.Name)
	client.WriteFixedString(1+len(body.Body), body.Body)
}
