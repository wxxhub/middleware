package config

import "fmt"

type Config interface {
	Map() map[string]interface{}
	Scan(v interface{}) error
}

type Watcher interface {
}

type Type string

var (
	LocalConfig Type = "local"
)

/*
var (
	// Default Config Manager
	DefaultConfig, _ = NewConfig()
)
*/

func NewConfig(configType Type, opts ...Option) (Config, Watcher, error) {
	switch configType {
	case LocalConfig:
		return newLocalConfig(opts...)
	default:
		return nil, nil, fmt.Errorf("umcomplete %s config", configType)
	}
}
