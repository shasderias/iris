package evt

import (
	"github.com/shasderias/iris/iter"
)

type EventGroupOption struct {
	Group []int
}

func OEventGroup(g ...int) EventGroupOption {
	return EventGroupOption{Group: g}
}

func (o EventGroupOption) ApplyRotationEventGroup(g *RotationEventGroup) {
	for _, gid := range o.Group {
		if !iter.SliceHas(g.Group, gid) {
			g.Group = append(g.Group, gid)
		}
	}
}
func (o EventGroupOption) ApplyColorEventGroup(g *ColorEventGroup) {
	for _, gid := range o.Group {
		if !iter.SliceHas(g.Group, gid) {
			g.Group = append(g.Group, gid)
		}
	}
}

type DistributionType int

const (
	DistributionTypeWave DistributionType = 1
	DistributionTypeStep DistributionType = 2
)

type BeatDistribution struct {
	Type  DistributionType
	Param float64
}

func OBeatDistWave(beats float64) BeatDistribution {
	return BeatDistribution{
		Type:  DistributionTypeWave,
		Param: beats,
	}
}

func OBeatDistStep(beats float64) BeatDistribution {
	return BeatDistribution{
		Type:  DistributionTypeStep,
		Param: beats,
	}
}

func (d BeatDistribution) ApplyColorEventBox(b *ColorEventBox)       { b.BeatDistribution = d }
func (d BeatDistribution) ApplyRotationEventBox(b *RotationEventBox) { b.BeatDistribution = d }

// DistributionOption represents the type and value of a brightness, rotation or translation
// distribution, depending on the context.
type DistributionOption struct {
	Type  DistributionType
	Value float64
}

func ODistWave(v float64) DistributionOption {
	return DistributionOption{
		Type:  DistributionTypeWave,
		Value: v,
	}
}

func ODistStep(v float64) DistributionOption {
	return DistributionOption{
		Type:  DistributionTypeStep,
		Value: v,
	}
}

func (d DistributionOption) ApplyColorEventBox(b *ColorEventBox) {
	b.BrightnessDistribution.Type = d.Type
	b.BrightnessDistribution.Param = d.Value
}

func (d DistributionOption) ApplyRotationEventBox(b *RotationEventBox) {
	b.RotationDistribution.Type = d.Type
	b.RotationDistribution.Param = d.Value
}

type distAffectFirstOption struct {
	AffectFirst bool
}

func ODistAffectFirst(b bool) distAffectFirstOption {
	return distAffectFirstOption{
		AffectFirst: b,
	}
}

func (d distAffectFirstOption) ApplyColorEventBox(b *ColorEventBox) {
	b.BrightnessDistribution.AffectFirst = d.AffectFirst
}
func (d distAffectFirstOption) ApplyRotationEventBox(b *RotationEventBox) {
	b.RotationDistribution.AffectFirst = d.AffectFirst
}

var AffectFirst = ODistAffectFirst(true)

type distFlipOption struct {
	Flip bool
}

func ODistFlip(b bool) distFlipOption {
	return distFlipOption{
		Flip: b,
	}
}

func (d distFlipOption) ApplyRotationEventBox(b *RotationEventBox) {
	b.RotationDistribution.Flip = d.Flip
}

var Flip = ODistFlip(true)

type Axis int

const (
	AxisX Axis = 0
	AxisY Axis = 1
	AxisZ Axis = 2
)

func (a Axis) ApplyRotationEventBox(b *RotationEventBox) { b.Axis = a }

type Easing int

const (
	EasingNone      Easing = -1
	EasingLinear    Easing = 0
	EasingInQuad    Easing = 1
	EasingOutQuad   Easing = 2
	EasingInOutQuad Easing = 3
)

func (e Easing) ApplyColorEventBox(b *ColorEventBox)       { b.Easing = e }
func (e Easing) ApplyRotationEventBox(b *RotationEventBox) { b.Easing = e }
func (e Easing) ApplyRotationEvent(evt *RotationEvent)     { evt.Easing = e }

type TransitionType struct {
	Color    ColorEventTransitionType
	Rotation RotationEventTransitionType
}

var (
	Instant    = TransitionType{ColorEventTransitionInstant, -1}
	Transition = TransitionType{ColorEventTransitionTransition, RotationTransitionTypeTransition}
	Extend     = TransitionType{ColorEventTransitionExtend, RotationTransitionTypeExtend}
)

func (t TransitionType) ApplyColorEvent(b *ColorEvent) {
	b.TransitionType = t.Color
}

func (t TransitionType) ApplyRotationEvent(b *RotationEvent) {
	if t.Rotation == -1 {
		panic("invalid transition type for rotation event")
	}
	b.TransitionType = t.Rotation
}
