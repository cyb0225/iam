/**
@author: yeebing
@date: 2022/10/2
**/

package redis

type tokens struct {
	*redisCache
}

func newTokens(rc *redisCache) *tokens {
	return &tokens{redisCache: rc}
}
