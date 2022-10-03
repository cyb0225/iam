/**
@author: yeebing
@date: 2022/10/3
**/

package gocache

import (
	"context"
	"github.com/cyb0225/iam/internal/apiserver/cache"
	gocache "github.com/cyb0225/iam/pkg/cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	tk *tokens
)

func TestMain(m *testing.M) {
	c, _ := gocache.New()
	tk = newTokens(&cacheStore{ca: c})
	tk.ca.Flush()
	m.Run()
}

func TestTokens(t *testing.T) {
	tk.ca.Flush()
	token := "token"
	val := &cache.TokenValue{UserID: 10}
	timeout := time.Second * 5

	// create
	err := tk.Create(context.Background(), token, val, timeout)
	assert.Equal(t, nil, err)

	// get
	t.Run("get success", func(t *testing.T) {
		got, err := tk.Get(context.Background(), token)
		assert.Equal(t, nil, err)
		assert.Equal(t, val.UserID, got.UserID)
	})

	t.Run("record not found", func(t *testing.T) {
		_, err := tk.Get(context.Background(), token+token)
		assert.NotEqual(t, nil, err)
	})

	// delete
	err = tk.Delete(context.Background(), token)
	assert.Equal(t, nil, err)

	_, err = tk.Get(context.Background(), token)
	assert.NotEqual(t, nil, err)

	// timeout
	//time.Sleep(timeout)
	//_, err = tk.Get(context.Background(), token)
	//assert.NotEqual(t, nil, err)
}
