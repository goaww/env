package env

import (
	"log"
	"reflect"
)

type Binder interface {
	Bind(interface{})
}

type BinderImpl struct {
	Env Environment
}

func NewBinderImpl() *BinderImpl {
	return &BinderImpl{NewEnvironmentImpl()}
}

func (b *BinderImpl) Bind(clazz interface{}) {
	t := reflect.TypeOf(clazz)
	v := reflect.ValueOf(clazz)
	log.Print(t, v)
}
