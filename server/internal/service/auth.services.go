// Package service contains a services which is the core business login for our app
package service

import (
	"errors"
	"time"

	"github.com/jirugutema/kaishare/internal/dto"
	"github.com/jirugutema/kaishare/internal/repository"
	"github.com/jirugutema/kaishare/pkg"
)
var (
	ErrUserExists      = errors.New("user already exists")
	ErrIDGenerating    = errors.New("couldn't generate unique id")
	ErrEmailOrPassword = errors.New("wrong email or password")
	ErrValidatingToken = errors.New("error validation token")
)

func LoginService(user dto.LoginDTO) (dto.User, string, string, error) {
	res, hashedPassword, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		return res, "", "", err
	}

	if !pkg.ComparePasswordHash(user.Password, hashedPassword) {
		return res, "", "", ErrEmailOrPassword
	}

	accessToken, err := pkg.GenerateJWT(res.ID)
	if err != nil {
		return res, "", "", err
	}

	refreshToken, err := pkg.GenerateRefreshToken(res.ID)
	if err != nil {
		return res, "", "", err
	}

	return res, accessToken, refreshToken, nil
}

func RegisterService(user dto.RegisterDTO) (dto.User, string, string, error) {
	var res dto.User

	id, err := pkg.IDGenerator()
	if err != nil {
		return res, "", "", ErrIDGenerating
	}

	res.ID = id
	res.Username = user.Username
	res.Email = user.Email
	res.CreatedAt = time.Now()
	res.UpdatedAt = res.CreatedAt

	exists, err := repository.UserExists(user.Email, user.Username)
	if err != nil {
		return res, "", "", err
	}
	if exists {
		return res, "", "", ErrUserExists
	}

	passwordHash, err := pkg.HashPassword(user.Password)
	if err != nil {
		return res, "", "", err
	}

	err = repository.CreateUser(res, passwordHash)
	if err != nil {
		return res, "", "", err
	}

	accessToken, err := pkg.GenerateJWT(res.ID)
	if err != nil {
		return res, "", "", err
	}

	refreshToken, err := pkg.GenerateRefreshToken(res.ID)
	if err != nil {
		return res, "", "", err
	}

	return res, accessToken, refreshToken, nil
}

func GetAccessTokenService(refreshToken string) (string, error) {
	token, err := pkg.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	userID, err := pkg.GetUserIDFromToken(token)
	if err != nil {
		return "", err
	}

	return pkg.GenerateJWT(userID)
}

