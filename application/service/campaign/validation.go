package campaign

import (
	"context"
)

func CheckAIOAccountStatus(ctx context.Context, aioAccountID int64, accountType client.AccountType) error {
	account, err := rpc.GetClientInfo(ctx, aioAccountID)
	if err != nil {
		logs.CtxError(ctx, "failed to get info for account id: %v", aioAccountID)
		return err
	}
	if account.GetAccountType() != accountType {
		return errors.NewEMPErrorWithMsg(creative_one.CommonParamCheckError, "AIO account type does not match expected account type")
	}
	return nil
}

func GetAIOAccount(ctx context.Context, aioAccountID int64) (account *client.AIOAccount, err error) {
	account, err = rpc.GetClientInfo(ctx, aioAccountID)
	if err != nil {
		logs.CtxError(ctx, "failed to get info for account id: %v", aioAccountID)
		return nil, err
	}
	return
}
