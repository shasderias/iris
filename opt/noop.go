package opt

import "github.com/shasderias/iris/evt"

var Nop = nop{}

type nop struct{}

func (o nop) ApplyBasicEvent(e *evt.BasicEvent)                 {}
func (o nop) ApplyColorEventGroup(e *evt.ColorEventGroup)       {}
func (o nop) ApplyColorEventBox(e *evt.ColorEventBox)           {}
func (o nop) ApplyColorEvent(e *evt.ColorEvent)                 {}
func (o nop) ApplyRotationEventGroup(e *evt.RotationEventGroup) {}
func (o nop) ApplyRotationEventBox(e *evt.RotationEventBox)     {}
func (o nop) ApplyRotationEvent(e *evt.RotationEvent)           {}
