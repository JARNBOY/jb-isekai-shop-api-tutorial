package service

import(
	_itemShopModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/model"
)

type ItemShopService interface{
	Listing() ([]*_itemShopModel.Item, error)
}