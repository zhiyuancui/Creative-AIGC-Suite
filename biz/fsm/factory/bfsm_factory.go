package factory

const MatchParamsKey = "matchParams"

// MatchParamsVal condition 判断条件，以map <key, MatchParamsVal>传入，key 指定未`matchParams`
type MatchParamsVal struct {
	IsClientSelect     bool
	NeedPartnerInvoice bool
}

func NewPartnerCampaignBFSM(agent entity.IAgentEntity) (FSM *bfsm.FSM, err error) {
	FSM, err = bfsm.NewFSM(agentBase, agent.GetCampaignStatus().String())
	if err != nil {
		return
	}
	return
}
