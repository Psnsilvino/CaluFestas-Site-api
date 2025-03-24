package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Nome      string             `json:"nome"`
	Email     string             `json:"email"`
	Senha     string             `json:"senha"`
}
