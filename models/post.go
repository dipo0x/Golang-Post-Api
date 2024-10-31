package models

import (
	"time"
	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id"`
	Title     string    `bson:"title" json:"title" validate:"required"`
	Content   string    `bson:"content" json:"content" validate:"required"`
	Author    string    `bson:"author" json:"author" validate:"required"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
