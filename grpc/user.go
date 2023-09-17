package grpc

import (
	"github.com/stewie1520/elasticpmapi/core"
	proto_user "github.com/stewie1520/elasticpmapi/proto/src"
	"google.golang.org/grpc"
)

type UserServiceServer struct {
	*proto_user.UnimplementedUserServiceServer
	app core.App
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
