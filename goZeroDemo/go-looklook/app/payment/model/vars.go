package model

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Executable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var ErrNotFound = sqlx.ErrNotFound
