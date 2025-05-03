package service

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/entities"
	_itemShopModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/model"
	_itemShopRepository "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopRepositoryImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing(itemFiler *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFiler)
	if err != nil {
		return nil, err
	}

	itemCounting, err := s.itemShopRepository.Counting(itemFiler)
	if err != nil {
		return nil, err
	}

	totalPage := s.totalPageCalculation(itemCounting, itemFiler.Paginate.Size)

	result := s.toItemResultResponse(itemList, itemFiler.Paginate.Page, totalPage)

	return result, nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems int64, size int64) int64 {
	totolPage := totalItems / size 
	if totalItems%size != 0 {
		totolPage++
	}
	return totolPage
}

func (s *itemShopServiceImpl) toItemResultResponse(itemEntityList []*entities.Item, page int64, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make ([]*_itemShopModel.Item,0)
	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page: page,
			TotalPage: totalPage,
		},
	}
}