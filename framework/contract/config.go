package contract

import "time"

const ConfigKey = "wm:config"

type Config interface {
	IsExist(key string) bool

	Get(key string) interface{}

	GetBool(key string) bool

	GetString(key string) string

	GetInt(key string) int

	GetFloat64(key string) float64

	GetTime(key string) time.Time

	GetIntSlice(key string) []int

	GetStringSlice(key string) []string

	GetStringMap(key string) map[string]interface{}

	GetStringMapString(key string) map[string]string

	GetStringMapStringSlice(key string) map[string][]string

	Load(key string,val interface{}) error
}
