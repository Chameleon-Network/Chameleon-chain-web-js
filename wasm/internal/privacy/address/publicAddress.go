package address

import (
	"chameleon-chain/privacy/operation"
)

type PublicAddress struct {
	publicSpend *operation.Point
	publicView  *operation.Point
}

func (this *PublicAddress) GetPublicSpend() *operation.Point {
	return this.publicSpend
}

func (this *PublicAddress) GetPublicView() *operation.Point {
	return this.publicView
}
