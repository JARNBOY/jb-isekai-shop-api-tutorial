package service

import (
	_itemShopRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/repository"
	_itemShopModel 		"github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/model"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopRepositoryImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing(itemFiler *_itemShopModel.ItemFilter) ([]*_itemShopModel.Item, error) {
	itemList, err := s.itemShopRepository.Listing(itemFiler)
	if err != nil {
		return nil, err
	}

	itemModelList := make ([]*_itemShopModel.Item, 0)
	for _, item := range itemList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return itemModelList, nil
}