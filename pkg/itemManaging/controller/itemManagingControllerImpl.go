package controller

import (
	_itemManagingService "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/service"
)

type itemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(
	itemManagingService  _itemManagingService.ItemManagingService,
) itemManagingController {
	return &itemManagingControllerImpl{
		itemManagingService: itemManagingService,
	}
}