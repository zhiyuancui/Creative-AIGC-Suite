package callback

func ApproveCallbackAfterTran(transitionCtx *bfsm.TransitionCtx) {
	ctx := transitionCtx.Ctx
	args := transitionCtx.Args
	campaignEntity := args[1].(entity.ICampaignEntity)

	if err := rpc.CreateChargingTicketForPartnerCampaign(ctx, campaignEntity); err != nil {
		transitionCtx.Cancel(err)
		logs(ctx, "CallbackAfterTran error=%v partnerCampaign=%+v", err, campaignEntity)
	}

	err := NotifyClientCampaignCompleted(ctx, campaignEntity)
	if err != nil {
		logs(ctx, "fail to send tto email notification for %d with err: %+v", campaignEntity.GetCampaignID(), err)
	}
}
