package env

import (
	"reflect"
	"strconv"
	"strings"
)

type Binder interface {
	Bind(interface{})
}

type BinderImpl struct {
	Prefix string
	Env    Environment
}

func NewBinderImpl(prefix string) *BinderImpl {
	return &BinderImpl{Prefix: prefix, Env: NewEnvironmentImpl()}
}

func (b *BinderImpl) Bind(clazz interface{}) {

	if reflect.ValueOf(clazz).Kind() == reflect.Ptr && reflect.Indirect(reflect.ValueOf(clazz)).Kind() == reflect.Struct {
		v := reflect.Indirect(reflect.ValueOf(clazz))

		for i := 0; i < v.NumField(); i++ {
			key := strings.ToLower(b.Prefix + "." + v.Type().Field(i).Name)
			value := b.Env.Get(key)

			switch v.Field(i).Kind() {

			case reflect.Int8, reflect.Uint8:
				b.parseInt8(value, v, i)

			case reflect.Int16, reflect.Uint16:
				b.parseInt16(value, v, i)

			case reflect.Int32, reflect.Uint32:
				b.parseInt32(value, v, i)

			case reflect.Int64, reflect.Uint64:
				b.parseInt64(value, v, i)

			case reflect.Int, reflect.Uint:
				b.parseInt(value, v, i)

			case reflect.Float32:
				b.parseFloat32(value, v, i)

			case reflect.Float64:
				b.parseFloat64(value, v, i)

			case reflect.String:
				b.parseString(value, v, i)
			default:
				panic("Unsupported type")
			}
		}
	} else {
		panic("unsupported type, type must be pointer of struct")
	}
}

func (b *BinderImpl) parseInt8(value string, v reflect.Value, index int) {
	data, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		panic(err.Error())
	}
	v.Field(index).SetInt(data)
}

func (b *BinderImpl) parseInt16(value string, v reflect.Value, index int) {
	data, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		panic(err.Error())
	}
	v.Field(index).SetInt(data)
}

func (b *BinderImpl) parseInt32(value string, v reflect.Value, index int) {
	data, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic(err.Error())
	}
	v.Field(index).SetInt(data)
}

func (b *BinderImpl) parseInt64(value string, v reflect.Value, index int) {
	data, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	v.Field(index).SetInt(data)
}

func (b *BinderImpl) parseInt(value string, v reflect.Value, index int) {
	if strconv.IntSize == 32 {
		data, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			panic(err.Error())
		}
		v.Field(index).SetInt(data)
	} else if strconv.IntSize == 64 {
		data, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		v.Field(index).SetInt(data)
	} else {
		panic("Unsupported type")
	}
}

func (b *BinderImpl) parseFloat32(value string, v reflect.Value, index int) {
	data, err := strconv.ParseFloat(value, 32)
	if err != nil {
		panic(err.Error())
	}
	v.Field(index).SetFloat(data)
}

func (b *BinderImpl) parseFloat64(value string, v reflect.Value, index int) {
	data, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err.Error())
	}
	v.Field(index).SetFloat(data)
}

func (b *BinderImpl) parseString(value string, v reflect.Value, index int) {
	v.Field(index).SetString(value)
}
