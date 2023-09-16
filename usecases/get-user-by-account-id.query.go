package usecases

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/stewie1520/elasticpmapi/core"
	"github.com/stewie1520/elasticpmapi/daos/dao_user"
	"github.com/stewie1520/elasticpmapi/tools/types"
)

var _ Query[GetUserByAccountIDResponse] = (*GetUserByAccountIDQuery)(nil)

func NewGetUserByAccountIDQuery(app core.App) *GetUserByAccountIDQuery {
	return &GetUserByAccountIDQuery{
		app: app,
		dao: app.Dao().User,
	}
}

type GetUserByAccountIDQuery struct {
	app core.App
	dao *dao_user.Queries

	AccountID string `json:"accountId"`
}

type GetUserByAccountIDResponse struct {
	ID        string         `json:"id"`
	AccountId string         `json:"accountId"`
	CreatedAt types.DateTime `json:"createdAt"`
	UpdatedAt types.DateTime `json:"updatedAt"`
	FullName  string         `json:"fullName"`
}

// Execute implements Query.
func (q *GetUserByAccountIDQuery) Execute() (GetUserByAccountIDResponse, error) {
	if err := q.Validate(); err != nil {
		return GetUserByAccountIDResponse{}, err
	}

	user, err := q.dao.GetUserByAccountID(context.Background(), q.AccountID)
	return GetUserByAccountIDResponse{
		ID:        user.ID.String(),
		AccountId: user.AccountId,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FullName:  user.FullName.String,
	}, err
}

// Validate implements Query.
func (q *GetUserByAccountIDQuery) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(&q.AccountID, validation.Required, is.UUIDv4),
	)
}
