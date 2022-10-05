/**
@author: yeebing
@date: 2022/10/2
**/

package mysql

import (
	"context"
	"github.com/cyb0225/iam/internal/apiserver/store"
	"github.com/cyb0225/iam/pkg/db"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
)

var (
	u *users
)

func TestMain(m *testing.M) {
	opts := db.Option{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		Database: "iam_test",
		LogFile:  "stdout",
	}

	DB, err := db.New(opts)
	if err != nil {
		log.Fatal(err)
	}
	u = newUsers(&datastore{db: DB})
	m.Run()
	err = u.db.Exec("DELETE FROM user").Error
	if err != nil {
		log.Fatal(err)
	}
}

func TestUsers_Create(t *testing.T) {
	clear(t)
	user := &store.User{
		Account:  "account",
		Password: "password",
		Email:    "email",
		Nick:     "nick",
	}
	t.Run("return nil", func(t *testing.T) {
		err := u.Create(context.Background(), user)
		assert.Equal(t, nil, err)
		if user.ID == 0 {
			t.Fatalf("can not get user's ID after create the record.")
		}
	})

	t.Run("user already existed", func(t *testing.T) {
		err := u.Create(context.Background(), user)
		if err == nil {
			t.Fatalf("want get an error(create user twice), but got nil.\n")
		}
	})

}

type UserDemo struct {
	ID    uint64
	Nick  string
	Email string
}

func TestUsers_Get(t *testing.T) {
	clear(t)

	user := &store.User{
		Account:  "account",
		Password: "password",
		Email:    "email",
		Nick:     "nick",
	}

	err := u.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	t.Run("return nil", func(t *testing.T) {
		got := &UserDemo{}
		err := u.Get(context.Background(), user.ID, got)
		assert.Equal(t, nil, err)
		assert.Equal(t, user.ID, got.ID)
		assert.Equal(t, user.Nick, got.Nick)
	})

	t.Run("record not found", func(t *testing.T) {
		got := UserDemo{}
		err := u.Get(context.Background(), user.ID+1, &got)
		t.Log("record not found:", err)
		assert.NotEqual(t, nil, err)
	})

}

func TestUsers_GetByAccount(t *testing.T) {
	clear(t)

	user := &store.User{
		Account:  "account",
		Password: "password",
		Email:    "email",
		Nick:     "nick",
	}

	err := u.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	t.Run("return nil", func(t *testing.T) {
		got := &store.User{}
		err := u.GetByAccount(context.Background(), user.Account, got)
		assert.Equal(t, nil, err)
		assert.Equal(t, user.Account, got.Account)
	})

	t.Run("record not found", func(t *testing.T) {
		got := &store.User{}
		err := u.GetByAccount(context.Background(), user.Account+user.Account, got)
		assert.NotEqual(t, nil, err)
	})
}

func TestUsers_Delete(t *testing.T) {
	clear(t)

	user := &store.User{
		Account:  "account",
		Password: "password",
		Email:    "email",
		Nick:     "nick",
	}

	err := u.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("unexpected error: %v\n", err)
	}

	t.Run("return nil", func(t *testing.T) {
		err := u.Delete(context.Background(), user.ID)
		assert.Equal(t, nil, err)
	})
}

func TestUsers_List(t *testing.T) {
	clear(t)
	length := 10
	for i := 0; i < length; i++ {
		user := &store.User{
			Account:  "account" + strconv.Itoa(i),
			Password: "password" + strconv.Itoa(i),
			Email:    "email" + strconv.Itoa(i),
			Nick:     "nick" + strconv.Itoa(i),
		}
		err := u.Create(context.Background(), user)
		assert.Equal(t, nil, err)
	}

	var list []UserDemo
	err := u.List(context.Background(), &list)
	assert.Equal(t, nil, err)
	assert.Equal(t, length, len(list))
	for i := 0; i < len(list); i++ {
		//t.Logf("%+v\n", list[i])
	}
	clear(t)
}

type UpdateUser struct {
	ID uint64 `json:"ID"`
	//Nick string `json:"nick"`
	Introduction string `json:"introduction"`
	University   string `json:"university"`
	Company      string `json:"company"`
	Blog         string `json:"blog"`
	Github       string `json:"github"`
}

func TestUsers_Update(t *testing.T) {
	clear(t)
	user := &store.User{
		Account:  "account",
		Password: "password",
		Email:    "email",
		Nick:     "nick",
	}
	err := u.Create(context.Background(), user)
	assert.Equal(t, nil, err)

	t.Run("return nil", func(t *testing.T) {
		want := &UpdateUser{
			//Nick: "nick",
			Introduction: "introduction",
			Blog:         "blog",
			Github:       "github",
		}
		err := u.Update(context.Background(), user.ID, want)
		assert.Equal(t, nil, err)

		got := &UpdateUser{}
		err = u.Get(context.Background(), user.ID, got)
		assert.Equal(t, nil, err)

		assert.Equal(t, want.Introduction, got.Introduction)

	})
}

func clear(t *testing.T) {
	t.Helper()
	err := u.db.Exec("DELETE FROM user").Error
	assert.Equal(t, err, nil)
}
