// Copyright (c) 2023, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package econfig

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"

	"golang.org/x/exp/maps"
)

// TestSubConfig is a sub-struct with special params
type TestSubConfig struct {

	// number of patterns to create
	NPats int `default:"10"`

	// proportion activity of created params
	Sparseness float32 `default:"0.15"`
}

// TestConfig is a testing config
type TestConfig struct {

	// specify include files here, and after configuration, it contains list of include files added
	Includes []string

	// open the GUI -- does not automatically run -- if false, then runs automatically and quits
	GUI bool `default:"true"`

	// use the GPU for computation
	GPU bool `default:"true"`

	// log debugging information
	Debug bool

	// important for testing . notation etc
	PatParams TestSubConfig

	// network parameters applied after built-in params -- use toml map format: '{key = val, key2 = val2}' where key is 'selector:path' (e.g., '.PFCLayer:Layer.Inhib.Layer.Gi' where '.PFCLayer' is a class) and values should be strings to be consistent with standard params format
	Network map[string]any

	// ParamSet name to use -- must be valid name as listed in compiled-in params or loaded params
	ParamSet string

	// Name of the JSON file to input saved parameters from.
	ParamFile string

	// Name of the file to output all parameter data. If not empty string, program should write file(s) and then exit
	ParamDocFile string

	// extra tag to add to file names and logs saved from this run
	Tag string

	// user note -- describe the run params etc -- like a git commit message for the run
	Note string `default:"testing is fun"`

	// starting run number -- determines the random seed -- runs counts from there -- can do all runs in parallel by launching separate jobs with each run, runs = 1
	Run int `default:"0"`

	// total number of runs to do when running Train
	Runs int `default:"10"`

	// total number of epochs per run
	Epochs int `default:"100"`

	// total number of trials per epoch.  Should be an even multiple of NData.
	NTrials int `default:"128"`

	// number of data-parallel items to process in parallel per trial -- works (and is significantly faster) for both CPU and GPU.  Results in an effective mini-batch of learning.
	NData int `default:"16"`

	// if true, save final weights after each run
	SaveWts bool

	// if true, save train epoch log to file, as .epc.tsv typically
	EpochLog bool `default:"true"`

	// if true, save run log to file, as .run.tsv typically
	RunLog bool `default:"true"`

	// if true, save train trial log to file, as .trl.tsv typically. May be large.
	TrialLog bool `default:"true"`

	// if true, save testing epoch log to file, as .tst_epc.tsv typically.  In general it is better to copy testing items over to the training epoch log and record there.
	TestEpochLog bool `default:"false"`

	// if true, save testing trial log to file, as .tst_trl.tsv typically. May be large.
	TestTrialLog bool `default:"false"`

	// if true, save network activation etc data from testing trials, for later viewing in netview
	NetData bool

	// can set these values by string representation if stringer and registered as an enum with kit
	Enum TestEnum

	// ] test slice case
	Slice []float32 `default:"[1, 2.14, 3.14]"`

	// ] test string slice case
	StrSlice []string `default:"['cat','dog one','dog two']"`
}

func (cfg *TestConfig) IncludesPtr() *[]string { return &cfg.Includes }

func TestDefaults(t *testing.T) {
	cfg := &TestConfig{}
	SetFromDefaults(cfg)
	if cfg.Epochs != 100 || cfg.EpochLog != true || cfg.Note != "testing is fun" {
		t.Errorf("Main defaults failed to set")
	}
	if cfg.PatParams.NPats != 10 || cfg.PatParams.Sparseness != 0.15 {
		t.Errorf("PatParams defaults failed to set")
	}
	// fmt.Printf("%#v\n", cfg.Slice)
	if len(cfg.Slice) != 3 || cfg.Slice[2] != 3.14 {
		t.Errorf("Slice defaults failed to set")
	}
	if len(cfg.StrSlice) != 3 || cfg.StrSlice[1] != "dog one" {
		t.Errorf("StrSlice defaults failed to set")
	}
}

func TestArgsPrint(t *testing.T) {
	t.Skip("prints all possible args")

	cfg := &TestConfig{}
	allArgs := make(map[string]reflect.Value)
	FieldArgNames(cfg, allArgs)

	keys := maps.Keys(allArgs)
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})
	fmt.Println("Args:")
	fmt.Println(strings.Join(keys, "\n"))
}

func TestArgs(t *testing.T) {
	cfg := &TestConfig{}
	SetFromDefaults(cfg)
	// note: cannot use "-Includes=testcfg.toml",
	args := []string{"-save-wts", "-nogui", "-no-epoch-log", "--NoRunLog", "--runs=5", "--run", "1", "--TAG", "nice", "--PatParams.Sparseness=0.1", "--Network", "{'.PFCLayer:Layer.Inhib.Gi' = '2.4', '#VSPatchPath:Path.Learn.LRate' = '0.01'}", "-Enum=TestValue2", "-Slice=[3.2, 2.4, 1.9]", "leftover1", "leftover2"}
	allArgs := make(map[string]reflect.Value)
	FieldArgNames(cfg, allArgs)
	leftovers, err := ParseArgs(cfg, args, allArgs, true)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(leftovers)
	if cfg.Runs != 5 || cfg.Run != 1 || cfg.Tag != "nice" || cfg.PatParams.Sparseness != 0.1 || cfg.SaveWts != true || cfg.GUI != false || cfg.EpochLog != false || cfg.RunLog != false {
		t.Errorf("args not set properly: %#v", cfg)
	}
	if cfg.Enum != TestValue2 {
		t.Errorf("args enum from string not set properly: %#v", cfg)
	}
	if len(cfg.Slice) != 3 || cfg.Slice[2] != 1.9 {
		t.Errorf("args Slice not set properly: %#v", cfg)
	}

	// if cfg.Network != nil {
	// 	mv := cfg.Network
	// 	for k, v := range mv {
	// 		fmt.Println(k, " = ", v)
	// 	}
	// }
}

func TestOpen(t *testing.T) {
	IncludePaths = []string{".", "testdata"}
	cfg := &TestConfig{}
	err := OpenWithIncludes(cfg, "testcfg.toml")
	if err != nil {
		t.Errorf(err.Error())
	}

	// fmt.Println("includes:", cfg.Includes)

	// if cfg.Network != nil {
	// 	mv := cfg.Network
	// 	for k, v := range mv {
	// 		fmt.Println(k, " = ", v)
	// 	}
	// }

	if cfg.GUI != true || cfg.Tag != "testing" {
		t.Errorf("testinc.toml not parsed\n")
	}
	if cfg.Epochs != 500 || cfg.GPU != true {
		t.Errorf("testinc2.toml not parsed\n")
	}
	if cfg.Note != "something else" {
		t.Errorf("testinc3.toml not parsed\n")
	}
	if cfg.Runs != 8 {
		t.Errorf("testinc3.toml didn't overwrite testinc2\n")
	}
	if cfg.NTrials != 32 {
		t.Errorf("testinc.toml didn't overwrite testinc2\n")
	}
	if cfg.NData != 12 {
		t.Errorf("testcfg.toml didn't overwrite testinc3\n")
	}
	if cfg.Enum != TestValue2 {
		t.Errorf("testinc.toml Enum value not parsed\n")
	}
}

func TestUsage(t *testing.T) {
	t.Skip("prints usage string")
	cfg := &TestConfig{}
	us := Usage(cfg)
	fmt.Println(us)
}

func TestSave(t *testing.T) {
	// t.Skip("prints usage string")
	IncludePaths = []string{".", "testdata"}
	cfg := &TestConfig{}
	OpenWithIncludes(cfg, "testcfg.toml")
	Save(cfg, "testdata/testwrite.toml")
}

func TestConfigOpen(t *testing.T) {
	// t.Skip("prints usage string")
	IncludePaths = []string{".", "testdata"}
	cfg := &TestConfig{}
	_, err := Config(cfg)
	// no errors for missing config fiels
	// if err == nil {
	// 	t.Errorf("should have Config error")
	// 	// } else {
	// 	// 	fmt.Println(err)
	// }
	_, err = Config(cfg, "aldfkj.toml")
	// if err == nil {
	// 	t.Errorf("should have Config error")
	// 	// } else {
	// 	// 	fmt.Println(err)
	// }
	_, err = Config(cfg, "aldfkj.toml", "testcfg.toml")
	if err != nil {
		t.Error(err)
	}
}

// TestIncConfig is a testing config with Include instead of Includes
type TestIncConfig struct {

	// specify include file here, and after configuration, it contains list of include files added
	Include string

	// open the GUI -- does not automatically run -- if false, then runs automatically and quits
	GUI bool `default:"true"`

	// extra tag to add to file names and logs saved from this run
	Tag string

	// starting run number -- determines the random seed -- runs counts from there -- can do all runs in parallel by launching separate jobs with each run, runs = 1
	Run int `default:"0"`
}

func (cfg *TestIncConfig) IncludePtr() *string { return &cfg.Include }

func TestIncOpen(t *testing.T) {
	IncludePaths = []string{".", "testdata"}
	cfg := &TestIncConfig{}
	err := OpenWithIncludes(cfg, "testcfginc.toml")
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("include:", cfg.Include)

	if cfg.GUI != true {
		t.Errorf("testcfginc.toml not parsed\n")
	}
	if cfg.Tag != "initial" {
		t.Errorf("testincinc.toml not parsed\n")
	}
}
