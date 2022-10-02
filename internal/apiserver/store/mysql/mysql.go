/**
@author: yeebing
@date: 2022/10/2
**/

package mysql

import (
	"github.com/cyb0225/iam/internal/apiserver/store"
	"gorm.io/gorm"
)

var _ store.Factory = (*datastore)(nil)

type datastore struct {
	db *gorm.DB
}

func New(db *gorm.DB) store.Factory {
	return &datastore{db: db}
}

func (ds *datastore) User() store.UserStore {
	return newUsers(ds)
}
