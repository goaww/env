package env

import (
	"log"
	"testing"
)

func TestNewBinderImpl(t *testing.T) {
	var env = struct {
		Port int
		Host string
	}{}
	log.Print(env)

	binder := NewBinderImpl()

	binder.Bind(nil)

}
