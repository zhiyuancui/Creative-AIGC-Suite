package custom

import (
	"context"
)

func CreateUpdateCampaign(ctx context.Context, campaignModel *model.CollaborationCampaignInfo) error {
	if campaignModel.ID == 0 {
		campaignID, err := idgenerator.NewIdLong2(ctx)
		if err != nil {
			logs.CtxError(ctx, "CreateUpdateCampaign id generate fail, err is %+v", err)
			return err
		}
		campaignModel.ID = int64(campaignID)
	}
	return query.CreativeSession(ctx).CollaborationCampaignInfo.WithContext(ctx).Save(campaignModel)
}

func GetCampaignByIDWithClient(ctx context.Context, campaignID int64, AIOClientID int64) (campaignModel *model.CollaborationCampaignInfo, err error) {
	if AIOClientID == 0 {
		return GetCampaignByIDWithoutClient(ctx, campaignID)
	}
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	return q.WithContext(ctx).Where(q.ID.Eq(campaignID)).Where(q.AioClientID.Eq(AIOClientID)).First()
}

func GetCampaignByIDWithoutClient(ctx context.Context, campaignID int64) (campaignModel *model.CollaborationCampaignInfo, err error) {
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	return q.WithContext(ctx).Where(q.ID.Eq(campaignID)).First()
}

func MGetCampaignByIDs(ctx context.Context, campaignIDs []int64) (campaignModel []*model.CollaborationCampaignInfo, err error) {
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	return q.WithContext(ctx).Where(q.ID.In(campaignIDs...)).Find()
}

func DeleteCampaignByID(ctx context.Context, campaignID int64) (err error) {
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	_, err = q.WithContext(ctx).Where(q.ID.Eq(campaignID)).Unscoped().Delete()
	return
}

func GetCampaignListWithPageByCondition(ctx context.Context, campaignID *int64, AIOClientID *int64, matchMode *collaboration.MatchMode, seatType *collaboration.SeatType, statusList []int32) (campaignModelList []*model.CollaborationCampaignInfo, err error) {
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	search := q.WithContext(ctx)
	if campaignID != nil {
		search = search.Where(q.ID.Eq(*campaignID))
	}
	if AIOClientID != nil {
		search = search.Where(q.AioClientID.Eq(*AIOClientID))
	}
	if matchMode != nil {
		search = search.Where(q.MatchingMode.Eq(matchMode))
	}
	if seatType != nil {
		search = search.Where(q.SeatType.Eq(int32(*seatType)))
	}
	if len(statusList) != 0 {
		search = search.Where(q.Status.In(statusList...))
	}
	return search.Find()
}

func GetCampaignCountGroupByStatus(ctx context.Context, aioClientID int64, matchingMode *collaboration.MatchMode, groupStatus []int32) (ret map[collaboration.CampaignStatus]int32, err error) {
	type result struct {
		Status int32
		Total  int32
	}
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	rows, err := q.WithContext(ctx).Select(q.Status, q.ID.Count()).Group(q.Status).
		Where(q.AioClientID.Eq(aioClientID)).
		Where(q.MatchingMode.Eq(matchingMode)).
		Where(q.Status.In(groupStatus...)).Rows()
	if err != nil {
		return nil, err
	}

	ret = make(map[collaboration.CampaignStatus]int32)
	for rows.Next() {
		var r result
		err = rows.Scan(&r.Status, &r.Total)
		if err != nil {
			return nil, err
		}
		ret[collaboration.CampaignStatus(r.Status)] = r.Total
	}
	return
}

func GetCampaignIDsByPage(ctx context.Context, page, limit int32) ([]int64, error) {
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	rows, err := q.WithContext(ctx).Where(q.ID.Gt(0)).Order(q.ID).Limit(int(limit)).Offset(int((page - 1) * limit)).Find()
	if err != nil {
		return nil, err
	}
	campaignIDs := make([]int64, 0, len(rows))
	for _, row := range rows {
		campaignIDs = append(campaignIDs, row.ID)
	}
	return campaignIDs, nil
}

func GetEUSolutionYCampaignsByPage(ctx context.Context, euCountries []string, page, limit int) (campaigns []*model.CollaborationCampaignInfo, err error) {
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	rows, _, err := q.WithContext(ctx).
		Where(q.SolutionType.Eq(gptr.Of(collaboration.SolutionType_SolutionY))).
		Where(q.CampaignCountry.In(euCountries...)).
		FindByPage((page-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	campaigns = make([]*model.CollaborationCampaignInfo, 0, len(rows))
	for _, row := range rows {
		campaigns = append(campaigns, row)
	}
	return
}

func GetSolutionXRunningOACCampaignsByPage(ctx context.Context, offset, limit int) (campaigns []*model.CollaborationCampaignInfo, err error) {
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	campaigns, _, err = q.WithContext(ctx).
		Where(q.SolutionType.Eq(gptr.Of(collaboration.SolutionType_SolutionX))).
		Where(q.Status.Eq(int32(collaboration.CampaignStatus_Running))).
		Where(q.MatchingMode.Eq(gptr.Of(collaboration.MatchMode_DIAndOAC))).
		FindByPage(offset, limit)
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

//func IsCampaignExists(ctx context.Context, campaignID int64, aioClientID int64) (isExists bool, err error) {
//	q := query.CreativeSession(ctx).CollaborationCampaignInfo
//
//	f := q.WithContext(ctx).Select(q.ID).Where(q.ID.Eq(campaignID))
//	if aioClientID != 0 {
//		f = f.Where(q.AioClientID.Eq(aioClientID))
//	}
//	ret, err := f.First()
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		log.V2.Error().With(ctx).Str("IsCampaignExists error").KV("ret", ret).Error(err).Emit()
//		err = error1.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, err.Error()).WithStack()
//		return
//	}
//	return ret != nil, nil
//}

func GetCampaignClientID(ctx context.Context, campaignID int64) int64 {
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	campaignModel, err := q.WithContext(ctx).Select(q.AioClientID).Where(q.ID.Eq(campaignID)).First()
	if err != nil {
		log.V2.Warn().With(ctx).Str("[GetCampaignClientID]error").Error(err).Emit()
		return 0
	}
	return campaignModel.AioClientID
}

func ScanCampaignByCursor(ctx context.Context, campaignID int64, limit int) (campaigns []*model.CollaborationCampaignInfo, err error) {
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	return q.WithContext(ctx).Where(q.ID.Gte(campaignID)).Order(q.ID).Limit(limit).Find()
}

func UpdateCampaignDataCenter(ctx context.Context, campaignID int64, dataCenter common_param.DataCenter, write bool) error {
	if !write {
		log.V2.Warn().With(ctx).Str("[write=false]campaign update data center").KVs("campaign id", campaignID, "data center", dataCenter).Emit()
		return nil
	}
	q := query.CreativeSession(ctx).CollaborationCampaignInfo
	_, err := q.WithContext(ctx).Where(q.ID.Eq(campaignID)).Update(q.DataCenter, dataCenter)
	if err != nil {
		return err
	}
	return nil
}
