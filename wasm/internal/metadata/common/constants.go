package common

import (
	"strconv"
)

const (
	InvalidMeta = 1

	IssuingRequestMeta     = 24
	IssuingResponseMeta    = 25
	ContractingRequestMeta = 26
	IssuingETHRequestMeta  = 80
	IssuingETHResponseMeta = 81

	ShardBlockReward             = 36
	AcceptedBlockRewardInfoMeta  = 37
	ShardBlockSalaryResponseMeta = 38
	BeaconRewardRequestMeta      = 39
	BeaconSalaryResponseMeta     = 40
	ReturnStakingMeta            = 41
	IncDAORewardRequestMeta      = 42
	ShardBlockRewardRequestMeta  = 43
	WithDrawRewardRequestMeta    = 44
	WithDrawRewardResponseMeta   = 45

	//staking
	ShardStakingMeta    = 63
	StopAutoStakingMeta = 127
	BeaconStakingMeta   = 64
	UnStakingMeta       = 210

	// chameleon -> Ethereum bridge
	BeaconSwapConfirmMeta = 70
	BridgeSwapConfirmMeta = 71
	BurningRequestMeta    = 27
	BurningRequestMetaV2  = 240
	BurningConfirmMeta    = 72
	BurningConfirmMetaV2  = 241

	// pde
	PDEContributionMeta                   = 90
	PDETradeRequestMeta                   = 91
	PDETradeResponseMeta                  = 92
	PDEWithdrawalRequestMeta              = 93
	PDEWithdrawalResponseMeta             = 94
	PDEContributionResponseMeta           = 95
	PDEPRVRequiredContributionRequestMeta = 204
	PDECrossPoolTradeRequestMeta          = 205
	PDECrossPoolTradeResponseMeta         = 206
	PDEFeeWithdrawalRequestMeta           = 207
	PDEFeeWithdrawalResponseMeta          = 208
	PDETradingFeesDistributionMeta        = 209

	// erc20/bep20 for prv token
	IssuingPRVERC20RequestMeta  = 270
	IssuingPRVERC20ResponseMeta = 271
	IssuingPRVBEP20RequestMeta  = 272
	IssuingPRVBEP20ResponseMeta = 273
	BurningPRVERC20RequestMeta  = 274
	BurningPRVERC20ConfirmMeta  = 150
	BurningPRVBEP20RequestMeta  = 275
	BurningPRVBEP20ConfirmMeta  = 151

	// erc20/bep20 for pdex token
	IssuingPDEXERC20RequestMeta  = 320
	IssuingPDEXERC20ResponseMeta = 321
	IssuingPDEXBEP20RequestMeta  = 322
	IssuingPDEXBEP20ResponseMeta = 323
	BurningPDEXERC20RequestMeta  = 324
	BurningPDEXERC20ConfirmMeta  = 152
	BurningPDEXBEP20RequestMeta  = 325
	BurningPDEXBEP20ConfirmMeta  = 153

	// pancake integration
	BurningPBSCForDepositToSCRequestMeta = 326

	// pDEX v3
	Pdexv3ModifyParamsMeta                  = 280
	Pdexv3AddLiquidityRequestMeta           = 281
	Pdexv3AddLiquidityResponseMeta          = 282
	Pdexv3WithdrawLiquidityRequestMeta      = 283
	Pdexv3WithdrawLiquidityResponseMeta     = 284
	Pdexv3TradeRequestMeta                  = 285
	Pdexv3TradeResponseMeta                 = 286
	Pdexv3AddOrderRequestMeta               = 287
	Pdexv3AddOrderResponseMeta              = 288
	Pdexv3WithdrawOrderRequestMeta          = 289
	Pdexv3WithdrawOrderResponseMeta         = 290
	Pdexv3UserMintNftRequestMeta            = 291
	Pdexv3UserMintNftResponseMeta           = 292
	Pdexv3MintNftRequestMeta                = 293
	Pdexv3MintNftResponseMeta               = 294
	Pdexv3StakingRequestMeta                = 295
	Pdexv3StakingResponseMeta               = 296
	Pdexv3UnstakingRequestMeta              = 297
	Pdexv3UnstakingResponseMeta             = 298
	Pdexv3WithdrawLPFeeRequestMeta          = 299
	Pdexv3WithdrawLPFeeResponseMeta         = 300
	Pdexv3WithdrawProtocolFeeRequestMeta    = 301
	Pdexv3WithdrawProtocolFeeResponseMeta   = 302
	Pdexv3MintPDEXGenesisMeta               = 303
	Pdexv3MintBlockRewardMeta               = 304
	Pdexv3DistributeStakingRewardMeta       = 305
	Pdexv3WithdrawStakingRewardRequestMeta  = 306
	Pdexv3WithdrawStakingRewardResponseMeta = 307

	// portal
	PortalCustodianDepositMeta                  = 100
	PortalRequestPortingMeta                    = 101
	PortalUserRequestPTokenMeta                 = 102
	PortalCustodianDepositResponseMeta          = 103
	PortalUserRequestPTokenResponseMeta         = 104
	PortalExchangeRatesMeta                     = 105
	PortalRedeemRequestMeta                     = 106
	PortalRedeemRequestResponseMeta             = 107
	PortalRequestUnlockCollateralMeta           = 108
	PortalCustodianWithdrawRequestMeta          = 110
	PortalCustodianWithdrawResponseMeta         = 111
	PortalLiquidateCustodianMeta                = 112
	PortalLiquidateCustodianResponseMeta        = 113
	PortalLiquidateTPExchangeRatesMeta          = 114
	PortalExpiredWaitingPortingReqMeta          = 116
	PortalRewardMeta                            = 117
	PortalRequestWithdrawRewardMeta             = 118
	PortalRequestWithdrawRewardResponseMeta     = 119
	PortalRedeemFromLiquidationPoolMeta         = 120
	PortalRedeemFromLiquidationPoolResponseMeta = 121
	PortalCustodianTopupMeta                    = 122
	PortalCustodianTopupResponseMeta            = 123
	PortalTotalRewardCustodianMeta              = 124
	PortalPortingResponseMeta                   = 125
	PortalReqMatchingRedeemMeta                 = 126
	PortalPickMoreCustodianForRedeemMeta        = 128
	PortalCustodianTopupMetaV2                  = 129
	PortalCustodianTopupResponseMetaV2          = 130

	// Portal v3
	PortalCustodianDepositMetaV3                  = 131
	PortalCustodianWithdrawRequestMetaV3          = 132
	PortalRewardMetaV3                            = 133
	PortalRequestUnlockCollateralMetaV3           = 134
	PortalLiquidateCustodianMetaV3                = 135
	PortalLiquidateByRatesMetaV3                  = 136
	PortalRedeemFromLiquidationPoolMetaV3         = 137
	PortalRedeemFromLiquidationPoolResponseMetaV3 = 138
	PortalCustodianTopupMetaV3                    = 139
	PortalTopUpWaitingPortingRequestMetaV3        = 140
	PortalRequestPortingMetaV3                    = 141
	PortalRedeemRequestMetaV3                     = 142
	PortalUnlockOverRateCollateralsMeta           = 143

	// chameleon => Ethereum's SC for portal
	PortalCustodianWithdrawConfirmMetaV3         = 170
	PortalRedeemFromLiquidationPoolConfirmMetaV3 = 171
	PortalLiquidateRunAwayCustodianConfirmMetaV3 = 172

	//Note: don't use this metadata type for others
	PortalResetPortalDBMeta = 199

	// relaying
	RelayingBNBHeaderMeta = 200
	RelayingBTCHeaderMeta = 201

	PortalTopUpWaitingPortingRequestMeta  = 202
	PortalTopUpWaitingPortingResponseMeta = 203

	// chameleon mode for smart contract
	BurningForDepositToSCRequestMeta   = 96
	BurningForDepositToSCRequestMetaV2 = 242
	BurningConfirmForDepositToSCMeta   = 97
	BurningConfirmForDepositToSCMetaV2 = 243

	InitTokenRequestMeta  = 244
	InitTokenResponseMeta = 245

	// chameleon mode for bsc
	IssuingBSCRequestMeta  = 250
	IssuingBSCResponseMeta = 251
	BurningPBSCRequestMeta = 252
	BurningBSCConfirmMeta  = 253

	// chameleon mode for polygon
	IssuingPLGRequestMeta  = 327
	IssuingPLGResponseMeta = 328
	BurningPLGRequestMeta  = 329
	BurningPLGConfirmMeta  = 153

	// pPLG ( Polygon )
	BurningPLGForDepositToSCRequestMeta = 330
	BurningPLGConfirmForDepositToSCMeta = 154

	// chameleon mode for Fantom
	IssuingFantomRequestMeta  = 331
	IssuingFantomResponseMeta = 332
	BurningFantomRequestMeta  = 333
	BurningFantomConfirmMeta  = 155

	// pFantom ( Fantom )
	BurningFantomForDepositToSCRequestMeta = 334
	BurningFantomConfirmForDepositToSCMeta = 156

	// portal v4
	PortalV4UnshieldRequestMeta = 262
)

var minerCreatedMetaTypes = []int{
	ShardBlockReward,
	BeaconSalaryResponseMeta,
	IssuingResponseMeta,
	IssuingETHResponseMeta,
	IssuingBSCResponseMeta,
	ReturnStakingMeta,
	WithDrawRewardResponseMeta,
	PDETradeResponseMeta,
	PDECrossPoolTradeResponseMeta,
	PDEWithdrawalResponseMeta,
	PDEFeeWithdrawalResponseMeta,
	PDEContributionResponseMeta,
	PortalUserRequestPTokenResponseMeta,
	PortalCustodianDepositResponseMeta,
	PortalRedeemRequestResponseMeta,
	PortalCustodianWithdrawResponseMeta,
	PortalLiquidateCustodianResponseMeta,
	PortalRequestWithdrawRewardResponseMeta,
	PortalRedeemFromLiquidationPoolResponseMeta,
	PortalCustodianTopupResponseMeta,
	PortalCustodianTopupResponseMetaV2,
	PortalPortingResponseMeta,
	PortalTopUpWaitingPortingResponseMeta,
	PortalRedeemFromLiquidationPoolResponseMetaV3,
	InitTokenResponseMeta,
	Pdexv3AddLiquidityResponseMeta,
	Pdexv3MintNftResponseMeta,
	Pdexv3WithdrawLiquidityResponseMeta,
	Pdexv3TradeResponseMeta,
	Pdexv3AddOrderResponseMeta,
	Pdexv3WithdrawOrderResponseMeta,
}

// Special rules for shardID: stored as 2nd param of instruction of BeaconBlock
const (
	AllShards  = -1
	BeaconOnly = -2
)

/*var (*/
//// if the blockchain is running in Docker container
//// then using GETH_NAME env's value (aka geth container name)
//// otherwise using localhost
//EthereumLightNodeHost     = utils.GetEnv("GETH_NAME", "127.0.0.1")
//EthereumLightNodeProtocol = utils.GetEnv("GETH_PROTOCOL", "http")
//EthereumLightNodePort     = utils.GetEnv("GETH_PORT", "8545")
/*)*/

const (
	StopAutoStakingAmount = 0
	EVMConfirmationBlocks = 15
)

var AcceptedWithdrawRewardRequestVersion = []int{0, 1}

var portalMetaTypesV3 = []int{
	PortalCustodianDepositMeta,
	PortalRequestPortingMeta,
	PortalUserRequestPTokenMeta,
	PortalCustodianDepositResponseMeta,
	PortalUserRequestPTokenResponseMeta,
	PortalExchangeRatesMeta,
	PortalRedeemRequestMeta,
	PortalRedeemRequestResponseMeta,
	PortalRequestUnlockCollateralMeta,
	PortalCustodianWithdrawRequestMeta,
	PortalCustodianWithdrawResponseMeta,
	PortalLiquidateCustodianMeta,
	PortalLiquidateCustodianResponseMeta,
	PortalLiquidateTPExchangeRatesMeta,
	PortalExpiredWaitingPortingReqMeta,
	PortalRewardMeta,
	PortalRequestWithdrawRewardMeta,
	PortalRequestWithdrawRewardResponseMeta,
	PortalRedeemFromLiquidationPoolMeta,
	PortalRedeemFromLiquidationPoolResponseMeta,
	PortalCustodianTopupMeta,
	PortalCustodianTopupResponseMeta,
	PortalTotalRewardCustodianMeta,
	PortalPortingResponseMeta,
	PortalReqMatchingRedeemMeta,
	PortalPickMoreCustodianForRedeemMeta,
	PortalCustodianTopupMetaV2,
	PortalCustodianTopupResponseMetaV2,

	// Portal v3
	PortalCustodianDepositMetaV3,
	PortalCustodianWithdrawRequestMetaV3,
	PortalRewardMetaV3,
	PortalRequestUnlockCollateralMetaV3,
	PortalLiquidateCustodianMetaV3,
	PortalLiquidateByRatesMetaV3,
	PortalRedeemFromLiquidationPoolMetaV3,
	PortalRedeemFromLiquidationPoolResponseMetaV3,
	PortalCustodianTopupMetaV3,
	PortalTopUpWaitingPortingRequestMetaV3,
	PortalRequestPortingMetaV3,
	PortalRedeemRequestMetaV3,
	PortalCustodianWithdrawConfirmMetaV3,
	PortalRedeemFromLiquidationPoolConfirmMetaV3,
	PortalLiquidateRunAwayCustodianConfirmMetaV3,
	PortalResetPortalDBMeta,

	PortalTopUpWaitingPortingRequestMeta,
	PortalTopUpWaitingPortingResponseMeta,
}

var portalRelayingMetaTypes = []int{
	RelayingBNBHeaderMeta,
	RelayingBTCHeaderMeta,
}

var bridgeMetas = []string{
	strconv.Itoa(BeaconSwapConfirmMeta),
	strconv.Itoa(BridgeSwapConfirmMeta),
	strconv.Itoa(BurningConfirmMeta),
	strconv.Itoa(BurningConfirmForDepositToSCMeta),
	strconv.Itoa(BurningConfirmMetaV2),
	strconv.Itoa(BurningConfirmForDepositToSCMetaV2),
	strconv.Itoa(BurningBSCConfirmMeta),
}
