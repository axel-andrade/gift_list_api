package markgiftuc

import (
	"go_gift_list_api/src/entities"
	common_interactor "go_gift_list_api/src/usecases/common"
)

type MarkGiftGateway interface {
	common_interactor.DefaultInteractorGateway
	CheckExistsCategory(categoryID entities.UniqueEntityID) bool
	CreateMarkedGift(markedGift *entities.MarkedGift) (*entities.MarkedGift, error)
	FindGiftByID(giftID entities.UniqueEntityID) (*entities.Gift, error)
	UpdateGift(gift *entities.Gift) (*entities.Gift, error)
}

type MarkGiftInputDTO struct {
	CategoryID entities.UniqueEntityID `json:"category_id"`
	GiftID     entities.UniqueEntityID `json:"gift_id"`
	PersonName string                  `json:"person_name"`
	Quantity   int64                   `json:"quantity"`
}

type MarkGiftOutputDTO struct{}
