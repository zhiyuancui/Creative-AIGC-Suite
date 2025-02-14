package service

var campaignDraftDomainService ICampaignDraftService

func InitDomain() {
	campaignDraftDomainService = &campaignDraftServiceImpl{}
}
