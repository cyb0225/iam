/**
@author: yeebing
@date: 2022/9/28
**/

package app

import "github.com/spf13/pflag"

type CliOption interface {
	Validate() []error
	Flags() *pflag.FlagSet
}
