package db

import (
	"fmt"

	"github.com/favecode/reflect-core/pkg/setting"
	"github.com/go-pg/pg"
)

func Connect() (*pg.DB){ 
	db := pg.Connect(&pg.Options{
		Addr: setting.Get().Database.Host  + ":" + fmt.Sprint(setting.Get().Database.Port),
		User: setting.Get().Database.User,
		Password: setting.Get().Database.Password, 
		Database: setting.Get().Database.Name,
	})
	fmt.Println("Connected to",db.Options().Addr)
	return db
}
