package pdexv3

import (
	"encoding/json"

	metadataCommon "chameleon-chain/metadata/common"
)

type AddLiquidityRequest struct {
	poolPairID  string // only "" for the first contribution of pool
	pairHash    string
	otaReceiver string // receive nfct
	tokenID     string
	nftID       string
	tokenAmount uint64
	amplifier   uint // only set for the first contribution
	metadataCommon.MetadataBase
}

func NewAddLiquidity() *AddLiquidityRequest {
	return &AddLiquidityRequest{}
}

func NewAddLiquidityRequestWithValue(
	poolPairID, pairHash,
	otaReceiver,
	tokenID, nftID string, tokenAmount uint64, amplifier uint,
) *AddLiquidityRequest {
	metadataBase := metadataCommon.MetadataBase{
		Type: metadataCommon.Pdexv3AddLiquidityRequestMeta,
	}
	return &AddLiquidityRequest{
		poolPairID:   poolPairID,
		pairHash:     pairHash,
		otaReceiver:  otaReceiver,
		tokenID:      tokenID,
		nftID:        nftID,
		tokenAmount:  tokenAmount,
		amplifier:    amplifier,
		MetadataBase: metadataBase,
	}
}

func (request *AddLiquidityRequest) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(struct {
		PoolPairID  string `json:"PoolPairID"` // only "" for the first contribution of pool
		PairHash    string `json:"PairHash"`
		OtaReceiver string `json:"OtaReceiver"` // receive nfct
		TokenID     string `json:"TokenID"`
		NftID       string `json:"NftID"`
		TokenAmount uint64 `json:"TokenAmount"`
		Amplifier   uint   `json:"Amplifier"` // only set for the first contribution
		metadataCommon.MetadataBase
	}{
		PoolPairID:   request.poolPairID,
		PairHash:     request.pairHash,
		OtaReceiver:  request.otaReceiver,
		TokenID:      request.tokenID,
		NftID:        request.nftID,
		TokenAmount:  request.tokenAmount,
		Amplifier:    request.amplifier,
		MetadataBase: request.MetadataBase,
	})
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (request *AddLiquidityRequest) UnmarshalJSON(data []byte) error {
	temp := struct {
		PoolPairID  string                      `json:"PoolPairID"` // only "" for the first contribution of pool
		PairHash    string                      `json:"PairHash"`
		OtaReceiver string                      `json:"OtaReceiver"` // receive nfct
		TokenID     string                      `json:"TokenID"`
		NftID       string                      `json:"NftID"`
		TokenAmount metadataCommon.Uint64Reader `json:"TokenAmount"`
		Amplifier   uint                        `json:"Amplifier"` // only set for the first contribution
		metadataCommon.MetadataBase
	}{}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	request.poolPairID = temp.PoolPairID
	request.pairHash = temp.PairHash
	request.otaReceiver = temp.OtaReceiver
	request.tokenID = temp.TokenID
	request.nftID = temp.NftID
	request.tokenAmount = uint64(temp.TokenAmount)
	request.amplifier = temp.Amplifier
	request.MetadataBase = temp.MetadataBase
	return nil
}

func (request *AddLiquidityRequest) PoolPairID() string {
	return request.poolPairID
}

func (request *AddLiquidityRequest) PairHash() string {
	return request.pairHash
}

func (request *AddLiquidityRequest) OtaReceiver() string {
	return request.otaReceiver
}

func (request *AddLiquidityRequest) TokenID() string {
	return request.tokenID
}

func (request *AddLiquidityRequest) TokenAmount() uint64 {
	return request.tokenAmount
}

func (request *AddLiquidityRequest) Amplifier() uint {
	return request.amplifier
}

func (request *AddLiquidityRequest) NftID() string {
	return request.nftID
}
