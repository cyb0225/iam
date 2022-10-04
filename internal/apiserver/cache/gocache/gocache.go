/**
@author: yeebing
@date: 2022/10/2
**/

package gocache

import (
	"github.com/cyb0225/iam/internal/apiserver/cache"
	gocache "github.com/patrickmn/go-cache"
)

var _ cache.Cache = (*cacheStore)(nil)

type cacheStore struct {
	ca *gocache.Cache
}

func New(ca *gocache.Cache) cache.Cache {
	return &cacheStore{ca: ca}
}

func (c *cacheStore) Code() cache.CodeCache {
	return newCodes(c)
}

func (c *cacheStore) Token() cache.TokenCache {
	return newTokens(c)
}
