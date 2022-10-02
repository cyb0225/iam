/**
@author: yeebing
@date: 2022/10/2
**/

package redis

import "github.com/go-redis/redis"

type redisCache struct {
	db *redis.Client
}

func (rc *redisCache) Token() {

}
