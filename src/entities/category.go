package entities

import (
	"errors"
	ERROR "go_gift_list_api/src/shared/errors"
	"time"
)

type Category struct {
	Base
	Name string `json:"name" bson:"name"`
}

func BuildCategory(name string, email string, password string, id UniqueEntityID) (*Category, error) {
	category := &Category{
		Base: Base{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: name,
	}

	if err := category.validate(); err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) validate() error {
	length := len(c.Name)

	if length <= 0 {
		return errors.New(ERROR.NAME_IS_EMPTY)
	}

	return nil
}
