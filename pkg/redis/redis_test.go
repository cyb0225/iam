/**
@author: yeebing
@date: 2022/9/25
**/

package redis

import "testing"

func TestNew(t *testing.T) {
	opts := Option{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}
	_, err := New(opts)
	if err != nil {
		t.Fatalf("unexpectd error or redis not configured: %v", err)
	}
}
