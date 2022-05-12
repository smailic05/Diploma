package repository

import (
	"crypto/sha256"
	"encoding/hex"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/smailic05/diploma/internal/model"
)

type Auth struct {
	*gorm.DB
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{db}
}

func (db *Auth) GetUser(username string, password string) (*model.User, error) {
	var u = &model.User{
		Name:     username,
		Password: hash(password),
	}
	db.First(u)
	if u.Email == "" {
		u.Email = "default@email.com"
	}
	return u, nil
}

func (db *Auth) SaveToken(userID uuid.UUID, token string) error {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	db.Create(&model.Tokens{ID: tokenID, Token: token, UserID: userID})
	if err != nil {
		return err
	}
	return nil
}

func hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
