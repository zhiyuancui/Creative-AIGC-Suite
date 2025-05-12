package custom

import (
	"context"
	"database/sql/driver"

	"gorm.io/gen"
)

func CreateUpdatePartnerCampaign(ctx context.Context, campaignModel *model.CollaborationPartnerCampaign) (err error) {
	if campaignModel.ID == 0 {
		opportunityID, err := idgenerator.NewIdLong2(ctx)
		if err != nil {
			logs.CtxError(ctx, "[CreateUpdatePartnerCampaign] id generate fail, err is %+v", err)
			return err
		}
		campaignModel.ID = int64(opportunityID)
	}
	return query.CreativeSession(ctx).CollaborationPartnerCampaign.WithContext(ctx).Save(campaignModel)
}

func BatchCreateUpdatePartnerCampaign(ctx context.Context, campaignModelList []*model.CollaborationPartnerCampaign) (err error) {
	for _, campaignModel := range campaignModelList {
		if campaignModel.ID == 0 {
			partnerCampaignID, err := idgenerator.NewIdLong2(ctx)
			if err != nil {
				logs.CtxError(ctx, "[CreateUpdatePartnerCampaign] id generate fail, err is %+v", err)
				return err
			}
			campaignModel.ID = int64(partnerCampaignID)
		}
	}
	return query.CreativeSession(ctx).CollaborationPartnerCampaign.WithContext(ctx).Save(campaignModelList...)
}

func GetPartnerCampaignByCampaignID(ctx context.Context, campaignID int64) (campaign *model.CollaborationPartnerCampaign, err error) {
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	return q.WithContext(ctx).Where(q.CampaignID.Eq(campaignID)).First()
}

func GetPartnerCampaignByID(ctx context.Context, id int64) (campaign *model.CollaborationPartnerCampaign, err error) {
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	return q.WithContext(ctx).Where(q.ID.Eq(id)).First()
}

func GetPartnerCampaignMapByIDs(ctx context.Context, ids []int64) (partnerCampaignMap map[int64]*model.CollaborationPartnerCampaign, err error) {
	if len(ids) == 0 {
		return map[int64]*model.CollaborationPartnerCampaign{}, nil
	}
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	partnerCampaignList, err := q.WithContext(ctx).Where(q.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	partnerCampaignMap = make(map[int64]*model.CollaborationPartnerCampaign, len(partnerCampaignList))
	for _, partnerCampaign := range partnerCampaignList {
		partnerCampaignMap[partnerCampaign.ID] = partnerCampaign
	}
	return partnerCampaignMap, nil
}

func GetPartnerCampaignWithPartner(ctx context.Context, campaignID, partnerID int64) (campaign *model.CollaborationPartnerCampaign, err error) {
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	return q.WithContext(ctx).Where(q.CampaignID.Eq(campaignID)).Where(q.AioPartnerID.Eq(partnerID)).First()
}

func GetPartnerCampaignWithClient(ctx context.Context, campaignID, clientID int64) (campaign *model.CollaborationPartnerCampaign, err error) {
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	return q.WithContext(ctx).Where(q.CampaignID.Eq(campaignID)).Where(q.AioClientID.Eq(clientID)).First()
}

func GetPartnerCampaignWithPartnerAndClient(ctx context.Context, campaignID, partnerID, clientID int64) (campaign *model.CollaborationPartnerCampaign, err error) {
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	return q.WithContext(ctx).
		Where(q.CampaignID.Eq(campaignID)).
		Where(q.AioClientID.Eq(clientID)).
		Where(q.AioPartnerID.Eq(partnerID)).
		First()
}

func GetPartnerCampaignWithCampaignAndStatus(ctx context.Context, campaignID int64, statusList []collaboration.PartnerCampaignStatus) (parterCampaignList []*model.CollaborationPartnerCampaign, err error) {
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	filter := q.WithContext(ctx).Where(q.CampaignID.Eq(campaignID))
	var statusFieldPtr []driver.Valuer
	for _, status := range statusList {
		statusFieldPtr = append(statusFieldPtr, &status)
	}
	if len(statusFieldPtr) > 0 {
		filter = filter.Where(q.Status.In(statusFieldPtr...))
	}
	return filter.Find()
}

func GetPartnerCampaignByPartnerAndStatus(ctx context.Context, aioPartnerIds []int64, statusList []collaboration.PartnerCampaignStatus) (parterCampaignList []*model.CollaborationPartnerCampaign, err error) {
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	filter := q.WithContext(ctx).Where(q.AioPartnerID.In(aioPartnerIds...))
	var statusFieldPtr []driver.Valuer
	for _, status := range statusList {
		statusFieldPtr = append(statusFieldPtr, &status)
	}
	if len(statusFieldPtr) > 0 {
		filter = filter.Where(q.Status.In(statusFieldPtr...))
	}
	return filter.Find()
}

func SearchPartnerCampaignWithAcceptRate(ctx context.Context, campaignID int64, statusList []collaboration.PartnerCampaignStatus, acceptRate100kLt *int64, acceptRate100kGt *int64) (parterCampaignList []*model.CollaborationPartnerCampaign, err error) {
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	filter := q.WithContext(ctx).Where(q.CampaignID.Eq(campaignID))
	var statusFieldPtr []driver.Valuer
	for _, status := range statusList {
		statusFieldPtr = append(statusFieldPtr, &status)
	}
	if len(statusFieldPtr) > 0 {
		filter = filter.Where(q.Status.In(statusFieldPtr...))
	}
	if acceptRate100kGt != nil {
		filter = filter.Where(q.AcceptRate100K.Gt(*acceptRate100kGt))
	}
	if acceptRate100kLt != nil {
		filter = filter.Where(q.AcceptRate100K.Lt(*acceptRate100kLt))
	}
	return filter.Find()
}

func ScanPartnerCampaignByCampaignID(ctx context.Context, campaignID int64) ([]*model.CollaborationPartnerCampaign, error) {
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	batchSize := 200
	result := make([]*model.CollaborationPartnerCampaign, 0, batchSize)
	totalResult := make([]*model.CollaborationPartnerCampaign, 0)
	err := q.WithContext(ctx).Where(q.CampaignID.Eq(campaignID)).FindInBatches(&result, batchSize, func(tx gen.Dao, batch int) error {
		totalResult = append(totalResult, result...)
		return nil
	})
	return totalResult, err
}

func UpdatePartnerCampaignDataCenter(ctx context.Context, id int64, dataCenter common_param.DataCenter, write bool) error {
	if !write {
		log.V2.Warn().With(ctx).Str("[write=false]partner campaign update data center").KVs("partner campaign id", id, "data center", dataCenter).Emit()
		return nil
	}
	q := query.CreativeSession(ctx).CollaborationPartnerCampaign
	_, err := q.WithContext(ctx).Where(q.ID.Eq(id)).Update(q.DataCenter, dataCenter)
	if err != nil {
		return err
	}
	return nil
}
