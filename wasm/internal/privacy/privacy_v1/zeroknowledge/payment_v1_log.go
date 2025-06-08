package zkp

import (
	"chameleon-chain/common"
	agg "chameleon-chain/privacy/privacy_v1/zeroknowledge/aggregatedrange"
	oom "chameleon-chain/privacy/privacy_v1/zeroknowledge/oneoutofmany"
	snn "chameleon-chain/privacy/privacy_v1/zeroknowledge/serialnumbernoprivacy"
	snp "chameleon-chain/privacy/privacy_v1/zeroknowledge/serialnumberprivacy"
	utils "chameleon-chain/privacy/privacy_util"
)

type PaymentV1Logger struct {
	Log common.Logger
}

func (logger *PaymentV1Logger) Init(inst common.Logger) {
	logger.Log = inst
	agg.Logger.Init(inst)
	oom.Logger.Init(inst)
	snn.Logger.Init(inst)
	snp.Logger.Init(inst)
	utils.Logger.Init(inst)
}

// Global instant to use
var Logger = PaymentV1Logger{}