package storage

import (
	"github.com/mediocregopher/radix.v2/redis"
	"strings"
	"github.com/Sirupsen/logrus"
)

type redisStorage struct {
	connection *redis.Client
	credentials Credentials
}

func GetRedisStorage(credentials *Credentials) *redisStorage {

	host := ""

	if credentials.Host == "" {

		host = strings.Join(credentials.Hosts, ",")

	} else {

		host = credentials.Host

	}

	conn, err := redis.Dial("tcp", host)

	if err != nil {

		logrus.Error("Can`t connect to Redis: ", err.Error())

		return nil
	}

	return &redisStorage{
		connection: conn,
		credentials: credentials,
	}
}

func (r redisStorage) Get(key string) (string, error) {

	return r.connection.Cmd("GET", key).Str()

}

func (r redisStorage) Set(key string, value string) error {

	return r.connection.Cmd("SET", key, value).Err

}

func (r redisStorage) Delete(key string) error {

	return r.connection.Cmd("DELETE", key).Err

}