package repository

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/databases"
	"github.com/JARNBOY/jb-isekai-shop-tutorial/entities"
	"github.com/labstack/echo/v4"
	_adminException "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/admin/exception"
)



type AdminRepositoryImpl struct {
	db databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(
	db databases.Database, 
	logger echo.Logger,
) AdminRepository {
	return &AdminRepositoryImpl{
		db: db,
		logger: logger,
	}
}

func (r *AdminRepositoryImpl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.Connect().Create(adminEntity).Scan(admin).Error; err != nil {
		r.logger.Error("Creating Admin Fail: %s", err.Error())
		return nil, &_adminException.AdminCreating{AdminID: adminEntity.ID}
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)
	
	if err := r.db.Connect().Where("id = ?", adminID).First(admin).Error; err != nil {
		r.logger.Errorf("Failed to find admin by ID: %s", err.Error())
		return nil, &_adminException.AdminNotFound{AdminID: adminID}
	}

	return admin, nil
}