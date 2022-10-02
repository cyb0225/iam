/**
@author: yeebing
@date: 2022/10/2
**/

package redis

type codes struct {
	*redisCache
}

func newCodes(rc *redisCache) *codes {
	return &codes{redisCache: rc}
}
