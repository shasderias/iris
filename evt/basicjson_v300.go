package evt

import "github.com/shasderias/iris/beatsaber"

func (e *BasicEvent) BasicEventV300() []beatsaber.BasicEventV300 {
	return []beatsaber.BasicEventV300{{
		Beat:       e.Beat,
		Type:       e.Type,
		Value:      e.Value,
		FloatValue: e.FloatValue,
	}}
}

func (e *BasicEvent) BasicEventV320() []beatsaber.BasicEventV320 {
	return []beatsaber.BasicEventV320{{
		Beat:       beatsaber.Float64(e.Beat),
		Type:       e.Type,
		Value:      e.Value,
		FloatValue: beatsaber.Float64(e.FloatValue),
	}}
}
