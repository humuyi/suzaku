package api_auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"suzaku/internal/dto/dto_api"
	"suzaku/internal/rpc/grpc_client"
	"suzaku/pkg/common/config"
	"suzaku/pkg/constant"
	"suzaku/pkg/http"
	pb_auth "suzaku/pkg/proto/auth"
	"suzaku/pkg/utils"
)

func UserRegister(c *gin.Context) {
	var (
		params     dto_api.UserRegisterReq
		err        error
		clientConn *grpc.ClientConn
		client     pb_auth.AuthClient
		req        *pb_auth.UserRegisterReq
		reply      *pb_auth.UserRegisterResp
		tokenReq   *pb_auth.UserTokenReq
		replyToken *pb_auth.UserTokenResp
		resp       *dto_api.UserRegisterResp
	)
	if err = c.BindJSON(&params); err != nil {
		http.Error(c, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}
	if params.Secret != config.Config.Secret {
		http.Error(c, err, http.ErrorCodeHttpReqNotAuthorized)
		return
	}
	req = &pb_auth.UserRegisterReq{}
	utils.CopyStructFields(req, params)
	//clientConn = getcdv3.GetConn(config.Config.Etcd.Schema, strings.Join(config.Config.Etcd.Address, ","), config.Config.RPCRegisterName.AuthName)
	clientConn = grpc_client.ClientConn(config.Config.RPCRegisterName.AuthName)
	client = pb_auth.NewAuthClient(clientConn)
	reply, _ = client.UserRegister(context.Background(), req)
	if reply.Common != nil && reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	tokenReq = &pb_auth.UserTokenReq{}
	utils.CopyStructFields(tokenReq, reply)
	replyToken, _ = client.UserToken(context.Background(), tokenReq)
	if replyToken.Common != nil && replyToken.Common.Code > 0 {
		http.Err(c, replyToken.Common.Msg, replyToken.Common.Code)
		return
	}
	resp = &dto_api.UserRegisterResp{
		PlatformId: reply.PlatformId,
		UserId:     reply.UserId,
		Token:      constant.HttpKeyJwt + replyToken.Token,
		Expire:     replyToken.Expire,
	}
	http.Success(c, resp)
}

func UserToken(c *gin.Context) {

}
