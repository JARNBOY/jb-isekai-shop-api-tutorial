package controller

import(
	_itemShopService "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/service"
)

type itemShopControllerImpl struct{
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(
	itemShopService _itemShopService.ItemShopService,
) itemShopController {
	return &itemShopControllerImpl{itemShopService}
}