// Copyright (c) 2022, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package looper

import (
	"github.com/emer/emergent/env"
	"github.com/emer/emergent/etime"
)

// Stack contains one stack of nested loops, associated with one EvalMode
type Stack struct {
	Mode  string                   `desc:"eval mode for this stack"`
	Env   env.Env                  `desc:"environment used by default for loop iteration, stopping, if set"`
	Order []etime.ScopeKey         `desc:"order of the loops"`
	Loops map[etime.ScopeKey]*Loop `desc:"the loops by scope"`
	Step  Step                     `desc:"stepping state"`
}

func NewStack(mode etime.EvalModes, times ...etime.Times) *Stack {
	ord := make([]etime.ScopeKey, len(times))
	for i, t := range times {
		ord[i] = etime.Scope(mode, t)
	}
	return NewStackScope(ord...)
}

func NewStackScope(scopes ...etime.ScopeKey) *Stack {
	st := &Stack{}
	st.Order = scopes
	md, _ := st.Order[0].ModesAndTimes()
	st.Mode = md[0]
	st.Loops = make(map[etime.ScopeKey]*Loop, len(st.Order))
	for _, sc := range st.Order {
		st.Loops[sc] = NewLoop(sc)
	}
	return st
}

// Scope returns the top-level scope for this stack
func (st *Stack) Scope() etime.ScopeKey {
	if len(st.Order) > 0 {
		return st.Order[0]
	}
	return etime.ScopeKey("")
}

// Loop returns loop for given time
func (st *Stack) Loop(time etime.Times) *Loop {
	sc := etime.ScopeStr(st.Mode, time.String())
	return st.Loops[sc]
}

// Level returns loop for given level in order
func (st *Stack) Level(lev int) *Loop {
	if lev < 0 || lev >= len(st.Order) {
		return nil
	}
	return st.Loops[st.Order[lev]]
}

// StopCheck checks if it is time to stop, based on set.StopFlag,
// Env counters (if set), and loop Stop functions
func (st *Stack) StopCheck(set *Set, lp *Loop) bool {
	if st.Env != nil {
		// todo
	}
	return lp.Stop.Run()
}

// StepCheck checks if it is time to stop based on stepping
func (st *Stack) StepCheck(lp *Loop) bool {
	return st.Step.StopCheck(lp.Scope)
}

func (st *Stack) Run(set *Set) {
	lev := 0
	lp := st.Level(lev)
	lp.Start.Run()
	var nlp *Loop
	for {
		lp.Pre.Run()
		lev++
		nlp = st.Level(lev)
		if nlp != nil {
			lp = nlp
			lp.Start.Run()
			continue
		}
		lev--
	post:
		lp.Post.Run()
		stop := st.StopCheck(set, lp)
		if stop {
			lp.End.Run()
			lev--
			nlp = st.Level(lev)
			if nlp == nil {
				break
			}
			lp = nlp
			goto post
		} else {
			if st.StepCheck(lp) {
				break
			}
			if set.StopFlag {
				break
			}
		}
	}
}

// SetStep sets the stepping scope and n -- 0 = no stepping
// resets counter.
func (st *Stack) SetStep(time etime.Times, n int) {
	sc := etime.ScopeStr(st.Mode, time.String())
	st.SetStepScope(sc, n)
}

// SetStepScope sets the stepping scope and n -- 0 = no stepping
// resets counter.
func (st *Stack) SetStepScope(scope etime.ScopeKey, n int) {
	st.Step.Set(scope, n)
}

// StepClear resets stepping
func (st *Stack) StepClear() {
	st.Step.Clear()
}
