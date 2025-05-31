package repository

import "github.com/JARNBOY/jb-isekai-shop-tutorial/entities"

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (*entities.Admin, error)
	FindByID(adminID string) (*entities.Admin, error)
}

