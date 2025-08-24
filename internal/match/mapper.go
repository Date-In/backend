package match

import "dating_service/internal/model"

func ToMatchPreviewDTOs(matches []model.Match, currentUserID uint) []MatchPreviewDTO {
	dtos := make([]MatchPreviewDTO, 0, len(matches))
	for _, match := range matches {
		dtos = append(dtos, toMatchPreviewDTO(match, currentUserID))
	}
	return dtos
}

func toMatchPreviewDTO(match model.Match, currentUserID uint) MatchPreviewDTO {
	return MatchPreviewDTO{
		MatchID:     match.ID,
		Partner:     mapPartnerToDTO(match, currentUserID),
		LastMessage: mapLastMessageToDTO(match.LastMessage, currentUserID),
	}
}

func mapPartnerToDTO(match model.Match, currentUserID uint) PartnerDTO {
	var partner model.User
	if match.User1ID == currentUserID {
		partner = match.User2
	} else {
		partner = match.User1
	}
	var avatar PhotoDto
	if partner.Avatar == nil {
		avatar = PhotoDto{}
	} else {
		avatar = PhotoDto{
			ID:  partner.Avatar.ID,
			Url: partner.Avatar.Url,
		}
	}
	return PartnerDTO{
		ID:     partner.ID,
		Name:   partner.Name,
		Avatar: avatar,
	}
}

func mapLastMessageToDTO(lastMessage *model.Message, currentUserID uint) *LastMessageDTO {
	if lastMessage == nil {
		return nil
	}

	return &LastMessageDTO{
		Text:      lastMessage.MessageText,
		CreatedAt: lastMessage.CreatedAt,
		IsRead:    lastMessage.IsRead || (lastMessage.SenderID == currentUserID),
	}
}
