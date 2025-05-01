package repository

import "github.com/JARNBOY/jb-isekai-shop-tutorial/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}