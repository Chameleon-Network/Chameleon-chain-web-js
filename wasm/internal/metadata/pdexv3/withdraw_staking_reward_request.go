package pdexv3

import (
	"chameleon-chain/common"
	metadataCommon "chameleon-chain/metadata/common"
	"chameleon-chain/privacy"
)

type WithdrawalStakingRewardRequest struct {
	metadataCommon.MetadataBase
	StakingPoolID string                              `json:"StakingPoolID"`
	NftID         common.Hash                         `json:"NftID"`
	Receivers     map[common.Hash]privacy.OTAReceiver `json:"Receivers"`
}

func NewPdexv3WithdrawalStakingRewardRequest(
	metaType int,
	stakingToken string,
	nftID common.Hash,
	receivers map[common.Hash]privacy.OTAReceiver,
) (*WithdrawalStakingRewardRequest, error) {
	metadataBase := metadataCommon.NewMetadataBase(metaType)

	return &WithdrawalStakingRewardRequest{
		MetadataBase:  *metadataBase,
		StakingPoolID: stakingToken,
		NftID:         nftID,
		Receivers:     receivers,
	}, nil
}
