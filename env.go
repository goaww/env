package env

import (
	"os"
	"strings"
)

type Environment interface {
	Get() map[string]string
}

type EnvironmentImpl struct {
	envPair map[string]string
}

func NewEnvironmentImpl() *EnvironmentImpl {
	env := &EnvironmentImpl{make(map[string]string)}
	for _, pair := range os.Environ() {
		kv := strings.Split(pair, "=")
		env.envPair[kv[0]] = kv[1]
	}
	return env
}

func (e *EnvironmentImpl) Get() map[string]string {
	return e.envPair
}
