// Package service contains a services which is the core business login for our app
package service

import (
	"errors"
	"time"

	"github.com/jirugutema/gopastebin/internal/config"
	"github.com/jirugutema/gopastebin/internal/dto"
	"github.com/jirugutema/gopastebin/pkg"
)

var (
	ErrUserExists      = errors.New("user already exists")
	ErrIDGenerating    = errors.New("couldn't generate unique id")
	ErrEmailOrPassword = errors.New("wrong email or password")
	ErrHashingPassword = errors.New("error hashing password")
)

// LoginService returns user, accessToken, refreshToken, error
func LoginService(user dto.LoginDTO) (dto.LoginResponse, string, string, error) {
	query := `SELECT id, email,password_hash, username, created_at, updated_at 
          FROM users 
          WHERE email = $1`
	row := config.DB.QueryRow(query, user.Email)
	var res dto.LoginResponse
	var hashedPassword string
	err := row.Scan(&res.ID, &res.Email, &hashedPassword, &res.Username, &res.UpdatedAt, &res.UpdatedAt)
	if err != nil {
		return res, "", "", err
	}
	match := pkg.ComparePasswordHash(user.Password, hashedPassword)
	if !match {
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

func RegisterServices(user dto.RegisterDTO) (dto.RegisterResponse, string, string, error) {
	var res dto.RegisterResponse
	var err error
	id, e := pkg.IDGenerator()
	if e != nil {
		return res, "", "", ErrIDGenerating
	}

	res.CreatedAt = time.Now()
	res.UpdatedAt = time.Now()
	res.ID = id
	res.Username = user.Username
	res.Email = user.Email

	var exists bool
	err = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 OR username = $2)", user.Email, user.Username).Scan(&exists)
	if err != nil {
		return res, "", " ", err
	}
	if exists {
		return res, "", " ", ErrUserExists
	}

	passwordHash, err := pkg.HashPassword(user.Password)
	if err != nil {
		return res, "", " ", err
	}
	query := "INSERT INTO users (id, username, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = config.DB.Exec(query, id, user.Username, user.Email, passwordHash, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return res, "", " ", err
	}
	accessToken, err := pkg.GenerateJWT(res.ID)
	if err != nil {
		return res, "", " ", err
	}

	refreshToken, err := pkg.GenerateRefreshToken(res.ID)
	if err != nil {
		return res, "", "", err
	}
	return res, accessToken, refreshToken, err
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
	accessToken, err := pkg.GenerateJWT(userID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}


