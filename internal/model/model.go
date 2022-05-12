package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

type Tokens struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	Token  string
	UserID uuid.UUID
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Error struct {
	Error string `json:"error"`
}
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
