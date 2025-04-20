package main

import (
	"github.com/JARNBOY/jb-isekai-shop-tutorial/config"
	"github.com/JARNBOY/jb-isekai-shop-tutorial/databases"
	"github.com/JARNBOY/jb-isekai-shop-tutorial/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectingGetting())

	server.Start()
}