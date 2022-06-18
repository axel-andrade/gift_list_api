package find_gifts_interactor

import (
	"go_gift_list_api/src/entities"
)

type FindGiftsGateway interface {
	FindGiftsPaginate(pagination *entities.PaginationOptions) ([]entities.Gift, int64, error)
}

type FindsGiftsInputDTO struct {
	PaginationOptions *entities.PaginationOptions
}

type FindsGiftsOutputDTO struct {
	Gifts      []entities.Gift
	TotalGifts int64
}
