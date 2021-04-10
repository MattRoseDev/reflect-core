package db

import (
	"fmt"

	"github.com/favecode/reflect-core/pkg/setting"
	"github.com/go-pg/pg"
)

var DB *pg.DB

func Connect() { 
	opt, err := pg.ParseURL(setting.Get().Database.URI)
	if err != nil {
 	  panic(err)
	}

	db := pg.Connect(opt)
	fmt.Println("Connected to",db.Options().Addr)
	DB = db
}
