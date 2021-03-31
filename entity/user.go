package entity

import "time"

type User struct {
	tableName struct{}   `sql:"user"`
	Id        string     `pg:"id"`
	Username  string     `pg:"username"`
	Email     string     `pg:"email"`
	Fullname  *string    `pg:"fullname"`
	Bio       *string    `pg:"bio"`
	Admin     bool       `pg:"admin"`
	CreatedAt *time.Time `pg:"created_at"`
	UpdatedAt *time.Time `pg:"updated_at"`
	DeletedAt *time.Time `pg:"deleted_at"`
}
