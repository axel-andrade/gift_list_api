package find_gifts_interactor

import (
	"go_gift_list_api/src/entities"
)

type FindGiftsGateway interface {
	CheckExistsCategory(categoryID entities.UniqueEntityID) bool
	FindGiftsPaginate(categoryID entities.UniqueEntityID, pagination *entities.PaginationOptions) ([]entities.Gift, int64, error)
}

type FindsGiftsInputDTO struct {
	CategoryID        *entities.UniqueEntityID
	PaginationOptions *entities.PaginationOptions
}

type FindsGiftsOutputDTO struct {
	Gifts      []entities.Gift
	TotalGifts int64
}
