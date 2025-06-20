package proof

import (
	// "chameleon-chain/common"
	"chameleon-chain/privacy/coin"
	errhandler "chameleon-chain/privacy/errorhandler"
	// "chameleon-chain/privacy/key"
	"chameleon-chain/privacy/proof/agg_interface"
)

// Paymentproof
type Proof interface {
	GetVersion() uint8
	Init()
	GetInputCoins() []coin.PlainCoin
	GetOutputCoins() []coin.Coin
	GetAggregatedRangeProof() agg_interface.AggregatedRangeProof

	SetInputCoins([]coin.PlainCoin) error
	SetOutputCoins([]coin.Coin) error

	Bytes() []byte
	SetBytes(proofBytes []byte) *errhandler.PrivacyError

	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error

	IsPrivacy() bool
	// ValidateSanity(interface{}) (bool, error)

	// Verify(boolParams map[string]bool, pubKey key.PublicKey, fee uint64, shardID byte, tokenID *common.Hash, additionalData interface{}) (bool, error)
}

