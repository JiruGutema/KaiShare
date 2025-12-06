package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jirugutema/kaishare/internal/dto"
	"github.com/jirugutema/kaishare/internal/repository"
	"github.com/jirugutema/kaishare/pkg"
)

var (
	ErrUserNotExist    = errors.New("user does not exist")
	ErrPasteExpired    = errors.New("paste has expired")
	ErrCantDeletePaste = errors.New("unauthorized access to delete paste")
)

func CreatePasteService(paste dto.PasteDTO) (uuid.UUID, error) {
	paste.CreatedAt = time.Now()

	pasteID, err := pkg.IDGenerator()
	if err != nil {
		return uuid.Nil, err
	}
	paste.ID = pasteID

	if paste.UserID != nil {
		exists, err := repository.PasteUserExists(*paste.UserID)
		if err != nil {
			return uuid.Nil, err
		}
		if !exists {
			return uuid.Nil, ErrUserNotExist
		}
	}

	if err := repository.CreatePaste(paste); err != nil {
		return uuid.Nil, err
	}

	return paste.ID, nil
}

func GetPasteService(pasteID uuid.UUID) (dto.PasteDTO, error) {
	res, err := repository.GetPaste(pasteID)
	if err != nil {
		return res, err
	}

	_ = repository.IncrementViews(pasteID)

	if res.BurnAfterRead {
		_ = repository.DeletePaste(pasteID)
	}

	if res.ExpiresAt != nil && res.ExpiresAt.Before(time.Now()) {
		return dto.PasteDTO{}, ErrPasteExpired
	}

	return res, nil
}

func GetMyPastesService(userID uuid.UUID) (dto.MyPastesDTO, error) {
	myPastes, err := repository.GetMyPastes(userID)
	if err != nil {
		return dto.MyPastesDTO{}, err
	}

	return myPastes, nil
}

func DeletePasteService(pasteID uuid.UUID, userID uuid.UUID) error {
	res, err := repository.GetPaste(pasteID)
	if err != nil {
		return err
	}

	if *res.UserID != userID {
		return ErrCantDeletePaste
	}

	e := repository.DeletePaste(pasteID)
	return e
}
