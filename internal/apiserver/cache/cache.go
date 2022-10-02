/**
@author: yeebing
@date: 2022/10/2
**/

package cache

type Cache interface {
	Code() CodeCache
	Token() TokenCache
}
