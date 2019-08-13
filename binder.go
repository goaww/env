package env

type Binder interface {
	Bind(*interface{})
}

type BinderImpl struct {
	Env Environment
}

func NewBinderImpl() *BinderImpl {
	return &BinderImpl{NewEnvironmentImpl()}
}

func (b *BinderImpl) Bind(class *interface{}) {

}
