package types

import (
	"github.com/google/uuid"
)
type IPost struct {
	ID  uuid.UUID `bson:"_id,omitempty" json:"id"`
    Title   string `json:"title" validate:"required"`
    Content string `json:"content" validate:"required"`
    Author  string `json:"author" validate:"required"`
}