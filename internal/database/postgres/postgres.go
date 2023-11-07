package postgres

import (


	"github.com/jmoiron/sqlx"
)

type UrlModel struct {
	DB *sqlx.DB
}