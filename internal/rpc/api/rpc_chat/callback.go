package rpc_chat

import (
	"suzaku/pkg/common/config"
	"suzaku/pkg/constant"
	"suzaku/pkg/http"
	"suzaku/pkg/model/callback"
	pb_chat "suzaku/pkg/proto/chart"
)

func copyCallbackCommonReq(msg *pb_chat.SendMsgReq) callback.CommonCallbackReq {
	return callback.CommonCallbackReq{
		SendID:           msg.MsgData.SendId,
		ServerMsgID:      msg.MsgData.ServerMsgId,
		ClientMsgID:      msg.MsgData.ClientMsgId,
		OperationID:      msg.OperationId,
		SenderPlatformID: msg.MsgData.SenderPlatformId,
		SenderNickname:   msg.MsgData.SenderNickname,
		SessionType:      msg.MsgData.SessionType,
		MsgFrom:          msg.MsgData.MsgFrom,
		ContentType:      msg.MsgData.ContentType,
		Status:           msg.MsgData.Status,
		CreateTime:       msg.MsgData.CreateTime,
		Content:          string(msg.MsgData.Content),
	}
}

func callbackWordFilter(msg *pb_chat.SendMsgReq) (canSend bool, err error) {
	var (
		commonCallbackReq callback.CommonCallbackReq
		req               callback.CallbackWordFilterReq
		resp              *callback.CallbackWordFilterResp
	)
	if !config.Config.Callback.CallbackWordFilter.Enable || msg.MsgData.ContentType != constant.Text {
		return
	}

	commonCallbackReq = copyCallbackCommonReq(msg)
	req = callback.CallbackWordFilterReq{
		CommonCallbackReq: commonCallbackReq,
	}
	resp = &callback.CallbackWordFilterResp{CommonCallbackResp: callback.CommonCallbackResp{}}

	if err = http.PostReturn(config.Config.Callback.CallbackURL, req, resp, config.Config.Callback.CallbackBeforeSendSingleMsg.CallbackTimeOut); err != nil {
		if !config.Config.Callback.CallbackBeforeSendSingleMsg.CallbackFailedContinue {
			canSend = false
			return
		}
		canSend = true
		return
	} else {
		if resp.ActionCode == constant.ActionForbidden && resp.ErrCode == constant.CallbackHandleSuccess {
			canSend = true
			return
		}
		if resp.ErrCode == constant.CallbackHandleSuccess {
			msg.MsgData.Content = []byte(resp.Content)
		}
	}
	canSend = true
	return
}

func callbackBeforeSendSingleMsg(msg *pb_chat.SendMsgReq) (canSend bool, err error) {
	if !config.Config.Callback.CallbackBeforeSendSingleMsg.Enable {
		canSend = true
		return
	}

	var (
		commonCallbackReq callback.CommonCallbackReq
		req               callback.CallbackBeforeSendSingleMsgReq
		resp              *callback.CallbackBeforeSendSingleMsgResp
	)
	if !config.Config.Callback.CallbackWordFilter.Enable || msg.MsgData.ContentType != constant.Text {
		return
	}

	commonCallbackReq = copyCallbackCommonReq(msg)
	commonCallbackReq.CallbackCommand = constant.CallbackBeforeSendSingleMsgCommand

	req = callback.CallbackBeforeSendSingleMsgReq{
		CommonCallbackReq: commonCallbackReq,
		RecvID:            msg.MsgData.RecvId,
	}
	resp = &callback.CallbackBeforeSendSingleMsgResp{
		CommonCallbackResp: callback.CommonCallbackResp{},
	}

	if err = http.PostReturn(config.Config.Callback.CallbackURL, req, resp, config.Config.Callback.CallbackBeforeSendSingleMsg.CallbackTimeOut); err != nil {
		if !config.Config.Callback.CallbackBeforeSendSingleMsg.CallbackFailedContinue {
			canSend = false
			return
		}
		canSend = true
		return
	} else {
		if resp.ActionCode == constant.ActionForbidden && resp.ErrCode == constant.CallbackHandleSuccess {
			canSend = true
			return
		}
	}
	return
}

func callbackAfterSendSingleMsg(msg *pb_chat.SendMsgReq) error {
	if !config.Config.Callback.CallbackAfterSendSingleMsg.Enable {
		return nil
	}
	commonCallbackReq := copyCallbackCommonReq(msg)
	commonCallbackReq.CallbackCommand = constant.CallbackAfterSendSingleMsgCommand
	req := callback.CallbackAfterSendSingleMsgReq{
		CommonCallbackReq: commonCallbackReq,
		RecvID:            msg.MsgData.RecvId,
	}
	resp := &callback.CallbackAfterSendSingleMsgResp{CommonCallbackResp: callback.CommonCallbackResp{}}

	if err := http.PostReturn(config.Config.Callback.CallbackURL, req, resp, config.Config.Callback.CallbackAfterSendSingleMsg.CallbackTimeOut); err != nil {
		return err
	}
	return nil
}
