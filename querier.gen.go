// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package o11y

import (
	"context"
)

type Querier interface {
	GetAPIRequests(ctx context.Context) (RemoteApiRequest, error)
	StoreAPIRequest(ctx context.Context, arg StoreAPIRequestParams) error
}

var _ Querier = (*Queries)(nil)
