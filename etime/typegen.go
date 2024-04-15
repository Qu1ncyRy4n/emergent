// Code generated by "core generate -add-types"; DO NOT EDIT.

package etime

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/etime.Modes", IDName: "modes", Doc: "Modes are evaluation modes (Training, Testing, etc)"})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/etime.ScopeKey", IDName: "scope-key", Doc: "ScopeKey the associated string representation of a scope or scopes.\nThey include one or more Modes and one or more Times.\nIt is fully extensible with arbitrary mode and time strings --\nthe enums are a convenience for standard cases.\nUltimately a single mode, time pair is used concretely, but the\nAll* cases and lists of multiple can be used as a convenience\nto specify ranges"})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/etime.Times", IDName: "times", Doc: "Times the enum"})
