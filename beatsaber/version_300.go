package beatsaber

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/shasderias/iris/internal/swallowjson"
)

func OpenDifficultyV300(info *Info, path string) (Difficulty, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var diff DifficultyV300

	err = json.Unmarshal(f, &diff)
	if err != nil {
		return nil, err
	}

	diff.info = info
	diff.filepath = path

	return &diff, nil
}

type DifficultyV300 struct {
	info     *Info
	filepath string

	Version                  string                     `json:"version"`
	LightColorEventGroups    []LightColorEventGroupV300 `json:"lightColorEventBoxGroups"`
	LightRotationEventGroups []LightRotationGroupV300   `json:"lightRotationEventBoxGroups"`
	BasicEvent               []BasicEventV300           `json:"basicBeatmapEvents"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (d *DifficultyV300) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(d, "Extra", raw)
}

func (d DifficultyV300) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(d, "Extra")
}

func (d *DifficultyV300) Save() error {
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

func (d *DifficultyV300) SetEvents(events []any) {
	d.BasicEvent = []BasicEventV300{}
	d.LightColorEventGroups = []LightColorEventGroupV300{}
	d.LightRotationEventGroups = []LightRotationGroupV300{}

	for _, e := range events {
		switch e := e.(type) {
		case BasicEventV300er:
			d.BasicEvent = append(d.BasicEvent, e.BasicEventV300())
		case LightColorEventGroupV300er:
			d.LightColorEventGroups = append(d.LightColorEventGroups, e.LightColorEventGroupV300())
		case LightRotationGroupV300er:
			d.LightRotationEventGroups = append(d.LightRotationEventGroups, e.LightRotationGroupV300())
		default:
			fmt.Printf("warning: unsupported event type %T\n", e)
		}
	}
}

func (d *DifficultyV300) DifficultyVersion() DifficultyVersion {
	return DifficultyVersion_3_0_0
}

type LightColorEventGroupV300 struct {
	Beat float64 `json:"-"`

	Time    Time                `json:"b"`
	GroupID int                 `json:"g"`
	Boxes   []LightColorBoxV300 `json:"e"`
}

type LightColorBoxV300 struct {
	IndexFilter IndexFilterV300 `json:"f"`

	BeatDistributionParam             float64 `json:"w"`
	BeatDistributionType              int     `json:"d"`
	BrightnessDistributionParam       float64 `json:"r"`
	BrightnessDistributionType        int     `json:"t"`
	BrightnessDistributionAffectFirst int     `json:"b"`

	Events []LightColorEventV300 `json:"e"`
}

type LightColorEventV300 struct {
	Beat                 float64 `json:"b"`
	TransitionType       int     `json:"i"`
	EnvironmentColorType int     `json:"c"`
	Brightness           float64 `json:"s"`
	StrobeFrequency      float64 `json:"f"`
}

type LightRotationGroupV300 struct {
	Beat float64 `json:"-"`

	Time    Time                   `json:"b"`
	GroupID int                    `json:"g"`
	Boxes   []LightRotationBoxV300 `json:"e"`
}

type LightRotationBoxV300 struct {
	IndexFilter IndexFilterV300 `json:"f"`

	BeatDistributionParam     float64 `json:"w"`
	BeatDistributionType      int     `json:"d"`
	RotationDistributionParam float64 `json:"s"`
	RotationDistributionType  int     `json:"t"`
	Axis                      int     `json:"a"` // 0 - X, 1 - Y
	FlipRotation              int     `json:"r"`
	AffectFirst               int     `json:"b"`

	Events []LightRotationEventV300 `json:"l"`
}

type LightRotationEventV300 struct {
	Beat                          float64 `json:"b"`
	UsePreviousEventRotationValue int     `json:"p"`
	EaseType                      int     `json:"e"`
	LoopsCount                    int     `json:"l"`
	Rotation                      float64 `json:"r"`
	RotationDirection             int     `json:"o"`
}

type IndexFilterV300 struct {
	FilterType int `json:"f"`
	ParamP     int `json:"p"`
	ParamT     int `json:"t"`
	Reverse    int `json:"r"`
}

type BasicEventV300 struct {
	Beat float64 `json:"-"`

	Time       Time    `json:"b"`
	Type       int     `json:"et"`
	Value      int     `json:"i"`
	FloatValue float64 `json:"f"`

	CustomData json.RawMessage `json:"_customData,omitempty"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (e *BasicEventV300) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(e, "Extra", raw)
}

func (e BasicEventV300) MarshalJSON() ([]byte, error) {
	return swallowjson.MarshalWith(e, "Extra")
}

type LightColorEventGroupV300er interface {
	LightColorEventGroupV300() LightColorEventGroupV300
}

type LightRotationGroupV300er interface {
	LightRotationGroupV300() LightRotationGroupV300
}

type BasicEventV300er interface {
	BasicEventV300() BasicEventV300
}
