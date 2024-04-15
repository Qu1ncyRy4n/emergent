// Code generated by "core generate -add-types"; DO NOT EDIT.

package ringidx

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/ringidx.FIx", IDName: "f-ix", Doc: "FIx is a fixed-length ring index structure -- does not grow\nor shrink dynamically.", Directives: []types.Directive{{Tool: "gosl", Directive: "start", Args: []string{"ringidx"}}}, Fields: []types.Field{{Name: "Zi", Doc: "the zero index position -- where logical 0 is in physical buffer"}, {Name: "Len", Doc: "the length of the buffer -- wraps around at this modulus"}, {Name: "pad"}, {Name: "pad1"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/ringidx.Index", IDName: "index", Doc: "Index is the ring index structure, maintaining starting index and length\ninto a ring-buffer with maximum length Max.  Max must be > 0 and Len <= Max.\nWhen adding new items would overflow Max, starting index is shifted over\nto overwrite the oldest items with the new ones.  No moving is ever\nrequired -- just a fixed-length buffer of size Max.", Directives: []types.Directive{{Tool: "go", Directive: "generate", Args: []string{"core", "generate", "-add-types"}}}, Fields: []types.Field{{Name: "StIndex", Doc: "the starting index where current data starts -- the oldest data is at this index, and continues for Len items, wrapping around at Max, coming back up at most to StIndex-1"}, {Name: "Len", Doc: "the number of items stored starting at StIndex.  Capped at Max"}, {Name: "Max", Doc: "the maximum number of items that can be stored in this ring"}}})
