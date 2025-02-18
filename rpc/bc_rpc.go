package rpc

import (
	"context"
)

func GetBusinessCenterListInfo(ctx context.Context, userId int64) ([]*i18n_bc_rpc.BusinessCenterUserInfo, error) {
	reqCtx := &rpc.BaseContext{
		UID:    &userId,
		Source: 4,
	}
	req := &rpc.GetBusinessCenterListReq{
		Context: reqCtx,
		Page:    1,
		Limit:   50,
	}
	logs.info(ctx, "GetBusinessCenterListReq req: %s", req)
	advResp, err := BcClient.GetBusinessCenterList(ctx, req)
	logs.info(ctx, "GetBusinessCenterListReq response: %+v", advResp)
	err = middleware.TransformError(advResp, err)
	if err != nil {
		logs.error(ctx, "[GetBusinessCenterListReq] get bc info error: %+v", err)
		return nil, errors.newError(err)
	}
	if len(advResp.BusinessCenterUserList) == 0 {
		logs.info(ctx, "uid [%d] doesn't have bc account", userId)
		return nil, nil
	}
	return advResp.BusinessCenterUserList, nil
}
