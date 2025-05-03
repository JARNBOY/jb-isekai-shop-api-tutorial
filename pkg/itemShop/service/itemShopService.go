package service

import(
	_itemShopModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/model"
)

type ItemShopService interface{
	Listing(itemFiler *_itemShopModel.ItemFilter) ([]*_itemShopModel.Item, error)
}