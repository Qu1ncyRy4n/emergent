// Code generated by "goki generate -add-types"; DO NOT EDIT.

package chem

import (
	"goki.dev/gti"
)

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/chem.Buffer", IDName: "buffer", Doc: "Buffer provides a soft buffering driving deltas relative to a target N\nwhich can be set by concentration and volume.", Fields: []gti.Field{{Name: "K", Doc: "rate of buffering (akin to permeability / conductance of a channel)"}, {Name: "Target", Doc: "buffer target concentration -- drives delta relative to this"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/chem.Diffuse", IDName: "diffuse", Doc: "Diffuse models diffusion between two compartments A and B as\na function of concentration in each and potentially asymmetric\nrate constants: A Kf -> B and B Kb -> A\ncomputes the difference between each direction and applies to each", Fields: []gti.Field{{Name: "Kf", Doc: "A -> B forward diffusion rate constant, sec-1"}, {Name: "Kb", Doc: "B -> A backward diffusion rate constant, sec-1"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/chem.Enz", IDName: "enz", Doc: "Enz models an enzyme-catalyzed reaction based on the Michaelis-Menten kinetics\nthat transforms S = substrate into P product via SE-bound C complex\n\n\tK1         K3\n\nS + E --> C(SE) ---> P + E\n\n\t<-- K2\n\nS = substrate, E = enzyme, C = SE complex, P = product\nThe source K constants are in terms of concentrations μM-1 and sec-1\nbut calculations take place using N's, and the forward direction has\ntwo factors while reverse only has one, so a corrective volume factor needs\nto be divided out to set the actual forward factor.", Fields: []gti.Field{{Name: "K1", Doc: "S+E forward rate constant, in μM-1 msec-1"}, {Name: "K2", Doc: "SE backward rate constant, in μM-1 msec-1"}, {Name: "K3", Doc: "SE -> P + E catalyzed rate constant, in μM-1 msec-1"}, {Name: "Km", Doc: "Michaelis constant = (K2 + K3) / K1"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/chem.EnzRate", IDName: "enz-rate", Doc: "EnzRate models an enzyme-catalyzed reaction based on the Michaelis-Menten kinetics\nthat transforms S = substrate into P product via SE bound C complex\n\n\tK1         K3\n\nS + E --> C(SE) ---> P + E\n\n\t<-- K2\n\nS = substrate, E = enzyme, C = SE complex, P = product\nThis version does NOT consume the E enzyme or directly use the C complex\nas an accumulated factor: instead it directly computes an overall rate\nfor the end-to-end S <-> P reaction based on the K constants:\nrate = S * E * K3 / (S + Km)\nThis amount is added to the P and subtracted from the S, and recorded\nin the C complex variable as rate / K3 -- it is just directly set.\nIn some situations this C variable can be used for other things.\nThe source K constants are in terms of concentrations μM-1 and sec-1\nbut calculations take place using N's, and the forward direction has\ntwo factors while reverse only has one, so a corrective volume factor needs\nto be divided out to set the actual forward factor.", Fields: []gti.Field{{Name: "K1", Doc: "S+E forward rate constant, in μM-1 msec-1"}, {Name: "K2", Doc: "SE backward rate constant, in μM-1 msec-1"}, {Name: "K3", Doc: "SE -> P + E catalyzed rate constant, in μM-1 msec-1"}, {Name: "Km", Doc: "Michaelis constant = (K2 + K3) / K1 -- goes into the rate"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/chem.Paramer", IDName: "paramer", Doc: "The Paramer interface defines functions implemented for Params\nstructures, containing chem React, Enz, etc functions.\nThis interface is largely for documentation purposes."})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/chem.React", IDName: "react", Doc: "React models a basic chemical reaction:\n\n\tKf\n\nA + B --> AB\n\n\t<-- Kb\n\nwhere Kf is the forward and Kb is the backward time constant.\nThe source Kf and Kb constants are in terms of concentrations μM-1 and sec-1\nbut calculations take place using N's, and the forward direction has\ntwo factors while reverse only has one, so a corrective volume factor needs\nto be divided out to set the actual forward factor.", Directives: []gti.Directive{{Tool: "go", Directive: "generate", Args: []string{"goki", "generate", "-add-types"}}}, Fields: []gti.Field{{Name: "Kf", Doc: "forward rate constant for N / sec assuming 2 forward factors"}, {Name: "Kb", Doc: "backward rate constant for N / sec assuming 1 backward factor"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/chem.SimpleEnz", IDName: "simple-enz", Doc: "SimpleEnz models a simple enzyme-catalyzed reaction\nthat transforms S = substrate into P product via E which is not consumed\nassuming there is much more E than S and P -- E effectively acts as a\nrate constant multiplier\n\n\tKf*E\n\nS ----> P\n\nS = substrate, E = enzyme, P = product, Kf is the rate of the reaction", Fields: []gti.Field{{Name: "Kf", Doc: "S->P forward rate constant, in μM-1 msec-1"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/chem.Stater", IDName: "stater", Doc: "The Stater interface defines the functions implemented for State\nstructures containing chem state variables.\nThis interface is largely for documentation purposes."})
