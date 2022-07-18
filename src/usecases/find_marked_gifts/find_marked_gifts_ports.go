package find_marked_gifts_interactor

import (
	"go_gift_list_api/src/entities"
)

type FindMarkedGiftsGateway interface {
	FindMarkedGiftsPaginate(pagination *entities.PaginationOptions) ([]entities.MarkedGift, int64, error)
}

type FindMarkedGiftsInputDTO struct {
	PaginationOptions *entities.PaginationOptions
}

type FindMarkedGiftsOutputDTO struct {
	MarkedGifts      []entities.MarkedGift
	TotalMarkedGifts int64
}
