package metadata

import (
	"errors"
	"strconv"

	"chameleon-chain/common"
)

type StopAutoStakingMetadata struct {
	MetadataBaseWithSignature
	CommitteePublicKey string
}

func (meta *StopAutoStakingMetadata) Hash() *common.Hash {
	record := strconv.Itoa(meta.Type)
	data := []byte(record)
	data = append(data, meta.Sig...)
	hash := common.HashH(data)
	return &hash
}

func (meta *StopAutoStakingMetadata) HashWithoutSig() *common.Hash {
	return meta.MetadataBase.Hash()
}

func NewStopAutoStakingMetadata(stopStakingType int, committeePublicKey string) (*StopAutoStakingMetadata, error) {
	if stopStakingType != StopAutoStakingMeta {
		return nil, errors.New("invalid stop staking type")
	}
	metadataBase := NewMetadataBaseWithSignature(stopStakingType)
	return &StopAutoStakingMetadata{
		MetadataBaseWithSignature:       *metadataBase,
		CommitteePublicKey: committeePublicKey,
	}, nil
}

func (stopAutoStakingMetadata StopAutoStakingMetadata) GetType() int {
	return stopAutoStakingMetadata.Type
}