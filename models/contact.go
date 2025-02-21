package models

import (
	"time"
)

type Contact struct {
	ID             int        `bson:"id"`
	PhoneNumber    *string    `bson:"phoneNumber,omitempty"`
	Email          *string    `bson:"email,omitempty"`
	LinkedID       *int       `bson:"linkedId,omitempty"`
	LinkPrecedence string     `bson:"linkPrecedence"`
	CreatedAt      time.Time  `bson:"createdAt"`
	UpdatedAt      time.Time  `bson:"updatedAt"`
	DeletedAt      *time.Time `bson:"deletedAt,omitempty"`
}
