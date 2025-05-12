package callback

func RollbackCoupon(transitionCtx *bfsm.TransitionCtx) {
	ctx := transitionCtx.Ctx
	args := transitionCtx.Args
	campaignEntity := args[1].(entity.ICampaignEntity)

	payment := campaignEntity.GetCampaignPayment()
	if payment == nil {
		logs.CtxError(ctx, "missing payment information for campaign %d", campaignEntity.GetCampaignID())
		return
	}

	err := rpc.RollbackCoupon(ctx, payment.GetCouponID(), 1, campaignEntity.GetOperatorID())
	if err != nil {
		logs.CtxError(ctx, "rollback coupon %d with error: %v", payment.GetCouponID(), err)
	}
}
