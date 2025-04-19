package service

import (
	_itemShopRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/repository"
)

type ItemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopRepositoryImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemShopService {
	return &ItemShopServiceImpl{itemShopRepository}
}