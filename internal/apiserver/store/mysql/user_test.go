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
	"time"
)

var (
	u *users
)

func TestMain(m *testing.M) {
	// init mysql gorm
	logOpts := db.LogOption{
		LogLevel:                  4,
		SlowThreshold:             time.Millisecond * 200, // 200ms
		LogFile:                   "stdout",
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}

	opts := db.Option{
		Host:                  "127.0.0.1",
		Port:                  "3306",
		Username:              "root",
		Password:              "123456",
		Database:              "iam_test",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: 1 * time.Minute,
		LogOpt:                logOpts,
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
		got, err := u.Get(context.Background(), user.ID)
		assert.Equal(t, nil, err)
		assert.Equal(t, got.ID, user.ID)
	})

	t.Run("record not existed", func(t *testing.T) {
		_, err := u.Get(context.Background(), user.ID+1)
		if err == nil {
			t.Fatalf("want error(record not found) got nil\n")
		}
	})

}

func TestUsers_List(t *testing.T) {
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

	list, err := u.List(context.Background())
	assert.Equal(t, nil, err)
	assert.Equal(t, length, len(list.Items))
	clear(t)
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
		want := user
		want.Introduction = "introduction"
		err := u.Update(context.Background(), user.ID, want)
		assert.Equal(t, nil, err)
	})
}

func clear(t *testing.T) {
	t.Helper()
	err := u.db.Exec("DELETE FROM user").Error
	assert.Equal(t, err, nil)
}
