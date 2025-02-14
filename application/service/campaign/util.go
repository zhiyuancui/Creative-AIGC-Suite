package campaign

import (
	"context"

	"code.byted.org/ad/creative_one_partner_business/domain/campaign/campaign_aggregate/entity"
	"code.byted.org/ad/creative_one_partner_business/kitex_gen/ad/creative_one/partner_business"
)

func ConvertCampaignInfoToNewEntity(ctx context.Context, campaignInfo *partner_business.CampaignInfo, AIOClientID int64) (campaignEntity entity.ICampaignEntity, err error) {
	return nil, nil
}

func ConvertCampaignEntityToDTO(ctx context.Context, campaignEntity entity.ICampaignEntity) (campaignInfo *partner_business.CampaignInfo, err error) {
	return nil, nil
}
