package trigger

import (
	"context"
)

func TransferCampaignWorkFlow(ctx context.Context, agentEntity entity.IAgentEntity, matchParams map[string]interface{}, params ...interface{}) (valid bool, err error) {
	fsm, err := factory.NewPartnerCampaignBFSM(campaignEntity)
	if err != nil {
		return
	}
	callbackParams := []interface{}{ctx, campaignEntity}
	callbackParams = append(callbackParams, params...)
	err = fsm.Fire(ctx, campaignEntity.GetCampaignOperationEvent().String(), matchParams, callbackParams...)
	if err != nil {
		return
	}
	valid = true
	return
}
