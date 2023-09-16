package daos

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stewie1520/elasticpmapi/daos/dao_user"
)

type Dao struct {
	db *pgxpool.Pool

	Builder squirrel.StatementBuilderType
	User    *dao_user.Queries
}

func New(db *pgxpool.Pool) *Dao {
	return &Dao{
		db:      db,
		User:    dao_user.New(db),
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (dao *Dao) DB() *pgxpool.Pool {
	return dao.db
}
