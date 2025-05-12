package callback

import (
	"context"
)

func ClientCampaignCompleted(ctx context.Context, campaignEntity entity.ICampaignEntity) error {
	//Email Notification
	placeHolder := eventbus.TTONotificationPlaceHolder{
		ClientName:  gptr.Of(campaignEntity.GetDemanderName()),
		ProjectName: gptr.Of(campaignEntity.GetName()),
		CampaignID:  gptr.Of(campaignEntity.GetCampaignID()),
		PartnerName: gptr.Of(campaignEntity.GetPartnerName()),
	}

	err := eventbus.SendEmailNotification(ctx, eventbus.REVIEW_PARTNER_CLIENT_EMAIL_NOTIFICATION, campaignEntity.GetCampaignID(), campaignEntity.GetAIOClientID(), placeHolder)
	if err != nil {
		logs.CtxError(ctx, "fail to send tto email notification for %d with err: %+v", campaignEntity.GetCampaignID(), err)
		return err
	}
	return nil
}

func ClientCampaignCompletedCallBack(transitionCtx *bfsm.TransitionCtx) {
	ctx := transitionCtx.Ctx
	args := transitionCtx.Args
	campaignEntity := args[1].(entity.ICampaignEntity)

	err := NotifyClientCampaignCompleted(ctx, campaignEntity)
	if err != nil {
		logs.CtxError(ctx, "fail to send tto email notification for %d with err: %+v", campaignEntity.GetCampaignID(), err)
	}
}

func ClientPartnerCompleteCampaign(transitionCtx *bfsm.TransitionCtx) {
	ctx := transitionCtx.Ctx
	args := transitionCtx.Args
	campaignEntity := args[1].(entity.ICampaignEntity)

	//Email Notification
	placeHolder := eventbus.TTONotificationPlaceHolder{
		ClientName:  gptr.Of(campaignEntity.GetDemanderName()),
		ProjectName: gptr.Of(campaignEntity.GetName()),
		CampaignID:  gptr.Of(campaignEntity.GetCampaignID()),
		PartnerName: gptr.Of(campaignEntity.GetPartnerName()),
	}
	err := eventbus.SendEmailNotification(ctx, eventbus.PARTNER_COMPLETE_CAMPAIGN_CLIENT_EMAIL_NOTIFICATION, campaignEntity.GetCampaignID(), campaignEntity.GetAIOClientID(), placeHolder)
	if err != nil {
		logs.CtxError(ctx, "fail to send tto email notification for %d with err: %+v", campaignEntity.GetCampaignID(), err)
	}
}

func ClientPartnerAcceptInvite(transitionCtx *bfsm.TransitionCtx) {
	ctx := transitionCtx.Ctx
	args := transitionCtx.Args
	campaignEntity := args[1].(entity.ICampaignEntity)

	//Email Notification
	placeHolder := eventbus.TTONotificationPlaceHolder{
		ClientName:  gptr.Of(campaignEntity.GetDemanderName()),
		ProjectName: gptr.Of(campaignEntity.GetName()),
		CampaignID:  gptr.Of(campaignEntity.GetCampaignID()),
		PartnerName: gptr.Of(campaignEntity.GetPartnerName()),
	}

	err := eventbus.SendEmailNotification(ctx, eventbus.PARTNER_ACCEPT_ASSIGN_CLIENT_EMAIL_NOTIFICATION, campaignEntity.GetCampaignID(), campaignEntity.GetAIOClientID(), placeHolder)
	if err != nil {
		logs.CtxError(ctx, "fail to send tto email notification for %d with err: %+v", campaignEntity.GetCampaignID(), err)
	}
}

func NotifyClientPartnerRejectInvite(transitionCtx *bfsm.TransitionCtx) {
	ctx := transitionCtx.Ctx
	args := transitionCtx.Args
	campaignEntity := args[1].(entity.ICampaignEntity)

	//Email Notification
	placeHolder := eventbus.TTONotificationPlaceHolder{
		ClientName:  gptr.Of(campaignEntity.GetDemanderName()),
		ProjectName: gptr.Of(campaignEntity.GetName()),
		CampaignID:  gptr.Of(campaignEntity.GetCampaignID()),
		PartnerName: gptr.Of(campaignEntity.GetPartnerName()),
	}

	err := eventbus.SendEmailNotification(ctx, eventbus.PARTNER_DENY_ASSIGN_CLIENT_EMAIL_NOTIFICATION, campaignEntity.GetCampaignID(), campaignEntity.GetAIOClientID(), placeHolder)
	if err != nil {
		logs.CtxError(ctx, "fail to send tto email notification for %d with err: %+v", campaignEntity.GetCampaignID(), err)
	}
}
