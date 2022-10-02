/**
@author: yeebing
@date: 2022/9/25
**/

package store

type Factory interface {
	User() UserStore
}
