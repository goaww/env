package env

import (
	"os"
	"strings"
)

type Environment interface {
	GetAll() map[string]string
	Get(string) string
}

type EnvironmentImpl struct {
	envPair map[string]string
}

func (e *EnvironmentImpl) GetAll() map[string]string {
	return e.envPair
}

func (e *EnvironmentImpl) Get(key string) string {
	return e.envPair[key]
}

func NewEnvironmentImpl() *EnvironmentImpl {
	env := &EnvironmentImpl{make(map[string]string)}
	for _, pair := range os.Environ() {
		kv := strings.Split(pair, "=")
		env.envPair[strings.ToLower(kv[0])] = strings.ToLower(kv[1])
	}
	return env
}
