/**
@author: yeebing
@date: 2022/9/24
**/

package redis

import (
	"github.com/go-redis/redis"
)

var DB *redis.Client

type Option struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"` // if password is "", it means no password set.
	DB       int    `yaml:"DB"`
}

func (opts *Option) Valid() []error {
	var err []error

	return err
}

// New returns a client to the Redis Server specified by opts.
func New(opts Option) (*redis.Client, error) {
	DB = redis.NewClient(&redis.Options{
		Addr:     opts.Addr,
		Password: opts.Password,
		DB:       opts.DB,
	})

	// check connect
	_, err := DB.Ping().Result()
	if err != nil {
		return nil, err
	}

	return DB, nil
}
