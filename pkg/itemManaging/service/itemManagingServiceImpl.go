package service

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/entities"
	_itemManagingModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/model"
	_itemManagingRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/repository"
	_itemShopModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/model"
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

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	itemEntity := &entities.Item{
		Name: itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture: itemCreatingReq.Picture,
		Price: itemCreatingReq.Price,
	}

	itemEntityResult, err := s.ItemManagingRepository.Creating(itemEntity)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}