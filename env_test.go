package env

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewEnvironmentImpl(t *testing.T) {
	env := NewEnvironmentImpl()

	log.Print(env)

	assert.NotEmpty(t, env.envPair)

}
