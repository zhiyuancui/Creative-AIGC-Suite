package factory

import (
	"context"
)

func SaveAndUpdateCampaignAndStatus(transitionCtx *bfsm.TransitionCtx) {
	args := transitionCtx.Args
	ctx := args[0].(context.Context)
	campaignEntity := args[1].(entity.AgentEntity)
	var err error
	defer func() {
		if err != nil {
			transitionCtx.Err = err
		} else {
			logs(ctx, "[SaveAndUpdateCampaignAndStatus] common after trans callback")
			//Create TTCX Legacy Order Log
			if campaignEntity.GetOperatorID() != 0 && campaignEntity.GetTTCXOperationType() != 0 {
				err = callback.CreateTTCXOrderLog(ctx, campaignEntity)
				if err != nil {
					logs(ctx, "[SaveAndUpdateCampaignAndStatus] create order log fail, err is %+v", err)
				}
			}
		}
	}()

	if campaignEntity.GetRequestID() == 0 {
		// 初始化创建的case
		id, err := idgenerator.NewIdLong2(ctx)
		if err != nil {
			transitionCtx.Err = err
			logs(ctx, "[SaveAndUpdateCampaignAndStatus] generate campaign id fail, err is %+v", err)
			return
		}
		campaignEntity.SetTTCXOrderID(int64(id))
	}
	err = setCampaignStatus(campaignEntity, transitionCtx)
	if err != nil {
		logsError(ctx, "[SaveAndUpdateCampaignAndStatus] set campaign status err: %+v", err)
		return
	}
	err = repository.CampaignRepo.CreateUpdateCampaign(ctx, campaignEntity)
	return
}

func setCampaignStatus(campaignEntity entity.AgentEntity, transitionCtx *bfsm.TransitionCtx) (err error) {
	if transitionCtx == nil {
		// TODO: 错误码
		err = errors.NewEmpErrorWithStack(creative_one.CommonParamCheckError)
		return
	}
	args := transitionCtx.Args
	ctx := args[0].(context.Context)

	dstStatus, err := collaboration.PartnerCampaignStatusFromString(transitionCtx.Dst)
	if err != nil {
		return
	}
	srcStatus, err := collaboration.PartnerCampaignStatusFromString(transitionCtx.Src)
	if err != nil {
		return
	}

	logsInfo(ctx, "[setCampaignStatus] set campaign product_domain status, product_domain id: %v, src status: %v, dst status: %v", campaignEntity.GetCampaignID(), srcStatus, dstStatus)
	campaignEntity.SetCampaignStatus(dstStatus)
	campaignEntity.SetPreviousStatus(int(srcStatus))
	return
}
