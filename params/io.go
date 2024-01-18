// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"cogentcore.org/core/gi"
	"cogentcore.org/core/glop/indent"
	"cogentcore.org/core/grows/jsons"
	"cogentcore.org/core/grows/tomls"
	"golang.org/x/exp/maps"
)

// WriteGoPrelude writes the start of a go file in package main that starts a
// variable assignment to given variable -- for start of SaveGoCode methods.
func WriteGoPrelude(w io.Writer, varNm string) {
	w.Write([]byte("// File generated by params.SaveGoCode\n\n"))
	w.Write([]byte("package main\n\n"))
	w.Write([]byte(`import "github.com/emer/emergent/v2/params"`))
	w.Write([]byte("\n\nvar " + varNm + " = "))
}

// OpenJSON opens params from a JSON-formatted file.
func (pr *Params) OpenJSON(filename gi.Filename) error {
	*pr = make(Params) // reset
	return jsons.Open(pr, string(filename))
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Params) SaveJSON(filename gi.Filename) error {
	return jsons.Save(pr, string(filename))
}

// OpenTOML opens params from a TOML-formatted file.
func (pr *Params) OpenTOML(filename gi.Filename) error {
	*pr = make(Params) // reset
	return tomls.Open(pr, string(filename))
}

// SaveTOML saves params to a TOML-formatted file.
func (pr *Params) SaveTOML(filename gi.Filename) error {
	return tomls.Save(pr, string(filename))
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Params) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte("params.Params{\n"))
	depth++
	paths := make([]string, len(*pr)) // alpha-sort paths for consistent output
	ctr := 0
	for pt := range *pr {
		paths[ctr] = pt
		ctr++
	}
	sort.StringSlice(paths).Sort()
	for _, pt := range paths {
		pv := (*pr)[pt]
		w.Write(indent.TabBytes(depth))
		w.Write([]byte(fmt.Sprintf("%q: %q,\n", pt, pv)))
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("}"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Params) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Params) SaveGoCode(filename gi.Filename) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParams")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Hypers

// OpenJSON opens hypers from a JSON-formatted file.
func (pr *Hypers) OpenJSON(filename gi.Filename) error {
	*pr = make(Hypers) // reset
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves hypers to a JSON-formatted file.
func (pr *Hypers) SaveJSON(filename gi.Filename) error {
	return jsons.Save(pr, string(filename))
}

// OpenTOML opens params from a TOML-formatted file.
func (pr *Hypers) OpenTOML(filename gi.Filename) error {
	*pr = make(Hypers) // reset
	return tomls.Open(pr, string(filename))
}

// SaveTOML saves params to a TOML-formatted file.
func (pr *Hypers) SaveTOML(filename gi.Filename) error {
	return tomls.Save(pr, string(filename))
}

// WriteGoCode writes hypers to corresponding Go initializer code.
func (pr *Hypers) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte("params.Hypers{\n"))
	depth++
	paths := maps.Keys(*pr)
	sort.StringSlice(paths).Sort()
	for _, pt := range paths {
		pv := (*pr)[pt]
		w.Write(indent.TabBytes(depth))
		w.Write([]byte(fmt.Sprintf("%q: {", pt)))
		ks := maps.Keys(pv)
		sort.StringSlice(ks).Sort()
		for _, k := range ks {
			v := pv[k]
			w.Write([]byte(fmt.Sprintf("%q: %q,", k, v)))
		}
		w.Write([]byte("},\n"))
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("}"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Hypers) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves hypers to corresponding Go initializer code.
func (pr *Hypers) SaveGoCode(filename gi.Filename) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedHypers")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Sel

// OpenJSON opens params from a JSON-formatted file.
func (pr *Sel) OpenJSON(filename gi.Filename) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Sel) SaveJSON(filename gi.Filename) error {
	return jsons.Save(pr, string(filename))
}

// OpenTOML opens params from a TOML-formatted file.
func (pr *Sel) OpenTOML(filename gi.Filename) error {
	return tomls.Open(pr, string(filename))
}

// SaveTOML saves params to a TOML-formatted file.
func (pr *Sel) SaveTOML(filename gi.Filename) error {
	return tomls.Save(pr, string(filename))
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Sel) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte(fmt.Sprintf("Sel: %q, Desc: %q,\n", pr.Sel, pr.Desc)))
	depth++
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("Params: "))
	pr.Params.WriteGoCode(w, depth)
	if len(pr.Hypers) > 0 {
		w.Write([]byte(", Hypers: "))
		pr.Hypers.WriteGoCode(w, depth)
	}
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Sel) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Sel) SaveGoCode(filename gi.Filename) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParamsSel")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Sheet

// OpenJSON opens params from a JSON-formatted file.
func (pr *Sheet) OpenJSON(filename gi.Filename) error {
	*pr = make(Sheet, 0) // reset
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Sheet) SaveJSON(filename gi.Filename) error {
	return jsons.Save(pr, string(filename))
}

// OpenTOML opens params from a TOML-formatted file.
func (pr *Sheet) OpenTOML(filename gi.Filename) error {
	*pr = make(Sheet, 0) // reset
	return tomls.Open(pr, string(filename))
}

// SaveTOML saves params to a TOML-formatted file.
func (pr *Sheet) SaveTOML(filename gi.Filename) error {
	return tomls.Save(pr, string(filename))
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Sheet) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte("{\n"))
	depth++
	for _, pv := range *pr {
		w.Write(indent.TabBytes(depth))
		w.Write([]byte("{"))
		pv.WriteGoCode(w, depth)
		w.Write([]byte("},\n"))
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("},\n"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Sheet) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Sheet) SaveGoCode(filename gi.Filename) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParamsSheet")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Sheets

// OpenJSON opens params from a JSON-formatted file.
func (pr *Sheets) OpenJSON(filename gi.Filename) error {
	*pr = make(Sheets) // reset
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Sheets) SaveJSON(filename gi.Filename) error {
	return jsons.Save(pr, string(filename))
}

// OpenTOML opens params from a TOML-formatted file.
func (pr *Sheets) OpenTOML(filename gi.Filename) error {
	*pr = make(Sheets) // reset
	return tomls.Open(pr, string(filename))
}

// SaveTOML saves params to a TOML-formatted file.
func (pr *Sheets) SaveTOML(filename gi.Filename) error {
	return tomls.Save(pr, string(filename))
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Sheets) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte("params.Sheets{\n"))
	depth++
	nms := make([]string, len(*pr)) // alpha-sort names for consistent output
	ctr := 0
	for nm := range *pr {
		nms[ctr] = nm
		ctr++
	}
	sort.StringSlice(nms).Sort()
	for _, nm := range nms {
		pv := (*pr)[nm]
		w.Write(indent.TabBytes(depth))
		w.Write([]byte(fmt.Sprintf("%q: &", nm)))
		pv.WriteGoCode(w, depth)
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("}"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Sheets) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Sheets) SaveGoCode(filename gi.Filename) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParamsSheets")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Set

// OpenJSON opens params from a JSON-formatted file.
func (pr *Set) OpenJSON(filename gi.Filename) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Set) SaveJSON(filename gi.Filename) error {
	return jsons.Save(pr, string(filename))
}

// OpenTOML opens params from a TOML-formatted file.
func (pr *Set) OpenTOML(filename gi.Filename) error {
	return tomls.Open(pr, string(filename))
}

// SaveTOML saves params to a TOML-formatted file.
func (pr *Set) SaveTOML(filename gi.Filename) error {
	return tomls.Save(pr, string(filename))
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Set) WriteGoCode(w io.Writer, depth int, name string) {
	w.Write([]byte(fmt.Sprintf("Name: %q, Desc: %q, Sheets: ", name, pr.Desc)))
	pr.Sheets.WriteGoCode(w, depth)
}

/////////////////////////////////////////////////////////
//   Sets

// OpenJSON opens params from a JSON-formatted file.
func (pr *Sets) OpenJSON(filename gi.Filename) error {
	*pr = make(Sets) // reset
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Sets) SaveJSON(filename gi.Filename) error {
	return jsons.Save(pr, string(filename))
}

// OpenTOML opens params from a TOML-formatted file.
func (pr *Sets) OpenTOML(filename gi.Filename) error {
	*pr = make(Sets) // reset
	return tomls.Open(pr, string(filename))
}

// SaveTOML saves params to a TOML-formatted file.
func (pr *Sets) SaveTOML(filename gi.Filename) error {
	return tomls.Save(pr, string(filename))
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Sets) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte("params.Sets{\n"))
	depth++
	for nm, st := range *pr {
		w.Write(indent.TabBytes(depth))
		w.Write([]byte("{"))
		st.WriteGoCode(w, depth, nm)
		w.Write([]byte("},\n"))
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("}\n"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Sets) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Sets) SaveGoCode(filename gi.Filename) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParamsSets")
	pr.WriteGoCode(fp, 0)
	return nil
}

/*
var ParamsProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"Save", ki.PropSlice{
			{"SaveTOML", ki.Props{
				"label": "Save As TOML...",
				"desc":  "save to TOML formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"SaveJSON", ki.Props{
				"label": "Save As JSON...",
				"desc":  "save to JSON formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
			{"SaveGoCode", ki.Props{
				"label": "Save Code As...",
				"desc":  "save to Go-formatted initializer code in file",
				"icon":  "go",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".go",
					}},
				},
			}},
		}},
		{"Open", ki.PropSlice{
			{"OpenTOML", ki.Props{
				"label": "Open...",
				"desc":  "open from TOML formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"OpenJSON", ki.Props{
				"label": "Open...",
				"desc":  "open from JSON formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
	},
}

var HypersProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"Save", ki.PropSlice{
			{"SaveTOML", ki.Props{
				"label": "Save As TOML...",
				"desc":  "save to TOML formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"SaveJSON", ki.Props{
				"label": "Save As JSON...",
				"desc":  "save to JSON formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
			{"SaveGoCode", ki.Props{
				"label": "Save Code As...",
				"desc":  "save to Go-formatted initializer code in file",
				"icon":  "go",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".go",
					}},
				},
			}},
		}},
		{"Open", ki.PropSlice{
			{"OpenTOML", ki.Props{
				"label": "Open...",
				"desc":  "open from TOML formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"OpenJSON", ki.Props{
				"label": "Open...",
				"desc":  "open from JSON formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
	},
}

var SelProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"Save", ki.PropSlice{
			{"SaveTOML", ki.Props{
				"label": "Save As TOML...",
				"desc":  "save to TOML formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"SaveJSON", ki.Props{
				"label": "Save As JSON...",
				"desc":  "save to JSON formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
			{"SaveGoCode", ki.Props{
				"label": "Save Code As...",
				"desc":  "save to Go-formatted initializer code in file",
				"icon":  "go",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".go",
					}},
				},
			}},
		}},
		{"Open", ki.PropSlice{
			{"OpenTOML", ki.Props{
				"label": "Open...",
				"desc":  "open from TOML formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"OpenJSON", ki.Props{
				"label": "Open...",
				"desc":  "open from JSON formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
	},
}

var SheetProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"Save", ki.PropSlice{
			{"SaveTOML", ki.Props{
				"label": "Save As TOML...",
				"desc":  "save to TOML formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"SaveJSON", ki.Props{
				"label": "Save As JSON...",
				"desc":  "save to JSON formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
			{"SaveGoCode", ki.Props{
				"label": "Save Code As...",
				"desc":  "save to Go-formatted initializer code in file",
				"icon":  "go",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".go",
					}},
				},
			}},
		}},
		{"Open", ki.PropSlice{
			{"OpenTOML", ki.Props{
				"label": "Open...",
				"desc":  "open from TOML formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"OpenJSON", ki.Props{
				"label": "Open...",
				"desc":  "open from JSON formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
	},
}

var SheetsProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"Save", ki.PropSlice{
			{"SaveTOML", ki.Props{
				"label": "Save As TOML...",
				"desc":  "save to TOML formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"SaveJSON", ki.Props{
				"label": "Save As JSON...",
				"desc":  "save to JSON formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
			{"SaveGoCode", ki.Props{
				"label": "Save Code As...",
				"desc":  "save to Go-formatted initializer code in file",
				"icon":  "go",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".go",
					}},
				},
			}},
		}},
		{"Open", ki.PropSlice{
			{"OpenTOML", ki.Props{
				"label": "Open...",
				"desc":  "open from TOML formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"OpenJSON", ki.Props{
				"label": "Open...",
				"desc":  "open from JSON formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
		{"sep-diffs", ki.BlankProp{}},
		{"DiffsWithin", ki.Props{
			"desc":        "reports where the same param path is being set to different values within this set (both within the same Sheet and betwen sheets)",
			"icon":        "search",
			"show-return": true,
		}},
	},
}

var SetProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"Save", ki.PropSlice{
			{"SaveTOML", ki.Props{
				"label": "Save As TOML...",
				"desc":  "save to TOML formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"SaveJSON", ki.Props{
				"label": "Save As JSON...",
				"desc":  "save to JSON formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"Open", ki.PropSlice{
			{"OpenTOML", ki.Props{
				"label": "Open...",
				"desc":  "open from TOML formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"OpenJSON", ki.Props{
				"label": "Open...",
				"desc":  "open from JSON formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"sep-diffs", ki.BlankProp{}},
		{"DiffsWithin", ki.Props{
			"desc":        "reports where the same param path is being set to different values within this set (both within the same Sheet and betwen sheets)",
			"icon":        "search",
			"show-return": true,
		}},
	},
}

var SetsProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"Save", ki.PropSlice{
			{"SaveTOML", ki.Props{
				"label": "Save As TOML...",
				"desc":  "save to TOML formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"SaveJSON", ki.Props{
				"label": "Save As JSON...",
				"desc":  "save to JSON formatted file",
				"icon":  "file-save",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
			{"SaveGoCode", ki.Props{
				"label": "Save Code As...",
				"desc":  "save to Go-formatted initializer code in file",
				"icon":  "go",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".go",
					}},
				},
			}},
		}},
		{"Open", ki.PropSlice{
			{"OpenTOML", ki.Props{
				"label": "Open...",
				"desc":  "open from TOML formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".toml",
					}},
				},
			}},
			{"OpenJSON", ki.Props{
				"label": "Open...",
				"desc":  "open from JSON formatted file",
				"icon":  "file-open",
				"Args": ki.PropSlice{
					{"File Name", ki.Props{
						"ext": ".json",
					}},
				},
			}},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
		{"sep-diffs", ki.BlankProp{}},
		{"DiffsAll", ki.Props{
			"desc":        "between all sets, reports where the same param path is being set to different values",
			"icon":        "search",
			"show-return": true,
		}},
		{"DiffsFirst", ki.Props{
			"desc":        "between first set (e.g., the Base set) and rest of sets, reports where the same param path is being set to different values",
			"icon":        "search",
			"show-return": true,
		}},
		{"DiffsWithin", ki.Props{
			"desc":        "reports all the cases where the same param path is being set to different values within different sheets in given set",
			"icon":        "search",
			"show-return": true,
			"Args": ki.PropSlice{
				{"Set Name", ki.Props{}},
			},
		}},
	},
}
*/
