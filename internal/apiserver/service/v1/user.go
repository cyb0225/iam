/**
@author: yeebing
@date: 2022/10/3
**/

package v1

import (
	"context"
	"errors"
	"github.com/cyb0225/iam/internal/apiserver/cache"
	"github.com/cyb0225/iam/internal/apiserver/service/v1/model"
	"github.com/cyb0225/iam/internal/apiserver/store"
	"github.com/cyb0225/iam/internal/pkg/code"
	"github.com/cyb0225/iam/internal/pkg/util"
	"github.com/cyb0225/iam/pkg/email"
	"github.com/cyb0225/iam/pkg/errno"
)

type UserSrv interface {
	Register(ctx context.Context, req *model.RegisterRequest) (*model.RegisterResponse, error)
	Get(ctx context.Context, userID uint64) (*model.UserGetResponse, error)
	GetCode(ctx context.Context, toEmail string) (*model.GetCodeResponse, error)
	Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	List(ctx context.Context) (*model.UserListResponse, error)
	ChangePassword(ctx context.Context, req *model.ChangePasswordRequest) error
	ChangeEmail(ctx context.Context, req *model.ChangeEmailRequest) error
	Logout(ctx context.Context) error
	Update(ctx context.Context, request *model.UserUpdateRequest) error
	UploadAvatar(ctx context.Context, avatar string) error
}

var _ UserSrv = (*userService)(nil)

type userService struct {
	s store.Factory
	c cache.Cache
}

func newUsers(srv *service) *userService {
	return &userService{s: srv.s, c: srv.c}
}

// Register create user.
func (u *userService) Register(ctx context.Context, req *model.RegisterRequest) (*model.RegisterResponse, error) {
	// check the code.
	if _, err := u.c.Code().Get(ctx, req.Code); err != nil {
		return nil, err
	}

	// check the password and the email
	if ok := util.JudgeEmail(req.Email); !ok {
		return nil, errno.WithCode(code.ErrEmailRequired, errors.New("email does not meet requirements"))
	}
	if err := util.JudgePassword(req.Password); err != nil {
		return nil, err
	}

	// create user.
	user := &store.User{
		Account:  req.Account,
		Password: util.PasswordEncrypt(req.Password),
		Nick:     req.Nick,
		Email:    req.Email,
	}
	if err := u.s.User().Create(ctx, user); err != nil {
		return nil, err
	}

	// create token
	token := util.GenerateToken(user.Account)
	// stored token
	_ = u.c.Token().Create(ctx, token, &cache.TokenValue{UserID: user.ID})
	return &model.RegisterResponse{Token: token}, nil
}

// Get user's information.
func (u *userService) Get(ctx context.Context, userID uint64) (*model.UserGetResponse, error) {
	res := &model.UserGetResponse{}
	if err := u.s.User().Get(ctx, userID, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetCode send code to the email
func (u *userService) GetCode(ctx context.Context, toEmail string) (*model.GetCodeResponse, error) {
	// check the email
	if ok := util.JudgeEmail(toEmail); !ok {
		return nil, errno.WithCode(code.ErrEmailRequired, errors.New("email does not meet requirements"))
	}

	// generate code.
	vcode := util.GenerateCode()

	// send email
	subject := "邮箱验证码"
	text := []byte("登录验证码为: " + vcode)
	if err := email.Send([]string{toEmail}, subject, text); err != nil {
		return nil, errno.WithCode(code.ErrSendEmail, err)
	}

	// stored code, it will not return an error
	_ = u.c.Code().Create(ctx, vcode, &cache.CodeValue{Email: vcode})
	return &model.GetCodeResponse{Code: vcode}, nil
}

func (u *userService) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	// check account and password
	en := util.PasswordEncrypt(req.Password)
	user := &store.User{}
	if err := u.s.User().GetByAccount(ctx, req.Account, user); err != nil {
		return nil, err
	}
	if en != user.Password {
		return nil, errno.WithCode(code.ErrPassword, errors.New("error password"))
	}

	// created a token
	token := util.GenerateToken(user.Account)
	// stored token
	_ = u.c.Token().Create(ctx, token, &cache.TokenValue{UserID: user.ID})
	return &model.LoginResponse{Token: token}, nil
}

func (u *userService) List(ctx context.Context) (*model.UserListResponse, error) {
	res := &model.UserListResponse{}
	err := u.s.User().List(ctx, &res.List)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userService) ChangePassword(ctx context.Context, req *model.ChangePasswordRequest) error {
	userID := ctx.Value("id").(uint64)
	if userID == 0 {
		return errno.WithCode(code.ErrGetUserIDFromCtx, errors.New("get user id from context failed"))
	}

	// check the old password.
	user := &store.User{}
	if err := u.s.User().Get(ctx, userID, user); err != nil {
		return err
	}
	if util.PasswordEncrypt(req.OldPassword) != user.Password {
		return errno.WithCode(code.ErrPassword, errors.New("error password"))
	}

	// check the strength of new password.
	if err := util.JudgePassword(req.OldPassword); err != nil {
		return err
	}

	// update user's password field.
	updateUser := &store.User{
		Password: util.PasswordEncrypt(req.NewPassword),
	}
	return u.s.User().Update(ctx, userID, updateUser)
}

func (u *userService) ChangeEmail(ctx context.Context, req *model.ChangeEmailRequest) error {
	userID := ctx.Value("id").(uint64)
	if userID == 0 {
		return errno.WithCode(code.ErrGetUserIDFromCtx, errors.New("get user id from context failed"))
	}

	// check the code.
	if _, err := u.c.Code().Get(ctx, req.Code); err != nil {
		return err
	}

	// updates user's email field.
	user := &store.User{
		Email: req.NewEmail,
	}
	return u.s.User().Update(ctx, userID, user)
}

func (u *userService) Logout(ctx context.Context) error {
	token := ctx.Value("token").(string)
	if len(token) == 0 {
		return errno.WithCode(code.ErrGetTokenFromCtx, errors.New("get token from context failed"))
	}

	// delete token in cache.
	_ = u.c.Token().Delete(ctx, token)
	return nil
}

func (u *userService) Update(ctx context.Context, request *model.UserUpdateRequest) error {
	userID := ctx.Value("id").(uint64)
	if userID == 0 {
		return errno.WithCode(code.ErrGetUserIDFromCtx, errors.New("get user id from context failed"))
	}

	// updates the user's information
	return u.s.User().Update(ctx, userID, request)
}

func (u *userService) UploadAvatar(ctx context.Context, avatar string) error {

	return nil
}
