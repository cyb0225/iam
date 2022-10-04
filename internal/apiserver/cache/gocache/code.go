/**
@author: yeebing
@date: 2022/10/2
**/

package gocache

import (
	"context"
	"errors"
	"github.com/cyb0225/iam/internal/apiserver/cache"

	errcode "github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/pkg/errno"
	"time"
)

type codes struct {
	*cacheStore
}

var _ cache.CodeCache = (*codes)(nil)
var codeTimeout = time.Second * 120

func newCodes(store *cacheStore) *codes {
	return &codes{cacheStore: store}
}

func (c *codes) Create(ctx context.Context, code string, value *cache.CodeValue) error {
	c.ca.Set(code, value, codeTimeout)
	return nil
}

func (c *codes) Delete(ctx context.Context, code string) error {
	c.ca.Delete(code)
	return nil
}

func (c *codes) Get(ctx context.Context, code string) (*cache.CodeValue, error) {
	value, found := c.ca.Get(code)
	if !found {
		return nil, errno.WithCode(errcode.ErrCodeNotExisted, errors.New("code not existed"))
	}

	cv, ok := value.(*cache.CodeValue)
	if !ok {
		return nil, errno.WithCode(errcode.ErrTypeAssertion, errors.New("type assert to CodeValue failed"))
	}

	return cv, nil
}
