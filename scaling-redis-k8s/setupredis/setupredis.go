package setupredis

import (
	"errors"

	retro "github.com/codeship/go-retro"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

var (
	ErrRetrySetup = retro.NewBackoffRetryableError(errors.New("retriable redis setup"), 30)
)

func New(addr, password string) (*pool.Pool, error) {
	p, err := pool.NewCustom("tcp", addr, 10,
		func(network, addr string) (*redis.Client, error) {
			client, err := redis.Dial(network, addr)
			if err != nil {
				return nil, err
			}

			if password != "" {
				if err = client.Cmd("AUTH", password).Err; err != nil {
					client.Close()
					return nil, err
				}
			}

			return client, nil
		})

	if err != nil {
		return nil, err
	}

	return p, nil
}

func NewWait(addr, password string) (*pool.Pool, error) {
	var redisPool *pool.Pool
	err := retro.DoWithRetry(func() error {
		p, err := New(addr, password)

		if err != nil {
			return ErrRetrySetup
		}

		redisPool = p

		return nil
	})

	if err != nil {
		return nil, err
	}

	return redisPool, nil
}
