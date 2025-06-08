package pdexv3

import (
	"chameleon-chain/common"
	metadataCommon "chameleon-chain/metadata/common"
	"chameleon-chain/privacy"
)

type WithdrawalProtocolFeeRequest struct {
	metadataCommon.MetadataBaseWithSignature
	PoolPairID string                              `json:"PoolPairID"`
	Receivers  map[common.Hash]privacy.OTAReceiver `json:"Receivers"`
}

type WithdrawalProtocolFeeContent struct {
	PoolPairID string                       `json:"PoolPairID"`
	TokenID    common.Hash                  `json:"TokenID"`
	Receivers  map[common.Hash]ReceiverInfo `json:"Receivers"`
	TxReqID    common.Hash                  `json:"TxReqID"`
	ShardID    byte                         `json:"ShardID"`
}

type WithdrawalProtocolFeeStatus struct {
	Status    int                          `json:"Status"`
	Receivers map[common.Hash]ReceiverInfo `json:"Receivers"`
}

func NewPdexv3WithdrawalProtocolFeeRequest(
	metaType int,
	pairID string,
	receivers map[common.Hash]privacy.OTAReceiver,
) (*WithdrawalProtocolFeeRequest, error) {
	metadataBase := metadataCommon.NewMetadataBaseWithSignature(metaType)

	return &WithdrawalProtocolFeeRequest{
		MetadataBaseWithSignature: *metadataBase,
		PoolPairID:                pairID,
		Receivers:                 receivers,
	}, nil
}
