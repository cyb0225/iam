/**
@author: yeebing
@date: 2022/10/2
**/

package mysql

import (
	"context"
	"errors"
	"github.com/cyb0225/iam/internal/apiserver/store"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/pkg/errno"
	"gorm.io/gorm"
)

type users struct {
	*datastore
}

var _ store.UserStore = (*users)(nil)

func newUsers(ds *datastore) *users {
	return &users{
		datastore: ds,
	}
}

func (u *users) Create(ctx context.Context, user *store.User) error {
	return u.db.Create(&user).Error
}

func (u *users) Delete(ctx context.Context, userID uint64) error {
	user := &store.User{}
	err := u.db.Where("id = ?", userID).Delete(&user).Error
	if err != nil {
		return errno.WithCode(code.ErrDatabase, err)
	}
	return nil
}

// Update user's Information
func (u *users) Update(ctx context.Context, userID uint64, user *store.User) error {
	err := u.db.Model(&store.User{}).Where("id = ?", userID).Updates(user).Error
	if err != nil {
		return errno.WithCode(code.ErrDatabase, err)
	}
	return nil
}

// Get get user's information by id.
func (u *users) Get(ctx context.Context, userID uint64) (*store.User, error) {
	user := &store.User{}
	err := u.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.WithCode(code.ErrUserNotFound, err)
		}
		return nil, errno.WithCode(code.ErrDatabase, err)
	}
	return user, nil
}

func (u *users) List(ctx context.Context) (*store.UserList, error) {
	ret := &store.UserList{}
	err := u.db.Find(&ret.Items).Error
	if err != nil {
		return nil, errno.WithCode(code.ErrDatabase, err)
	}
	return ret, err
}
