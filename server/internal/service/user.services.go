package service

import (
	"github.com/google/uuid"
	"github.com/jirugutema/kaishare/internal/dto"
	"github.com/jirugutema/kaishare/internal/repository"
)


func DeleteUser(userID uuid.UUID) (dto.User, error) {
	user, _, err := repository.GetUserByID(userID)
	if err != nil {
		return dto.User{}, ErrUserNotExist
	}
	err = repository.DeleteUser(userID)
	if err != nil {
		return dto.User{}, err
	}
	return user, nil
}
