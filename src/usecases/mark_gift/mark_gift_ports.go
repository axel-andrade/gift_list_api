package markgiftuc

import (
	"go_gift_list_api/src/entities"
	common_interactor "go_gift_list_api/src/usecases/common"
)

type MarkGiftGateway interface {
	common_interactor.DefaultInteractorGateway
	CreateMarkedGift(markedGift *entities.MarkedGift) (*entities.MarkedGift, error)
	FindGiftByID(giftID entities.UniqueEntityID) (*entities.Gift, error)
	UpdateGift(gift *entities.Gift) (*entities.Gift, error)
}

type MarkGiftInputDTO struct {
	GiftID     entities.UniqueEntityID `json:"gift_id"`
	PersonName string                  `json:"person_name"`
}

type MarkGiftOutputDTO struct{}
