/**
@author: yeebing
@date: 2022/10/2
**/

package cache

import "context"

type TokenCache interface {
	Create(ctx context.Context)
	Get(ctx context.Context)
	Delete(ctx context.Context)
}
