package daos

import (
	"database/sql"

	"github.com/stewie1520/elasticpmapi/daos/dao_user"
)

type Dao struct {
	db *sql.DB

	UserQueries *dao_user.Queries
}

func New(db *sql.DB) *Dao {
	return &Dao{
		db:          db,
		UserQueries: dao_user.New(db),
	}
}

func (dao *Dao) DB() *sql.DB {
	return dao.db
}
