// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: user.query.sql

package dao_user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/stewie1520/elasticpmapi/tools/types"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "users" ("id", "fullName", "accountId", "createdAt", "updatedAt") VALUES ($1, $2, $3, $4, $5) RETURNING id, "fullName", "accountId", "createdAt", "updatedAt"
`

type CreateUserParams struct {
	ID        uuid.UUID      `json:"id"`
	FullName  sql.NullString `json:"fullName"`
	AccountId string         `json:"accountId"`
	CreatedAt types.DateTime `json:"createdAt"`
	UpdatedAt types.DateTime `json:"updatedAt"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.FullName,
		arg.AccountId,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.AccountId,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT id, "fullName", "accountId", "createdAt", "updatedAt" FROM "users" WHERE "users"."id" = $1 LIMIT 1
`

func (q *Queries) FindUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, findUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.AccountId,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByAccountID = `-- name: GetUserByAccountID :one
SELECT id, "fullName", "accountId", "createdAt", "updatedAt" FROM "users" WHERE "users"."accountId" = $1 LIMIT 1
`

func (q *Queries) GetUserByAccountID(ctx context.Context, accountid string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByAccountID, accountid)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.AccountId,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
