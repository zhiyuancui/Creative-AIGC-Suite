package factory

const AgentBase = "agent_base"

var CreatorAgentFsmConfig = map[string]bfsm.BizDesc{
	AgentBase: {
		CommonDstCallback: SaveAndUpdateCampaignAndStatus,
		TransDescList: []bfsm.TransDesc{
			{
				Src:                 []string{collaboration.CreatorAgentStatus_ClientSelectingPartner.String(), collaboration.CreatorAgentStatus_PSOAssignPartner.String()},
				Event:               partner_business.CampaignOperationEvent_ClientCancelCampaign.String(),
				BeforeTransCallback: nil,
				SrcCallback:         nil,
				Matchers: []bfsm.Matcher{
					{
						Dst:                collaboration.CreatorAgentStatus_Cancelled.String(),
						AfterTransCallback: callback.RollbackCoupon,
					},
				},
			},
			{
				Src:                 []string{collaboration.CreatorAgentStatus_PartnerReviewingAssignment.String()},
				Event:               partner_business.CampaignOperationEvent_PartnerAcceptInvitation.String(), //partner approve
				BeforeTransCallback: nil,
				SrcCallback:         nil,
				Matchers: []bfsm.Matcher{
					{
						Dst:                collaboration.CreatorAgentStatus_PartnerMakingProduct.String(),
						AfterTransCallback: callback.NotifyClientPartnerAcceptInvite,
					},
				},
			},
			{
				Src:                 []string{collaboration.CreatorAgentStatus_PartnerReviewingAssignment.String()},
				Event:               partner_business.CampaignOperationEvent_PartnerRejectInvitation.String(), //partner reject
				BeforeTransCallback: nil,
				SrcCallback:         nil,
				Matchers: []bfsm.Matcher{
					{
						Dst:                collaboration.CreatorAgentStatus_ClientSelectingPartner.String(),
						Condition:          "matchParams.IsClientSelect",
						AfterTransCallback: callback.NotifyClientPartnerRejectInvite,
					},
					{
						Dst:                collaboration.CreatorAgentStatus_PSOAssignPartner.String(),
						Condition:          "!matchParams.IsClientSelect",
						AfterTransCallback: callback.NotifyClientPartnerRejectInvite,
					},
				},
			},
			{
				Src:                 []string{collaboration.CreatorAgentStatus_PartnerMakingProduct.String(), collaboration.CreatorAgentStatus_ProductUpload.String()},
				Event:               partner_business.CampaignOperationEvent_PartnerUploadCreatives.String(),
				BeforeTransCallback: nil,
				SrcCallback:         nil,
				Matchers: []bfsm.Matcher{
					{
						Dst: collaboration.CreatorAgentStatus_ProductUpload.String(),
					},
				},
			},
			{
				Src:                 []string{collaboration.CreatorAgentStatus_ClientConfirmingProduct.String()},
				Event:               partner_business.CampaignOperationEvent_PartnerUploadCreatives.String(),
				BeforeTransCallback: nil,
				SrcCallback:         nil,
				Matchers: []bfsm.Matcher{
					{
						Dst: collaboration.CreatorAgentStatus_ClientConfirmingProduct.String(),
					},
				},
			},
			{
				Src:                 []string{collaboration.CreatorAgentStatus_ProductComplete.String()},
				Event:               partner_business.CampaignOperationEvent_PartnerUploadCreatives.String(),
				BeforeTransCallback: nil,
				SrcCallback:         nil,
				Matchers: []bfsm.Matcher{
					{
						Dst: collaboration.CreatorAgentStatus_ProductComplete.String(),
					},
				},
			},
			{
				Src:                 []string{collaboration.CreatorAgentStatus_ProductUpload.String()},
				Event:               partner_business.CampaignOperationEvent_PartnerCompleteCampaign.String(),
				BeforeTransCallback: nil,
				SrcCallback:         nil,
				Matchers: []bfsm.Matcher{
					{
						Dst:                collaboration.CreatorAgentStatus_ClientConfirmingProduct.String(),
						AfterTransCallback: callback.NotifyClientPartnerCompleteCampaign,
					},
				},
			},
			{
				Src:                 []string{collaboration.CreatorAgentStatus_ClientConfirmingProduct.String()},
				Event:               partner_business.CampaignOperationEvent_ClientRejectCompletion.String(), //partner complete order
				BeforeTransCallback: nil,
				SrcCallback:         nil,
				Matchers: []bfsm.Matcher{
					{
						Dst:                collaboration.CreatorAgentStatus_ProductUpload.String(),
						AfterTransCallback: nil,
					},
				},
			},
			{
				Src:                 []string{collaboration.CreatorAgentStatus_ClientConfirmingProduct.String()},
				Event:               partner_business.CampaignOperationEvent_ClientCompleteCampaign.String(), //partner complete order
				BeforeTransCallback: nil,
				SrcCallback:         nil,
				Matchers: []bfsm.Matcher{
					{
						Dst:                collaboration.CreatorAgentStatus_ProductComplete.String(),
						Condition:          "!matchParams.NeedPartnerInvoice",
						AfterTransCallback: callback.NotifyClientCampaignCompletedCallBack,
					},
					{
						Dst:                collaboration.CreatorAgentStatus_ProductComplete.String(),
						Condition:          "matchParams.NeedPartnerInvoice",
						AfterTransCallback: callback.CreatorAgentApproveCallbackAfterTran,
					},
				},
			},
		},
	},
}
