package mint

import (
	"context"
	"fmt"
	"math/big"

	"github.com/spolu/settle/mint/model"
)

// AssetResource is the representation of an asset in the mint API.
type AssetResource struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`

	Name   string `json:"name"`
	Issuer string `json:"issuer"`
	Code   string `json:"code"`
	Scale  int8   `json:"scale"`
}

// NewAssetResource generates a new resource.
func NewAssetResource(
	ctx context.Context,
	asset *model.Asset,
	issuer *model.User,
	host string,
) AssetResource {
	return AssetResource{
		ID: fmt.Sprintf(
			"%s@%s[%s]", issuer.Username, host, asset.Token),
		Created: asset.Created.UnixNano() / (1000 * 1000),
		Name: fmt.Sprintf(
			"%s@%s[%s.%d]",
			issuer.Username, host, asset.Code, asset.Scale,
		),
		Issuer: fmt.Sprintf(
			"%s@%s", issuer.Username, host,
		),
		Code:  asset.Code,
		Scale: asset.Scale,
	}
}

// OperationResource is the representation of an operation in the mint API.
type OperationResource struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`

	Asset       AssetResource `json:"asset"`
	Source      *string       `json:"source"`
	Destination *string       `json:"destination"`
	Amount      *big.Int      `json:"amount"`
}

// NewOperationResource generates a new resource.
func NewOperationResource(
	ctx context.Context,
	operation *model.Operation,
	assetResource AssetResource,
) OperationResource {
	return OperationResource{
		ID: fmt.Sprintf(
			"%s[%s]", assetResource.Issuer, operation.Token),
		Created:     operation.Created.UnixNano() / (1000 * 1000),
		Asset:       assetResource,
		Source:      operation.Source,
		Destination: operation.Destination,
		Amount:      (*big.Int)(&operation.Amount),
	}
}

// OfferResource is the representation of an offer in the mint API.
type OfferResource struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`

	Pair   string   `json:"pair"`
	Price  string   `json:"price"`
	Amount *big.Int `json:"amount"`
	Status string   `json:"status"`
}

// NewOfferResource generates a new resource.
func NewOfferResource(
	ctx context.Context,
	offer *model.Offer,
) OfferResource {
	return OfferResource{
		ID: fmt.Sprintf(
			"%s[%s]", offer.Owner, offer.Token),
		Created: offer.Created.UnixNano() / (1000 * 1000),

		Pair: fmt.Sprintf("%s/%s", offer.BaseAsset, offer.QuoteAsset),
		Price: fmt.Sprintf(
			"%s/%s",
			(*big.Int)(&offer.BasePrice).String(),
			(*big.Int)(&offer.QuotePrice).String()),
		Amount: (*big.Int)(&offer.Amount),
		Status: string(offer.Status),
	}
}
