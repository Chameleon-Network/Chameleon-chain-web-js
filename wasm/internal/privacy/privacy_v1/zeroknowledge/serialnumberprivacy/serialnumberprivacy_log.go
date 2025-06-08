package serialnumberprivacy

import "chameleon-chain/common"

type SerialnumberprivacyLogger struct {
	Log common.Logger
}

func (logger *SerialnumberprivacyLogger) Init(inst common.Logger) {
	logger.Log = inst
}

// Global instant to use
var Logger = SerialnumberprivacyLogger{}
