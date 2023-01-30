package opt

import (
	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/evt"
)

type colBoxOnly []evt.ColorEventBoxOption

func (options colBoxOnly) ApplyColorEventBox(e *evt.ColorEventBox) { panic("upgrade me") }
func (options colBoxOnly) ApplyColorEventBoxContext(ctx context.Context, b *evt.ColorEventBox) {
	for _, opt := range options {
		switch o := opt.(type) {
		case evt.ColorEventBoxContextOption:
			o.ApplyColorEventBoxContext(ctx, b)
		case evt.ColorEventBoxOption:
			o.ApplyColorEventBox(b)
		}
	}
}

func ColorBoxOnly(options ...evt.ColorEventBoxOption) evt.ColorEventBoxOption {
	return colBoxOnly(options)
}

type rotBoxOnly []evt.RotationEventBoxOption

func (options rotBoxOnly) ApplyRotationEventBox(e *evt.RotationEventBox) { panic("upgrade me") }
func (options rotBoxOnly) ApplyRotationEventBoxContext(ctx context.Context, b *evt.RotationEventBox) {
	for _, opt := range options {
		switch o := opt.(type) {
		case evt.RotationEventBoxContextOption:
			o.ApplyRotationEventBoxContext(ctx, b)
		case evt.RotationEventBoxOption:
			o.ApplyRotationEventBox(b)
		}
	}
}

func RotationBoxOnly(options ...evt.RotationEventBoxOption) evt.RotationEventBoxOption {
	return rotBoxOnly(options)
}
