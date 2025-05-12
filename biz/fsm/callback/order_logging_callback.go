package callback

import (
	"context"
)

func CreateTTCXOrderLog(ctx context.Context, campaignEntity entity.ICampaignEntity) error {
	return order_dal.CreateOrderLogModel(ctx, campaignEntity.GetTTCXOrderID(), campaignEntity.GetOperatorID(), campaignEntity.GetUserRole(), campaignEntity.GetTTCXOperationType(), campaignEntity.GetPreviousStatus(), int(campaignEntity.GetCampaignStatus()), campaignEntity.GetContent(), model.CustomInfo{}, campaignEntity.GetIsTest())
}
