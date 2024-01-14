// Code generated by "goki generate -add-types"; DO NOT EDIT.

package decoder

import (
	"goki.dev/gti"
)

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/decoder.ActivationFunc", IDName: "activation-func"})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/decoder.Linear", IDName: "linear", Doc: "Linear is a linear neural network, which can be configured with a custom\nactivation function. By default it will use the identity function.\nIt learns using the delta rule for each output unit.", Fields: []gti.Field{{Name: "LRate", Doc: "learning rate"}, {Name: "Layers", Doc: "layers to decode"}, {Name: "Units", Doc: "unit values -- read this for decoded output"}, {Name: "NInputs", Doc: "number of inputs -- total sizes of layer inputs"}, {Name: "NOutputs", Doc: "number of outputs -- total sizes of layer inputs"}, {Name: "Inputs", Doc: "input values, copied from layers"}, {Name: "ValsTsrs", Doc: "for holding layer values"}, {Name: "Weights", Doc: "synaptic weights: outer loop is units, inner loop is inputs"}, {Name: "ActivationFn", Doc: "activation function"}, {Name: "PoolIndex", Doc: "which pool to use within a layer"}, {Name: "Comm", Doc: "mpi communicator -- MPI users must set this to their comm -- do direct assignment"}, {Name: "MPIDWts", Doc: "delta weight changes: only for MPI mode -- outer loop is units, inner loop is inputs"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/decoder.Layer", IDName: "layer", Doc: "Layer is the subset of emer.Layer that is used by this code"})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/decoder.LinearUnit", IDName: "linear-unit", Doc: "LinearUnit has variables for Linear decoder unit", Fields: []gti.Field{{Name: "Target", Doc: "target activation value -- typically 0 or 1 but can be within that range too"}, {Name: "Act", Doc: "final activation = sum x * w -- this is the decoded output"}, {Name: "Net", Doc: "net input = sum x * w"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/decoder.SoftMax", IDName: "soft-max", Doc: "SoftMax is a softmax decoder, which is the best choice for a 1-hot classification\nusing the widely-used SoftMax function: https://en.wikipedia.org/wiki/Softmax_function", Fields: []gti.Field{{Name: "Lrate", Doc: "learning rate"}, {Name: "Layers", Doc: "layers to decode"}, {Name: "NCats", Doc: "number of different categories to decode"}, {Name: "Units", Doc: "unit values"}, {Name: "Sorted", Doc: "sorted list of indexes into Units, in descending order from strongest to weakest -- i.e., Sortedhas the most likely categorization, and its activity is Units].Act"}, {Name: "NInputs", Doc: "number of inputs -- total sizes of layer inputs"}, {Name: "Inputs", Doc: "input values, copied from layers"}, {Name: "Target", Doc: "current target index of correct category"}, {Name: "ValsTsrs", Doc: "for holding layer values"}, {Name: "Weights", Doc: "synaptic weights: outer loop is units, inner loop is inputs"}, {Name: "Comm", Doc: "mpi communicator -- MPI users must set this to their comm -- do direct assignment"}, {Name: "MPIDWts", Doc: "delta weight changes: only for MPI mode -- outer loop is units, inner loop is inputs"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/decoder.SoftMaxUnit", IDName: "soft-max-unit", Doc: "SoftMaxUnit has variables for softmax decoder unit", Fields: []gti.Field{{Name: "Act", Doc: "final activation = e^Ge / sum e^Ge"}, {Name: "Net", Doc: "net input = sum x * w"}, {Name: "Exp", Doc: "exp(Net)"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/decoder.softMaxForSerialization", IDName: "soft-max-for-serialization", Fields: []gti.Field{{Name: "Weights"}}})
