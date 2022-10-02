/**
@author: yeebing
@date: 2022/9/25
**/

package store

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint64         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Account      string         `gorm:"column:account;type:varchar(255);comment:账号，唯一;NOT NULL" json:"account"`
	Password     string         `gorm:"column:password;type:varchar(255);comment:加密后的密码;NOT NULL" json:"password"`
	Nick         string         `gorm:"column:nick;type:varchar(255);comment:用户昵称;NOT NULL" json:"nick"`
	Email        string         `gorm:"column:email;type:varchar(255);comment:邮箱;NOT NULL" json:"email"`
	Introduction string         `gorm:"column:introduction;type:varchar(255);comment:个人简介" json:"introduction"`
	University   string         `gorm:"column:university;type:varchar(255);comment:大学" json:"university"`
	Company      string         `gorm:"column:company;type:varchar(255);comment:公司" json:"company"`
	Blog         string         `gorm:"column:blog;type:varchar(255);comment:博客地址" json:"blog"`
	Github       string         `gorm:"column:github;type:varchar(255);comment:github 地址" json:"github"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;NOT NULL" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;NOT NULL" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (m *User) TableName() string {
	return "user"
}

type UserList struct {
	Items []*User
}

type UserStore interface {
	Create(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID uint64) error
	Update(ctx context.Context, userID uint64, user *User) error
	Get(ctx context.Context, userID uint64) (*User, error)
	List(ctx context.Context) (*UserList, error)
}
