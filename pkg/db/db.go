package db

import (
	"fmt"

	"github.com/favecode/reflect-core/pkg/setting"
	"github.com/go-pg/pg"
)

func Connect() (*pg.DB){ 
	opt, err := pg.ParseURL(setting.Get().Database.URI)
	if err != nil {
 	  panic(err)
	}

	db := pg.Connect(opt)
	fmt.Println("Connected to",db.Options().Addr)
	return db
}
