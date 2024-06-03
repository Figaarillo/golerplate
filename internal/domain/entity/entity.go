package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(str string) (ID, error) {
	id, err := uuid.Parse(str)
	return ID(id), err
}

type Response struct {
	body    interface{}
	message string
	count   int
	err     bool
}
