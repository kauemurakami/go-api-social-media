package models

import (
	sec "api-social-media/internal/core/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	Pass      string    `json:"pass"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}
	if err := user.formatter(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("o nome deve ser inserido")
	}
	if user.Email == "" {
		return errors.New("o email deve ser inserido")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("o email inserido é inválido")
	}
	if user.Nick == "" {
		return errors.New("o nick deve ser inserido")
	}
	if step == "cadastro" && user.Pass == "" {
		return errors.New("o password deve ser inserido")
	}

	return nil
}

func (user *User) formatter(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Pass = strings.TrimSpace(user.Pass)
	if step == "cadastro" {
		senhaComHash, err := sec.Hash(user.Pass)
		if err != nil {
			return err
		}
		user.Pass = string(senhaComHash)
	}
	return nil
}
