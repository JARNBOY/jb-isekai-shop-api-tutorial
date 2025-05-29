package service

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/entities"
	_itemManagingModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/model"
	_itemManagingRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemManaging/repository"
	_itemShopModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/model"
	_itemShopRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/repository"
)

type itemManagingServiceImpl struct {
	ItemManagingRepository _itemManagingRepository.ItemManagingRepository
	ItemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemManagingServiceImpl(
	ItemManagingRepository _itemManagingRepository.ItemManagingRepository,
	ItemShopRepository _itemShopRepository.ItemShopRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{
		ItemManagingRepository: ItemManagingRepository,
		ItemShopRepository: ItemShopRepository,
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

func (s *itemManagingServiceImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error) {
	itemIdValidEditing, err := s.ItemManagingRepository.Editing(itemID, itemEditingReq)
	if err != nil {
		return nil, err
	}

	itemEntityResult, err := s.ItemShopRepository.FindById(itemIdValidEditing)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}