package domain

import "github.com/google/uuid"

type UUIDValueObject struct {
	value uuid.UUID
}

func (u *UUIDValueObject) Value() string {
	return u.value.String()
}

func NewUUIDValueObject() (*UUIDValueObject, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &UUIDValueObject{value: id}, nil
}

func NewUUIDValueObjectFromString(value string) (*UUIDValueObject, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return nil, err
	}

	return &UUIDValueObject{value: id}, nil
}
