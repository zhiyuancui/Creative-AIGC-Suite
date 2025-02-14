package repository

import (
	"code.byted.org/ad/creative_one_partner_business/domain/campaign/campaign_aggregate/entity"
	"context"
)

var (
	BriefRepo BriefRepository
)

type BriefRepository interface {
	CreateUpdateCampaignBrief(ctx context.Context, brief entity.ICampaignEntity) (err error)
	GetCampaignBriefDetail(ctx context.Context, brief entity.ICampaignEntity) (err error)
}
