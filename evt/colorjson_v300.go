package evt

//func (g *ColorEventGroup) LightColorEventGroupV300() beatsaber.LightColorEventGroupV300 {
//	return beatsaber.LightColorEventGroupV300{
//		Beat:    g.Beat,
//		GroupID: g.Group,
//		Boxes: iter.MapSlice(g.Boxes, func(b *ColorEventBox) beatsaber.LightColorBoxV300 {
//			return b.LightColorEventBoxV300()
//		}),
//	}
//}
//func (b *ColorEventBox) LightColorEventBoxV300() beatsaber.LightColorBoxV300 {
//	return beatsaber.LightColorBoxV300{
//		IndexFilter:                       b.IndexFilter.IndexFilterV300(),
//		BeatDistributionParam:             b.BeatDistribution.Value,
//		BeatDistributionType:              int(b.BeatDistribution.Type),
//		BrightnessDistributionParam:       b.BrightnessDistribution.Value,
//		BrightnessDistributionType:        int(b.BrightnessDistribution.Type),
//		BrightnessDistributionAffectFirst: boolToIntBool(b.BrightnessDistribution.AffectFirst),
//		Events: iter.MapSlice(b.Events, func(e *ColorEvent) beatsaber.LightColorEventV300 {
//			return e.LightColorEventV300()
//		}),
//	}
//}
//
//func (e *ColorEvent) LightColorEventV300() beatsaber.LightColorEventV300 {
//	return beatsaber.LightColorEventV300{
//		Beat:                 e.Beat,
//		TransitionType:       int(e.TransitionType),
//		EnvironmentColorType: int(e.Color),
//		Brightness:           e.Brightness,
//		StrobeFrequency:      e.StrobeFrequency,
//	}
//}
