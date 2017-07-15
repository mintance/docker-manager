package storage

import (
	"github.com/mediocregopher/radix.v2/redis"
	"strings"
	"github.com/Sirupsen/logrus"
)

type redisStorage struct {
	Storage
	connection *redis.Client
	credentials *Credentials
}

func GetRedisStorage(credentials *Credentials) Storage {

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

	return redisStorage{
		connection: conn,
		credentials: credentials,
	}
}

func (r redisStorage) Get(key string) (string, error) {

	return r.connection.Cmd("GET", key).Str()

}

func (r redisStorage) GetMap(key string) (map[string]string, error) {

	data, err := r.Get(key)

	data_map := map[string]string{}

	if err != nil {

		return data_map, err

	}

	data_array := strings.Split(data, "&")

	for _, item := range data_array {

		item_array := strings.Split(item, "=")

		data_map[item_array[0]] = item_array[1]

	}

	return data_map, nil
}

func (r redisStorage) Set(key string, value string) error {

	return r.connection.Cmd("SET", key, value).Err

}

func (r redisStorage) SetMap(key string, value map[string]string) error {

	data_array := []string{}

	for key, value := range value {

		data_array = append(data_array, strings.Join([]string{key, value}, "="))

	}

	return r.Set(key, strings.Join(data_array, "&"))
}

func (r redisStorage) Delete(key string) error {

	return r.connection.Cmd("DELETE", key).Err

}