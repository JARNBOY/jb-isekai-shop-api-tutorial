package repository

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/databases"
	"github.com/JARNBOY/jb-isekai-shop-tutorial/entities"
	"github.com/labstack/echo/v4"
	_playerException "github.com/JARNBOY/jb-isekai-shop-tutorial/pkg/player/exception"
)

type PlayerRepositoryImpl struct {
	db databases.Database
	logger echo.Logger
}

func NewPlayerRepositoryImpl(
	db databases.Database, 
	Logger echo.Logger,
) PlayerRepository {
	return &PlayerRepositoryImpl{
		db: db,
		logger: Logger,
	}
}

func (r *PlayerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Create(playerEntity).Scan(player).Error; err != nil {
		r.logger.Error("Creating player Fail: %s", err.Error())
		return nil, &_playerException.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return player, nil
}

func (r *PlayerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)
	
	if err := r.db.Connect().Where("id = ?", playerID).First(player).Error; err != nil {
		r.logger.Errorf("Failed to find player by ID: %s", err.Error())
		return nil, &_playerException.PlayerNotFound{PlayerID: playerID}
	}

	return player, nil
}