/**
@author: yeebing
@date: 2022/10/2
**/

package cache

type CodeCache interface {
	Create()
	Delete()
	Get()
}
