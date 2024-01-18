// Code generated by "core generate -add-types"; DO NOT EDIT.

package etime

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"cogentcore.org/core/enums"
)

var _ModesValues = []Modes{0, 1, 2, 3, 4, 5, 6}

// ModesN is the highest valid value
// for type Modes, plus one.
const ModesN Modes = 7

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ModesNoOp() {
	var x [1]struct{}
	_ = x[NoEvalMode-(0)]
	_ = x[AllModes-(1)]
	_ = x[Train-(2)]
	_ = x[Test-(3)]
	_ = x[Validate-(4)]
	_ = x[Analyze-(5)]
	_ = x[Debug-(6)]
}

var _ModesNameToValueMap = map[string]Modes{
	`NoEvalMode`: 0,
	`noevalmode`: 0,
	`AllModes`:   1,
	`allmodes`:   1,
	`Train`:      2,
	`train`:      2,
	`Test`:       3,
	`test`:       3,
	`Validate`:   4,
	`validate`:   4,
	`Analyze`:    5,
	`analyze`:    5,
	`Debug`:      6,
	`debug`:      6,
}

var _ModesDescMap = map[Modes]string{
	0: ``,
	1: `AllModes indicates that the log should occur over all modes present in other items.`,
	2: `Train is when the network is learning`,
	3: `Test is when testing, typically without learning`,
	4: `Validate is typically for a special held-out testing set`,
	5: `Analyze is when analyzing the representations and behavior of the network`,
	6: `Debug is for recording info particularly useful for debugging`,
}

var _ModesMap = map[Modes]string{
	0: `NoEvalMode`,
	1: `AllModes`,
	2: `Train`,
	3: `Test`,
	4: `Validate`,
	5: `Analyze`,
	6: `Debug`,
}

// String returns the string representation
// of this Modes value.
func (i Modes) String() string {
	if str, ok := _ModesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Modes value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Modes) SetString(s string) error {
	if val, ok := _ModesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _ModesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Modes")
}

// Int64 returns the Modes value as an int64.
func (i Modes) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Modes value from an int64.
func (i *Modes) SetInt64(in int64) {
	*i = Modes(in)
}

// Desc returns the description of the Modes value.
func (i Modes) Desc() string {
	if str, ok := _ModesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ModesValues returns all possible values
// for the type Modes.
func ModesValues() []Modes {
	return _ModesValues
}

// Values returns all possible values
// for the type Modes.
func (i Modes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ModesValues))
	for i, d := range _ModesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Modes.
func (i Modes) IsValid() bool {
	_, ok := _ModesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Modes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Modes) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Modes.UnmarshalText:", err)
	}
	return nil
}

var _TimesValues = []Times{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

// TimesN is the highest valid value
// for type Times, plus one.
const TimesN Times = 20

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _TimesNoOp() {
	var x [1]struct{}
	_ = x[NoTime-(0)]
	_ = x[AllTimes-(1)]
	_ = x[Cycle-(2)]
	_ = x[FastSpike-(3)]
	_ = x[GammaCycle-(4)]
	_ = x[Phase-(5)]
	_ = x[BetaCycle-(6)]
	_ = x[AlphaCycle-(7)]
	_ = x[ThetaCycle-(8)]
	_ = x[Event-(9)]
	_ = x[Trial-(10)]
	_ = x[Tick-(11)]
	_ = x[Sequence-(12)]
	_ = x[Epoch-(13)]
	_ = x[Block-(14)]
	_ = x[Condition-(15)]
	_ = x[Run-(16)]
	_ = x[Expt-(17)]
	_ = x[Scene-(18)]
	_ = x[Episode-(19)]
}

var _TimesNameToValueMap = map[string]Times{
	`NoTime`:     0,
	`notime`:     0,
	`AllTimes`:   1,
	`alltimes`:   1,
	`Cycle`:      2,
	`cycle`:      2,
	`FastSpike`:  3,
	`fastspike`:  3,
	`GammaCycle`: 4,
	`gammacycle`: 4,
	`Phase`:      5,
	`phase`:      5,
	`BetaCycle`:  6,
	`betacycle`:  6,
	`AlphaCycle`: 7,
	`alphacycle`: 7,
	`ThetaCycle`: 8,
	`thetacycle`: 8,
	`Event`:      9,
	`event`:      9,
	`Trial`:      10,
	`trial`:      10,
	`Tick`:       11,
	`tick`:       11,
	`Sequence`:   12,
	`sequence`:   12,
	`Epoch`:      13,
	`epoch`:      13,
	`Block`:      14,
	`block`:      14,
	`Condition`:  15,
	`condition`:  15,
	`Run`:        16,
	`run`:        16,
	`Expt`:       17,
	`expt`:       17,
	`Scene`:      18,
	`scene`:      18,
	`Episode`:    19,
	`episode`:    19,
}

var _TimesDescMap = map[Times]string{
	0:  `NoTime represents a non-initialized value, or a null result`,
	1:  `AllTimes indicates that the log should occur over all times present in other items.`,
	2:  `Cycle is the finest time scale -- typically 1 msec -- a single activation update.`,
	3:  `FastSpike is typically 10 cycles = 10 msec (100hz) = the fastest spiking time generally observed in the brain. This can be useful for visualizing updates at a granularity in between Cycle and GammaCycle.`,
	4:  `GammaCycle is typically 25 cycles = 25 msec (40hz)`,
	5:  `Phase is typically a Minus or Plus phase, where plus phase is bursting / outcome that drives positive learning relative to prediction in minus phase. It can also be used for other time scales involving multiple Cycles.`,
	6:  `BetaCycle is typically 50 cycles = 50 msec (20 hz) = one beta-frequency cycle. Gating in the basal ganglia and associated updating in prefrontal cortex occurs at this frequency.`,
	7:  `AlphaCycle is typically 100 cycles = 100 msec (10 hz) = one alpha-frequency cycle.`,
	8:  `ThetaCycle is typically 200 cycles = 200 msec (5 hz) = two alpha-frequency cycles. This is the modal duration of a saccade, the update frequency of medial temporal lobe episodic memory, and the minimal predictive learning cycle (perceive an Alpha 1, predict on 2).`,
	9:  `Event is the smallest unit of naturalistic experience that coheres unto itself (e.g., something that could be described in a sentence). Typically this is on the time scale of a few seconds: e.g., reaching for something, catching a ball.`,
	10: `Trial is one unit of behavior in an experiment -- it is typically environmentally defined instead of endogenously defined in terms of basic brain rhythms. In the minimal case it could be one ThetaCycle, but could be multiple, and could encompass multiple Events (e.g., one event is fixation, next is stimulus, last is response)`,
	11: `Tick is one step in a sequence -- often it is useful to have Trial count up throughout the entire Epoch but also include a Tick to count trials within a Sequence`,
	12: `Sequence is a sequential group of Trials (not always needed).`,
	13: `Epoch is used in two different contexts. In machine learning, it represents a collection of Trials, Sequences or Events that constitute a &#34;representative sample&#34; of the environment. In the simplest case, it is the entire collection of Trials used for training. In electrophysiology, it is a timing window used for organizing the analysis of electrode data.`,
	14: `Block is a collection of Trials, Sequences or Events, often used in experiments when conditions are varied across blocks.`,
	15: `Condition is a collection of Blocks that share the same set of parameters. This is intermediate between Block and Run levels. Aggregation of stats at this level is based on the last 5 rows by default.`,
	16: `Run is a complete run of a model / subject, from training to testing, etc. Often multiple runs are done in an Expt to obtain statistics over initial random weights etc. Aggregation of stats at this level is based on the last 5 rows by default.`,
	17: `Expt is an entire experiment -- multiple Runs through a given protocol / set of parameters.`,
	18: `Scene is a sequence of events that constitutes the next larger-scale coherent unit of naturalistic experience corresponding e.g., to a scene in a movie. Typically consists of events that all take place in one location over e.g., a minute or so. This could be a paragraph or a page or so in a book.`,
	19: `Episode is a sequence of scenes that constitutes the next larger-scale unit of naturalistic experience e.g., going to the grocery store or eating at a restaurant, attending a wedding or other &#34;event&#34;. This could be a chapter in a book.`,
}

var _TimesMap = map[Times]string{
	0:  `NoTime`,
	1:  `AllTimes`,
	2:  `Cycle`,
	3:  `FastSpike`,
	4:  `GammaCycle`,
	5:  `Phase`,
	6:  `BetaCycle`,
	7:  `AlphaCycle`,
	8:  `ThetaCycle`,
	9:  `Event`,
	10: `Trial`,
	11: `Tick`,
	12: `Sequence`,
	13: `Epoch`,
	14: `Block`,
	15: `Condition`,
	16: `Run`,
	17: `Expt`,
	18: `Scene`,
	19: `Episode`,
}

// String returns the string representation
// of this Times value.
func (i Times) String() string {
	if str, ok := _TimesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Times value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Times) SetString(s string) error {
	if val, ok := _TimesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _TimesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Times")
}

// Int64 returns the Times value as an int64.
func (i Times) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Times value from an int64.
func (i *Times) SetInt64(in int64) {
	*i = Times(in)
}

// Desc returns the description of the Times value.
func (i Times) Desc() string {
	if str, ok := _TimesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// TimesValues returns all possible values
// for the type Times.
func TimesValues() []Times {
	return _TimesValues
}

// Values returns all possible values
// for the type Times.
func (i Times) Values() []enums.Enum {
	res := make([]enums.Enum, len(_TimesValues))
	for i, d := range _TimesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Times.
func (i Times) IsValid() bool {
	_, ok := _TimesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Times) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Times) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Times.UnmarshalText:", err)
	}
	return nil
}
