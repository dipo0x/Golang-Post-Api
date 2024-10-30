package models

import (
	"time"
	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id"`
	Title     string    `bson:"title" json:"title"`
	Content   string    `bson:"content" json:"content"`
	Author    string    `bson:"author" json:"author"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
