package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jirugutema/kaishare/internal/dto"
	"github.com/jirugutema/kaishare/internal/repository"
	"github.com/jirugutema/kaishare/pkg"
)

var (
	ErrUserNotExist       = errors.New("user does not exist")
	ErrPasteExpired       = errors.New("paste has expired")
	ErrCantDeletePaste    = errors.New("unauthorized access to delete paste")
	ErrWrongPassword      = errors.New("wrong password is provided for the paste")
	ErrYouCantUpdatePaste = errors.New("you are not unauthorized to update the paste")
	ErrPasteIDIsRequired  = errors.New("paste id is required")
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

	if paste.Password != nil && *paste.Password != ""  {
		hashedPassword, err := pkg.HashPassword(*paste.Password)
		if err != nil {
			return uuid.Nil, err
		}
		paste.Password = &hashedPassword
	} else {
		paste.Password = nil
	}
	if err := repository.CreatePaste(paste); err != nil {
		fmt.Print("here")
		return uuid.Nil, err
	}

	return paste.ID, nil
}

func GetPasteService(pasteID uuid.UUID, password string, unlocked bool) (dto.PasteResponse, error) {
	res, err := repository.GetPaste(pasteID)
	paste := dto.PasteResponse{}

	if err != nil {
		return paste, err
	}

	if res.Password != nil && !unlocked {
		match := pkg.ComparePasswordHash(password, *res.Password)
		if !match {
			paste.RequiresPassword = true
			return paste, ErrWrongPassword
		}
	}
	_ = repository.IncrementViews(pasteID)

	if res.BurnAfterRead {
		_ = repository.DeletePaste(pasteID)
	}

	if res.ExpiresAt != nil && res.ExpiresAt.Before(time.Now()) {
		_ = repository.DeletePaste(pasteID)
		return dto.PasteResponse{}, ErrPasteExpired
	}

	paste.ID = res.ID
	paste.Title = res.Title
	paste.Content = res.Content
	paste.Language = res.Language
	paste.BurnAfterRead = res.BurnAfterRead
	paste.ExpiresAt = res.ExpiresAt
	paste.CreatedAt = res.CreatedAt
	paste.Views = res.Views
	paste.UserID = res.UserID
	paste.IsPublic = res.IsPublic

	return paste, nil
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

func UpdatePasteService(paste dto.UpdatePasteDTO, userID uuid.UUID) (dto.PasteResponse, error) {
	empty := dto.PasteResponse{}

	pasteID := paste.ID
	if pasteID == uuid.Nil {
		return empty, errors.New("paste id is required")
	}

	p, err := repository.GetPaste(pasteID)
	if err != nil {
		return empty, errors.New("paste id is required")
	}

	if p.UserID != nil && *p.UserID != userID {
		return empty, errors.New(" id is required")
	}
	res, err := repository.UpdatePaste(paste)
	if err != nil {
		return empty, err
	}
	return *res, err
}
