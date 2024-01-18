// Code generated by "core generate -add-types"; DO NOT EDIT.

package stepper

import (
	"cogentcore.org/core/gti"
)

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/stepper.RunState", IDName: "run-state", Directives: []gti.Directive{{Tool: "enums", Directive: "enum"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/stepper.StopCheckFn", IDName: "stop-check-fn", Doc: "A StopCheckFn is a callback to check whether an arbitrary condition has been matched.\nIf a StopCheckFn returns true, the program is suspended with a RunState of Paused,\nand will remain so until the RunState changes to Stepping, Running, or Stopped.\nAs noted below for the PauseNotifyFn, the StopCheckFn is called with the Stepper's lock held."})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/stepper.PauseNotifyFn", IDName: "pause-notify-fn", Doc: "A PauseNotifyFn is a callback that will be invoked if the program enters the Paused state.\nNOTE! The PauseNotifyFn is called with the Stepper's lock held, so it must not call any Stepper methods\nthat try to take the lock on entry, or a deadlock will result."})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/stepper.Stepper", IDName: "stepper", Doc: "The Stepper struct contains all of the state info for stepping a program, enabling step points.\nwhere the running application can be suspended with no loss of state.", Fields: []gti.Field{{Name: "RunState", Doc: "current run state"}, {Name: "StepGrain", Doc: "granularity of one step. No enum type here so clients can define their own"}, {Name: "StepsPer", Doc: "number of steps to execute before returning"}, {Name: "PauseNotifyFn", Doc: "function to deal with any changes on client side when paused after stepping"}, {Name: "StopCheckFn", Doc: "function to test for special stopping conditions"}, {Name: "stateMut", Doc: "mutex for RunState"}, {Name: "stateChange", Doc: "state change condition variable"}, {Name: "stepsLeft", Doc: "number of steps yet to execute before returning"}, {Name: "waitTimer", Doc: "watchdog timer channel"}, {Name: "initOnce", Doc: "this ensures that global initialization only happens once"}}})
