/**
@author: yeebing
@date: 2022/10/2
**/

package cache

import (
	"context"
)

type TokenValue struct {
	UserID uint64
}

type TokenCache interface {
	Create(ctx context.Context, token string, value *TokenValue) error
	Get(ctx context.Context, token string) (*TokenValue, error)
	Delete(ctx context.Context, token string) error
}
