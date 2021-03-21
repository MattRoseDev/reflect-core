package db

import (
	"github.com/favecode/reflect-core/pkg/setting"
	"github.com/go-pg/pg"
)

func Connect() (*pg.DB){ 
	db := pg.Connect(&pg.Options{
		Addr: setting.Get().Database.Host  + ":" + setting.Get().Database.Name,
		User: setting.Get().Database.User,
		Password: setting.Get().Database.Password, 
		Database: setting.Get().Database.Name, 
	})

	return db
}