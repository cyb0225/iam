/**
@author: yeebing
@date: 2022/9/25
**/

package user

import (
	"github.com/cyb0225/iam/internal/apiserver/cache"
	v1 "github.com/cyb0225/iam/internal/apiserver/service/v1"
	"github.com/cyb0225/iam/internal/apiserver/store"
)

type User struct {
	srv v1.Service
}

func New(factory store.Factory, cache cache.Cache) *User {
	return &User{
		srv: v1.NewService(factory, cache),
	}
}
