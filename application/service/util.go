package service

import (
	"context"
	basicJSON "encoding/json"
)

func ConvertBriefModelToCampaignEntity(ctx context.Context, brief *BriefDraft) (campaignInfo *CampaignInfo, err error) {

}

func ConvertCampaignEntityToBriefModel(ctx context.Context, campaignInfo *CampaignInfo, param *ClientCommonParam) (brief *ttcxmodel.BriefDraft) {

}

// OverwriteFields overwrites the fields of target with the same fields of source
func OverwriteFields(ctx context.Context, target interface{}, source interface{}) error {
	logs.CtxDebug(ctx, "OverwriteFields, target: %v, source: %v", target, source)
	if source == nil {
		return nil
	}

	sourceData, err := basicJSON.Marshal(source)
	if err != nil {
		logs.CtxError(ctx, "overwrite fields failed. unable to marshal the source, err: %v", err)
		return err
	}
	err = basicJSON.Unmarshal(sourceData, target)
	if err != nil {
		logs.CtxError(ctx, "overwrite fields failed. unable to unmarshal the target, err: %v", err)
		return err
	}
	return nil
}
