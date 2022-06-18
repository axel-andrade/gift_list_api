package repositories

import (
	"go_gift_list_api/src/entities"
)

type GiftRepository interface {
	BaseRepository
	FindGiftByID(id entities.UniqueEntityID) (*entities.Gift, error)
	FindGiftsPaginate(filters *entities.PaginationOptions) ([]entities.Gift, int64, error)
}
