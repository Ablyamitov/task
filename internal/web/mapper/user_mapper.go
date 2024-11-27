package mapper

import (
	"github.com/Ablyamitov/task/internal/storage/model"
	"github.com/Ablyamitov/task/internal/web/dto"
)

func MapUserToUserDTO(user *model.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:        user.ID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
		Phone:     user.Phone,
	}
}

func MapUserDTOToUser(dto *dto.UserDTO) *model.User {
	return &model.User{
		ID:        dto.ID,
		LastName:  dto.LastName,
		FirstName: dto.FirstName,
		Gender:    dto.Gender,
		BirthDate: dto.BirthDate,
		Phone:     dto.Phone,
	}
}
