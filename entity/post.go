package entity

import "time"

type Post struct {
	tableName struct{}   `sql:"post"`
	Id        string     `pg:"id"`
	UserId string     `pg:"userId"`
	Content string       `pg:"content"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt *time.Time `pg:"deleted_at"`
}
