package handlers

import (
	"strings"
	"yamul-gateway/internal/dtos/commands"
	"yamul-gateway/internal/interfaces"
	"yamul-gateway/internal/transport/multima/listeners"
)

func unicodeSpeechRequestHandler(client interfaces.ClientConnection) { // 0xA0
	result := unicodeSpeechRequestReadBuffer(client)

	listeners.OnUnicodeSpeechRequest.Trigger(client, result)

}

func unicodeSpeechRequestReadBuffer(client interfaces.ClientConnection) commands.UnicodeSpeechSelected {
	_ = client.ReadUShort() // size
	mode := client.ReadByte()
	hue := client.ReadUShort()
	font := client.ReadUShort()
	language := client.ReadFixedString(4)
	var keywords []uint16
	if mode&0xC0 != 0 {
		prevByte := client.ReadByte()
		currentByte := client.ReadByte()
		nrKeywords := int((prevByte << 4) | (currentByte >> 4))
		keywords = make([]uint16, nrKeywords)
		nextKeywordNeedsFullByte := true
		for i := 0; i < nrKeywords; i++ {
			prevByte = currentByte
			currentByte = client.ReadByte()
			if nextKeywordNeedsFullByte {
				keywords[i] = (uint16(prevByte)&0xF)<<8 | uint16(currentByte)
			} else {
				prevByte = currentByte
				currentByte = client.ReadByte()
				keywords[i] = uint16(prevByte)<<4 | uint16(currentByte&0xF0)>>4
			}
			nextKeywordNeedsFullByte = !nextKeywordNeedsFullByte
		}
	}
	var sb strings.Builder
	for {
		t := client.ReadByte()
		if t == 0 {
			break
		}
		sb.WriteByte(t)
	}

	result := commands.UnicodeSpeechSelected{
		Mode:     mode &^ 0xC0,
		Hue:      hue,
		Font:     font,
		Language: language,
		Keywords: keywords,
		Text:     sb.String(),
	}
	return result
}
