package db

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	Pool *redis.Pool
)

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitRedis(redisURI string) {
	Pool = newPool(redisURI)
	cleanupHook()
}

func CloseRedis() error {
	if Pool != nil {
		return Pool.Close()
	}
	return nil
}
