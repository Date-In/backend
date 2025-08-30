package chat

import "dating_service/internal/model"

func MessagesToMessagesDto(messages []*model.Message) []MessageDto {
	var msgDtos []MessageDto
	for _, message := range messages {
		msgDtos = append(msgDtos, MessageDto{
			ID:          message.ID,
			UpdatedAt:   message.UpdatedAt,
			MessageText: message.MessageText,
			IsRead:      message.IsRead,
			MatchID:     message.MatchID,
			SenderID:    message.SenderID,
		})
	}
	return msgDtos
}
