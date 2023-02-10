package beatsaber

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/shasderias/iris/internal/swallowjson"
)

func OpenDifficultyV320(info *Info, path string) (Difficulty, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var diff DifficultyV320

	err = json.Unmarshal(f, &diff)
	if err != nil {
		return nil, err
	}

	diff.info = info
	diff.filepath = path

	return &diff, nil
}

type DifficultyV320 struct {
	info     *Info
	filepath string

	Version                  string                     `json:"version"`
	LightColorEventGroups    []LightColorEventGroupV320 `json:"lightColorEventBoxGroups"`
	LightRotationEventGroups []LightRotationGroupV320   `json:"lightRotationEventBoxGroups"`
	BasicEvent               []BasicEventV320           `json:"basicBeatmapEvents"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (d *DifficultyV320) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(d, "Extra", raw)
}

func (d DifficultyV320) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(d, "Extra")
}

func (d *DifficultyV320) Save() error {
	for i := range d.BasicEvent {
		d.BasicEvent[i].Time = Time(d.BasicEvent[i].Beat)
	}
	for i := range d.LightColorEventGroups {
		d.LightColorEventGroups[i].Time = Time(d.LightColorEventGroups[i].Beat)
	}
	for i := range d.LightRotationEventGroups {
		d.LightRotationEventGroups[i].Time = Time(d.LightRotationEventGroups[i].Beat)
	}

	sort.Slice(d.BasicEvent, func(i, j int) bool {
		return d.BasicEvent[i].Time < d.BasicEvent[j].Time
	})
	sort.Slice(d.LightColorEventGroups, func(i, j int) bool {
		return d.LightColorEventGroups[i].Time < d.LightColorEventGroups[j].Time
	})
	sort.Slice(d.LightRotationEventGroups, func(i, j int) bool {
		return d.LightRotationEventGroups[i].Time < d.LightRotationEventGroups[j].Time
	})

	bytes, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return os.WriteFile(d.filepath, bytes, 0644)
}

func (d *DifficultyV320) SetEvents(events []any) {
	d.BasicEvent = []BasicEventV320{}
	d.LightColorEventGroups = []LightColorEventGroupV320{}
	d.LightRotationEventGroups = []LightRotationGroupV320{}

	for _, e := range events {
		switch e := e.(type) {
		case BasicEventV320er:
			d.BasicEvent = append(d.BasicEvent, e.BasicEventV320()...)
		case LightColorEventGroupV320er:
			d.LightColorEventGroups = append(d.LightColorEventGroups, e.LightColorEventGroupV320()...)
		case LightRotationGroupV320er:
			d.LightRotationEventGroups = append(d.LightRotationEventGroups, e.LightRotationEventGroupV320()...)
		default:
			fmt.Printf("warning: unsupported event type %T\n", e)
		}
	}
}

func (d *DifficultyV320) DifficultyVersion() DifficultyVersion {
	return DifficultyVersion_3_0_0
}

type LightColorEventGroupV320 struct {
	Beat Float64 `json:"-"`

	Time    Time                `json:"b"`
	GroupID int                 `json:"g"`
	Boxes   []LightColorBoxV320 `json:"e"`
}

type LightColorBoxV320 struct {
	IndexFilter IndexFilterV320 `json:"f"`

	BeatDistributionParam             Float64 `json:"w"`
	BeatDistributionType              int     `json:"d"`
	BrightnessDistributionParam       Float64 `json:"r"`
	BrightnessDistributionType        int     `json:"t"`
	BrightnessDistributionAffectFirst int     `json:"b"`
	BrightnessDistributionEasing      int     `json:"i"` // 0 - Linear, 1 - In Quad, 2 - Out Quad,  3 - In Out Quad

	Events []LightColorEventV320 `json:"e"`
}

type LightColorEventV320 struct {
	Beat                 Float64 `json:"b"`
	TransitionType       int     `json:"i"`
	EnvironmentColorType int     `json:"c"`
	Brightness           Float64 `json:"s"`
	StrobeFrequency      Float64 `json:"f"`
}

type LightRotationGroupV320 struct {
	Beat Float64 `json:"-"`

	Time    Time                   `json:"b"`
	GroupID int                    `json:"g"`
	Boxes   []LightRotationBoxV320 `json:"e"`
}

type LightRotationBoxV320 struct {
	IndexFilter IndexFilterV320 `json:"f"`

	BeatDistributionParam     Float64 `json:"w"`
	BeatDistributionType      int     `json:"d"`
	RotationDistributionParam Float64 `json:"s"`
	RotationDistributionType  int     `json:"t"`
	Axis                      int     `json:"a"` // 0 - X, 1 - Y, 2 - Z
	FlipRotation              int     `json:"r"`
	AffectFirst               int     `json:"b"`
	Easing                    int     `json:"i"` // 0 - Linear, 1 - In Quad, 2 - Out Quad,  3 - In Out Quad

	Events []LightRotationEventV320 `json:"l"`
}

type LightRotationEventV320 struct {
	Beat                          Float64 `json:"b"`
	UsePreviousEventRotationValue int     `json:"p"`
	EaseType                      int     `json:"e"`
	LoopsCount                    int     `json:"l"`
	Rotation                      Float64 `json:"r"`
	RotationDirection             int     `json:"o"`
}

type IndexFilterV320 struct {
	FilterType   int     `json:"f"`
	ParamP       int     `json:"p"`
	ParamT       int     `json:"t"`
	Reverse      int     `json:"r"`
	Chunks       int     `json:"c"`
	Order        int     `json:"n"` // 0 - in order, 2 - random order, 3 - random starting index
	Seed         int     `json:"s"`
	Limit        Float64 `json:"l"`
	LimitAffects int     `json:"d"` // 0 - no other, 1 - duration, 2 - distribution, 3 - duration and distribution
}

type BasicEventV320 struct {
	Beat Float64 `json:"-"`

	Time       Time    `json:"b"`
	Type       int     `json:"et"`
	Value      int     `json:"i"`
	FloatValue Float64 `json:"f"`

	CustomData json.RawMessage `json:"_customData,omitempty"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (e *BasicEventV320) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(e, "Extra", raw)
}

func (e BasicEventV320) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(e, "Extra")
}

type LightColorEventGroupV320er interface {
	LightColorEventGroupV320() []LightColorEventGroupV320
}

type LightRotationGroupV320er interface {
	LightRotationEventGroupV320() []LightRotationGroupV320
}

type BasicEventV320er interface {
	BasicEventV320() []BasicEventV320
}
