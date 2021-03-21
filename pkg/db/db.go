package db

import (
	"fmt"

	"github.com/favecode/reflect-core/pkg/setting"
	"github.com/go-pg/pg"
)

var db *pg.DB

func Connect(){ 
	db := pg.Connect(&pg.Options{
		Addr: setting.Get().Database.Host  + ":" + setting.Get().Database.Name,
		User: setting.Get().Database.User,
		Password: setting.Get().Database.Password, 
		Database: setting.Get().Database.Name, 
	})
	
	fmt.Println("Connected to",db.Options().Addr)
}

func GetDB() (*pg.DB) {
	return db
}