package repository

import (
	"github.com/google/uuid"
	"github.com/jirugutema/kaishare/internal/config"
	"github.com/jirugutema/kaishare/internal/dto"
)

func GetUserByEmail(email string) (dto.User, string, error) {
	var res dto.User
	var hashedPassword string
	query := `SELECT id, email, password_hash, username, created_at, updated_at FROM users WHERE email=$1`
	err := config.DB.QueryRow(query, email).Scan(
		&res.ID,
		&res.Email,
		&hashedPassword,
		&res.Username,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return res, hashedPassword, err
}


func GetUserByID(userID uuid.UUID) (dto.User, string, error) {
	var res dto.User
	var hashedPassword string
	query := `SELECT id, email, password_hash, username, created_at, updated_at FROM users WHERE id=$1`
	err := config.DB.QueryRow(query, userID).Scan(
		&res.ID,
		&res.Email,
		&hashedPassword,
		&res.Username,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return res, hashedPassword, err
}

func UserExists(email, username string) (bool, error) {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1 OR username=$2)", email, username).Scan(&exists)
	return exists, err
}

func CreateUser(user dto.User, passwordHash string) error {
	query := "INSERT INTO users (id, username, email, password_hash, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)"
	_, err := config.DB.Exec(query, user.ID, user.Username, user.Email, passwordHash, user.CreatedAt, user.UpdatedAt)
	return err
}

func DeleteUser(userID uuid.UUID) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := config.DB.Exec(query, userID)
	return err
}
