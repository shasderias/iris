package opt

import "github.com/shasderias/iris/context"

type KVOpt struct {
	Key   string
	Value any
}

func Set(key string, value any) KVOpt {
	return KVOpt{key, value}
}

func Get[T any](key string, defaultValue T, ctx context.Context, options []any) T {
	options = append(options, ctx.Options()...)
	v, _ := get[T](key, defaultValue, options)
	return v
}

func get[T any](key string, defaultValue T, options []any) (T, bool) {
	for _, opt := range options {
		switch o := opt.(type) {
		case KVOpt:
			if o.Key == key {
				return o.Value.(T), true
			}
		case Combined:
			if v, ok := get[T](key, defaultValue, o); ok {
				return v, true
			}
		}
	}
	return defaultValue, false
}
