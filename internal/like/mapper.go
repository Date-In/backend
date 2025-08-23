package like

import "dating_service/internal/model"

func LikeToDto(like []model.Like) []LikeDto {
	var result []LikeDto
	for _, v := range like {
		photos := v.User.Photos
		var gallery []PhotoDto
		for _, photo := range photos {
			gallery = append(gallery, PhotoDto{
				photo.ID,
				photo.Url,
			})
		}
		result = append(result, LikeDto{
			ID: v.ID,
			User: LikeUserDto{
				ID:   v.UserID,
				Name: v.User.Name,
				Age:  v.User.Age,
				Bio:  v.User.Bio,
				City: v.User.City,
				Avatar: &PhotoDto{
					ID: v.User.Avatar.ID,
				},
				Gallery: gallery,
			},
		})
	}
	return result
}
