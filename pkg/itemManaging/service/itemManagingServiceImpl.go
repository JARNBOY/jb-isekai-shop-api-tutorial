package service

import (
	_itemManagingRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/repository"
)

type itemManagingServiceImpl struct {
	ItemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(
	ItemManagingRepository _itemManagingRepository.ItemManagingRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{
		ItemManagingRepository: ItemManagingRepository,
	}
}