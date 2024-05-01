// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package efuns has misc functions, such as Gaussian and Logistic,
// that are used in neural models, and do not have a home elsewhere.
package efuns

//go:generate core generate -add-types

import (
	"cogentcore.org/core/math32"
)

// GaussVecDistNoNorm returns the gaussian of the distance between two 2D vectors
// using given sigma standard deviation, without normalizing area under gaussian
// (i.e., max value is 1 at dist = 0)
func GaussVecDistNoNorm(a, b math32.Vector2, sigma float32) float32 {
	dsq := a.DistanceToSquared(b)
	return math32.FastExp((-0.5 * dsq) / (sigma * sigma))
}

// Gauss1DNoNorm returns the gaussian of a given x value, without normalizing
// (i.e., max value is 1 at x = 0)
func Gauss1DNoNorm(x, sig float32) float32 {
	x /= sig
	return math32.FastExp(-0.5 * x * x)
}
