// Copyright (c) 2022, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package egui

import (
	"fmt"
	"log"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/colors"
	"cogentcore.org/core/plot/plotview"
	"cogentcore.org/core/tensor/tensorview"
	"github.com/emer/emergent/v2/elog"
	"github.com/emer/emergent/v2/etime"
)

// AddPlots adds plots based on the unique tables we have,
// currently assumes they should always be plotted
func (gui *GUI) AddPlots(title string, lg *elog.Logs) {
	gui.Plots = make(map[etime.ScopeKey]*plotview.PlotView)
	// for key, table := range Log.Tables {
	for _, key := range lg.TableOrder {
		modes, times := key.ModesAndTimes()
		time := times[0]
		mode := modes[0]
		lt := lg.Tables[key] // LogTable struct
		if doplot, has := lt.Meta["Plot"]; has {
			if doplot == "false" {
				continue
			}
		}

		plt := gui.NewPlotTab(key, mode+" "+time+" Plot")
		plt.SetTable(lt.Table)
		plt.UpdatePlot()
		plt.Params.FromMetaMap(lt.Meta)

		ConfigPlotFromLog(title, plt, lg, key)
	}
}

func ConfigPlotFromLog(title string, plt *plotview.PlotView, lg *elog.Logs, key etime.ScopeKey) {
	_, times := key.ModesAndTimes()
	time := times[0]
	lt := lg.Tables[key] // LogTable struct

	for _, item := range lg.Items {
		_, ok := item.Write[key]
		if !ok {
			continue
		}
		cp := plt.SetColParams(item.Name, item.Plot, item.FixMin, item.Range.Min, item.FixMax, item.Range.Max)

		if item.Color != "" {
			cp.Color = errors.Log1(colors.FromString(item.Color, nil))
		}
		cp.TensorIndex = item.TensorIndex
		cp.ErrColumn = item.ErrCol

		plt.Params.Title = title + " " + time + " Plot"
		plt.Params.XAxisColumn = time
		if xaxis, has := lt.Meta["XAxisColumn"]; has {
			plt.Params.XAxisColumn = xaxis
		}
		if legend, has := lt.Meta["LegendColumn"]; has {
			plt.Params.LegendColumn = legend
		}
	}
	plt.ColumnsFromMetaMap(lt.Table.MetaData)
	plt.ColumnsFromMetaMap(lt.Meta)
	plt.Update()
}

// Plot returns plot for mode, time scope
func (gui *GUI) Plot(mode etime.Modes, time etime.Times) *plotview.PlotView {
	return gui.PlotScope(etime.Scope(mode, time))
}

// PlotScope returns plot for given scope
func (gui *GUI) PlotScope(scope etime.ScopeKey) *plotview.PlotView {
	if !gui.Active {
		return nil
	}
	plot, ok := gui.Plots[scope]
	if !ok {
		// fmt.Printf("egui Plot not found for scope: %s\n", scope)
		return nil
	}
	return plot
}

// SetPlot stores given plot in Plots map
func (gui *GUI) SetPlot(scope etime.ScopeKey, plt *plotview.PlotView) {
	if gui.Plots == nil {
		gui.Plots = make(map[etime.ScopeKey]*plotview.PlotView)
	}
	gui.Plots[scope] = plt
}

// UpdatePlot updates plot for given mode, time scope
func (gui *GUI) UpdatePlot(mode etime.Modes, time etime.Times) *plotview.PlotView {
	plot := gui.Plot(mode, time)
	if plot != nil {
		plot.GoUpdatePlot()
	}
	return plot
}

// UpdatePlotScope updates plot at given scope
func (gui *GUI) UpdatePlotScope(scope etime.ScopeKey) *plotview.PlotView {
	plot := gui.PlotScope(scope)
	if plot != nil {
		plot.GoUpdatePlot()
	}
	return plot
}

// UpdateCyclePlot updates cycle plot for given mode.
// only updates every CycleUpdateInterval
func (gui *GUI) UpdateCyclePlot(mode etime.Modes, cycle int) *plotview.PlotView {
	plot := gui.Plot(mode, etime.Cycle)
	if plot == nil {
		return plot
	}
	if (gui.CycleUpdateInterval > 0) && (cycle%gui.CycleUpdateInterval == 0) {
		plot.GoUpdatePlot()
	}
	return plot
}

// NewPlotTab adds a new plot with given key for Plots lookup
// and using given tab label.  For ad-hoc plots, you can
// construct a ScopeKey from any two strings using etime.ScopeStr.
func (gui *GUI) NewPlotTab(key etime.ScopeKey, tabLabel string) *plotview.PlotView {
	plt := plotview.NewSubPlot(gui.Tabs.NewTab(tabLabel))
	gui.Plots[key] = plt
	return plt
}

// AddTableView adds a table view of given log,
// typically particularly useful for Debug logs.
func (gui *GUI) AddTableView(lg *elog.Logs, mode etime.Modes, time etime.Times) *tensorview.TableView {
	if gui.TableViews == nil {
		gui.TableViews = make(map[etime.ScopeKey]*tensorview.TableView)
	}

	key := etime.Scope(mode, time)
	lt, ok := lg.Tables[key]
	if !ok {
		log.Printf("ERROR: in egui.AddTableView, log: %s not found\n", key)
		return nil
	}

	tt := gui.Tabs.NewTab(mode.String() + " " + time.String() + " ")
	tv := tensorview.NewTableView(tt)
	gui.TableViews[key] = tv
	tv.SetReadOnly(true)
	tv.SetTable(lt.Table)
	return tv
}

// TableView returns TableView for mode, time scope
func (gui *GUI) TableView(mode etime.Modes, time etime.Times) *tensorview.TableView {
	if !gui.Active {
		return nil
	}
	scope := etime.Scope(mode, time)
	tv, ok := gui.TableViews[scope]
	if !ok {
		fmt.Printf("egui TableView not found for scope: %s\n", scope)
		return nil
	}
	return tv
}

// UpdateTableView updates TableView for given mode, time scope
func (gui *GUI) UpdateTableView(mode etime.Modes, time etime.Times) *tensorview.TableView {
	tv := gui.TableView(mode, time)
	if tv != nil {
		tv.GoUpdateView()
	}
	return tv
}
