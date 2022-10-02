/**
@author: yeebing
@date: 2022/10/2
**/

package gocache

import (
	"context"
	"errors"
	"github.com/cyb0225/iam/internal/apiserver/cache"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/pkg/errno"
	"time"
)

var _ cache.TokenCache = (*tokens)(nil)

type tokens struct {
	*cacheStore
}

func newTokens(store *cacheStore) *tokens {
	return &tokens{cacheStore: store}
}

func (t *tokens) Create(ctx context.Context, token string, value *cache.TokenValue, timeout time.Duration) error {
	t.ca.Set(token, value, timeout)
	return nil
}

func (t *tokens) Get(ctx context.Context, token string) (*cache.TokenValue, error) {
	value, found := t.ca.Get(token)
	if !found {
		return nil, errno.WithCode(code.ErrTokenNotExisted, errors.New("token not existed"))
	}

	tv, ok := value.(*cache.TokenValue)
	if !ok {
		return nil, errno.WithCode(code.ErrTypeAssertion, errors.New("type assert to TokenValue failed"))
	}

	return tv, nil
}

func (t *tokens) Delete(ctx context.Context, token string) error {
	t.ca.Delete(token)
	return nil
}
