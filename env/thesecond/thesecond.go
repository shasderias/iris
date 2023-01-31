package thesecond

import "github.com/shasderias/iris/evt"

var (
	Logo       = evt.OType(1)
	BRunway    = evt.OType(4)
	LeftFlags  = evt.OType(2)
	RightFlags = evt.OType(3)
	RingZoom   = evt.OType(9)
	Buildings  = evt.OType(0)
)

var (
	TopLasersLeftBottom     = evt.OEventGroup(6)
	TopLasersRightBottom    = evt.OEventGroup(7)
	TopLasersLeftTop        = evt.OEventGroup(8)
	TopLasersRightTop       = evt.OEventGroup(9)
	TopLasers               = evt.OEventGroup(6, 7, 8, 9)
	TopLasersLeft           = evt.OEventGroup(6, 8)
	TopLasersRight          = evt.OEventGroup(7, 9)
	TopLasersBottom         = evt.OEventGroup(6, 7)
	TopLasersTop            = evt.OEventGroup(8, 9)
	BottomLasersLeftBottom  = evt.OEventGroup(10)
	BottomLasersRightBottom = evt.OEventGroup(11)
	BottomLasersLeftTop     = evt.OEventGroup(12)
	BottomLasersRightTop    = evt.OEventGroup(13)
	BottomLasers            = evt.OEventGroup(10, 11, 12, 13)
	BottomLasersLeft        = evt.OEventGroup(10, 12)
	BottomLasersRight       = evt.OEventGroup(11, 13)
	BottomLasersBottom      = evt.OEventGroup(10, 11)
	BottomLasersTop         = evt.OEventGroup(12, 13)
	BigRing                 = evt.OEventGroup(0)
	SmallRing               = evt.OEventGroup(1)
	RunwayLeft              = evt.OColorEventGroup(2) // no rotation group
	RunwayRight             = evt.OColorEventGroup(3) // no rotation group
	Runway                  = evt.OColorEventGroup(2, 3)
	SpotlightLeft           = evt.OEventGroup(4)
	SpotlightRight          = evt.OEventGroup(5)
	Spotlight               = evt.OEventGroup(4, 5)
)
