package rpc_user

import (
	"context"
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	pb_user "suzaku/pkg/proto/user"
	pb_com "suzaku/pkg/proto/pb_com"
	"suzaku/pkg/common/config"
)

type userRpc struct {
	pb_user.UnimplementedUserServer
	rpc_category.Rpc
}

func NewRpcUserServer(port int) *userRpc {
	return &userRpc{
		Rpc: rpc_category.NewRpcServer(port,config.Config.RPCRegisterName.UserName),
	}
}

func (rpc *userRpc) Run() {
	var (
		server *grpc.Server
	)
	server = grpc.NewServer()
	pb_user.RegisterUserServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *userRpc) UserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, err error) {
	var (
		common = &pb_com.CommonResp{}
	)
	resp = &pb_user.UserInfoResp{Common: common}
	resp.UserId = req.UserId
	return
}