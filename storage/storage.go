package storage

type Storage interface {
	Get(key string) (string, error)
	GetMap(key string) (map[string]interface{}, error)
	Set(key string, value string) error
	SetMap(key string, data map[string]interface{}) error
	Delete(key string) error
}

type Credentials struct {
	Host string
	Hosts []string
	Port string
	User string
	Password string
}

func GetStorage(storage_name string, credentials interface{}) Storage {
	switch storage_name {
	case "redis":
		return GetRedisStorage(credentials.(*Credentials))
	}

	return nil
}