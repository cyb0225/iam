/**
@author: yeebing
@date: 2022/10/2
**/

package cache

import (
	"context"
	"time"
)

type CodeValue struct {
	ID uint64
}

type CodeCache interface {
	Create(ctx context.Context, code string, value *CodeValue, timeout time.Duration) error
	Delete(ctx context.Context, code string) error
	Get(ctx context.Context, code string) (*CodeValue, error)
}
