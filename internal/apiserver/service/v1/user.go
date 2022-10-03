/**
@author: yeebing
@date: 2022/10/3
**/

package v1

import (
	"context"
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

func (u *userService) Register(ctx context.Context) {

}

func (u *userService) Get(ctx context.Context, userID uint64) {

}

// GetCode send code to the email
func (u *userService) GetCode(ctx context.Context, email string) {
	// generate code.

	// send email

	// stored code
}

func (u *userService) ChangePassword(ctx context.Context) {
	// check the code.

	// check the old password

	// update user's password field.

}

func (u *userService) ChangeEmail(ctx context.Context) {
	// check the code.

	// updates user's email field.

}

func (u *userService) Login(ctx context.Context) {
	// check account and password

	// created a token

	// stored token in the cache.

}

func (u *userService) Logout(ctx context.Context) {
	// delete token in cache.

}

func (u *userService) Update(ctx context.Context) {
	// updates the user's information

}

func (u *userService) UploadAvatar(ctx context.Context) {

}

func (u *userService) List(ctx context.Context) {

}
