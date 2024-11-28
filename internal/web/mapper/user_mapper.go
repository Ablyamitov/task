package mapper

import (
	"github.com/Ablyamitov/task/internal/storage/model"
	"github.com/Ablyamitov/task/internal/web/response"
)

func MapUserToUserDTO(user *model.User) *response.UserDTO {
	return &response.UserDTO{
		ID:        user.ID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
		Phone:     user.Phone,
	}
}

func MapUserDTOToUser(dto *response.UserDTO) *model.User {
	return &model.User{
		ID:        dto.ID,
		LastName:  dto.LastName,
		FirstName: dto.FirstName,
		Gender:    dto.Gender,
		BirthDate: dto.BirthDate,
		Phone:     dto.Phone,
	}
}
