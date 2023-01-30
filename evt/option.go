package evt

import (
	"fmt"

	"github.com/shasderias/iris/context"
)

func ApplyOptions[T any](ctx context.Context, event any, options ...T) {
	switch e := event.(type) {
	case *BasicEvent:
		for _, opt := range options {
			switch o := any(opt).(type) {
			case BasicEventContextOption:
				o.ApplyBasicEventContext(ctx, e)
			case BasicEventOption:
				o.ApplyBasicEvent(e)
			}
		}
	case *ColorEventGroup:
		for _, opt := range options {
			switch o := any(opt).(type) {
			case ColorEventGroupContextOption:
				o.ApplyColorEventGroupContext(ctx, e)
			case ColorEventGroupOption:
				o.ApplyColorEventGroup(e)
			}
		}
	case *ColorEventBox:
		for _, opt := range options {
			switch o := any(opt).(type) {
			case ColorEventBoxContextOption:
				o.ApplyColorEventBoxContext(ctx, e)
			case ColorEventBoxOption:
				o.ApplyColorEventBox(e)
			}
		}
	case *ColorEvent:
		for _, opt := range options {
			switch o := any(opt).(type) {
			case ColorEventContextOption:
				o.ApplyColorEventContext(ctx, e)
			case ColorEventOption:
				o.ApplyColorEvent(e)
			}
		}
	case *RotationEventGroup:
		for _, opt := range options {
			switch o := any(opt).(type) {
			case RotationEventGroupContextOption:
				o.ApplyRotationEventGroupContext(ctx, e)
			case RotationEventGroupOption:
				o.ApplyRotationEventGroup(e)
			}
		}
	case *RotationEventBox:
		for _, opt := range options {
			switch o := any(opt).(type) {
			case RotationEventBoxContextOption:
				o.ApplyRotationEventBoxContext(ctx, e)
			case RotationEventBoxOption:
				o.ApplyRotationEventBox(e)
			}
		}
	case *RotationEvent:
		for _, opt := range options {
			switch o := any(opt).(type) {
			case RotationEventContextOption:
				o.ApplyRotationEventContext(ctx, e)
			case RotationEventOption:
				o.ApplyRotationEvent(e)
			}
		}
	default:
		panic(fmt.Sprintf("unsupported event type: %T", event))
	}
}
