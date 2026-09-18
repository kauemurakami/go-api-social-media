package models

import "github.com/google/uuid"

type Followers struct {
	ID         uuid.UUID `json:"id,omitempty"`
	FollowerID uuid.UUID `json:"follower_id"`
	UserID     uuid.UUID `json:"user_id"`
}
