/**
@author: yeebing
@date: 2022/10/2
**/

package cache

import (
	gocache "github.com/patrickmn/go-cache"
	"time"
)

var (
	Ca *gocache.Cache
)

func New() (*gocache.Cache, error) {
	Ca := gocache.New(5*time.Minute, 10*time.Minute)
	return Ca, nil
}
