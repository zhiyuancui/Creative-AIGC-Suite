package global_framework

import (
	"code.byted.org/ad/creative_one_partner_business/application/service"
	"code.byted.org/ad/creative_one_partner_business/infrastructure/rpc"
	"code.byted.org/ad/creative_one_partner_business/kitex_gen/ad/creative_one/partner_business"
	"code.byted.org/ad/creative_one_partner_business/kitex_gen/base"
	"code.byted.org/ad/creative_tcpp_server_i18n/biz/clients"
	ttcxrpc "code.byted.org/ad/creative_tcpp_server_i18n/biz/dals/rpc"
	"code.byted.org/gopkg/lang/conv"
	"code.byted.org/lang/gg/gptr"
	"code.byted.org/lang/gg/gslice"
	"context"
)

func ClientCampaignServicePackages(ctx context.Context, req *partner_business.ClientCampaignServicePackageReq) (resp *partner_business.ClientCampaignServicePackageResp, err error) {
	resp = &partner_business.ClientCampaignServicePackageResp{BaseResp: base.NewBaseResp()}

	servicePackages, err := rpc.GetBriefAvailableSPV2s(
		ctx, req.BriefRegion, req.CoreUserID, req.SelectedAdvertiserID, conv.BoolPtr(false), req.IsSubscription, req.IndustryType)
	if err != nil {
		return
	}
	isJBP, err := ttcxrpc.IsJBPAdvertiser(ctx, req.GetSelectedAdvertiserID())
	if err != nil {
		return
	}
	if isJBP && !gslice.Contains(clients.DynamicConfigGetter.JBPVCPDisableMinimumSpendRegion(), req.BriefRegion) {
		for _, pck := range append(servicePackages.NetNewPackages, servicePackages.RemixPackages...) {
			pck.MinSpend = 0
			pck.SubscriptionMinSpend = gptr.Of(0.0)
			pck.OriginalMinSpend = gptr.Of(0.0)
			pck.AffiliatedMediaFee = gptr.Of(0.0)
		}
		for header, pcks := range servicePackages.GetServicePackages() {
			for _, pck := range pcks {
				pck.MinSpend = 0
				pck.SubscriptionMinSpend = gptr.Of(0.0)
				pck.OriginalMinSpend = gptr.Of(0.0)
				pck.AffiliatedMediaFee = gptr.Of(0.0)
				if header == "net_new_packages" || header == "remix_packages" {
					pck.Price = 0
				}
			}
		}
	}

	servicePackageMap := map[string][]*partner_business.ServicePackageCardInfo{}
	for key, value := range servicePackages.ServicePackages {
		servicePackageList := []*partner_business.ServicePackageCardInfo{}
		for _, item := range value {
			servicePackageList = append(servicePackageList, service.ConvertServicePackageCardInfo(item))
		}

		servicePackageMap[key] = servicePackageList
	}

	maps := &partner_business.BriefAvailableSPV2sRespMaps{
		HeaderMap:           map[string]*partner_business.PackageHeaderDesc{},
		PackageMap:          map[string][]string{},
		BusinessTypeDescMap: map[string]*partner_business.BusinessTypeDesc{},
	}

	if servicePackages.Maps != nil {
		maps.SetPackageMap(servicePackages.GetMaps().PackageMap)
		headerMap := map[string]*partner_business.PackageHeaderDesc{}
		businessTypeDescMap := map[string]*partner_business.BusinessTypeDesc{}
		if servicePackages.GetMaps().HeaderMap != nil {
			for key, value := range servicePackages.Maps.HeaderMap {
				headerMap[key] = service.ConvertPackageHeaderDesc(value)
			}
		}

		if servicePackages.GetMaps().BusinessTypeDescMap != nil {
			for key, value := range servicePackages.Maps.BusinessTypeDescMap {
				businessTypeDescMap[key] = &partner_business.BusinessTypeDesc{
					TitleDesc:       value.TitleDesc,
					ContentDesc:     value.ContentDesc,
					DisclaimerDesc:  value.DisclaimerDesc,
					StartingPrice:   value.StartingPrice,
					CrossedOutPrice: value.CrossedOutPrice,
					Currency:        value.Currency,
				}
			}
		}
		maps.SetHeaderMap(headerMap)
		maps.SetBusinessTypeDescMap(businessTypeDescMap)
	}

	resp.SetServicePackages(servicePackageMap)
	resp.SetMaps(maps)
	return
}
