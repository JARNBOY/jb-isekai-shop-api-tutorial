package repository

import "github.com/JARNBOY/jb-isekai-shop-tutorial/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindByID(playerID string) (*entities.Player, error)
}

