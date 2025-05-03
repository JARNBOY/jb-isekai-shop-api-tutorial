package repository

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/entities"
	_itemShopModel "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFiler *_itemShopModel.ItemFilter) ([]*entities.Item, error)
}