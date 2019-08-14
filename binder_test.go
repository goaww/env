package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewBinderImpl(t *testing.T) {
	var env = struct {
		Port int
		Host string
	}{}

	_ = os.Setenv("test.Host", "localhost")
	_ = os.Setenv("test.Port", "8080")

	binder := NewBinderImpl("test")

	binder.Bind(&env)

	assert.Equal(t, "localhost", env.Host)
	assert.Equal(t, 8080, env.Port)

}
