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
	if err := u.db.Create(&user).Error; err != nil {
		return errno.WithCode(code.ErrUserAlreadyExisted, err)
	}
	return nil
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
func (u *users) Update(ctx context.Context, userID uint64, val any) error {
	err := u.db.Model(&store.User{}).Where("id = ?", userID).Updates(val).Error
	if err != nil {
		return errno.WithCode(code.ErrDatabase, err)
	}
	return nil
}

// Get  user's information by id.
// val is used to return the result.
func (u *users) Get(ctx context.Context, userID uint64, val any) error {
	err := u.db.Model(&store.User{}).Where("id = ?", userID).First(val).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.WithCode(code.ErrUserNotFound, err)
		}
		return errno.WithCode(code.ErrDatabase, err)
	}
	return nil
}

func (u *users) GetByAccount(ctx context.Context, account string, val any) error {
	err := u.db.Model(&store.User{}).Where("account = ?", account).First(val).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.WithCode(code.ErrUserNotFound, err)
		}
		return errno.WithCode(code.ErrDatabase, err)
	}
	return nil
}

func (u *users) List(ctx context.Context, val any) error {
	err := u.db.Model(&store.User{}).Find(val).Error
	if err != nil {
		return errno.WithCode(code.ErrDatabase, err)
	}

	return nil
}
