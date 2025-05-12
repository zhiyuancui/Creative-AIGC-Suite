package campaign

import (
	"context"
	"sort"
	"strconv"
	"time"
)

func ClientDraftCampaignDetail(ctx context.Context, aioClientID int64, campaignID int64) (resp *rpcClientCampaignDraftDetailResp, err error) {
	resp = &rpcClientCampaignDraftDetailResp{BaseResp: base.NewBaseResp()}
	if aioClientID == 0 || campaignID == 0 {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonParamCheckError, "Invalid input").WithStack()
		return
	}

	briefModel, err := custom.GetCampaignDraftByID(ctx, campaignID, aioClientID)
	if err != nil {
		logs.CtxError(ctx, "fail to get campaign draft of campaign id %d with err: %+v", campaignID, err)
		err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "Invalid input").WithStack()
		return
	}

	campaignInfo, err := service.ConvertBriefModelToCampaignResp(ctx, briefModel, nil)
	if err != nil {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "fail to convert").WithStack()
		return
	}

	resp.SetCampaignInfo(campaignInfo)
	return
}

func ClientCreateCampaignDraft(ctx context.Context, req *rpcClientCreateCampaignDraftReq) (resp *rpcClientCreateCampaignDraftResp, err error) {
	resp = &rpcClientCreateCampaignDraftResp{BaseResp: base.NewBaseResp()}

	if req.GetCampaignInfo() == nil {
		logs.CtxError(ctx, "campaign draft info is null")
		err = errors.NewEMPErrorWithMsg(creative_one.CommonParamCheckError, "empty campaign").WithStack()
		return
	}
	req.GetCampaignInfo().SetBriefID(nil)
	req.GetCampaignInfo().SetCampaignID(nil)
	briefID, ttoCampaignID, updatedAt, err := ClientCreateUpdateCampaignDraft(ctx, req.GetCampaignInfo(), req.GetParam(), req.GetAioClientID(), req.GetCoreUserID())
	if err != nil {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "fail to create draft").WithStack()
		return
	}

	resp.SetBriefID(briefID)
	resp.SetCampaignID(ttoCampaignID)
	resp.SetUpdatedAt(updatedAt.Unix())

	return
}

func ClientUpdateCampaignDraft(ctx context.Context, req *rpcClientUpdateCampaignDraftReq) (resp *rpcClientUpdateCampaignDraftResp, err error) {
	resp = &rpcClientUpdateCampaignDraftResp{BaseResp: base.NewBaseResp()}

	briefID, ttoCampaignID, updatedAt, err := ClientCreateUpdateCampaignDraft(ctx, req.GetCampaignInfo(), req.GetParam(), req.GetAioClientID(), req.GetCoreUserID())
	if err != nil {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "fail to update draft").WithStack()
		return
	}

	resp.SetBriefID(briefID)
	resp.SetCampaignID(ttoCampaignID)
	resp.SetUpdatedAt(updatedAt.Unix())

	return
}

func ClientCreateUpdateCampaignDraft(ctx context.Context, campaign *rpcCampaignInfo, param *common_param.ClientCommonParam, aioClientID int64, coreUserID int64) (briefID int64, ttoCampaignID int64, updatedAt time.Time, err error) {
	campaign.SetAioClientID(&aioClientID)
	campaign.SetCoreUserID(&coreUserID)

	//TODO: wrap to separate method
	if len(campaign.CampaignName) == 0 {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonParamCheckError, "invalid brief").WithStack()
		return
	}

	var createCampaign = campaign.GetCampaignID() == 0
	if campaign.GetCampaignID() != 0 {
		existingBrief, err := custom.GetCampaignDraftByID(ctx, campaign.GetCampaignID(), campaign.GetAioClientID())
		if err == nil {
			createCampaign = false
			campaign.SetBriefID(gptr.Of(existingBrief.ID))
		}
	}

	briefModel := service.ConvertCampaignEntityToBriefModel(ctx, campaign, param)

	q := query.Session(clients.LegoDBCli)
	err = q.Transaction(func(tx *query.Query) error {
		if createCampaign {
			ttoCampaignID, err = RPC.ClientCreateCampaignDraft(ctx, campaign, param.IsTest)
			if err != nil {
				logs.CtxError(ctx, "fail to create campaign in tto with err: %+v", err)
				err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "fail to create/update tto campaign draft").WithStack()
				return err
			}
			briefModel.TtoCampaignID = ttoCampaignID

		} else {
			ttoCampaignID = briefModel.TtoCampaignID
			err = RPC.ClientUpdateCampaignDraft(ctx, briefModel.TtoCampaignID, campaign, param.IsTest)
			if err != nil {
				logs.CtxError(ctx, "fail to create campaign in tto with err: %+v", err)
				err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "fail to create/update tto campaign draft").WithStack()
				return err
			}
		}
		err = custom.CreateUpdateCampaignDraft(ctx, tx, briefModel)
		if err != nil {
			logs.CtxError(ctx, "fail to update/create brief with err: %+v", err)
			err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "fail to update/create brief").WithStack()
			return err
		}

		err = RPC.InnerUpdatetCampaign(ctx, aioClientID, briefModel.TtoCampaignID, nil, &briefModel.ID, nil, &briefModel.OrderName, nil)
		if err != nil {
			logs.CtxError(ctx, "fail to sync campaign %d in tto, err: %+v", briefModel.TtoCampaignID, err)
			err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "fail to sync tto campaign draft").WithStack()
			return err
		}

		return nil
	})
	if err != nil {
		return
	}

	briefModel, err = custom.GetCampaignDraftByBriefID(ctx, briefModel.ID)
	if err != nil {
		updatedAt = time.Now()
	} else {
		updatedAt = briefModel.UpdatedAt
	}

	return briefModel.ID, ttoCampaignID, updatedAt, nil
}

func ClientSubmitCampaign(ctx context.Context, req *rpcClientSubmittCampaignReq) (resp *rpcClientSubmittCampaignResp, err error) {
	resp = &rpcClientSubmittCampaignResp{BaseResp: base.NewBaseResp()}

	brief, err := custom.GetCampaignDraftByID(ctx, req.GetCampaignID(), req.GetAioClientID())
	if err != nil {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "Invalid input").WithStack()
		return
	}

	campaign, err := service.ConvertBriefModelToCampaignResp(ctx, brief, nil)
	if err != nil {
		return
	}
	if !service.CheckCampaignKeyField(ctx, campaign) {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonParamCheckError, "Invalid input").WithStack()
		return
	}

	orderID, isTest, err := handlers.SubmitTTOCampaign(ctx, brief.ID, req.GetCoreUserID(), req.GetCouponID(), 0, true, req.Param.GetIsTest())
	if err != nil && orderID == 0 {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonInternalSystemError, "fail to submit campaign").WithStack()
		return
	}

	resp.SettOrderID(orderID)
	resp.SetIsTest(isTest)
	resp.SetCampaignID(brief.TtoCampaignID)

	return
}

func ClientCampaignIndustries(ctx context.Context) (resp *rpcClientCampaignIndustriesResp, err error) {
	resp = &rpcClientCampaignIndustriesResp{BaseResp: base.NewBaseResp()}

	industryConf := clients.DynamicConfigGetter.IndustryConf()

	var industries []*rpcIndustryConfig
	for _, item := range industryConf.FirstIndustries {
		industry := rpcIndustryConfig{
			IndustryID: item.IndustryID,
		}
		industries = append(industries, &industry)
	}
	resp.SetIndustries(industries)
	return
}

func ClientCampaignLanguages(ctx context.Context, req *rpcClientCampaignLanguagesReq) (resp *rpcClientCampaignLanguagesResp, err error) {
	resp = &rpcClientCampaignLanguagesResp{BaseResp: base.NewBaseResp()}

	partnerInfos, err := RPC.PartnerGetProfileInfoList(ctx, req.GetRegion(), req.Param.GetIsTest(), 1, 1000)
	if err != nil {
		logs.CtxError(ctx, "[ClientCampaignLanguages], failed to get partner infos from es. error: %v", err)
		return
	}

	var languageMap = make(map[string]bool)
	for _, partner := range partnerInfos {
		if partner.FormData == nil || len(partner.FormData.CreativeSupportedLanguages) == 0 {
			continue
		}
		for _, language := range partner.FormData.CreativeSupportedLanguages {
			languageMap[language] = true
		}
	}

	var allLanguagesSupported []string
	for language := range languageMap {
		allLanguagesSupported = append(allLanguagesSupported, language)
	}
	sort.Strings(allLanguagesSupported)
	resp.SetLanguages(allLanguagesSupported)

	return
}

func ClientCampaignMarkets(ctx context.Context, req *rpcClientCampaignMarketsReq) (resp *rpcClientCampaignMarketsResp, err error) {
	resp = &rpcClientCampaignMarketsResp{BaseResp: base.NewBaseResp()}

	if req.GetRegion() == "" {
		return
	}

	rpcResp, err := rpc.GetBriefSupportedMarkets(ctx, req.GetRegion(), req.Param.GetIsTest())
	if err != nil {
		return
	}

	markets := []*rpcSupportedRegionAndMarkets{}

	for _, item := range rpcResp.GetSupportedMarkets() {
		market := rpcSupportedRegionAndMarkets{
			SupportedRegion: item.GetSupportedRegion(),
			Markets:         item.GetMarkets(),
		}
		markets = append(markets, &market)
	}

	resp.SetSupportedMarkets(markets)
	return
}

func ClientCampaignAdvertiserInfo(ctx context.Context, req *rpcClientCampaignAdvertiserInfoReq) (resp *rpcClientCampaignAdvertiserInfoResp, err error) {
	resp = &rpcClientCampaignAdvertiserInfoResp{BaseResp: base.NewBaseResp()}

	if req.GetCoreUserID() == 0 {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonParamCheckError, "Invalid input").WithStack()
		return
	}

	advAllowedInfo, isAllowed, err := qualification.IsAdvQualified(ctx, req.GetCoreUserID())
	if err != nil {
		return
	}

	advList := []model.tAdvAccountInfo{}
	resp.SetRegion(gptr.Of("Unknown"))
	resp.SetIsWhitelisted(gptr.Of(false))
	if isAllowed {
		advList, err = account.GetClientAdvInfo(ctx, req.GetCoreUserID(), req.GetNeedCouponInfo(), true, isAllowed, advAllowedInfo.Region)
		resp.SetRegion(&advAllowedInfo.Region)
		resp.SetIsWhitelisted(gptr.Of(true))
		if err != nil {
			return
		}
	} else {
		advList, err = account.GetClientAdvInfo(ctx, req.GetCoreUserID(), req.GetNeedCouponInfo(), true, isAllowed, "")
		if err != nil {
			return
		}
	}

	if err = setIsJBPCoupon(ctx, advList); err != nil {
		return
	}

	if err = setCanSubmitSubscription(ctx, advList, req.GetCoreUserID()); err != nil {
		return
	}

	res := []*rpctAdvAccountInfo{}
	supportedRegions := []string{}
	for _, item := range advList {
		advInfo := rpctAdvAccountInfo{
			Id:                    item.ID,
			Name:                  item.Name,
			HasCoupon:             conv.BoolPtr(item.HasCoupon),
			IncentiveType:         conv.Int32Ptr(int32(item.IncentiveType)),
			BaseRegion:            conv.StringPtr(item.BaseRegion),
			IsJBP:                 conv.BoolPtr(item.IsJBP),
			CanSubmitSubscription: conv.BoolPtr(item.CanSubmitSubscription),
		}
		supportedRegions = append(supportedRegions, item.BaseRegion)
		res = append(res, &advInfo)
	}

	resp.SetSupportedGARegion(gslice.Uniq(supportedRegions))
	resp.SetAccountInfo(res)

	return
}

func setIsJBPCoupon(ctx context.Context, advList []model.tAdvAccountInfo) error {
	advIDs := gslice.Map(advList, func(info model.tAdvAccountInfo) int64 {
		advID, err := strconv.ParseInt(info.ID, 10, 64)
		if err != nil {
			logs.CtxError(ctx, "parse AdvID error: %+v", err)
		}
		return advID
	})

	advToCoupon := make(map[string][]*t.CreativeCoupon)

	for _, subAdvIDs := range gslice.Chunk(advIDs, 100) {
		couponMap, err := rpc.MGetAdvCouponList(ctx, subAdvIDs)
		if err != nil {
			return err
		}
		for key, value := range couponMap {
			advToCoupon[strconv.FormatInt(key, 10)] = value
		}
	}

	for i, adv := range advList {
		advList[i].IsJBP = len(advToCoupon[adv.ID]) > 0
	}

	return nil
}

func setCanSubmitSubscription(ctx context.Context, advList []model.tAdvAccountInfo, coreUserId int64) error {
	GARegions := clients.DynamicConfigGetter.FeatureGARegions()["subscription"]
	allowedIdMap, err := rpc.GetAllObjectIdsInAllowList(ctx, core.AllowListModuleType_tGeneralModule, core.AllowListEntityType_Subscription, 0)
	if err != nil {
		logs.CtxError(ctx, "[CheckUserInAllowList] get all users in allowList for subscription failed, err: [%+v]", err)
		return bizerr.ErrInternalError
	}

	var baseRegion string
	advertiserInfo, _, err := qualification.IsAdvQualified(ctx, coreUserId)
	if err != nil {
		return err
	}

	for i, adv := range advList {
		if advertiserInfo != nil {
			baseRegion = advertiserInfo.Region
		} else {
			baseRegion = adv.BaseRegion
		}

		if gslice.Contains(GARegions, baseRegion) {
			advList[i].CanSubmitSubscription = true
		} else if _, ok := allowedIdMap[adv.ID]; ok {
			advList[i].CanSubmitSubscription = true
		}
	}

	return nil
}

func ClientRecommendPartner(ctx context.Context, req *rpcClientRecommendPartnerReq) (resp *rpcClientRecommendPartnerResp, err error) {
	resp = &rpcClientRecommendPartnerResp{BaseResp: base.NewBaseResp()}

	var (
		recommendPartnerInfos  []params.PartnerProfile
		preSelectedPartnerInfo *params.PreSelectedPartnerProfile
	)
	if req.GetCampaignInfo() != nil && req.GetCampaignInfo().GettOrderID() != 0 {
		logs.CtxInfo(ctx, "[ClientRecommendPartner] advertiser reselect partner for campaign %d", req.GetCampaignInfo().CampaignID)
		orderID := req.GetCampaignInfo().GettOrderID()
		recommendPartnerInfos, preSelectedPartnerInfo, err = match_algo.GetRecommendProfilerImpl(ctx, orderID, false)
		if err != nil {
			logs.CtxError(ctx, "fail to get recommend partner for campaign %d with err: %+v", req.GetCampaignInfo().CampaignID, err)
			return
		}
		if preSelectedPartnerInfo != nil {
			resp.SetSelectedPartnerInfo(service.ConvertToSelectPartnerProfile(preSelectedPartnerInfo))
		}

		resp.SetRecommendPartnerInfo(service.ConvertToPartnerProfile(recommendPartnerInfos))
		return
	}

	if req.CampaignInfo == nil ||
		req.CampaignInfo.Region == nil ||
		req.CampaignInfo.AdvertiserID == nil ||
		req.CampaignInfo.GetServicePackageID() == 0 ||
		len(req.CampaignInfo.GetTargetMarkets()) == 0 ||
		req.CampaignInfo.IndustryCategory == nil ||
		len(req.CampaignInfo.Objective) == 0 ||
		len(req.CampaignInfo.CreativeLanguages) == 0 {
		err = errors.NewEMPErrorWithMsg(creative_one.CommonParamCheckError, "Invalid input").WithStack()
		return
	}

	if req.CampaignInfo.GetRegion() == "CNOB" {
		return
	}

	addOnIds := []int64{}
	for _, item := range req.CampaignInfo.AddOns {
		addOnIds = append(addOnIds, item.GetId())
	}

	objectives := []int64{}
	for _, item := range req.CampaignInfo.Objective {
		objectives = append(objectives, int64(item))
	}
	formData := &core.BriefFormData{
		AdvertiserId:      req.CampaignInfo.AdvertiserID,
		TargetMarkets:     req.CampaignInfo.TargetMarkets,
		CreativeLanguage:  req.CampaignInfo.CreativeLanguages,
		ServicePackageId:  req.CampaignInfo.ServicePackageID,
		IndustryCategory:  req.CampaignInfo.IndustryCategory,
		Objective:         objectives,
		SelectedPartnerId: req.CampaignInfo.SelectedPartnerID,
	}

	logs.CtxInfo(ctx, "[ClientRecommendPartner] ClientRecommendPartner request: [%+v]", req)

	recommendPartnerInfos, preSelectedPartnerInfo, err = match_algo.GetTTORecommendPartnersFromBriefImpl(ctx, false, addOnIds, req.GetIsFunded(), req.GetPreSelectedPartnerID(), req.GetCoreUserID(), req.Param.GetIsTest(), req.CampaignInfo.GetRegion(), formData)
	if err != nil {
		return
	}

	if preSelectedPartnerInfo != nil {
		resp.SetSelectedPartnerInfo(service.ConvertToSelectPartnerProfile(preSelectedPartnerInfo))
	}

	resp.SetRecommendPartnerInfo(service.ConvertToPartnerProfile(recommendPartnerInfos))

	return
}
