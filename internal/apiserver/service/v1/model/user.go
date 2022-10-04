/**
@author: yeebing
@date: 2022/10/3
**/

package model

type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

// UserGetResponse Information used by users to display
type UserGetResponse struct {
	ID           uint64 `json:"id"`
	Nick         string `json:"nick"`
	Account      string `json:"account"`
	Email        string `json:"email"`
	Introduction string `json:"introduction"`
	University   string `json:"university"`
	Company      string `json:"company"`
	Blog         string `json:"blog"`
	Github       string `json:"github"`
}

type GetCodeResponse struct {
	Code string `json:"code"`
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserListResponse struct {
	List []UserGetResponse `json:"list"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ChangeEmailRequest struct {
	Code     string `json:"code"`
	NewEmail string `json:"newEmail"`
}

type UserUpdateRequest struct {
}
