/**
@author: yeebing
@date: 2022/10/3
**/

package v1

import (
	"github.com/cyb0225/iam/internal/apiserver/cache"
	"github.com/cyb0225/iam/internal/apiserver/store"
)

type Service interface {
	User() UserSrv
}

var _ Service = (*service)(nil)

type service struct {
	s store.Factory
	c cache.Cache
}

func NewService(s store.Factory, c cache.Cache) Service {
	return &service{s: s, c: c}
}

func (s *service) User() UserSrv {
	return newUsers(s)
}
