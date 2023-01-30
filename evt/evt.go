package evt

import "github.com/shasderias/iris/context"

func getCtxOpts[T any](ctx context.Context) []T {
	opts := []T{}

	for _, opt := range ctx.Options() {
		if o, ok := opt.(T); ok {
			opts = append(opts, o)
		}
	}

	return opts
}

func getOptions[T any](ctx context.Context, argOptions []T) []T {
	opts := []T{}

	for _, opt := range ctx.Options() {
		if o, ok := opt.(T); ok {
			opts = append(opts, o)
		}
	}

	opts = append(opts, argOptions...)

	return opts
}

func optionsOf[T any](options ...any) []T {
	ret := []T{}
	for _, opt := range options {
		if o, ok := opt.(T); ok {
			ret = append(ret, o)
		}
	}
	return ret
}
