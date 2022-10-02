/**
@author: yeebing
@date: 2022/10/3
**/

package v1

import (
	"github.com/cyb0225/iam/internal/apiserver/cache"
	"github.com/cyb0225/iam/internal/apiserver/store"
)

type UserSrv interface {
}

var _ UserSrv = (*userService)(nil)

type userService struct {
	s store.Factory
	c cache.Cache
}

func newUsers(srv *service) *userService {
	return &userService{s: srv.s, c: srv.c}
}
