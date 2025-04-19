package databases

import "gorm.io/gorm"

type Database interface {
	ConnectingGetting() *gorm.DB
}