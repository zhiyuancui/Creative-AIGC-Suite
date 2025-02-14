package service

import (
	"code.byted.org/ad/creative_one_partner_business/domain/campaign/campaign_aggregate/entity"
	"context"
)

type ICampaignDraftService interface {
	CreateCampaignDraft(ctx context.Context, campaignDraftEntity entity.ICampaignEntity) (err error)
	UpdateCampaignDraft(ctx context.Context, campaignDraftEntity entity.ICampaignEntity) (err error)
	GetCampaignDraftDetail(ctx context.Context, campaignDraftID int64) (campaignDraftEntity entity.ICampaignEntity, err error)
}

func GetService() ICampaignDraftService {
	return campaignDraftDomainService
}
