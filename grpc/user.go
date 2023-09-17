package grpc

import (
	"context"
	"database/sql"

	"github.com/stewie1520/elasticpmapi/core"
	proto_user "github.com/stewie1520/elasticpmapi/proto/src"
	"github.com/stewie1520/elasticpmapi/usecases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServiceServer struct {
	*proto_user.UnimplementedUserServiceServer
	app core.App
}

// GetUserInfo implements proto_user.UserServiceServer.
func (srv *UserServiceServer) GetUserInfo(ctx context.Context, req *proto_user.GetUserInfoRequest) (*proto_user.GetUserInfoResponse, error) {
	q := usecases.NewGetUserByAccountIDQuery(srv.app)
	q.AccountID = req.AccountId

	user, err := q.Execute()

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto_user.GetUserInfoResponse{
		Id:        user.ID,
		FullName:  user.FullName,
		AccountId: user.AccountId,
		CreatedAt: timestamppb.New(user.CreatedAt.Time()),
		UpdatedAt: timestamppb.New(user.UpdatedAt.Time()),
	}, nil
}

func newUserServer(app core.App) *UserServiceServer {
	return &UserServiceServer{
		app: app,
	}
}

func NewUserServer(app core.App) (*grpc.Server, error) {
	gsrv := grpc.NewServer()
	srv := newUserServer(app)
	proto_user.RegisterUserServiceServer(gsrv, srv)

	return gsrv, nil
}
