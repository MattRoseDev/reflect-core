package entity

import "time"

type Password struct {
	tableName struct{}   `sql:"password"`
	Id        string     `pg:"id"`
	UserId string     `pg:"userId"`
	Password string       `pg:"password"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	DeletedAt *time.Time `pg:"deleted_at"`

}
