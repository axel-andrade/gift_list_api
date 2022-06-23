package markgiftuc

import (
	"errors"
	"go_gift_list_api/src/entities"
	ERROR "go_gift_list_api/src/shared/errors"
)

type MarkGiftInteractor struct {
	Gateway MarkGiftGateway
}

func BuildMarkGiftInteractor(g MarkGiftGateway) *MarkGiftInteractor {
	return &MarkGiftInteractor{Gateway: g}
}

func (bs *MarkGiftInteractor) Execute(input *MarkGiftInputDTO) error {
	var gift *entities.Gift
	var err error

	if gift, err = bs.Gateway.FindGiftByID(input.GiftID); err != nil {
		return err
	}

	if gift == nil {
		return errors.New(ERROR.GIFT_NOT_FOUND)
	}

	if !gift.Available {
		return errors.New(ERROR.GIFT_NOT_AVAILABLE)
	}

	var quantityThatWillRemain int64 = gift.Quantity - 1

	if quantityThatWillRemain < 0 {
		return errors.New(ERROR.GIFT_QUANTITY_NOT_AVAILABLE)
	}

	markedGiftToCreate := entities.MarkedGift{
		PersonName: input.PersonName,
		GiftID:     input.GiftID,
	}

	bs.Gateway.StartTransaction()

	if _, err = bs.Gateway.CreateMarkedGift(&markedGiftToCreate); err != nil {
		bs.Gateway.CancelTransaction()
		return err
	}

	giftToUpdate := gift
	giftToUpdate.Quantity = quantityThatWillRemain

	if quantityThatWillRemain == 0 {
		giftToUpdate.Available = false
	}

	if _, err = bs.Gateway.UpdateGift(giftToUpdate); err != nil {
		bs.Gateway.CancelTransaction()
		return err
	}

	bs.Gateway.CommitTransaction()

	return nil
}
