package service

import (
	"code.byted.org/ad/creative_one_partner_business/domain/campaign/campaign_aggregate/entity"
	"context"
)

type campaignDraftServiceImpl struct{}

func (c *campaignDraftServiceImpl) CreateCampaignDraft(ctx context.Context, campaignDraftEntity entity.ICampaignEntity) (err error) {
	return
}

func (c *campaignDraftServiceImpl) UpdateCampaignDraft(ctx context.Context, campaignDraftEntity entity.ICampaignEntity) (err error) {
	return
}

func (c *campaignDraftServiceImpl) GetCampaignDraftDetail(ctx context.Context, campaignDraftID int64) (campaignDraftEntity entity.ICampaignEntity, err error) {
	return
}
