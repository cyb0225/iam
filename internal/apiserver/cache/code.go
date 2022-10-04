/**
@author: yeebing
@date: 2022/10/2
**/

package cache

import (
	"context"
)

type CodeValue struct {
	Email string
}

type CodeCache interface {
	Create(ctx context.Context, code string, value *CodeValue) error
	Delete(ctx context.Context, code string) error
	Get(ctx context.Context, code string) (*CodeValue, error)
}
