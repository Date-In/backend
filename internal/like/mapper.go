package like

import "dating_service/internal/model"

func LikeToDto(like []model.Like) []LikeDto {
	if len(like) == 0 {
		return make([]LikeDto, 0)
	}
	var result []LikeDto
	for _, v := range like {
		photos := v.User.Photos
		var gallery []PhotoDto
		if len(photos) == 0 {
			gallery = make([]PhotoDto, 0)
		} else {
			for _, photo := range photos {
				gallery = append(gallery, PhotoDto{
					photo.ID,
					photo.Url,
				})
			}
		}
		var avatar PhotoDto
		if v.User.Avatar == nil {
			avatar = PhotoDto{}
		} else {
			avatar = PhotoDto{
				ID:  v.User.Avatar.ID,
				Url: v.User.Avatar.Url,
			}
		}
		result = append(result, LikeDto{
			ID: v.ID,
			User: LikeUserDto{
				ID:      v.UserID,
				Name:    v.User.Name,
				Age:     v.User.Age,
				Bio:     v.User.Bio,
				City:    v.User.City,
				Avatar:  &avatar,
				Gallery: gallery,
			},
		})
	}
	return result
}
